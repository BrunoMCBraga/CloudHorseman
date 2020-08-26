package menubuilders

import (
	"bytes"
	"fmt"
	"text/tabwriter"
)

func GetCFTeardownOptionsMenu(commandLineMap map[string]interface{}) string {

	var repoNameTemp string
	if repoName, ok := commandLineMap["repo_name"].(string); ok && repoName != "" {
		repoNameTemp = repoName
	} else {
		repoNameTemp = ""
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

	var formattedTableByteArray *bytes.Buffer = new(bytes.Buffer)

	formattedTableNewWriter := tabwriter.NewWriter(formattedTableByteArray, 0, 0, 1, ' ', 0)
	fmt.Fprintln(formattedTableNewWriter, "Module options (cfteardown):\t")
	fmt.Fprintln(formattedTableNewWriter, "\t")
	fmt.Fprintln(formattedTableNewWriter, "Name\tCurrent Setting\tRequired\tDescription")
	fmt.Fprintln(formattedTableNewWriter, "----\t---------------\t--------\t-----------")
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("ECRREPONAME\t%s\tyes\tThe name of ECR repo name (same as the one on CF template).", repoNameTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("STACKNAME\t%s\tyes\tCloudFormation stack name.", cfStackNameTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("S3BUCKETNAME\t%s\tno\tS3 bucket name for bucket used to host template file.", s3BucketNameTemp))
	formattedTableNewWriter.Flush()

	return formattedTableByteArray.String()

}

func GetCFTeardownInfoMenu(commandLineMap map[string]interface{}) string {

	cfBuilderOptionsString := `
Name: CloudFormation Teardown
Module: cfteardown
Version: 1

Provided by:
 Bruno Braga

Basic options:
 %s

Description:
 This module is used to teardown a CloudFormation deployment. This will delete all Kubernetes services, ECR repos and images (if part of CF template) and S3 bucket (optionally).

References:
 https://aws.amazon.com/cloudformation/
 https://docs.aws.amazon.com/sdk-for-go/api/service/cloudformation/
 
`
	return fmt.Sprintf(cfBuilderOptionsString, GetCFTeardownOptionsMenu(commandLineMap))

}
