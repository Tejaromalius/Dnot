# Dnot

[![Go Report Card](https://goreportcard.com/badge/github.com/Tejaromalius/Dnot)](https://goreportcard.com/report/github.com/Tejaromalius/Dnot)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/Tejaromalius/Dnot)](https://github.com/Tejaromalius/Dnot/releases/latest)
[![License: MIT](https://img.shields.io/badge/license-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Dnot** is a simple, interactive command-line tool written in Go that allows you to quickly find and run `.csproj` projects within the current directory. It uses a TUI (Terminal User Interface) built with the [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework.

## Features

*   **Interactive Project Selection:** Presents a clean, navigable list of `.csproj` files in the current directory.
*   **Easy Navigation:** Use arrow keys to navigate the list and Enter to select a project.
*   **Instant Execution:** Runs the selected project using `dotnet run --project`.
*   **Error Handling:** Gracefully handles errors, such as missing `.csproj` files or failures during project execution.

## Prerequisites

*   **Go:** You need Go (version 1.18 or later is recommended) installed on your system.  [Download Go](https://go.dev/dl/)
*   **dotnet SDK:** You need the .NET SDK installed to run the projects. [Download .NET SDK](https://dotnet.microsoft.com/en-us/download)

## Installation via `sh installer`

This script automates downloading the latest pre-built binary from GitHub and installing it to your `$HOME/.local/bin` directory.

```bash
curl -sSL https://raw.githubusercontent.com/Tejaromalius/Dnot/main/install.sh | sh
```

**Alternatively**, you can download the script, make it executable, and then run it:

```bash
curl -sSL https://raw.githubusercontent.com/Tejaromalius/Dnot/main/install.sh -o install.sh
chmod +x install.sh
./install.sh
```
**Important:**  Ensure that `$HOME/.local/bin` is in your `$PATH` environment variable.  You can usually add this to your shell's configuration file (e.g., `~/.bashrc`, `~/.zshrc`):

## Example

```
$ cd MyDotnetProjects
$ dnot

# A list of .csproj files appears.  Use arrow keys to navigate.

> 1. MyProject.csproj
  2. AnotherProject.csproj
  3. LibraryProject.csproj

# Press Enter on MyProject.csproj

# ... output from `dotnet run --project MyProject/MyProject.csproj` ...
```

## Why this tool?

Truth be told, I was tired of writing `dotnet run --project x.csproj` every time.

## Contributing

Contributions are welcome!  If you find a bug or have a feature request, please open an issue on GitHub.  If you'd like to contribute code, please fork the repository and submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE) - see the [LICENSE](LICENSE) file for details.
