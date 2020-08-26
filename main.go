package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/BrunoMCBraga/CloudHorseman/commandprocessors"
	"github.com/BrunoMCBraga/CloudHorseman/globalstringsproviders"
)

func main() {

	var setCommandRegex string = "use [a-zA-Z]+"

	var commandReader *bufio.Reader = bufio.NewReader(os.Stdin)

	useCompileResult, useCompileError := regexp.Compile(setCommandRegex)
	if useCompileError != nil {
		fmt.Println("|CloudHorseman->main->regexp.Compile(set):" + useCompileError.Error() + "|")
		return
	}

	fmt.Println(globalstringsproviders.GetMenuPictureString())

	for true {

		//fmt.Print(globalstringsproviders.GetFormattedInteractivePromptString(""))
		fmt.Print(globalstringsproviders.GetInteractivePromptString())
		commandReader.Reset(os.Stdin)
		readStringResult, readStringError := commandReader.ReadString('\n')
		if readStringError != nil {
			fmt.Println("|CloudHorseman->main->commandReader.ReadString:" + readStringError.Error() + "|")
			continue
		}

		command := strings.Trim(readStringResult, " \n")
		useCommandSplit := strings.Split(command, " ")

		switch {
		case useCompileResult.MatchString(command) && len(useCommandSplit) == 2:

			switch useCommandSplit[1] {
			case "avatar":
				processAvatarCommandInteractivelyError := commandprocessors.ProcessAvatarCommandInteractively()
				if processAvatarCommandInteractivelyError != nil {
					fmt.Println("|CloudHorseman->main->commandprocessors.ProcessAvatarCommandInteractively:" + processAvatarCommandInteractivelyError.Error() + "|")
					continue
				}
			case "cfbuilder":
				processCFBuilderCommandInteractivelyError := commandprocessors.ProcessCFBuilderCommandInteractively()
				if processCFBuilderCommandInteractivelyError != nil {
					fmt.Println("|CloudHorseman->main->commandprocessors.ProcessCFBuilderCommandInteractively:" + processCFBuilderCommandInteractivelyError.Error() + "|")
					continue
				}
			case "cfteardown":
				processCFTeardownCommandInteractivelyError := commandprocessors.ProcessCFTeardownCommandInteractively()
				if processCFTeardownCommandInteractivelyError != nil {
					fmt.Println("|CloudHorseman->main->commandprocessors.ProcessCFTeardownCommandInteractively:" + processCFTeardownCommandInteractivelyError.Error() + "|")
					continue
				}
			case "depteardown":
				processDepTeardownCommandInteractivelyError := commandprocessors.ProcessDepTeardownCommandInteractively()
				if processDepTeardownCommandInteractivelyError != nil {
					fmt.Println("|CloudHorseman->main->commandprocessors.ProcessDepTeardownCommandInteractively:" + processDepTeardownCommandInteractivelyError.Error() + "|")
					continue
				}
			case "handler":
				processPullerCommandInteractivelyError := commandprocessors.ProcessPullerCommandInteractively()
				if processPullerCommandInteractivelyError != nil {
					fmt.Println("|CloudHorseman->main->commandprocessors.ProcessPullerCommandInteractively:" + processPullerCommandInteractivelyError.Error() + "|")
					continue
				}
			case "pusher":
				processPusherCommandInteractivelyError := commandprocessors.ProcessPusherCommandInteractively()
				if processPusherCommandInteractivelyError != nil {
					fmt.Println("|CloudHorseman->main->commandprocessors.ProcessPusherCommandInteractively:" + processPusherCommandInteractivelyError.Error() + "|")
					continue
				}
			default:
				fmt.Printf("%s\n", "Unknown Module")
			}
		case command == "exit":
			return
		case command == "help" || command == "?":
			fmt.Println(globalstringsproviders.GetHelpString())
			continue
		case command == "version":
			fmt.Println(globalstringsproviders.GetVersionString())
			continue
		case command == "show":
			fmt.Println(globalstringsproviders.GetShowModulesString())
			continue
		default:
			fmt.Printf("%s\n", "Unknown Command")
		}
	}
}
