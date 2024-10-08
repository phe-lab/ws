# VSCode Workspace CLI

A simple CLI tool to quickly open VSCode workspaces.

https://github.com/user-attachments/assets/f4ce5609-51c8-4340-a0de-87d06491091f

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

# Set the logging level to "debug"
ws --debug
```
