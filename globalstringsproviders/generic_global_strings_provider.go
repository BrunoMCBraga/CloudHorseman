package globalstringsproviders

import (
	"fmt"
	"strings"
)

func GetCachedStackParametersWarningString() string {

	cachedStackParametersWarningString := `Cloud Horseman detected cached parameters. Select one of the following stack names so that parameters can be automatically set (you will still have the chance to review them):`

	return strings.Trim(cachedStackParametersWarningString, "\n")

}

func GetStackNamePromptString() string {

	stackNamePromptString := `Stack > `

	return stackNamePromptString

}

func GetCacheKubernetesdDeploymentParametersWarningString() string {

	cachedDeploymentParametersWarningString := `Cloud Horseman detected cached parameters. Select one of the following Kubernetes deployment names so that parameters can be automatically set (you will still have the chance to review them):`

	return cachedDeploymentParametersWarningString

}

func GetKubernetesDeploymentNamePromptString() string {

	stackNamePromptString := `Kubernetes Deployment > `

	return stackNamePromptString

}

func GetMenuPictureString() string {

	pictureString := `
   ________                __   __  __                                         
  / ____/ /___  __  ______/ /  / / / /___  _____________  ____ ___  ____ _____ 
 / /   / / __ \/ / / / __  /  / /_/ / __ \/ ___/ ___/ _ \/ __  __ \/ __  / __ \
/ /___/ / /_/ / /_/ / /_/ /  / __  / /_/ / /  (__  )  __/ / / / / / /_/ / / / /
\____/_/\____/\__,_/\__,_/  /_/ /_/\____/_/  /____/\___/_/ /_/ /_/\__,_/_/ /_/                                                                                
`
	return pictureString

}

func GetInteractivePromptString() string {

	commandPrompt := `Cloud Horseman > `

	return commandPrompt
}

func GetFormattedInteractivePromptString(suffix string) string {

	commandPrompt := `Cloud Horseman %s > `

	return strings.Trim(fmt.Sprintf(commandPrompt, suffix), "\n")
}

func GetPodFileRetrievalMaxAttemptString() string {

	podFileRetrievalMaxAttemptString := `The maximum number of attempts to retrieve the requested files from deployment pods has been reached. Do you with to reset the counter and retry?(y/any key for no):`

	return podFileRetrievalMaxAttemptString
}

func GetPodStateMaxAttemptString() string {

	podStateMaxAttemptString := `The maximum number of attempts to retrieve the pods state has been reached. It is likely that something went wrong with Kubernetes and due configurations. Do you with to reset the counter and retry?(y/any key for no):`

	return podStateMaxAttemptString
}

func GetWorkloadUploadString() string {

	workloadUploadString := `Do you wish to push the workload to the pods? This will trigger the execution (e.g. DDOS will start) (y/any key for no):`

	return workloadUploadString
}

func GetPodFileRetrievalSuccessfulPullString() string {

	podFileRetrievalSuccessfulPullString := `All the files were pulled. Do you wish to reset the counter and pull again? This can be useful if the output is constantly being written and the output file updated (y/any key for no):`

	return podFileRetrievalSuccessfulPullString
}

func GetPodFileUploadMaxAttemptString() string {

	podFileUploadMaxAttemptString := `The maximum number of attempts to upload the requested files to deployment pods has been reached. Do you with to reset the counter and retry?(y/any key for no):`

	return podFileUploadMaxAttemptString
}

func GetDeploymentStopString() string {

	deploymentStopStringString := `Would you like to stop the processing on all containers? (y/any key for no):`

	return deploymentStopStringString
}

func GetHelpString() string {

	helpString := `
Core Commands
=============

    Command       Description
    -------       -----------
    ?             Help menu
    set           Sets a context-specific variable to a value
    use           Used to switch to the context of a module (e.g. use avatar)
    exit          Exit the console
    help          Help menu
    version       Show the framework and console library version numbers


Module Commands
===============

	Command       Description
	-------       -----------
	back          Move back from the current context
	info          Displays information about one or more modules
	options       Displays global options or for one or more modules
	show          Displays all modules categorized
	load          Load set parameters from JSON file

`
	return helpString

}

func GetVersionString() string {

	versionString := `
Framework: 1.0 
Console: 1.0
`
	return strings.Trim(versionString, "\n")

}

func GetShowModulesString() string {

	showModulesString := `
Stack/Deployment Management
===========================

	#   Name                          Description
	-   ----                          -----------
	0   cfbuilder                     Deploys CloudFormation stack
	1   cfteardown                    Tears down CloudFormation stack
	2   depteardown                   Tears down Kubernetes deployment and service


Container Deployers
===================

	#   Name                          Description
	-   ----                          -----------
	0   avatar                        Generic container deployer


Pullers/Pushers
===============

	#   Name                          Description
	-   ----                          -----------
	0   puller                        Pulls (downloads) files from each pod under a deployment
	1   pusher                        Pushes (uploads) files to each pod under a deployment
	   
`
	return showModulesString

}
