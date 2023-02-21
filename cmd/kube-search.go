package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"

var (
	kubesearch = &cobra.Command{
		Use:   "kubectl-search",
		Short: "A cli to find commands related to kubectl",
		Long:  `A cli to quickly find kubectl commands and documentation`,
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			kubecommands := readYaml()
			if result, err := searchCommands(args, kubecommands); err != nil {
				return err
			} else {
				printDocumentation(result)
			}
			return nil
		},
	}
)

func printDocumentation(commands []Command) {
	for _, v := range commands {
		fmt.Print(colorGreen, v.Cmd)
		fmt.Print(colorReset, " --> "+v.Description)
		fmt.Println()
	}
}

func searchCommands(args []string, kubecommands *KubeCommands) ([]Command, error) {

	var result []Command
	if args == nil {
		return result, errors.New("no input argument provided")
	}
	for _, command := range kubecommands.Commands {
		for _, input := range args {
			if strings.Contains(strings.ToLower(command.Command.Description), strings.ToLower(input)) || strings.Contains(strings.ToLower(command.Command.Cmd), strings.ToLower(input)) {
				result = append(result, command.Command)
			}
		}

	}

	if len(result) == 0 {
		fmt.Println(colorRed, "no commands matched words from input. Try a different search term")
		fmt.Print(colorReset)
	}
	return result, nil
}
