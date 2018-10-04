# Identity CLI (ISVR)
The Identity Server CLI helps you manage your IdentityServer client and resource by calling the IdentityServer.ConfigurationApi

To install the CLI  you can either compile the program yourself using `go install` (see 2.) or getting the latest executable from the artifact folder.

## 1 - How to use the CLI ?

### 1.1 - Setting up which Configuration API to call
The first thing to do is to set up the address of the configuration API by setting the environment variable `ISVR_HOST`. The default value is set to: `localhost:5020`
To check you configuration you can use the following command :
```
isvr info
```

### 1.2 - Creating / Updating clients and resources
In order to create/ update new clients and resources you need first to declare them in a yaml the `AuthConfiguration`. Here a simple example that will declare a Mvc Application and a Web Api :
``` yml
clients:
  - id: "MvcApplication1"
    name: "MvcApplication1"
    allowed_grant_types:
      - "implicit"
    redirect_uris:
      - "http://localhost:5002/signin-oidc"
    post_logout_redirect_uris:
      - "http://localhost:5002/signout-callback-oidc"
resources:
  - name: "WebApi1"
    description: "A simple ASP.NET Core API"
    secrets:
      - "asecretvalue"
      - "anothersecretvalue"
    scopes:
      - name: "Administration"
        description: "Access to the administration interface"
      - name: "Finance"
        description: "Access to the finance data"
```
Other Samples files can be found under */samples*.
Once you file is defined you just need to run the following command to apply your configuration 
```
isvr apply -f <PATH_TO_CONFIGURATION_FILE>
```
The default path is set to *./AuthConfiguration.yml*

### 1.3 - Listing and Deleting existing Clients and Resources
To list the resources or clients use the following command

#### List elements with GET
```
isvr get clients
```
```
isvr get resources
```
**You can choose to return only IDs using `-q ` flag**

```
isvr get resources -q
```


#### Delete elements with DELETE
To delete an element you can use one of these commands.
```
isvr delete clients <CLIENT_ID>
```
```
isvr delete resources <RESOURCE_ID>
```

#### Delete All Resources/Clients
Becaus commands can be combined you can delete all resources or clients by using one of the following command:
```
isvr delete clients $(isvr get clients -q)
```
```
isvr delete resources $(isvr get resources -q)
```

### 1.4 - Showing help 
If you forget how to use a command you can use th `-h` flag or `--help` on any command to get more informations about it.


## 2 - Getting Started with Development
Identity CLI have been developed using Golang and Cobra (a package to help you to write powerful CLI). 

**Don't know Golang ?**
You don't know Golang don't worry it's very easy to pick up, you can have a look at the official [Tour of Go ](https://tour.golang.org/welcome/1) to help you start. 

### 2.1 - Setting Up your environment
First you need to install go by following theses instructions :
[Installing Go - Getting Started](https://golang.org/doc/install)

If you never developed with go i strongly recommend you to read this article first, it explains how worspaces are organised: 
[How to write Go Code](https://golang.org/doc/code.html)

### 2.2 - Develop and Debug using Visual Studio Code 
I recommend you to use Visual Studio Code. It will be easier to debug as the  
`launch.json` have already been configure to run the command of you choice.
Like usual just need to click on F5 to start debugging.

### 2.3 - Project Structure
As you will probably notice the project structure follow the same structure as the docker CLI. As i am not (yet) a Go Expert some mechanism are a direct fork of their code source, don't hesitate to refer to it.

### 2.3 - Contributing 
Please have a look at the CONTRIBUTING.md file before starting your development.






 
 








