package menubuilders

import (
	"bytes"
	"fmt"
	"text/tabwriter"
)

func GetPusherOptionsMenu(commandLineMap map[string]interface{}) string {

	var kubernetesLocalPodFileInputTemp string
	if kubernetesLocalPodFileInput, ok := commandLineMap["kubernetes_local_pod_file"].(string); ok {
		kubernetesLocalPodFileInputTemp = kubernetesLocalPodFileInput
	}

	var kubernetesRemotePodFileInputTemp string
	if kubernetesRemotePodFileInput, ok := commandLineMap["kubernetes_remote_pod_file"].(string); ok {
		kubernetesRemotePodFileInputTemp = kubernetesRemotePodFileInput
	}

	var kubernetesReplicasTemp int
	if kubernetesReplicas, ok := commandLineMap["kubernetes_replicas"].(int); ok {
		kubernetesReplicasTemp = kubernetesReplicas
	}

	var kubernetesDeploymentNameTemp string
	if kubernetesDeploymentName, ok := commandLineMap["kubernetes_deployment_name"].(string); ok {
		kubernetesDeploymentNameTemp = kubernetesDeploymentName
	}

	var kubernetesWorkloadSplitTemp string
	if kubernetesWorkloadSplit, ok := commandLineMap["kubernetes_workload_split"].(string); ok {
		kubernetesWorkloadSplitTemp = kubernetesWorkloadSplit
	}

	var formattedTableByteArray *bytes.Buffer = new(bytes.Buffer)
	formattedTableNewWriter := tabwriter.NewWriter(formattedTableByteArray, 0, 0, 1, ' ', 0)
	fmt.Fprintln(formattedTableNewWriter, "Module options (pusher):\t")
	fmt.Fprintln(formattedTableNewWriter, "\t")
	fmt.Fprintln(formattedTableNewWriter, "Name\tCurrent Setting\tRequired\tDescription")
	fmt.Fprintln(formattedTableNewWriter, "----\t---------------\t--------\t-----------")
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("REPLICAS\t%d\tyes\tNumber of Docker replicas.", kubernetesReplicasTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("DEPLOYMENTNAME\t%s\tyes\tKubernetes deployment name.", kubernetesDeploymentNameTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("REMOTEFILE\t%s\tyes\tPath for remote file on container.", kubernetesRemotePodFileInputTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("WORKLOADSPLIT\t%s\tno\tOne of the following options: ", kubernetesWorkloadSplitTemp))
	fmt.Fprintln(formattedTableNewWriter, "\t\t\tSPLITINPUT (e.g. hash cracking): LOCALFILE is split line by line and spread across replicas.")
	fmt.Fprintln(formattedTableNewWriter, "\t\t\tDUPLICATEINPUT : LOCALFILE is sent as it is to all replicas. Alternatively, the file can be shipped with the base container.")
	fmt.Fprintln(formattedTableNewWriter, "\t\t\tMULTIPLYINPUT (e.g. DDOS): same as SPLITINPUT but for each line REPLICAS are allocated to process the line (e.g. 5 lines * 3 replicas = 15 containers).")
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("LOCALFILE\t%s\tyes\tPath for remote file to split and upload.", kubernetesLocalPodFileInputTemp))
	formattedTableNewWriter.Flush()

	return formattedTableByteArray.String()

}

func GetPusherInfoMenu(commandLineMap map[string]interface{}) string {

	pusherOptionsString := `
Name: Pusher module
Module: pusher
Version: 1

Provided by:
 Bruno Braga

Basic options:
 %s

Description:
 Splits file into multiple subfiles and pushes them to remote Pods.

References:
 https://kubernetes.io/docs/reference/
 https://kubernetes.io/docs/reference/using-api/client-libraries/ 
 https://kubernetes.io/docs/tasks/tools/install-kubectl/

`
	return fmt.Sprintf(pusherOptionsString, GetPusherOptionsMenu(commandLineMap))

}
