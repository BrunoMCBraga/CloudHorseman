package commandprocessors

import (
	"errors"

	"github.com/BrunoMCBraga/HayMakerCF/commandlineprocessors"
	"github.com/BrunoMCBraga/CloudHorseman/commandbuilders"
)

func processDepTeardownCommand(parametersMap map[string]interface{}) error {

	haymakerCommandMap := make(map[string]interface{}, 0)
	haymakerCommandMap["option"] = "sd"
	haymakerCommandMap["kubernetes_deployment_name"] = parametersMap["kubernetes_deployment_name"].(string)

	//go run ./main.go -cm ds -dn haymaker
	processCommandLineError := commandlineprocessors.ProcessCommandLine(haymakerCommandMap)
	if processCommandLineError != nil {
		return errors.New("|CloudHorseman->commandprocessors->depteardown_command_processor->processDepTeardownCommand->ProcessCommandLine:" + processCommandLineError.Error() + "|")
	}

	return nil
}

func ProcessDepTeardownCommandInteractively() error {

	depTeardownCommandBuilderFunction := commandbuilders.BuildDepTeardownCommand
	depTeardownCommandBuilderFunctionResult, depTeardownCommandBuilderFunctionError := depTeardownCommandBuilderFunction()
	if depTeardownCommandBuilderFunctionError != nil {
		return errors.New("|CloudHorseman->commandprocessors->depteardown_command_processor->ProcessDepTeardownCommandInteractively->handlerCommandBuilderFunction:" + depTeardownCommandBuilderFunctionError.Error() + "|")
	} else if depTeardownCommandBuilderFunctionResult == nil {
		return nil
	}

	processDepTeardownCommandError := processDepTeardownCommand(depTeardownCommandBuilderFunctionResult)
	if processDepTeardownCommandError != nil {
		return errors.New("|CloudHorseman->commandprocessors->depteardown_command_processor->ProcessDepTeardownCommandInteractively->processDepTeardownCommand:" + processDepTeardownCommandError.Error() + "|")
	}

	return nil
}

func ProcessDepTeardownCommandFromParametersMap(parametersMap map[string]interface{}) error {
	processDepTeardownCommandError := processDepTeardownCommand(parametersMap)
	if processDepTeardownCommandError != nil {
		return errors.New("|CloudHorseman->commandprocessors->depteardown_command_processor->ProcessDepTeardownCommandFromParametersMap->processDepTeardownCommand:" + processDepTeardownCommandError.Error() + "|")
	}

	return nil
}
