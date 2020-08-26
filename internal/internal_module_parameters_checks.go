package internal

import (
	"errors"
	"strconv"
)

var guiToInternalHaymakerCFParameters map[string]string = map[string]string{
	"TEMPLATE":              "cf_template",
	"STACKNAME":             "cf_stack_name",
	"S3BUCKETNAME":          "s3_bucket_name",
	"S3TEMPLATEKEY":         "s3_cf_template_file_key",
	"EKSCLUSTERNAME":        "cluster_name",
	"ECRREPONAME":           "repo_name",
	"REGISTRYID":            "registry_id",
	"DOCKERFILEFOLDER":      "docker_file_folder",
	"DELETEIMAGESAFTERPUSH": "delete_local_images_after_push",
	"EXPOSEDPORT":           "kubernetes_port",
	"DEPLOYMENTNAME":        "kubernetes_deployment_name",
	"IMAGENAME":             "kubernetes_image_name",
	"EXPOSEDPORTPROTOCOL":   "kubernetes_protocol",
	"REPLICAS":              "kubernetes_replicas",
	"REMOTEOUTPUT":          "kubernetes_remote_pod_file_output",
	"LOCALOUTPUT":           "kubernetes_local_pod_folder_output",
	"REMOTEINPUT":           "kubernetes_remote_pod_file_input",
	"LOCALINPUT":            "kubernetes_local_pod_file_input",
	"WORKLOADSPLIT":         "kubernetes_workload_split",
	"REMOTESTOP":            "kubernetes_remote_pod_file_stop",
	"LOCALFOLDER":           "kubernetes_local_pod_file_folder",
	"REMOTEFILE":            "kubernetes_remote_pod_file",
	"LOCALFILE":             "kubernetes_local_pod_file"}

var booleanParameters map[string]bool = map[string]bool{
	"DELETEIMAGESAFTERPUSH": true}

var integerParameters map[string]bool = map[string]bool{
	"EXPOSEDPORT": true,
	"REPLICAS":    true}

func GetInternalHaymakerCFParameterForUIRepresentation(uiRepresentation string) (string, error) {

	if internalSetValue, ok := guiToInternalHaymakerCFParameters[uiRepresentation]; ok {
		return internalSetValue, nil
	}

	return "", errors.New("|CloudHorseman->internal->internal_module_parameters_checks->GetInternalHaymakerCFParameterForUIRepresentation: invalid key.|")

}

func IsUIParameterBool(uiParameter string) bool {

	if _, ok := booleanParameters[uiParameter]; ok {
		return true
	}

	return false

}

func IsUIParameterInteger(uiParameter string) bool {

	if _, ok := integerParameters[uiParameter]; ok {
		return true
	}

	return false

}

//Set option
func SetOptionsStub(commandArray []string, currentOptions map[string]interface{}) error {

	getInternalHaymakerCFParameterForUIRepresentationResult, getInternalHaymakerCFParameterForUIRepresentationError := GetInternalHaymakerCFParameterForUIRepresentation(commandArray[0])
	if getInternalHaymakerCFParameterForUIRepresentationError != nil {
		return errors.New("|CloudHorseman->stubs->structure_setting_stub->SetOptionsStub->GetInternalHaymakerCFParameterForUIRepresentation:" + getInternalHaymakerCFParameterForUIRepresentationError.Error() + "|")
	}

	if IsUIParameterBool(commandArray[0]) {
		parseBoolResult, parseBoolError := strconv.ParseBool(commandArray[1])

		if parseBoolError != nil {
			return errors.New("|CloudHorseman->stubs->structure_setting_stub->SetOptionsStub->strconv.ParseBool:" + parseBoolError.Error() + "|")
		}

		currentOptions[getInternalHaymakerCFParameterForUIRepresentationResult] = parseBoolResult
	} else if IsUIParameterInteger(commandArray[0]) {
		parseIntResult, parseIntError := strconv.ParseInt(commandArray[1], 10, 32)

		if parseIntError != nil {
			return errors.New("|CloudHorseman->stubs->structure_setting_stub->SetOptionsStub->strconv.ParseInt:" + parseIntError.Error() + "|")
		}

		currentOptions[getInternalHaymakerCFParameterForUIRepresentationResult] = int(parseIntResult)
	} else {
		currentOptions[getInternalHaymakerCFParameterForUIRepresentationResult] = commandArray[1]
	}

	return nil
}

func GenerateInternalConfigurationMapFromExternalConfigurationMap(configurationMap map[string]interface{}) (map[string]interface{}, error) {

	internalConfigurationMap := make(map[string]interface{}, 0)

	for parameter, value := range configurationMap {
		if internalConfigurationKey, ok := guiToInternalHaymakerCFParameters[parameter]; ok {

			if IsUIParameterBool(parameter) {
				internalConfigurationMap[internalConfigurationKey] = value.(bool)
			} else if IsUIParameterInteger(parameter) {
				internalConfigurationMap[internalConfigurationKey] = int(value.(float64))
			} else {
				internalConfigurationMap[internalConfigurationKey] = value
			}

		}
	}

	return internalConfigurationMap, nil
}
