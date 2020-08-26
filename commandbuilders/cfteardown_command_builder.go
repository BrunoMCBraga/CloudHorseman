package commandbuilders

import (
	"errors"
	"fmt"

	"github.com/BrunoMCBraga/CloudHorseman/commandcachegenerators"
	"github.com/BrunoMCBraga/CloudHorseman/internal"
	"github.com/BrunoMCBraga/CloudHorseman/menubuilders"
)

func runCFTeardown(currentOptions map[string]interface{}) {}

func loadCFTeardown(externalConfigurationMap map[string]interface{}, internalConfigurationMap map[string]interface{}) {

	for externalConfigurationKey, externalConfigurationValue := range externalConfigurationMap {
		if _, ok := internalConfigurationMap[externalConfigurationKey]; ok {
			internalConfigurationMap[externalConfigurationKey] = externalConfigurationValue
		}
	}
	return
}

func setOptionsCFTeardown(commandArray []string, currentOptions map[string]interface{}) error {

	setOptionsStubError := internal.SetOptionsStub(commandArray, currentOptions)
	if setOptionsStubError != nil {
		return errors.New("|CloudHorseman->commandbuilders->cfteardown_command_builder->setOptionsCFTeardown->internal.SetOptionsStub:" + setOptionsStubError.Error() + "|.")
	}

	return nil
}

func BuildCFTeardownCommand() (map[string]interface{}, error) {

	var currentOptions map[string]interface{} = make(map[string]interface{}, 0)

	currentOptions["cf_stack_name"] = ""
	currentOptions["s3_bucket_name"] = ""
	currentOptions["repo_name"] = ""

	listOfKeysToSet := make([]string, 0)

	for currentOptionsKey, _ := range currentOptions {
		listOfKeysToSet = append(listOfKeysToSet, currentOptionsKey)
	}

	setParametersBasedOnGlobalCacheAndStackNameError := commandcachegenerators.SetParametersBasedOnGlobalCacheAndStackName(listOfKeysToSet, currentOptions)
	if setParametersBasedOnGlobalCacheAndStackNameError != nil {
		fmt.Println("|CloudHorseman->commandbuilders->cfteardown_command_builder->BuildCFTeardownCommand->commandcachegenerators.SetParametersBasedOnGlobalCacheAndStackName:" + setParametersBasedOnGlobalCacheAndStackNameError.Error() + "|")
	}

	return BuildGenericCommand("cfteardown", currentOptions, menubuilders.GetCFTeardownOptionsMenu, menubuilders.GetCFTeardownInfoMenu, setOptionsCFTeardown, runCFTeardown, loadCFTeardown)

}
