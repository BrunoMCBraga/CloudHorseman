package commandbuilders

import (
	"errors"
	"fmt"

	"github.com/BrunoMCBraga/CloudHorseman/commandcachegenerators"
	"github.com/BrunoMCBraga/CloudHorseman/internal"
	"github.com/BrunoMCBraga/CloudHorseman/menubuilders"
)

func runPusher(currentOptions map[string]interface{}) {}

func loadPusherTeardown(externalConfigurationMap map[string]interface{}, internalConfigurationMap map[string]interface{}) {

	for externalConfigurationKey, externalConfigurationValue := range externalConfigurationMap {
		if _, ok := internalConfigurationMap[externalConfigurationKey]; ok {
			internalConfigurationMap[externalConfigurationKey] = externalConfigurationValue
		}
	}
	return
}

func setOptionsPusher(commandArray []string, currentOptions map[string]interface{}) error {

	setOptionsStubError := internal.SetOptionsStub(commandArray, currentOptions)
	if setOptionsStubError != nil {
		return errors.New("|CloudHorseman->commandbuilders->pusher_command_builder->setOptionsPusher->internal.SetOptionsStub:" + setOptionsStubError.Error() + "|")
	}

	return nil
}

func BuildPusherCommand() (map[string]interface{}, error) {

	var currentOptions map[string]interface{} = make(map[string]interface{}, 0)

	currentOptions["kubernetes_local_pod_file"] = ""
	currentOptions["kubernetes_remote_pod_file"] = ""
	currentOptions["kubernetes_replicas"] = 1 ///??? it should be defined. otherwise it will crash.
	currentOptions["kubernetes_deployment_name"] = ""
	currentOptions["kubernetes_workload_split"] = ""

	setParametersBasedOnGlobalCacheAndDeploymentNameError := commandcachegenerators.SetParametersBasedOnGlobalCacheAndDeploymentName("kubernetes_deployment_name", currentOptions)
	if setParametersBasedOnGlobalCacheAndDeploymentNameError != nil {
		fmt.Println("|CloudHorseman->commandbuilders->pusher_command_builder->BuildPusherCommand->util.GetInternalHaymakerCFParameterForUIRepresentation:" + setParametersBasedOnGlobalCacheAndDeploymentNameError.Error() + "|")
	}

	return BuildGenericCommand("pusher", currentOptions, menubuilders.GetPusherOptionsMenu, menubuilders.GetPusherInfoMenu, setOptionsPusher, runPusher, loadPusherTeardown)

}
