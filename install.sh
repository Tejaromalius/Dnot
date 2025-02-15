#!/bin/sh

REPO_OWNER="Tejaromalius"
REPO_NAME="Dnot"
BINARY_NAME="dnot"
INSTALL_DIR="$HOME/.local/bin"

get_latest_release_url() {
  curl -s "https://api.github.com/repos/${REPO_OWNER}/${REPO_NAME}/releases/latest" |
    jq -r '.assets[] | select(.name | contains("'"$BINARY_NAME"'")) | .browser_download_url'
}

if ! command -v curl >/dev/null 2>&1; then
  echo "Error: curl is required but not found." >&2
  exit 1
fi

if ! command -v jq >/dev/null 2>&1; then
  echo "Error: jq is required but not found." >&2
  exit 1
fi

if ! command -v mktemp >/dev/null 2>&1; then
  echo "Error: mktemp is required but not found." >&2
  exit 1
fi

mkdir -p "$INSTALL_DIR"

download_url=$(get_latest_release_url)

if [ -z "$download_url" ]; then
  echo "Error: Could not find the latest release URL." >&2
  exit 1
fi


temp_file=$(mktemp)
if ! curl -sSL "$download_url" -o "$temp_file"; then
  echo "Error: Failed to download the release binary." >&2
  rm -f "$temp_file"
  exit 1
fi

if ! chmod +x "$temp_file"; then
  echo "Error: Failed to make the binary executable." >&2
  rm -f "$temp_file"
  exit 1
fi

if ! mv "$temp_file" "$INSTALL_DIR/$BINARY_NAME"; then
  echo "Error: Failed to move the binary to $INSTALL_DIR." >&2
  exit 1
fi

echo "Dnot installed successfully to $INSTALL_DIR"
echo "Make sure $INSTALL_DIR is in your \$PATH."

exit 0