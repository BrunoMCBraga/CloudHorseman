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

const kubeCtlPullAttemptsSleepTime int = 60
const maxKubeCtlPullAttempts int = 5

func processPullerCommand(parametersMap map[string]interface{}) error {
	var commandReader *bufio.Reader = bufio.NewReader(os.Stdin)

	kubernetesGetPodsForDeploymentResult, kubernetesGetPodsForDeploymentError := haymakercfengines.KubernetesGetPodsForDeployment(parametersMap["kubernetes_deployment_name"].(string))
	if kubernetesGetPodsForDeploymentError != nil {
		return errors.New("|CloudHorseman->commandprocessors->puller_command_processor->processPullerCommand->KubernetesGetPodsForDeployment:" + kubernetesGetPodsForDeploymentError.Error() + "|")
	} else if kubernetesGetPodsForDeploymentResult == nil {
		return errors.New("|CloudHorseman->commandprocessors->puller_command_processor->processPullerCommand->KubernetesGetPodsForDeployment: returned null map of pods.|")
	}

	kubernetesGetPodsForDeploymentResultCopy := util.CloneStringBoolMap(kubernetesGetPodsForDeploymentResult)

	currentKubeCtlPullAttempt := 0

	var kubernetesCopyFileFromOrToRemotePodsResult map[string]string
	var kubernetesCopyFileFromOrToRemotePodsError error

	for true {
		kubernetesCopyFileFromOrToRemotePodsResult, kubernetesCopyFileFromOrToRemotePodsError = haymakercfengines.KubernetesCopyFileFromRemotePods(
			kubernetesGetPodsForDeploymentResultCopy,
			parametersMap["kubernetes_remote_pod_file"].(string),
			parametersMap["kubernetes_local_pod_file_folder"].(string))

		if kubernetesCopyFileFromOrToRemotePodsError != nil {
			fmt.Println("|CloudHorseman->commandprocessors->puller_command_processor->processPullerCommand->KubernetesCopyFileFromOrToRemotePods:" + kubernetesCopyFileFromOrToRemotePodsError.Error() + "|")
			continue
		} else if kubernetesCopyFileFromOrToRemotePodsResult == nil {
			fmt.Println("|CloudHorseman->commandprocessors->puller_command_processor->processPullerCommand->KubernetesCopyFileFromOrToRemotePods: returned null map of pods and local file paths.|")
			continue
		}

		//This is not good enough to make sure the file has been transferred....with cp to local system, sometimes an empty file is created...
		for podName, localFile := range kubernetesCopyFileFromOrToRemotePodsResult {
			if util.FileExists(localFile) {
				fmt.Println("|CloudHorseman->commandprocessors->puller_command_processor->processPullerCommand:" + fmt.Sprintf("Pod:%s --> File:%s", podName, localFile) + " |")
				delete(kubernetesGetPodsForDeploymentResultCopy, podName)
			}
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
				currentKubeCtlPullAttempt = 0
				kubernetesGetPodsForDeploymentResultCopy = util.CloneStringBoolMap(kubernetesGetPodsForDeploymentResult)
				continue
			} else {
				break
			}

		} else {
			time.Sleep(time.Duration(kubeCtlPullAttemptsSleepTime) * time.Second)
		}
		if currentKubeCtlPullAttempt > maxKubeCtlPullAttempts {
			fmt.Print(globalstringsproviders.GetPodFileRetrievalMaxAttemptString())
			commandReader.Reset(os.Stdin)
			readStringResult, readStringError := commandReader.ReadString('\n')
			if readStringError != nil {
				return errors.New("|CloudHorseman->commandprocessors->puller_command_processor->commandReader.ReadString:" + readStringError.Error() + "|")
			}

			command := strings.Trim(readStringResult, " \n")
			if command == "y" {
				currentKubeCtlPullAttempt = 0
				continue
			} else {
				break
			}
		} else {
			currentKubeCtlPullAttempt++
		}

	}

	return nil

}

func ProcessPullerCommandInteractively() error {

	pullerCommandBuilderFunction := commandbuilders.BuildPullerCommand
	pullerCommandBuilderFunctionResult, pullerCommandBuilderFunctionError := pullerCommandBuilderFunction()
	if pullerCommandBuilderFunctionError != nil {
		return errors.New("|CloudHorseman->commandprocessors->puller_command_processor->ProcessPullerCommandInteractively->pullerCommandBuilderFunction:" + pullerCommandBuilderFunctionError.Error() + "|")
	} else if pullerCommandBuilderFunctionResult == nil {
		return nil
	}

	processPullerCommandError := processPullerCommand(pullerCommandBuilderFunctionResult)
	if processPullerCommandError != nil {
		return errors.New("|CloudHorseman->commandprocessors->puller_command_processor->ProcessPullerCommandInteractively->processPullerCommand:" + processPullerCommandError.Error() + "|")
	}

	return nil
}

func ProcessPullerCommandFromParametersMap(parametersMap map[string]interface{}) error {
	processPullerCommandError := processPullerCommand(parametersMap)
	if processPullerCommandError != nil {
		return errors.New("|CloudHorseman->commandprocessors->puller_command_processor->ProcessPullerCommandFromParametersMap->processPullerCommand:" + processPullerCommandError.Error() + "|")
	}

	return nil
}
