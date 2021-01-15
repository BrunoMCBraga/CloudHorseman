# Cloud Horseman

## Description
Cloud Horseman is command-line tool built on top of https://github.com/BrunoMCBraga/HayMakerCF and was born out of frustration when choosing a new laptop. I wanted a powerful computer but did not want to spend my 401K on it. As such, I decided to build a tool with a Metasploit-like CLI to deploy distributed workloads on AWS using containers. 

You can now perform a global scan, DDOS the company that promoted the good-looking guy instead of you and even create a proxy on AWS with a couple of clicks.

## The Project Structure And Philosophy
- Util contains generic utility code such as JSON processors, random generators, structures, etc
- Commandbuilders, commandcachegenerators and commandprocessors deal with the CLI. They collect inputs, manage menus and provide some degree of caching so that you don't need to keep defining the same variables across modules
- Globalstringproviders: where all the strings are stored
- Internal contains code responsible for internal checks on variables and mappings from UI variables to internal fields in structures
- Types store custom types
- CloudFormationFiles is similar to HayMakerCF folder and contains adapted CloudFormation templates for Cloud Horseman

In order to simplify the code and increase modularity, I chose to couple the code with bash scripts and input/output files like so:
- DockerTemplates contains multiple folders for different workloads (i.e. tools) to run in AWS. 
- The main file in each folder is the Dockerfile that sets up basic environment and passes the ball to a bash script called exec.sh
- The bash script follows a more-or-less standard template: It starts by killing existing workloads with the same name, waits for the input file (copied by Cloud Horseman to containers), executes and stores output on output file which is then copied back from each container to the local host. The output represents the results of your workload (e.g. cracked hashes)

In order to speed up deployment I have defined the command "load [PATH_TO_JSON_TEMPLATE]" that you can use inside each module to set the module variables in batch. Check the folder LoadTemplates.

Using Bash scripts to extend functionality is a two-edge-sword: on the one hand you can easily copy paste and extend the tool easily. On the other hand, you need to specify ports and workload-specific fields (e.g. ports for udpflood) on the bash script. Not much of a hassle but something to bear in mind. 

Regarding CloudFormation templates, same principle applies as with HayMakerCF: If you choose to change things like cluster and repository names, you may or may not need to adjust the cloudformation_cluster.json template. As a matter of fact, If you plan on changing something like that, you need to change the json file first, deploy and then use the names you wish to use on Cloud Horseman CLI. HayMakerCF has no state of what is running in the cloud so If you deploy a CF template through the Web UI and then choose to use Cloud Horseman just to deploy on an existing cluster, you can without problem (you will need to generate a kubeconfig file either manually or using HayMakerCF). Knowledge on CF is important but you can use the out-of-the-box templates I provide and you will be fine. 

## How Are Workloads Splitted And How Is The Execution Workflow?
Cloud Horseman has currently three ways to split workloads across Docker instances (default is SPLITINPUT): 
- SPLITINPUT: Splits the input file (e.g. lists of hosts to be scanned) across the number of replicas specified on REPLICAS field. This can be useful to crack hashes.
- DUPLICATEINPUT: Sends the same input file to all pods/REPLICAS
- MULTIPLYINPUT: For each line on input file, n amount of REPLICAS are allocated. If you have 2 IPs on the input file and you set REPLICAS to 10, 20 containers will be allocated and 10 of them will have the first IP. The other 10 will have the second IP. This can be useful for DDOS. 

Cloud Horseman deploys a Docker container with a bash script that waits for input file (REMOTEINPUT and LOCALINPUT). The file is then sent once you decide so using a pusher module. Pusher uses kubectl to copy the files to each pod and trigger execution by the bash script that is waiting. Then, Cloud Horseman waits and tries to fetch the output file (using a puller module) resulting from processing (e.g. cracked hashes) using kubectl again (REMOTEOUTPUT and LOCALOUTPUT).

The execution script then goes to wait state once more and is ready to execute code again using some provided input file. If you need to use the same deployment with different inputs, you need to use a pusher and puller modules to push the inputs and get the results.

Cloud Horseman also has the concept of stop files. This is useful for DDOS attacks where you need to stop the attack. If you execute avatar fully, you will be asked If you want to stop execution. If not, you will need to push (using a pusher) a stop file to the proper folder. Alternatively you can use cfteardown to kill the entire CF deployment (i.e. ECR repository, Kubernetes cluster).


## How To Use It?
As mentioned above, Cloud Horseman is a light frontend for HayMakerCF which means you can use it for everything: deploying a CF template, create containers, execute workload and get results, or you can jump in the middle and assume an existing CF template executed through some other means. **However, I would strongly advise you to use Cloud Horseman for the full lifecycle (i.e. from deployment of CF to execution of containers and download of results). Jumping in the middle of any of these steps has caveats you must be aware of (e.g. If you try to deploy a container using an existing CF deployment, you will fail if you don't have the kubeconfig file which is automatically generated by Cloud Horseman using HayMakerCF after cfbuilder is ran).**  

A full workflow would look like:
- Create user named "user" on AWS and deploy cloudformation_iam.json to setup the proper permissions
- Start Cloud Horseman interactive CLI
- Type: **use cfbuilder** (this switches to the module responsible to deploy CloudFormation template). Set the proper variables or use **load [JSON_FILE]** to set the variables using a JSON file (e.g. the one in LoadTemplates). Type **run** to execute module and wait for CF deployment to finish
- Once the CF template is deployed, switch to avatar using **use avatar**.
- Inside the module, you either load one of the LoadTemplates or you can set the fields manually and type **run**
- Follow instructions on CLI (e.g. trigger input files to be sent, stop execution, get results, etc)
- Kill the deployment and/or the entire CF infrastructure using depteardown and/or cfteardown, respectively

If you use CFBuilder module, Cloud Horseman will use HayMakerCF to generate a local kubeconfig file. You can then use kubectl client to query the state of the pods and clusters (e.g. kubectl get pods).

## Current Modules
You can inspect the list of modules in the main menu by typing **show**. Currently, there are 5 modules:
- CFBuilder/CFTeardown: Deploys/Tears down CloudFormation infrastructure from JSON. The out-of-the-box template "cloudformation_iam.json" creates EKS cluster, ECR repository, Network ACLs, etc.
- Avatar/Depteardown: Currently, Cloud Horseman is flexible enough to allow any workload. This means that it only requires a module to deploy (Avatar) and another to teardown the Kubernetes deployment. Why Not DepBuilder instead of Avatar? Avatar is flexible enough to deploy anything but in the future, people may choose to extend and add specific modules for attack tools and push functionality to golang code instad of external scripts. There may be multiple modules but to create a deployment but Depteardown is good enough for any teardowns.
- Pusher/Puller: Used to send or download files to/from pods. These modules are used to split work across Kubernetes pods

## Dependencies
- Docker
- Kubectl

## The Current Features
- HayMakerCF features are implemented
- Offensive Security modules: john the ripper, udp flooding, nmap, ncrack and ssh proxy

## Potential Bugs And Things To Improve
- Since I could not get the file copy code using Kubernetes to work (i.e. compilation issues, not a lot of code available online) I am making calls to kubectl through Cloud Horseman

## Found this useful? Help me buy a new Ozone Layer to save the endangered polar bears:
https://www.paypal.com/donate/?hosted_button_id=UDFXULV3WV5GL
