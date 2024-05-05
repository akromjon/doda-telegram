#!/bin/bash

set -e

URL_PREFIX="https://github.com/akromjon/doda-telegram/releases/download/1.1.0"
INSTALL_DIR="/usr/local/bin"

case "$(uname -sm)" in
  "Darwin x86_64") FILENAME="doda-telegram-darwin-amd64" ;;
  "Darwin arm64") FILENAME="doda-telegram-darwin-arm64" ;;
  "Linux x86_64") FILENAME="doda-telegram-linux-amd64" ;;
  "Linux i686") FILENAME="doda-telegram-linux-386" ;;
  "Linux armv7l") FILENAME="doda-telegram-linux-arm" ;;
  "Linux aarch64") FILENAME="doda-telegram-linux-arm64" ;;
  *) echo "Unsupported architecture: $(uname -sm)" >&2; exit 1 ;;
esac

echo "Downloading $FILENAME from github releases"
if ! curl -sSLf "$URL_PREFIX/$FILENAME" -o "$INSTALL_DIR/doda-cli"; then
  echo "Failed to write to $INSTALL_DIR; try with sudo" >&2
  exit 1
fi

if ! chmod +x "$INSTALL_DIR/doda-cli"; then
  echo "Failed to set executable permission on $INSTALL_DIR/doda-cli" >&2
  exit 1
fi

echo "doda-cli is successfully installed"