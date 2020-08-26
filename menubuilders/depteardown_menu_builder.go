package menubuilders

import (
	"bytes"
	"fmt"
	"text/tabwriter"
)

func GetDepTeardownOptionsMenu(commandLineMap map[string]interface{}) string {

	var kubernetesDeploymentNameTemp string
	if kubernetesDeploymentName, ok := commandLineMap["kubernetes_deployment_name"].(string); ok {
		kubernetesDeploymentNameTemp = kubernetesDeploymentName
	}

	var formattedTableByteArray *bytes.Buffer = new(bytes.Buffer)

	formattedTableNewWriter := tabwriter.NewWriter(formattedTableByteArray, 0, 0, 1, ' ', 0)
	fmt.Fprintln(formattedTableNewWriter, "Module options (avatar):\t")
	fmt.Fprintln(formattedTableNewWriter, "\t")
	fmt.Fprintln(formattedTableNewWriter, "Name\tCurrent Setting\tRequired\tDescription")
	fmt.Fprintln(formattedTableNewWriter, "----\t---------------\t--------\t-----------")
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("DEPLOYMENTNAME\t%s\tyes\tKubernetes deployment name.", kubernetesDeploymentNameTemp))
	formattedTableNewWriter.Flush()

	return formattedTableByteArray.String()

}

func GetDepTeardownInfoMenu(commandLineMap map[string]interface{}) string {

	cfBuilderOptionsString := `
Name: Kubernetes Deployment Teardown
Module: cfteardown
Version: 1

Provided by:
 Bruno Braga

Basic options:
 %s

Description:
 This module is used to teardown a Kubernetes deployment (including service).

References:
 https://kubernetes.io/docs/reference/
 https://kubernetes.io/docs/reference/using-api/client-libraries/ 
 https://kubernetes.io/docs/tasks/tools/install-kubectl/
`
	return fmt.Sprintf(cfBuilderOptionsString, GetDepTeardownOptionsMenu(commandLineMap))

}
