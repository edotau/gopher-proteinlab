#!/bin/bash

set -euo pipefail

release_list="https://golang.org/dl/"
SOURCE="https://storage.googleapis.com/golang"

# Figures out what OS you are using
os=$(uname -s | tr "[:upper:]" "[:lower:]")

# Determines the architecture of your current machine
arch=$(uname -m)
case "$arch" in
    i*|x86_64)
        arch="amd64"
        ;;
    x86)
        arch="386"
        ;;
    aarch64|arm64)
        arch="arm64"
        ;;
    armv7l)
        echo 'armv7l is not supported, using armv6l'
        arch="armv6l"
        ;;
    *)
        echo "Unsupported architecture: $arch"
        exit 1
        ;;
esac

# Queries for the list of Golang releases and extracts the latest stable version
queryReleaseList() {
    local fetch_command="$1"
    local rl="${release_list}?mode=json"
    if command -v jq &>/dev/null; then
        $fetch_command "$rl" | jq -r '.[].files[].version' | sort -V | uniq | grep -v -E 'go[0-9\.]+(beta|rc)' | sed -e 's/go//'
    else
        $fetch_command "$release_list" | grep -Eo 'go[0-9]+\.[0-9]+(\.[0-9]+)?' | grep -v -E '(beta|rc)' | sort -V | uniq
    fi
}

# Fetches the latest stable Golang release or a specific one
fetchUpdate() {
    local fetch_command
    if command -v wget &>/dev/null; then
        fetch_command="wget -qO-"
    elif command -v curl &>/dev/null; then
        fetch_command="curl -s"
    else
        echo "Error: Neither wget nor curl is installed."
        exit 3
    fi

    if [[ -n "${1:-}" ]]; then
        release="$1"
        echo "Selected release: Go $release"
    else
        release=$(queryReleaseList "$fetch_command" | tail -n 1)
        if [[ -z "$release" ]]; then
            echo "Error: Unable to determine the latest Go release."
            exit 2
        fi
        echo "Latest release: Go $release"
    fi
}

# Function to list all available Golang releases
listReleases() {
    local fetch_command
    if command -v wget &>/dev/null; then
        fetch_command="wget -qO-"
    elif command -v curl &>/dev/null; then
        fetch_command="curl -s"
    else
        echo "Error: Neither wget nor curl is installed."
        exit 3
    fi

    echo "Available Go releases:"
    queryReleaseList "$fetch_command"
    exit 0
}

# Extract and install Go to the HOME directory
installGo() {
    local go_root="$HOME/go"
    local go_path="$HOME/go_projects"
    
    echo "Extracting Go to $go_root"
    mkdir -p "$go_root"
    tar -C "$HOME" -xzf "$FILENAME"
    
    # Setting up GOROOT and GOPATH
    echo "Setting up GOROOT and GOPATH"
    if [[ "$SHELL" =~ "bash" ]]; then
        profile_file="$HOME/.bashrc"
    elif [[ "$SHELL" =~ "zsh" ]]; then
        profile_file="$HOME/.zshrc"
    else
        profile_file="$HOME/.profile"
    fi

    if ! grep -q 'export GOROOT=' "$profile_file"; then
        echo "export GOROOT=$go_root" >> "$profile_file"
    fi
    if ! grep -q 'export GOPATH=' "$profile_file"; then
        echo "export GOPATH=$go_path" >> "$profile_file"
    fi
    if ! grep -q "export PATH=.*$GOROOT/bin.*" "$profile_file"; then
        echo "export PATH=$PATH:$GOROOT/bin:$GOPATH/bin" >> "$profile_file"
    fi

    echo "GOROOT set to $go_root"
    echo "GOPATH set to $go_path"
    echo "Go environment variables added to $profile_file"
    echo "Please restart your terminal or run 'source $profile_file' to apply changes."
}

# Set variables
label=""
FILENAME=""
URL=""

# Handle positional arguments for listing or selecting specific release
if [[ $# -gt 0 ]]; then
    case "$1" in
        list)
            listReleases
            ;;
        install)
            if [[ -n "${2:-}" ]]; then
                fetchUpdate "$2"
            else
                echo "Please specify the version to install (e.g., ./script.sh install 1.20.1)"
                exit 4
            fi
            ;;
        *)
            echo "Unknown argument: $1"
            echo "Usage: $0 [list | install <version>]"
            exit 5
            ;;
    esac
else
    # No arguments: Install the latest release
    fetchUpdate
fi

# Set download URL
label="go${release}.${os}-${arch}"
FILENAME="${label}.tar.gz"
URL="${SOURCE}/${FILENAME}"

# Download and decompress the Golang tarball
echo "Downloading from: $URL"
if command -v wget &>/dev/null; then
    wget "$URL" -O "$FILENAME"
elif command -v curl &>/dev/null; then
    curl -O "$URL"
fi

echo "Downloaded: $FILENAME"

# Extract and install Go
installGo
