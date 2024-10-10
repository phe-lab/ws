# VSCode Workspace CLI

A simple CLI tool to quickly open VSCode Workspace.

![Screen Recording 2024-10-06 at 17 40 44](https://github.com/user-attachments/assets/84ae5c52-aa15-4b9c-b4ce-000190667e2d)

## Installation

Using [Homebrew](https://brew.sh):

```bash
brew install phe-lab/tap/ws
```

If you have a Go environment:

```bash
go get -u -v github.com/phe-lab/ws
```

## Usage

```bash
# List the workspaces
ws

# Open the workspace with the filename "simple-scrollspy.code-workspace"
ws simple-scrollspy

# Execute CLI with the logging level set to "debug"
ws --debug
```
