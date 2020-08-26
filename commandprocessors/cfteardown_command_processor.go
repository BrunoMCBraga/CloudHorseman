package commandprocessors

import (
	"errors"

	"github.com/BrunoMCBraga/HayMakerCF/commandlineprocessors"
	"github.com/BrunoMCBraga/CloudHorseman/commandbuilders"
)

func processCFTeardownCommand(parametersMap map[string]interface{}) error {

	haymakerCommandMap := make(map[string]interface{}, 0)
	haymakerCommandMap["option"] = "tt"
	haymakerCommandMap["s3_bucket_name"] = parametersMap["s3_bucket_name"].(string)
	haymakerCommandMap["cf_stack_name"] = parametersMap["cf_stack_name"].(string)
	haymakerCommandMap["repo_name"] = parametersMap["repo_name"].(string)

	////go run ./main.go -cm tt -sn haymakerstack -bn haymakercfbucket -rn haymaker-docker-repo/haymaker-docker
	processCommandLineError := commandlineprocessors.ProcessCommandLine(haymakerCommandMap)
	if processCommandLineError != nil {
		return errors.New("|CloudHorseman->commandprocessors->cfteardown_command_processor->ProcessCFTeardownCommand->ProcessCommandLine:" + processCommandLineError.Error() + "|")
	}

	return nil
}

func ProcessCFTeardownCommandInteractively() error {

	cfTeardownCommandBuilderFunction := commandbuilders.BuildCFTeardownCommand
	cfTeardownCommandBuilderFunctionResult, cfTeardownCommandBuilderFunctionError := cfTeardownCommandBuilderFunction()
	if cfTeardownCommandBuilderFunctionError != nil {
		return errors.New("|CloudHorseman->commandprocessors->cfteardown_command_processor->ProcessCFTeardownCommandInteractively->handlerCommandBuilderFunction:" + cfTeardownCommandBuilderFunctionError.Error() + "|")
	} else if cfTeardownCommandBuilderFunctionResult == nil {
		return nil
	}

	processCFTeardownCommandError := processCFTeardownCommand(cfTeardownCommandBuilderFunctionResult)
	if processCFTeardownCommandError != nil {
		return errors.New("|CloudHorseman->commandprocessors->cfteardown_command_processor->ProcessCFTeardownCommandInteractively->processCFTeardownCommand:" + processCFTeardownCommandError.Error() + "|")
	}

	return nil
}

func ProcessCFTeardownCommandFromParametersMap(parametersMap map[string]interface{}) error {
	processCFTeardownCommandError := processCFTeardownCommand(parametersMap)
	if processCFTeardownCommandError != nil {
		return errors.New("|CloudHorseman->commandprocessors->cfteardown_command_processor->ProcessCFTeardownCommandFromParametersMap->processCFTeardownCommand:" + processCFTeardownCommandError.Error() + "|")
	}

	return nil
}
