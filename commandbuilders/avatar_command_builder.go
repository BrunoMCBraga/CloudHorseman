package commandbuilders

import (
	"errors"
	"fmt"

	"github.com/BrunoMCBraga/CloudHorseman/commandcachegenerators"
	"github.com/BrunoMCBraga/CloudHorseman/internal"
	"github.com/BrunoMCBraga/CloudHorseman/menubuilders"
)

func runAvatar(currentOptions map[string]interface{}) {
	commandcachegenerators.UpdateAvatarParametersCache(currentOptions)

}

func loadAvatar(externalConfigurationMap map[string]interface{}, internalConfigurationMap map[string]interface{}) {

	for externalConfigurationKey, externalConfigurationValue := range externalConfigurationMap {
		if _, ok := internalConfigurationMap[externalConfigurationKey]; ok {
			internalConfigurationMap[externalConfigurationKey] = externalConfigurationValue
		}
	}
	return
}

func setOptionsAvatar(commandArray []string, currentOptions map[string]interface{}) error {

	setOptionsStubError := internal.SetOptionsStub(commandArray, currentOptions)
	if setOptionsStubError != nil {
		return errors.New("|CloudHorseman->commandbuilders->avatar_command_builder->setOptionsAvatar->internal.SetOptionsStub:" + setOptionsStubError.Error() + "|.")
	}

	return nil
}

func BuildAvatarCommand() (map[string]interface{}, error) {

	var currentOptions map[string]interface{} = make(map[string]interface{}, 0)

	currentOptions["repo_name"] = ""
	currentOptions["cf_stack_name"] = ""
	currentOptions["registry_id"] = ""
	currentOptions["docker_file_folder"] = ""
	currentOptions["delete_local_images_after_push"] = true
	currentOptions["kubernetes_port"] = 0
	currentOptions["kubernetes_deployment_name"] = ""
	currentOptions["kubernetes_image_name"] = ""
	currentOptions["kubernetes_protocol"] = ""
	currentOptions["kubernetes_replicas"] = 1
	currentOptions["kubernetes_remote_pod_file_output"] = ""
	currentOptions["kubernetes_local_pod_folder_output"] = ""
	currentOptions["kubernetes_remote_pod_file_input"] = ""
	currentOptions["kubernetes_local_pod_file_input"] = ""
	currentOptions["kubernetes_remote_pod_file_stop"] = ""
	currentOptions["kubernetes_local_pod_file_stop"] = ""
	currentOptions["kubernetes_workload_split"] = "SPLITINPUT"

	listOfKeysToSet := make([]string, 0)

	for currentOptionsKey, _ := range currentOptions {
		listOfKeysToSet = append(listOfKeysToSet, currentOptionsKey)
	}

	setParametersBasedOnGlobalCacheAndStackNameError := commandcachegenerators.SetParametersBasedOnGlobalCacheAndStackName(listOfKeysToSet, currentOptions)
	if setParametersBasedOnGlobalCacheAndStackNameError != nil {
		fmt.Println("|CloudHorseman->commandbuilders->avatar_command_builder->BuildAvatarCommand->commandcachegenerators.setParametersBasedOnGlobalCacheAndStackNameError:" + setParametersBasedOnGlobalCacheAndStackNameError.Error() + "|")
	}

	return BuildGenericCommand("avatar", currentOptions, menubuilders.GetAvatarOptionsMenu, menubuilders.GetAvatarInfoMenu, setOptionsAvatar, runAvatar, loadAvatar)

}
