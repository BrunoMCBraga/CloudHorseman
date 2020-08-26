package commandbuilders

import (
	"errors"
	"fmt"

	"github.com/BrunoMCBraga/CloudHorseman/commandcachegenerators"
	"github.com/BrunoMCBraga/CloudHorseman/internal"
	"github.com/BrunoMCBraga/CloudHorseman/menubuilders"
)

var s3BucketDefaultPrefix string = "cloudhorseman"

func runCFBuilder(currentOptions map[string]interface{}) {

	if stackName, ok := currentOptions["cf_stack_name"]; ok && stackName != "" {
		if s3BucketName, ok := currentOptions["s3_bucket_name"]; ok && s3BucketName == "" {
			currentOptions["s3_bucket_name"] = fmt.Sprintf("%s-%s", s3BucketDefaultPrefix, stackName)
		}
		if s3CFTemplateFileKey, ok := currentOptions["s3_cf_template_file_key"]; ok && s3CFTemplateFileKey == "" {
			currentOptions["s3_cf_template_file_key"] = stackName
		}
	}
	commandcachegenerators.UpdateCFBuilderParametersCache(currentOptions)
}

func loadCFBuilder(externalConfigurationMap map[string]interface{}, internalConfigurationMap map[string]interface{}) {

	for externalConfigurationKey, externalConfigurationValue := range externalConfigurationMap {
		if _, ok := internalConfigurationMap[externalConfigurationKey]; ok {
			internalConfigurationMap[externalConfigurationKey] = externalConfigurationValue
		}
	}

	return
}

func setOptionsCFBuilder(commandArray []string, currentOptions map[string]interface{}) error {

	setOptionsStubError := internal.SetOptionsStub(commandArray, currentOptions)
	if setOptionsStubError != nil {
		return errors.New("|CloudHorseman->commandbuilders->cfbuilder_command_builder->setOptionsCFBuilder->internal.SetOptionsStub:" + setOptionsStubError.Error() + "|")
	}

	return nil
}

func BuildCFBuilderCommand() (map[string]interface{}, error) {

	var currentOptions map[string]interface{} = make(map[string]interface{}, 0)

	currentOptions["cf_template"] = ""
	currentOptions["cf_stack_name"] = ""
	currentOptions["s3_bucket_name"] = ""
	currentOptions["s3_cf_template_file_key"] = ""
	currentOptions["cluster_name"] = ""
	currentOptions["repo_name"] = ""

	return BuildGenericCommand("cfbuilder", currentOptions, menubuilders.GetCFBuilderOptionsMenu, menubuilders.GetCFBuilderInfoMenu, setOptionsCFBuilder, runCFBuilder, loadCFBuilder)

}
