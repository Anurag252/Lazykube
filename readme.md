# `kubectl-command-cli`

A CLI tool that allows users to quickly find `kubectl` commands and their documentation.

## How to use

After installing the tool, use the following command to search for a `kubectl` command:

`kubectl-command-cli kubectl-command`


This command will search for all the `kubectl` commands that match the given input and display their descriptions.

## Code structure

The codebase is divided into the following files:

- `main.go`: The entry point of the application.
- `kubesearch.go`: Contains the search command and related functions.
- `readcheatsheet.go`: Contains the function to read the YAML file that contains the `kubectl` commands and their descriptions.
- `root.go`: The root command of the application.

## Dependencies

- `cobra`: A library for creating powerful modern CLI applications in Go.
- `yaml.v2`: A library for parsing and encoding YAML files in Go.

## Installation

To install the `kubectl-command-cli` tool, run the following command:

`go install anurag252/kubectl-cli/cmd/kubectl-command-cli`


This will install the tool in your `$GOBIN` directory.

## Usage

To use the tool, run the following command:

`kubectl-command-cli <command>`

Replace `<command>` with the search term you want to use to find `kubectl` commands.
