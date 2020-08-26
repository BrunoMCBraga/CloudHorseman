package commandbuilders

import (
	"errors"
	"fmt"

	"github.com/BrunoMCBraga/CloudHorseman/commandcachegenerators"
	"github.com/BrunoMCBraga/CloudHorseman/internal"
	"github.com/BrunoMCBraga/CloudHorseman/menubuilders"
)

func runPuller(currentOptions map[string]interface{}) {}

func loadPullerTeardown(externalConfigurationMap map[string]interface{}, internalConfigurationMap map[string]interface{}) {

	for externalConfigurationKey, externalConfigurationValue := range externalConfigurationMap {
		if _, ok := internalConfigurationMap[externalConfigurationKey]; ok {
			internalConfigurationMap[externalConfigurationKey] = externalConfigurationValue
		}
	}

	return
}

func setOptionsPuller(commandArray []string, currentOptions map[string]interface{}) error {

	setOptionsStubError := internal.SetOptionsStub(commandArray, currentOptions)
	if setOptionsStubError != nil {
		return errors.New("|CloudHorseman->commandbuilders->puller_command_builder->setOptionsPuller->internal.SetOptionsStub:" + setOptionsStubError.Error() + "|")
	}

	return nil
}

func BuildPullerCommand() (map[string]interface{}, error) {

	var currentOptions map[string]interface{} = make(map[string]interface{}, 0)

	currentOptions["kubernetes_deployment_name"] = ""
	currentOptions["kubernetes_remote_pod_file"] = ""
	currentOptions["kubernetes_local_pod_file_folder"] = ""

	setParametersBasedOnGlobalCacheAndDeploymentNameError := commandcachegenerators.SetParametersBasedOnGlobalCacheAndDeploymentName("kubernetes_deployment_name", currentOptions)
	if setParametersBasedOnGlobalCacheAndDeploymentNameError != nil {
		fmt.Println("|CloudHorseman->commandbuilders->puller_command_builder->BuildPullerCommand->commandcachegenerators.SetParametersBasedOnGlobalCacheAndDeploymentName:" + setParametersBasedOnGlobalCacheAndDeploymentNameError.Error() + "|")
	}

	return BuildGenericCommand("puller", currentOptions, menubuilders.GetPullerOptionsMenu, menubuilders.GetPullerInfoMenu, setOptionsPuller, runPuller, loadPullerTeardown)

}
