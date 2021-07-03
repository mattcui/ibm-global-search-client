# Go API client for openapi

Search for resources with the global and shared resource properties repository integrated in the IBM Cloud platform. The search repository stores and searches cloud resources attributes, which categorize or classify resources. A resource is a physical or logical component that can be created or reserved for an application or service instance and is owned by resource providers, such as Cloud Foundry, IBM Kubernetes Service, or resource controller in IBM Cloud. Resources are uniquely identified by a Cloud Resource Name (CRN)  or by an IMS ID. The properties of a resource include tags and system properties. Both properties are defined in an IBM Cloud billing account, and span across many regions.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./openapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.global-search-tagging.cloud.ibm.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ResourceTypesApi* | [**GetSupportedTypes**](docs/ResourceTypesApi.md#getsupportedtypes) | **Get** /v2/resources/supported_types | DEPRECATED. Get all GhoST indices.
*SearchApi* | [**Search**](docs/SearchApi.md#search) | **Post** /v3/resources/search | Find instances of resources (v3)
*SearchApi* | [**V2Search**](docs/SearchApi.md#v2search) | **Post** /v2/resources/search | Find instances of resources (v2)


## Documentation For Models

 - [Error](docs/Error.md)
 - [Errors](docs/Errors.md)
 - [ResultItem](docs/ResultItem.md)
 - [ScanBody](docs/ScanBody.md)
 - [ScanResult](docs/ScanResult.md)
 - [SearchBody](docs/SearchBody.md)
 - [SearchResults](docs/SearchResults.md)
 - [SupportedTypesList](docs/SupportedTypesList.md)


## Documentation For Authorization



### IAM

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



