# cli-demo
Project for Demoing CLI Development using golang

### Use Case

Hi, I am Ram and I am devops professional employed by a pet store which deals with adoption of pets. 
Company has another employee named Shyam who works on store's front desk and deals with adopters. 
This employee is a volunteer who is not so tech savy and has requested an easy way to deal with day to day activities 
like getting pets, add new pets and so on and so forth.

Shyam asked Ram if he can give a very easy to use system to interact with backend database which is maintained by
Pet Adoption community. So Ram decided to give Shyam a very easy to use cli program which he can use to get information
he need easily and efficiently.

## Agenda

In this branch Ram will initialize go project and get dependencies for creating CLI.

He has decided to use following tools:

* [GoLang](https://go.dev/) : Language which Ram will use to create CLI.
* [Cobra](https://github.com/spf13/cobra) : It's an opensource module which provides framework to create CLI's in golang.
* [Viper](https://github.com/spf13/viper) : It's an opensource module which helps to manage configuration for any go program.
* [osapi-codegen](https://github.com/deepmap/oapi-codegen) : Helps to create client/server code based on open api specification.
* [go-pretty]([https://github.com/jedib0t/go-pretty) : helps in outputting text in cli application in go.

## Steps

### Step 1 : Go Module

Lets create go module and start with creating app directory which will host go code and then create module

```go
mkdir cli-demo
cd myapp
go mod init github.com/ankurnema/cli-demo
```

### Step 2: Download dependencies

Lets download go dependencies we will need to create our commands

```go
go get -u github.com/spf13/cobra@latest
go get github.com/spf13/viper
go get github.com/jedib0t/go-pretty/v6

# For generating code using osapi-codegen we need to install this tool
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

# For using cobra CLI we need to install this tool
go install github.com/spf13/cobra-cli@latest
```

### Step 3: Project Folder Structure

Lets create folder structure for organizing our code.

```text
cli-demo/           # Root Directory
├── cmd/            # Directory which contains all command high level code
├── pkg/            # Contain packages which are public in nature
│   └── petstore/   # Package containing petstore client code and type code
└── internal/       # dicrectory containing internal code.
```

### Step 4: Petstore Client

Since we want to create a wrapper around petstore using this cli tool, we will need a client for interacting with 
petstore api's. Lets create our client using oapi-codegen tool. For this to work we need open api specification of
petstore. We also need to access petstore server.

Petstore Server: https://petstore.swagger.io/

API Key: special-key

Petstore open api spec: https://petstore.swagger.io/v2/swagger.json 

Lets generate client using oapi-codegen

```shell
cd pkg/petstore
oapi-codegen -package petstore --generate types,client petstore.yaml > petstore.gen.go
```

Above command will generate petstore client in petstore package.