package commandprocessors

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/BrunoMCBraga/HayMakerCF/commandlineprocessors"
	"github.com/BrunoMCBraga/HayMakerCF/haymakercfengines"
	"github.com/BrunoMCBraga/CloudHorseman/commandbuilders"
	"github.com/BrunoMCBraga/CloudHorseman/globalstringsproviders"
	"github.com/BrunoMCBraga/CloudHorseman/util"
)

const kubeCtlWaitTimeBeforePulling int = 60
const kubeCtlWaitTimeBeforeCheckingPodsState int = 30
const kubeCtlMaxAttemptsAtCheckingPodsState int = 5

func processAvatarCommand(parametersMap map[string]interface{}) error {

	haymakerCommandMap := make(map[string]interface{}, 0)
	var commandReader *bufio.Reader = bufio.NewReader(os.Stdin)
	/*
		pi: HayMakerCF Docker Image Build And Push To ECR (e.g.
		go run ./main.go
		-cm pi
		-rn haymaker-docker-repo/haymaker-docker
		-df /Users/brubraga/go/src/github.com/haymakercf/Docker
		-di
		-ri 965440066241*/

	haymakerCommandMap["option"] = "pi"
	haymakerCommandMap["repo_name"] = parametersMap["repo_name"].(string)
	haymakerCommandMap["docker_file_folder"] = parametersMap["docker_file_folder"].(string)
	haymakerCommandMap["delete_local_images_after_push"] = parametersMap["delete_local_images_after_push"].(bool)
	haymakerCommandMap["registry_id"] = parametersMap["registry_id"].(string)

	processCommandLineError := commandlineprocessors.ProcessCommandLine(haymakerCommandMap)

	if processCommandLineError != nil {
		return errors.New("|CloudHorseman->commandprocessors->avatar_command_processor->ProcessAvatarCommand->ProcessCommandLine:" + processCommandLineError.Error() + "|")
	}

	var eCRGetRepositoryURIResult string
	var eCRGetRepositoryURIError error
	if kubernetesImageName, kubernetesImageNameOk := parametersMap["kubernetes_image_name"].(string); kubernetesImageNameOk && kubernetesImageName == "" {

		if parametersMap["registry_id"].(string) == "" {
			eCRGetRepositoryURIResult, eCRGetRepositoryURIError = haymakercfengines.ECRGetRepositoryURIWithDefaultRegistry(parametersMap["repo_name"].(string))
		} else {
			eCRGetRepositoryURIResult, eCRGetRepositoryURIError = haymakercfengines.ECRGetRepositoryURI(parametersMap["repo_name"].(string), parametersMap["registry_id"].(string))
		}

		if eCRGetRepositoryURIError != nil {
			fmt.Println("|CloudHorseman->commandprocessors->avatar_command_processor->ProcessAvatarCommand->ECRGetRepositoryURI:" + eCRGetRepositoryURIError.Error() + "|")
		} else {
			haymakerCommandMap["kubernetes_image_name"] = eCRGetRepositoryURIResult
		}

	} else {
		haymakerCommandMap["kubernetes_image_name"] = parametersMap["kubernetes_image_name"].(string)
	}

	var splitFileBasedOnSplitTypeResult map[string]bool
	var splitFileBasedOnSplitTypeError error
	if kubernetesLocalPodFileInput, kubernetesLocalPodFileInputOk := parametersMap["kubernetes_local_pod_file_input"].(string); kubernetesLocalPodFileInputOk && kubernetesLocalPodFileInput != "" {

		splitFileBasedOnSplitTypeResult, splitFileBasedOnSplitTypeError = util.SplitFileBasedOnSplitType(kubernetesLocalPodFileInput, parametersMap["kubernetes_replicas"].(int), parametersMap["kubernetes_workload_split"].(string))
		if splitFileBasedOnSplitTypeError != nil {
			return errors.New("|CloudHorseman->commandprocessors->avatar_command_processor->ProcessAvatarCommand->stubs.SplitFileBasedOnSplitType:" + splitFileBasedOnSplitTypeError.Error() + "|")
		}
		parametersMap["kubernetes_replicas"] = len(splitFileBasedOnSplitTypeResult)
	}

	/*
		sc: HayMakerCF Deploy Container And Create Service
		go run ./main.go
		-cm sc
		-kp 80
		-dn haymaker
		-in 965440066241.dkr.ecr.us-east-1.amazonaws.com/haymaker-docker-repo/haymaker-docker:latest
		-pr TCP
		-kr 2
	*/

	haymakerCommandMap["option"] = "sc"
	haymakerCommandMap["kubernetes_port"] = parametersMap["kubernetes_port"].(int)
	haymakerCommandMap["kubernetes_deployment_name"] = parametersMap["kubernetes_deployment_name"].(string)
	haymakerCommandMap["kubernetes_protocol"] = parametersMap["kubernetes_protocol"].(string)
	haymakerCommandMap["kubernetes_replicas"] = parametersMap["kubernetes_replicas"].(int)

	processCommandLineError = commandlineprocessors.ProcessCommandLine(haymakerCommandMap)

	if processCommandLineError != nil {
		return errors.New("|CloudHorseman->commandprocessors->avatar_command_processor->ProcessAvatarCommand->ProcessCommandLine:" + processCommandLineError.Error() + "|")
	}

	currentKubeCtlPodStateCheckAttempt := 0
	for true {
		kubernetesGetListOfPodsInReadyStateResult, kubernetesGetListOfPodsInReadyStateError := haymakercfengines.KubernetesGetListOfPodsInReadyState(parametersMap["kubernetes_deployment_name"].(string))
		if kubernetesGetListOfPodsInReadyStateError != nil {
			return errors.New("|CloudHorseman->commandprocessors->avatar_command_processor->ProcessAvatarCommand-haymakercfengines.KubernetesGetListOfPodsInReadyState:" + kubernetesGetListOfPodsInReadyStateError.Error() + "|")
		}

		if len(kubernetesGetListOfPodsInReadyStateResult) == parametersMap["kubernetes_replicas"].(int) {
			break
		} else if currentKubeCtlPodStateCheckAttempt <= kubeCtlMaxAttemptsAtCheckingPodsState {
			time.Sleep(time.Duration(kubeCtlWaitTimeBeforeCheckingPodsState) * time.Second)
			currentKubeCtlPodStateCheckAttempt++
		} else {
			fmt.Print(globalstringsproviders.GetPodStateMaxAttemptString())
			commandReader.Reset(os.Stdin)
			readStringResult, readStringError := commandReader.ReadString('\n')
			if readStringError != nil {
				return errors.New("|CloudHorseman->commandprocessors->puller_command_processor->commandReader.ReadString:" + readStringError.Error() + "|")
			}

			command := strings.Trim(readStringResult, " \n")
			if command == "y" {
				currentKubeCtlPodStateCheckAttempt = 0
				continue
			} else {
				break
			}
		}
	}

	kubernetesRemotePodFileInput, kubernetesRemotePodFileInputOk := parametersMap["kubernetes_remote_pod_file_input"].(string)
	kubernetesLocalPodFileInput, kubernetesLocalPodFileInputOk := parametersMap["kubernetes_local_pod_file_input"].(string)

	if kubernetesRemotePodFileInputOk && kubernetesRemotePodFileInput != "" && kubernetesLocalPodFileInputOk && kubernetesLocalPodFileInput != "" {

		fmt.Print(globalstringsproviders.GetWorkloadUploadString())
		commandReader.Reset(os.Stdin)
		readStringResult, readStringError := commandReader.ReadString('\n')
		if readStringError != nil {
			return errors.New("|CloudHorseman->commandprocessors->puller_command_processor->commandReader.ReadString:" + readStringError.Error() + "|")
		}

		command := strings.Trim(readStringResult, " \n")
		if command == "y" {

			if kubernetesRemotePodFileInputOk && kubernetesRemotePodFileInput != "" && kubernetesLocalPodFileInputOk && kubernetesLocalPodFileInput != "" {
				/// We upload the file

				pusherCommandParametersMap := make(map[string]interface{}, 0)
				pusherCommandParametersMap["kubernetes_replicas"] = parametersMap["kubernetes_replicas"].(int)
				pusherCommandParametersMap["kubernetes_deployment_name"] = parametersMap["kubernetes_deployment_name"].(string)
				pusherCommandParametersMap["kubernetes_remote_pod_file"] = parametersMap["kubernetes_remote_pod_file_input"].(string)
				pusherCommandParametersMap["kubernetes_workload_split"] = splitFileBasedOnSplitTypeResult

				processPusherCommandFromParametersMapError := ProcessPusherCommandFromParametersMap(pusherCommandParametersMap)
				if processPusherCommandFromParametersMapError != nil {
					return errors.New("|CloudHorseman->commandprocessors->avatar_command_processor->ProcessAvatarCommand->commandprocessors.ProcessPusherCommandFromParametersMap:" + processPusherCommandFromParametersMapError.Error() + "|")
				}

				time.Sleep(time.Duration(kubeCtlWaitTimeBeforePulling) * time.Second)
			}
		}
	}

	///We get the result....
	kubernetesRemotePodFileOutput, kubernetesRemotePodFileOutputOk := parametersMap["kubernetes_remote_pod_file_output"].(string)
	kubernetesLocalPodFolderOutput, kubernetesLocalPodFolderOutputOk := parametersMap["kubernetes_local_pod_folder_output"].(string)

	if kubernetesRemotePodFileOutputOk && kubernetesRemotePodFileOutput != "" && kubernetesLocalPodFolderOutputOk && kubernetesLocalPodFolderOutput != "" {

		pullerCommandParametersMap := make(map[string]interface{}, 0)
		pullerCommandParametersMap["kubernetes_deployment_name"] = parametersMap["kubernetes_deployment_name"].(string)
		pullerCommandParametersMap["kubernetes_remote_pod_file"] = parametersMap["kubernetes_remote_pod_file_output"].(string)
		pullerCommandParametersMap["kubernetes_local_pod_file_folder"] = parametersMap["kubernetes_local_pod_folder_output"].(string)
		processPullerCommandFromParametersMapError := ProcessPullerCommandFromParametersMap(pullerCommandParametersMap)
		if processPullerCommandFromParametersMapError != nil {
			fmt.Println("|CloudHorseman->commandprocessors->avatar_command_processor->ProcessAvatarCommand->commandprocessors.ProcessPullerCommandFromParametersMap:" + processPullerCommandFromParametersMapError.Error() + "|")
		}

	}

	kubernetesRemotePodFileStop, kubernetesRemotePodFileStopOk := parametersMap["kubernetes_remote_pod_file_stop"].(string)
	if kubernetesRemotePodFileStopOk && kubernetesRemotePodFileStop != "" {
		fmt.Print(globalstringsproviders.GetDeploymentStopString())
		commandReader.Reset(os.Stdin)
		readStringResult, readStringError := commandReader.ReadString('\n')
		if readStringError != nil {
			return errors.New("|CloudHorseman->commandprocessors->avatar_command_processor->ProcessAvatarCommand->ommandReader.ReadString:" + readStringError.Error() + "|")
		}

		command := strings.Trim(readStringResult, " \n")
		if command == "y" {

			createTemporaryFileResult, createTemporaryFileError := util.CreateTemporaryFile()
			if createTemporaryFileError != nil {
				return errors.New("|CloudHorseman->commandprocessors->pusher_command_processor->ProcessAvatarCommand->util.CreateTemporaryFile:" + createTemporaryFileError.Error() + "|")
			} else if createTemporaryFileResult == "" {
				return errors.New("|CloudHorseman->commandprocessors->pusher_command_processor->ProcessAvatarCommand->util.CreateTemporaryFile: returned empty path for stop file.|")
			}

			duplicateInputFileSplitterFunction := util.FileSplitterMappings["DUPLICATEINPUT"]
			duplicateInputFileSplitterFunctionResult, duplicateInputFileSplitterFunctionError := duplicateInputFileSplitterFunction(createTemporaryFileResult, parametersMap["kubernetes_replicas"].(int))
			if duplicateInputFileSplitterFunctionError != nil {
				return errors.New("|CloudHorseman->commandprocessors->pusher_command_processor->ProcessAvatarCommand->duplicateInputFileSplitterFunction:" + duplicateInputFileSplitterFunctionError.Error() + "|")
			} else if duplicateInputFileSplitterFunctionResult == nil {
				return errors.New("|CloudHorseman->commandprocessors->pusher_command_processor->ProcessAvatarCommand->duplicateInputFileSplitterFunction: returned null commands structure.|")
			}

			parametersMap["kubernetes_replicas"] = len(duplicateInputFileSplitterFunctionResult)

			pusherCommandParametersMap := make(map[string]interface{}, 0)
			pusherCommandParametersMap["kubernetes_replicas"] = parametersMap["kubernetes_replicas"].(int)
			pusherCommandParametersMap["kubernetes_deployment_name"] = parametersMap["kubernetes_deployment_name"].(string)
			pusherCommandParametersMap["kubernetes_remote_pod_file"] = kubernetesRemotePodFileStop
			pusherCommandParametersMap["kubernetes_workload_split"] = duplicateInputFileSplitterFunctionResult

			processPusherCommandFromParametersMapError := ProcessPusherCommandFromParametersMap(pusherCommandParametersMap)
			if processPusherCommandFromParametersMapError != nil {
				return errors.New("|CloudHorseman->commandprocessors->avatar_command_processor->ProcessAvatarCommand->commandprocessors.ProcessPusherCommandFromParametersMap:" + processPusherCommandFromParametersMapError.Error() + "|")
			}
			return nil
		} else {
			return nil
		}
	}

	return nil
}

func ProcessAvatarCommandInteractively() error {

	avatarCommandBuilderFunction := commandbuilders.BuildAvatarCommand
	avatarCommandBuilderFunctionResult, avatarCommandBuilderFunctionError := avatarCommandBuilderFunction()
	if avatarCommandBuilderFunctionError != nil {
		return errors.New("|CloudHorseman->commandprocessors->cfbuilder_command_processor->ProcessAvatarCommandInteractively->avatarCommandBuilderFunction:" + avatarCommandBuilderFunctionError.Error() + "|")
	} else if avatarCommandBuilderFunctionResult == nil {
		return nil
	}

	processAvatarCommandError := processAvatarCommand(avatarCommandBuilderFunctionResult)
	if processAvatarCommandError != nil {
		return errors.New("|CloudHorseman->commandprocessors->cfbuilder_command_processor->ProcessAvatarCommandInteractively->processAvatarCommand:" + processAvatarCommandError.Error() + "|")
	}

	return nil
}

func ProcessAvatarCommandFromParametersMap(parametersMap map[string]interface{}) error {
	processAvatarCommandError := processAvatarCommand(parametersMap)
	if processAvatarCommandError != nil {
		return errors.New("|CloudHorseman->commandprocessors->cfbuilder_command_processor->ProcessAvatarCommandFromParametersMap->processAvatarCommand:" + processAvatarCommandError.Error() + "|")
	}

	return nil
}
