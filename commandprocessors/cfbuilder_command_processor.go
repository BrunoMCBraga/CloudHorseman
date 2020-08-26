package commandprocessors

import (
	"errors"

	"github.com/BrunoMCBraga/CloudHorseman/commandbuilders"
	"github.com/BrunoMCBraga/HayMakerCF/commandlineprocessors"
)

func processCFBuilderCommand(parametersMap map[string]interface{}) error {

	/*go run ./main.go
	-cm td
	-t /Users/brubraga/go/src/github.com/haymakercf/CloudFormationFiles/cloudformation_cluster.json
	-sn haymakerstack
	-fk something
	-bn haymakerbucket
	-cn haymaker-eks*/

	haymakerCommandMap := make(map[string]interface{}, 0)

	haymakerCommandMap["option"] = "td"
	haymakerCommandMap["cf_template"] = parametersMap["cf_template"].(string)
	haymakerCommandMap["cf_stack_name"] = parametersMap["cf_stack_name"].(string)
	haymakerCommandMap["s3_bucket_name"] = parametersMap["s3_bucket_name"].(string)
	haymakerCommandMap["s3_cf_template_file_key"] = parametersMap["s3_cf_template_file_key"].(string)
	haymakerCommandMap["cluster_name"] = parametersMap["cluster_name"].(string)

	//(e.g. go run ./main.go -cm tt -sn haymakerstack -bn haymakercfbucket -rn haymaker-docker-repo/haymaker-docker)
	processCommandLineError := commandlineprocessors.ProcessCommandLine(haymakerCommandMap)
	if processCommandLineError != nil {
		return errors.New("|CloudHorseman->main->commandprocessors->cfbuilder_command_processor->processCFBuilderCommand->ProcessCommandLine:" + processCommandLineError.Error() + "|")
	}

	return nil
}

func ProcessCFBuilderCommandInteractively() error {

	cfBuilderCommandBuilderFunction := commandbuilders.BuildCFBuilderCommand
	cfBuilderCommandBuilderFunctionResult, cfBuilderCommandBuilderFunctionError := cfBuilderCommandBuilderFunction()
	if cfBuilderCommandBuilderFunctionError != nil {
		return errors.New("|CloudHorseman->commandprocessors->cfbuilder_command_processor->ProcessCFBuilderCommandInteractively->handlerCommandBuilderFunction:" + cfBuilderCommandBuilderFunctionError.Error() + "|")
	} else if cfBuilderCommandBuilderFunctionResult == nil {
		return nil
	}

	processCFBuilderCommandError := processCFBuilderCommand(cfBuilderCommandBuilderFunctionResult)
	if processCFBuilderCommandError != nil {
		return errors.New("|CloudHorseman->commandprocessors->cfbuilder_command_processor->ProcessCFBuilderCommandInteractively->processCFBuilderCommand:" + processCFBuilderCommandError.Error() + "|")
	}

	return nil
}

func ProcessCFBuilderCommandFromParametersMap(parametersMap map[string]interface{}) error {
	processCFBuilderCommandError := processCFBuilderCommand(parametersMap)
	if processCFBuilderCommandError != nil {
		return errors.New("|CloudHorseman->commandprocessors->cfbuilder_command_processor->ProcessCFBuilderCommandFromParametersMap->processCFBuilderCommand:" + processCFBuilderCommandError.Error() + "|")
	}

	return nil
}
