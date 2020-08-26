package menubuilders

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
)

func GetPullerOptionsMenu(commandLineMap map[string]interface{}) string {

	var kubernetesDeploymentNameTemp string
	if kubernetesDeploymentName, ok := commandLineMap["kubernetes_deployment_name"].(string); ok {
		kubernetesDeploymentNameTemp = kubernetesDeploymentName
	}

	var kubernetesRemotePodFileTemp string
	if kubernetesRemotePodFile, ok := commandLineMap["kubernetes_remote_pod_file"].(string); ok {
		kubernetesRemotePodFileTemp = kubernetesRemotePodFile
	}

	var kubernetesLocalFileFolderTemp string
	if kubernetesLocalFileFolder, ok := commandLineMap["kubernetes_local_pod_file_folder"].(string); ok {
		kubernetesLocalFileFolderTemp = kubernetesLocalFileFolder
	}

	var formattedTableByteArray *bytes.Buffer = new(bytes.Buffer)
	formattedTableNewWriter := tabwriter.NewWriter(formattedTableByteArray, 0, 0, 1, ' ', 0)
	fmt.Fprintln(formattedTableNewWriter, "Module options (puller):\t")
	fmt.Fprintln(formattedTableNewWriter, "\t")
	fmt.Fprintln(formattedTableNewWriter, "Name\tCurrent Setting\tRequired\tDescription")
	fmt.Fprintln(formattedTableNewWriter, "----\t---------------\t--------\t-----------")
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("DEPLOYMENTNAME\t%s\tyes\tKubernetes deployment name.", kubernetesDeploymentNameTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("REMOTEFILE\t%s\tyes\tPath for remote file on container.", kubernetesRemotePodFileTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("LOCALFOLDER\t%s\tyes\tFolder where to store file (if multiple files, each name will be prepended with pod name).", kubernetesLocalFileFolderTemp))
	formattedTableNewWriter.Flush()

	return formattedTableByteArray.String()

}

func GetPullerInfoMenu(commandLineMap map[string]interface{}) string {

	pullerOptionsString := `
Name: Puller module
Module: puller
Version: 1

Provided by:
 Bruno Braga

Basic options:
 %s

Description:
 Pulls file from each container running inside a Pod.

References:
 https://kubernetes.io/docs/reference/
 https://kubernetes.io/docs/reference/using-api/client-libraries/ 
 https://kubernetes.io/docs/tasks/tools/install-kubectl/

`
	return strings.Trim(fmt.Sprintf(pullerOptionsString, GetPullerOptionsMenu(commandLineMap)), "\n")

}
