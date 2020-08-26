package commandbuilders

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/BrunoMCBraga/CloudHorseman/globalstringsproviders"
	"github.com/BrunoMCBraga/CloudHorseman/internal"
	"github.com/BrunoMCBraga/CloudHorseman/types"
	"github.com/BrunoMCBraga/CloudHorseman/util"
)

func BuildGenericCommand(commandPromptArgument string,
	currentOptions map[string]interface{},
	moduleOptionsMenuFunction types.ModuleOptionsMenuFunction,
	moduleInfoMenuFunction types.ModuleInfoMenuFunction,
	moduleSetOptions types.ModuleSetOptions,
	moduleRun types.ModuleRun,
	moduleLoadOptions types.ModuleLoadOptions) (map[string]interface{}, error) {

	var commandReader *bufio.Reader = bufio.NewReader(os.Stdin)
	var loadCommandRegex string = "load [^ ]+"
	var setCommandRegex string = "set [^ ]+ [^ ]+"

	setCompileResult, setCompileError := regexp.Compile(setCommandRegex)
	if setCompileError != nil {
		return nil, errors.New("|CloudHorseman->commandbuilders->generic_command_builder->BuildGenericCommand->regexp.Compile(set):" + setCompileError.Error() + "|")
	}

	loadCompileResult, loadCompileError := regexp.Compile(loadCommandRegex)
	if loadCompileError != nil {
		return nil, errors.New("|CloudHorseman->commandbuilders->generic_command_builder->BuildGenericCommand->regexp.Compile(load):" + loadCompileError.Error() + "|")
	}

	for true {

		fmt.Print(globalstringsproviders.GetFormattedInteractivePromptString(commandPromptArgument))
		commandReader.Reset(os.Stdin)
		readStringResult, readStringError := commandReader.ReadString('\n')
		if readStringError != nil {
			fmt.Println("|CloudHorseman->commandbuilders->generic_command_builder->BuildGenericCommand->commandReader.ReadString:" + readStringError.Error() + "|")
			continue
		}
		command := strings.Trim(readStringResult, " \n")

		switch {
		case command == "show options":
			fmt.Println(moduleOptionsMenuFunction(currentOptions))
		case setCompileResult.MatchString(command):
			moduleSetOptionsError := moduleSetOptions(strings.Split(command, " ")[1:], currentOptions)

			if moduleSetOptionsError != nil {
				fmt.Println("|CloudHorseman->commandbuilders->generic_command_builder->BuildGenericCommand->moduleSetOptions:" + moduleSetOptionsError.Error() + "|")
				continue
			}
		case loadCompileResult.MatchString(command):

			splittedLoad := strings.Split(command, " ")
			if len(splittedLoad) != 2 {
				fmt.Println("|CloudHorseman->commandbuilders->generic_command_builder->BuildGenericCommand: invalid load command.|")
				continue
			}

			loadJSONFileIntoMapResult, loadJSONFileIntoMapError := util.LoadJSONFileIntoMap(splittedLoad[1])

			if loadJSONFileIntoMapError != nil {
				fmt.Println("|CloudHorseman->commandbuilders->generic_command_builder->BuildGenericCommand->util.LoadJSONFileIntoMap:" + loadJSONFileIntoMapError.Error() + "|")
				continue
			}

			generateInternalConfigurationMapFromExternalConfigurationMapResult, generateInternalConfigurationMapFromExternalConfigurationMapError := internal.GenerateInternalConfigurationMapFromExternalConfigurationMap(loadJSONFileIntoMapResult)
			if generateInternalConfigurationMapFromExternalConfigurationMapError != nil {
				fmt.Println("|CloudHorseman->commandbuilders->generic_command_builder->BuildGenericCommand->internal.GenerateInternalConfigurationMapFromExternalConfigurationMap:" + generateInternalConfigurationMapFromExternalConfigurationMapError.Error() + "|")
				continue
			}
			moduleLoadOptions(generateInternalConfigurationMapFromExternalConfigurationMapResult, currentOptions)

		case command == "run":
			moduleRun(currentOptions)
			return currentOptions, nil
		case command == "info":
			fmt.Println(moduleInfoMenuFunction(currentOptions))
		case command == "back":
			return nil, nil
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
	return nil, nil
}
