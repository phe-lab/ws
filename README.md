# VSCode Workspace CLI

A simple CLI tool to quickly open VSCode Workspace.

![Screen Recording 2024-10-06 at 17 40 44](https://github.com/user-attachments/assets/84ae5c52-aa15-4b9c-b4ce-000190667e2d)

## Installation

Using [Homebrew](https://brew.sh):

```bash
brew install teguru-labs/tap/ws
```

If you have a Go environment:

```bash
go get -u -v github.com/teguru-labs/ws
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

## Configuration

### Environment Variables

| Name              | Default             | Description                                                |
| ----------------- | ------------------- | ---------------------------------------------------------- |
| `VSCODE_WS_PATH`  | `~/code-workspaces` | The directory path containing the `*.code-workspace` files |
| `VSCODE_WS_DEBUG` |                     | Enable the debug mode: `VSCODE_WS_DEBUG=true`              |
