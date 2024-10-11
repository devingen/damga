package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

const descTG = `Usage: damga tg <add-target | remove-target>
Examples:
  # add a single target
  damga tg add-target --id abcd --target 1.2.3.4
  
  # remove a single target
  damga tg remove-target --id abcd --target 1.2.3.4
`

func handleTargetCommand() error {
	if len(os.Args) == 2 {
		fmt.Print(descTG)
		return nil
	}

	subCommand := os.Args[2]
	switch subCommand {
	case "add-target", "remove-target":
		tgCmd := flag.NewFlagSet("add", flag.ExitOnError)
		id := tgCmd.String("id", "", "id")
		target := tgCmd.String("target", "", "target")
		tgCmd.Parse(os.Args[3:])

		if *id == "" {
			return errors.New("Target group ID must be provided with --id")
		}
		if *target == "" {
			return errors.New("Target host must be provided with --target")
		}

		return addOrRemoveTargetsInTargetGroups(*id, *target, subCommand)
	default:
		fmt.Println("Expected 'add-target' or 'remove-target' but got '" + subCommand + "' instead.")
		os.Exit(1)
	}
	return nil
}

func addOrRemoveTargetsInTargetGroups(id, target, action string) error {
	responseBody := map[string]interface{}{}

	switch action {
	case "add-target", "remove-target":
		actionInPath := "add"
		if action == "remove-target" {
			actionInPath = "remove"
		}

		_, err := post(
			"/target-groups/"+id+"/targets/"+actionInPath,
			map[string]interface{}{"value": target},
			&responseBody,
		)
		if err != nil {
			return err
		}
		if errorMessage, hasError := responseBody["error"].(string); hasError {
			return errors.New(errorMessage)
		}
		return err
	default:
		return errors.New("Unexpected command '" + action + "'.")
	}
}
