package commandcachegenerators

func UpdateAvatarParametersCache(parametersMap map[string]interface{}) {

	var tempParametersMap map[string]interface{} = make(map[string]interface{}, 0)

	if repoName, ok := parametersMap["repo_name"].(string); ok && repoName != "" {
		tempParametersMap["repo_name"] = repoName
	}

	if registryId, ok := parametersMap["registry_id"].(string); ok && registryId != "" {
		tempParametersMap["registry_id"] = registryId
	}

	if kubernetesDeploymentName, ok := parametersMap["kubernetes_deployment_name"].(string); ok && kubernetesDeploymentName != "" {
		tempParametersMap["kubernetes_deployment_name"] = kubernetesDeploymentName
	}

	if cfStackName, ok := parametersMap["cf_stack_name"].(string); ok && cfStackName != "" {
		tempParametersMap["cf_stack_name"] = cfStackName
	}

	UpdateGlobalParametersCache(tempParametersMap)
}
