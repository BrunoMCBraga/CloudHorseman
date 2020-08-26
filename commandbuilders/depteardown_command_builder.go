package commandbuilders

import (
	"errors"
	"fmt"

	"github.com/BrunoMCBraga/CloudHorseman/commandcachegenerators"
	"github.com/BrunoMCBraga/CloudHorseman/internal"
	"github.com/BrunoMCBraga/CloudHorseman/menubuilders"
)

func runDepTeardown(currentOptions map[string]interface{}) {}

func loadDepTeardown(externalConfigurationMap map[string]interface{}, internalConfigurationMap map[string]interface{}) {

	for externalConfigurationKey, externalConfigurationValue := range externalConfigurationMap {
		if _, ok := internalConfigurationMap[externalConfigurationKey]; ok {
			internalConfigurationMap[externalConfigurationKey] = externalConfigurationValue
		}
	}

	return
}

func setOptionsDepTeardown(commandArray []string, currentOptions map[string]interface{}) error {

	setOptionsStubError := internal.SetOptionsStub(commandArray, currentOptions)
	if setOptionsStubError != nil {
		return errors.New("|CloudHorseman->commandbuilders->depteardown_command_builder->setOptionsDepTeardown->internal.SetOptionsStub:" + setOptionsStubError.Error() + "|")
	}

	return nil
}

func BuildDepTeardownCommand() (map[string]interface{}, error) {

	var currentOptions map[string]interface{} = make(map[string]interface{}, 0)

	currentOptions["kubernetes_deployment_name"] = ""

	setParametersBasedOnGlobalCacheAndDeploymentNameError := commandcachegenerators.SetParametersBasedOnGlobalCacheAndDeploymentName("kubernetes_deployment_name", currentOptions)
	if setParametersBasedOnGlobalCacheAndDeploymentNameError != nil {
		fmt.Println("|CloudHorseman->commandbuilders->depteardown_command_builder->BuildDepTeardownCommand->util.GetInternalHaymakerCFParameterForUIRepresentation:" + setParametersBasedOnGlobalCacheAndDeploymentNameError.Error() + "|")
	}

	return BuildGenericCommand("depteardown", currentOptions, menubuilders.GetDepTeardownOptionsMenu, menubuilders.GetDepTeardownInfoMenu, setOptionsDepTeardown, runDepTeardown, loadDepTeardown)

}
