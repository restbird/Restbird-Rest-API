# Restbird-Rest-API
This is the demo project for [APIs of Restbird](https://restbird.org/docs/API-overview.html)

## Getting started
To run the project

*  Download Restbird docker

~~~
  docker pull restbird/rest
~~~

* Start Restbird docker and mapping this project directory as Restbird working directory

~~~
  docker run -ti --user $(id -u):$(id -g) --net host -v {path-to-project}:/data/restbird restbird/rest
~~~

Open your browser to access http://{host-ip}:8080/ , use default user ***admin/admin*** to login.

Enter "Rest Project", "Mockserver" and "Task" menu, and test run all APIs directly.
You can also download [Restbird Debugger VSCode extension](https://marketplace.visualstudio.com/items?itemName=Restbird.vscode-restbird) to debug test scripts through [VSCode](https://code.visualstudio.com/).

> The default host is set to http://127.0.0.1:8080, so all APIs will request the docker's localhost:8080 which is the listening prot inside container.
 
  

## Project organization

- Rest Project
	-  Demo_Rest : Demo project of rest APIs, script written in Golang and Python
		- Rest_Golang
		- Rest_Python

	- Restbird API : Contains all Restful APIs of Restbird 
		- MockServer
		- RestProject
		- Task
- Mock Server : Demo project of Mock Server,  script written in Golang, Javascript and Python
	- Demo
		- Hello_GoLang
		- Hello_JS
		- Hello_Python
		
- Task : Demo project of Task, script written in Golang and Python
	- Demo_Task_Golang
	- Demo_Task_Python    