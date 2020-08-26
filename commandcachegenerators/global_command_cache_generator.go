package commandcachegenerators

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/BrunoMCBraga/CloudHorseman/globalstringsproviders"
	"github.com/go-errors/errors"
)

var globalCommandParametersCacheMap map[string]interface{} = make(map[string]interface{}, 0)

func UpdateGlobalParametersCache(parametersMap map[string]interface{}) {

	var tempParametersMap map[string]interface{}

	if cfStackName, ok := parametersMap["cf_stack_name"].(string); ok && cfStackName != "" {

		if globalParametersMapForStackName, ok := globalCommandParametersCacheMap[cfStackName]; ok {
			tempParametersMap = globalParametersMapForStackName.(map[string]interface{})
		} else {
			tempParametersMap = make(map[string]interface{}, 0)
			globalCommandParametersCacheMap[cfStackName] = tempParametersMap
		}

		if repoName, ok := parametersMap["repo_name"].(string); ok && repoName != "" {
			tempParametersMap["repo_name"] = repoName
		}

		if s3BucketName, ok := parametersMap["s3_bucket_name"].(string); ok && s3BucketName != "" {
			tempParametersMap["s3_bucket_name"] = s3BucketName
		}

		if clusterName, ok := parametersMap["cluster_name"].(string); ok && clusterName != "" {
			tempParametersMap["cluster_name"] = clusterName
		}

	}

}

func SetParametersBasedOnGlobalCacheAndStackName(parametersNamesToSet []string, parameters map[string]interface{}) error {

	if len(globalCommandParametersCacheMap) > 0 {

		var commandReader *bufio.Reader = bufio.NewReader(os.Stdin)

		fmt.Println(globalstringsproviders.GetCachedStackParametersWarningString())
		for stackName, _ := range globalCommandParametersCacheMap {
			fmt.Println(stackName)
		}
		fmt.Print(globalstringsproviders.GetStackNamePromptString())

		commandReader.Reset(os.Stdin)
		readStringResult, readStringError := commandReader.ReadString('\n')
		if readStringError != nil {
			return errors.New("|CloudHorseman->commandcachegenerators->global_command_cache_generator->SetParametersBasedOnGlobalCacheAndStackName->commandReader.ReadString:" + readStringError.Error() + "|")

		} else {
			command := strings.Trim(readStringResult, " \n")
			if getGetGlobalCommandParametersCacheEntry, ok := globalCommandParametersCacheMap[command]; ok {

				for _, parameterNameToSet := range parametersNamesToSet {
					if parameterNameToSet == "cf_stack_name" {
						parameters["cf_stack_name"] = command
					} else if parameterValue, ok := getGetGlobalCommandParametersCacheEntry.(map[string]interface{})[parameterNameToSet]; ok {
						parameters[parameterNameToSet] = parameterValue.(string) //here i assume that all parameters are strings
					}
				}
				return nil
			} else {

				return nil
			}
		}
	}
	return nil
}

//This is used to set kubernetes deployment name and nothing more.
func SetParametersBasedOnGlobalCacheAndDeploymentName(parametersNamesToSet string, parameters map[string]interface{}) error {

	if len(globalCommandParametersCacheMap) > 0 {

		var commandReader *bufio.Reader = bufio.NewReader(os.Stdin)

		for _, globalCommandsParameters := range globalCommandParametersCacheMap {
			if kubernetesDeploymentsNames, ok := globalCommandsParameters.(map[string]interface{})["kubernetes_deployment_name"].([]string); ok {
				fmt.Println(globalstringsproviders.GetCacheKubernetesdDeploymentParametersWarningString())
				for _, kubernetesDeploymentName := range kubernetesDeploymentsNames {
					fmt.Println(kubernetesDeploymentName)
				}
				fmt.Print(globalstringsproviders.GetKubernetesDeploymentNamePromptString())
			}
		}

		commandReader.Reset(os.Stdin)
		readStringResult, readStringError := commandReader.ReadString('\n')
		if readStringError != nil {
			return errors.New("|CloudHorseman->commandcachegenerators->global_command_cache_generator->SetParametersBasedOnGlobalCacheAndDeploymentName->commandReader.ReadString:" + readStringError.Error() + "|")

		} else {
			command := strings.Trim(readStringResult, " \n")
			parameters["kubernetes_deployment_name"] = command

		}

	}

	return nil
}
