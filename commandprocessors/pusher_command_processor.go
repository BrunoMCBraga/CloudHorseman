package commandprocessors

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/BrunoMCBraga/HayMakerCF/haymakercfengines"
	"github.com/BrunoMCBraga/CloudHorseman/commandbuilders"
	"github.com/BrunoMCBraga/CloudHorseman/globalstringsproviders"
	"github.com/BrunoMCBraga/CloudHorseman/util"
)

const maxKubeCtlPushAttempts int = 5
const kubeCtlPushAttemptsSleepTime int = 60

func processPusherCommand(parametersMap map[string]interface{}) error {

	var commandReader *bufio.Reader = bufio.NewReader(os.Stdin)
	var kubernetesWorkloadSplit map[string]bool = parametersMap["kubernetes_workload_split"].(map[string]bool)

	kubernetesGetPodsForDeploymentResult, kubernetesGetPodsForDeploymentError := haymakercfengines.KubernetesGetPodsForDeployment(parametersMap["kubernetes_deployment_name"].(string))
	if kubernetesGetPodsForDeploymentError != nil {
		return errors.New("|CloudHorseman->commandprocessors->pusher_command_processor->processPusherCommand->KubernetesGetPodsForDeployment:" + kubernetesGetPodsForDeploymentError.Error() + "|")
	} else if kubernetesGetPodsForDeploymentResult == nil {
		return errors.New("|CloudHorseman->commandprocessors->pusher_command_processor->processPusherCommand->KubernetesGetPodsForDeployment: returned null map of pods.|")
	}

	kubernetesGetPodsForDeploymentResultCopy := util.CloneStringBoolMap(kubernetesGetPodsForDeploymentResult)
	kubernetesWorkloadSplitCopy := util.CloneStringBoolMap(kubernetesWorkloadSplit)

	currentKubeCtlPushAttempt := 0

	for true {
		kubernetesCopyFilesToRemotePodsResult, kubernetesCopyFilesToRemotePodsError := haymakercfengines.KubernetesCopyFilesToRemotePods(
			kubernetesGetPodsForDeploymentResultCopy,
			parametersMap["kubernetes_remote_pod_file"].(string),
			kubernetesWorkloadSplitCopy)

		if kubernetesCopyFilesToRemotePodsError != nil {
			return errors.New("|CloudHorseman->commandprocessors->pusher_command_processor->processPusherCommand->haymakercfengines.KubernetesCopyFilesToRemotePods:" + kubernetesCopyFilesToRemotePodsError.Error() + "|")
		}

		for localFile, podName := range kubernetesCopyFilesToRemotePodsResult {
			delete(kubernetesGetPodsForDeploymentResultCopy, podName)
			delete(kubernetesWorkloadSplitCopy, localFile)

		}

		if len(kubernetesGetPodsForDeploymentResultCopy) == 0 {
			break
		} else {
			time.Sleep(time.Duration(kubeCtlPushAttemptsSleepTime) * time.Second)
		}

		if currentKubeCtlPushAttempt > maxKubeCtlPushAttempts {
			return errors.New("|CloudHorseman->commandprocessors->pusher_command_processor->processPusherCommand: maximum number of attempts reached. Failed to push all the files to deployment pods.|")
		} else {
			currentKubeCtlPushAttempt++
		}

		if len(kubernetesGetPodsForDeploymentResultCopy) == 0 {
			fmt.Print(globalstringsproviders.GetPodFileRetrievalSuccessfulPullString())
			commandReader.Reset(os.Stdin)
			readStringResult, readStringError := commandReader.ReadString('\n')
			if readStringError != nil {
				return errors.New("|CloudHorseman->commandprocessors->puller_command_processor->commandReader.ReadString:" + readStringError.Error() + "|")
			}

			command := strings.Trim(readStringResult, " \n")
			if command == "y" {
				currentKubeCtlPushAttempt = 0
				kubernetesGetPodsForDeploymentResultCopy = util.CloneStringBoolMap(kubernetesGetPodsForDeploymentResult)
				kubernetesWorkloadSplitCopy = util.CloneStringBoolMap(kubernetesWorkloadSplit)
				continue
			} else {
				break
			}

		} else {
			time.Sleep(time.Duration(kubeCtlPushAttemptsSleepTime) * time.Second)
		}
		if currentKubeCtlPushAttempt > currentKubeCtlPushAttempt {
			fmt.Print(globalstringsproviders.GetPodFileRetrievalMaxAttemptString())
			commandReader.Reset(os.Stdin)
			readStringResult, readStringError := commandReader.ReadString('\n')
			if readStringError != nil {
				return errors.New("|CloudHorseman->commandprocessors->pusher_command_processor->commandReader.ReadString:" + readStringError.Error() + "|")
			}

			command := strings.Trim(readStringResult, " \n")
			if command == "y" {
				currentKubeCtlPushAttempt = 0
				continue
			} else {
				break
			}
		} else {
			currentKubeCtlPushAttempt++
		}

	}

	return nil

}

func ProcessPusherCommandInteractively() error {

	pusherCommandBuilderFunction := commandbuilders.BuildPusherCommand
	pusherCommandBuilderFunctionResult, pusherCommandBuilderFunctionError := pusherCommandBuilderFunction()
	if pusherCommandBuilderFunctionError != nil {
		return errors.New("|CloudHorseman->commandprocessors->pusher_command_processor->ProcessPusherCommandInteractively->pusherCommandBuilderFunction:" + pusherCommandBuilderFunctionError.Error() + "|")
	} else if pusherCommandBuilderFunctionResult == nil {
		return nil
	}

	var splitFileBasedOnSplitTypeResult map[string]bool
	var splitFileBasedOnSplitTypeError error
	if kubernetesLocalPodFileInput, kubernetesLocalPodFileInputOk := pusherCommandBuilderFunctionResult["kubernetes_local_pod_file"].(string); kubernetesLocalPodFileInputOk && kubernetesLocalPodFileInput != "" {

		splitFileBasedOnSplitTypeResult, splitFileBasedOnSplitTypeError = util.SplitFileBasedOnSplitType(kubernetesLocalPodFileInput, pusherCommandBuilderFunctionResult["kubernetes_replicas"].(int), pusherCommandBuilderFunctionResult["kubernetes_workload_split"].(string))
		if splitFileBasedOnSplitTypeError != nil {
			return errors.New("|CloudHorseman->commandprocessors->pusher_command_processor->ProcessPusherCommandInteractively->stubs.SplitFileBasedOnSplitType:" + splitFileBasedOnSplitTypeError.Error() + "|")
		}

	}

	pusherCommandBuilderFunctionResult["kubernetes_workload_split"] = splitFileBasedOnSplitTypeResult
	processPusherCommandError := processPusherCommand(pusherCommandBuilderFunctionResult)
	if processPusherCommandError != nil {
		return errors.New("|CloudHorseman->commandprocessors->pusher_command_processor->ProcessPusherCommandInteractively->processPusherCommand:" + processPusherCommandError.Error() + "|")
	}

	return nil
}

func ProcessPusherCommandFromParametersMap(parametersMap map[string]interface{}) error {
	processPusherCommandError := processPusherCommand(parametersMap)
	if processPusherCommandError != nil {
		return errors.New("|CloudHorseman->commandprocessors->pusher_command_processor->ProcessPusherCommandFromParametersMap->processPusherCommand:" + processPusherCommandError.Error() + "|")
	}

	return nil
}
