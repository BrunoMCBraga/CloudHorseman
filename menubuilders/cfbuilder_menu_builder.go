package menubuilders

import (
	"bytes"
	"fmt"
	"text/tabwriter"
)

func GetCFBuilderOptionsMenu(commandLineMap map[string]interface{}) string {

	var cfTemplateTemp string
	if cfTemplate, ok := commandLineMap["cf_template"].(string); ok && cfTemplate != "" {
		cfTemplateTemp = cfTemplate
	} else {
		cfTemplateTemp = ""
	}

	var cfStackNameTemp string
	if cfStackName, ok := commandLineMap["cf_stack_name"].(string); ok && cfStackName != "" {
		cfStackNameTemp = cfStackName
	} else {
		cfStackNameTemp = ""
	}

	var s3BucketNameTemp string
	if s3BucketName, ok := commandLineMap["s3_bucket_name"].(string); ok && s3BucketName != "" {
		s3BucketNameTemp = s3BucketName
	} else {
		s3BucketNameTemp = ""
	}

	var s3CFTemplateFileKeyTemp string
	if s3CFTemplateFileKey, ok := commandLineMap["s3_cf_template_file_key"].(string); ok && s3CFTemplateFileKey != "" {
		s3CFTemplateFileKeyTemp = s3CFTemplateFileKey
	} else {
		s3CFTemplateFileKeyTemp = ""
	}

	var clusterNameTemp string
	if clusterName, ok := commandLineMap["cluster_name"].(string); ok && clusterName != "" {
		clusterNameTemp = clusterName
	} else {
		clusterNameTemp = ""
	}

	var repoNameTemp string
	if repoName, ok := commandLineMap["repo_name"].(string); ok && repoName != "" {
		repoNameTemp = repoName
	} else {
		repoNameTemp = ""
	}

	var formattedTableByteArray *bytes.Buffer = new(bytes.Buffer)

	formattedTableNewWriter := tabwriter.NewWriter(formattedTableByteArray, 0, 0, 1, ' ', 0)
	fmt.Fprintln(formattedTableNewWriter, "Module options (cfbuilder):\t")
	fmt.Fprintln(formattedTableNewWriter, "\t")
	fmt.Fprintln(formattedTableNewWriter, "Name\tCurrent Setting\tRequired\tDescription")
	fmt.Fprintln(formattedTableNewWriter, "----\t---------------\t--------\t-----------")
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("TEMPLATE\t%s\tyes\tCloudFormation template file path.", cfTemplateTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("ECRREPONAME\t%s\tyes\tThe name of ECR repo name (same as the one on CF template).", repoNameTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("STACKNAME\t%s\tyes\tCloudFormation stack name.", cfStackNameTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("S3BUCKETNAME\t%s\tno\tS3 bucket name for bucket used to host template file.", s3BucketNameTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("S3TEMPLATEKEY\t%s\tno\tKey used to index the template on the S3 bucket.", s3CFTemplateFileKeyTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("EKSCLUSTERNAME\t%s\tyes\tEKS Cluster Name (same as the one on CF template).", clusterNameTemp))
	formattedTableNewWriter.Flush()

	return formattedTableByteArray.String()

}

func GetCFBuilderInfoMenu(commandLineMap map[string]interface{}) string {

	cfBuilderOptionsString := `
Name: CloudFormation Builder
Module: cfbuilder
Version: 1

Provided by:
 Bruno Braga

Basic options:
 %s

Description:
 This module is used to deploy a CloudFormation template. 

References:
 https://aws.amazon.com/cloudformation/
 https://docs.aws.amazon.com/sdk-for-go/api/service/cloudformation/

`
	return fmt.Sprintf(cfBuilderOptionsString, GetCFBuilderOptionsMenu(commandLineMap))

}
