package menubuilders

import (
	"bytes"
	"fmt"
	"text/tabwriter"
)

func GetAvatarOptionsMenu(commandLineMap map[string]interface{}) string {

	var cfStackNameTemp string
	if cfStackName, ok := commandLineMap["cf_stack_name"].(string); ok && cfStackName != "" {
		cfStackNameTemp = cfStackName
	} else {
		cfStackNameTemp = ""
	}

	var repoNameTemp string
	if repoName, ok := commandLineMap["repo_name"].(string); ok && repoName != "" {
		repoNameTemp = repoName
	} else {
		repoNameTemp = ""
	}

	var registryIdTemp string
	if registryId, ok := commandLineMap["registry_id"].(string); ok && registryId != "" {
		registryIdTemp = registryId
	} else {
		registryIdTemp = ""
	}

	var dockerFileFolderTemp string
	if dockerFileFolder, ok := commandLineMap["docker_file_folder"].(string); ok && dockerFileFolder != "" {
		dockerFileFolderTemp = dockerFileFolder
	} else {
		dockerFileFolderTemp = ""
	}

	var deleteLocalImagesAfterPushTemp bool
	if deleteLocalImagesAfterPush, ok := commandLineMap["delete_local_images_after_push"].(bool); ok {
		deleteLocalImagesAfterPushTemp = deleteLocalImagesAfterPush
	}

	var kubernetesPortTemp int
	if kubernetesPort, ok := commandLineMap["kubernetes_port"].(int); ok {
		kubernetesPortTemp = kubernetesPort
	}

	var kubernetesDeploymentNameTemp string
	if kubernetesDeploymentName, ok := commandLineMap["kubernetes_deployment_name"].(string); ok {
		kubernetesDeploymentNameTemp = kubernetesDeploymentName
	}

	var kubernetesImageNameTemp string
	if kubernetesImageName, ok := commandLineMap["kubernetes_image_name"].(string); ok {
		kubernetesImageNameTemp = kubernetesImageName
	}

	var kubernetesProtocolTemp string
	if kubernetesProtocol, ok := commandLineMap["kubernetes_protocol"].(string); ok {
		kubernetesProtocolTemp = kubernetesProtocol
	}

	var kubernetesReplicasTemp int
	if kubernetesReplicas, ok := commandLineMap["kubernetes_replicas"].(int); ok {
		kubernetesReplicasTemp = kubernetesReplicas
	}

	var kubernetesRemotePodFileOutputTemp string
	if kubernetesRemotePodFileOutput, ok := commandLineMap["kubernetes_remote_pod_file_output"].(string); ok {
		kubernetesRemotePodFileOutputTemp = kubernetesRemotePodFileOutput
	}

	var kubernetesLocalPodFolderOutputTemp string
	if kubernetesLocalPodFolderOutput, ok := commandLineMap["kubernetes_local_pod_folder_output"].(string); ok {
		kubernetesLocalPodFolderOutputTemp = kubernetesLocalPodFolderOutput
	}

	var kubernetesRemotePodFileInputTemp string
	if kubernetesRemotePodFileInput, ok := commandLineMap["kubernetes_remote_pod_file_input"].(string); ok {
		kubernetesRemotePodFileInputTemp = kubernetesRemotePodFileInput
	}

	var kubernetesLocalPodFileInputTemp string
	if kubernetesLocalPodFileInput, ok := commandLineMap["kubernetes_local_pod_file_input"].(string); ok {
		kubernetesLocalPodFileInputTemp = kubernetesLocalPodFileInput
	}

	var kubernetesRemotePodFileStopTemp string
	if kubernetesRemotePodFileStop, ok := commandLineMap["kubernetes_remote_pod_file_stop"].(string); ok {
		kubernetesRemotePodFileStopTemp = kubernetesRemotePodFileStop
	}

	var kubernetesWorkloadSplitTemp string
	if kubernetesWorkloadSplit, ok := commandLineMap["kubernetes_workload_split"].(string); ok {
		kubernetesWorkloadSplitTemp = kubernetesWorkloadSplit
	}

	var formattedTableByteArray *bytes.Buffer = new(bytes.Buffer)

	formattedTableNewWriter := tabwriter.NewWriter(formattedTableByteArray, 0, 0, 1, ' ', 0)
	fmt.Fprintln(formattedTableNewWriter, "Module options (avatar):\t")
	fmt.Fprintln(formattedTableNewWriter, "\t")
	fmt.Fprintln(formattedTableNewWriter, "Name\tCurrent Setting\tRequired\tDescription")
	fmt.Fprintln(formattedTableNewWriter, "----\t---------------\t--------\t-----------")
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("STACKNAME\t%s\tyes\tCloudFormation stack name.", cfStackNameTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("ECRREPONAME\t%s\tyes\tThe name of ECR repo name.", repoNameTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("REGISTRYID\t%s\tno\tECR registry ID. If not provided, default will be used.", registryIdTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("DOCKERFILEFOLDER\t%s\tyes\tThe target address.", dockerFileFolderTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("DELETEIMAGESAFTERPUSH\t%t\tno\tDelete local Docker images after build and push to ECR.", deleteLocalImagesAfterPushTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("EXPOSEDPORT\t%d\tyes\tPort exposed by Kubernetes for container (0 means no port is exposed).", kubernetesPortTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("DEPLOYMENTNAME\t%s\tyes\tKubernetes deployment name.", kubernetesDeploymentNameTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("IMAGENAME\t%s\tno\tFull URL name for container on ECR.", kubernetesImageNameTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("EXPOSEDPORTPROTOCOL\t%s\tyes\tExposed port protocol (i.e. TCP, UDP).", kubernetesProtocolTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("REPLICAS\t%d\tno\tNumber of Docker replicas.", kubernetesReplicasTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("WORKLOADSPLIT\t%s\tno\tOne of the following options: ", kubernetesWorkloadSplitTemp))
	fmt.Fprintln(formattedTableNewWriter, "\t\t\tSPLITINPUT (e.g. hash cracking).")
	fmt.Fprintln(formattedTableNewWriter, "\t\t\tDUPLICATEINPUT (e.g. same file to all pods).")
	fmt.Fprintln(formattedTableNewWriter, "\t\t\tMULTIPLYINPUT (e.g. DDOS).")
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("REMOTEOUTPUT\t%s\tno\tPath for file on each container where output of processing is expected (e.g. passwords for input hashes).", kubernetesRemotePodFileOutputTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("LOCALOUTPUT\t%s\tno\tPath for local folder where REMOTEOUTPUT will be saved (e.g. passwords for input hashes).", kubernetesLocalPodFolderOutputTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("REMOTEINPUT\t%s\tno\tPath inside container where the workload data is expected (e.g. hashes to be cracked).", kubernetesRemotePodFileInputTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("LOCALINPUT\t%s\tno\tLocal path for file where the full workload data is (e.g. hashes to be cracked).", kubernetesLocalPodFileInputTemp))
	fmt.Fprintln(formattedTableNewWriter, fmt.Sprintf("REMOTESTOP\t%s\tno\tPath inside container where the stop file is expected.", kubernetesRemotePodFileStopTemp))
	formattedTableNewWriter.Flush()

	return formattedTableByteArray.String()

}

func GetAvatarInfoMenu(commandLineMap map[string]interface{}) string {

	avatarOptionsString := `
Name: Avatar module
Module: avatar
Version: 1

Provided by:
 Bruno Braga

Basic options:
 %s

Description:
 This module deployes a Docker container using a provided Dockerfile. 

References:
 https://docs.docker.com/engine/reference/builder/
 https://docs.docker.com/engine/api/sdk/
 https://kubernetes.io/docs/reference/
 https://kubernetes.io/docs/reference/using-api/client-libraries/ 
 https://kubernetes.io/docs/tasks/tools/install-kubectl/
 
`
	return fmt.Sprintf(avatarOptionsString, GetAvatarOptionsMenu(commandLineMap))

}
