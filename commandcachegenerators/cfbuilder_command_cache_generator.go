package commandcachegenerators

func UpdateCFBuilderParametersCache(parametersMap map[string]interface{}) {

	var tempParametersMap map[string]interface{} = make(map[string]interface{}, 0)

	if cfStackName, ok := parametersMap["cf_stack_name"].(string); ok && cfStackName != "" {
		tempParametersMap["cf_stack_name"] = cfStackName
	}

	if s3BucketName, ok := parametersMap["s3_bucket_name"].(string); ok && s3BucketName != "" {
		tempParametersMap["s3_bucket_name"] = s3BucketName
	}

	if clusterName, ok := parametersMap["cluster_name"].(string); ok && clusterName != "" {
		tempParametersMap["cluster_name"] = clusterName
	}

	if repoName, ok := parametersMap["repo_name"].(string); ok && repoName != "" {
		tempParametersMap["repo_name"] = repoName
	}

	UpdateGlobalParametersCache(tempParametersMap)
}
