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
  docker run -ti -p {host-port}:8080 -p 20001-20009:20001-20009 -v {path-to-project}:/data/restbird restbird/rest
~~~

> Here we reserve ports 20001 to 20009 for Mock Server demo

Open your browser to access http://{host-ip}:{host-port}/ , use default user ***admin/admin*** to login.

Enter "Rest Project" menu, and test run all APIs directly.
> The default host is set to http://127.0.0.1:8080, so all APIs will request the docker's localhost:8080 which is the listening prot inside container.
 
> Test server can be pointed to other location by change the {{host}} environment variable
  

## Project organization

- Rest Project
	- Restbird API : Contains Restbird APIs, runing the APIs directly 
		- MockServer
		- RestProject
		- Task
	-  Test_rest
		- Test
- Mock Server : Demo project of Mock Server
	- Demo
		- Hello_GoLang
		- Hello_JS
		- Hello_Python
- Task : Demo project of Task
	- DemoTask    