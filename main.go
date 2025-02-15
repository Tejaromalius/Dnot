package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	listHeight   = 14
	defaultWidth = 40
)

var (
	styles = struct {
		title        lipgloss.Style
		item         lipgloss.Style
		selectedItem lipgloss.Style
		pagination   lipgloss.Style
		help         lipgloss.Style
	}{
		title:        lipgloss.NewStyle().MarginLeft(2).Bold(true).Foreground(lipgloss.Color("#FAFAFA")),
		item:         lipgloss.NewStyle().PaddingLeft(4),
		selectedItem: lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170")).Bold(true),
		pagination:   list.DefaultStyles().PaginationStyle.PaddingLeft(4),
		help:         list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1),
	}
)

type item struct {
	path string
}

func (i item) FilterValue() string { return "" }

func (i item) String() string { return i.path }

type csprojDelegate struct{}

func (d csprojDelegate) Height() int                             { return 1 }
func (d csprojDelegate) Spacing() int                            { return 0 }
func (d csprojDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d csprojDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {

		fmt.Fprintf(os.Stderr, "ERROR: Invalid list item type: %T\n", listItem)
		return
	}

	s := fmt.Sprintf("%d. %s", index+1, filepath.Base(i.path))

	style := styles.item
	if index == m.Index() {
		style = styles.selectedItem
		s = "> " + s
	}

	fmt.Fprint(w, style.Render(s))
}

type model struct {
	list     list.Model
	selected string
	err      error
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:

		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":

			itm, ok := m.list.SelectedItem().(item)
			if !ok {

				m.err = fmt.Errorf("unexpected item type: %T", m.list.SelectedItem())
				return m, tea.Quit
			}
			m.selected = itm.path
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {

	if m.err != nil {
		return fmt.Sprintf("Error: %v\n", m.err)
	}

	if m.selected != "" {
		return ""
	}
	return "\n" + m.list.View()
}

func scanForProjects(root string) ([]list.Item, error) {
	var items []list.Item
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".csproj") {
			items = append(items, item{path: path})
		}
		return nil
	})
	return items, err
}

func runDotnetProject(projectPath string) error {
	cmd := exec.Command("dotnet", "run", "--project", projectPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start dotnet process: %w", err)
	}
	return cmd.Wait()
}

func main() {
	startPath, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	items, err := scanForProjects(startPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error scanning for projects: %v\n", err)
		os.Exit(1)
	}

	if len(items) == 0 {
		fmt.Println("No .csproj files found.")
		os.Exit(0)
	}

	l := list.New(items, csprojDelegate{}, defaultWidth, listHeight)
	l.Title = "Select .csproj to Run"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = styles.title
	l.Styles.PaginationStyle = styles.pagination
	l.Styles.HelpStyle = styles.help

	p := tea.NewProgram(model{list: l})
	m, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running program: %v\n", err)
		os.Exit(1)
	}

	if m, ok := m.(model); ok && m.selected != "" {
		if err := runDotnetProject(m.selected); err != nil {
			fmt.Fprintf(os.Stderr, "Error running project: %v\n", err)
			os.Exit(1)
		}
	} else if ok && m.err != nil {

		os.Exit(1)
	}
}
