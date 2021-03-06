# Go API client for tsuru

Open source, extensible and Docker-based Platform as a Service (PaaS)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.6.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./tsuru"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AppApi* | [**Create**](docs/AppApi.md#create) | **Post** /apps | 
*AppApi* | [**List**](docs/AppApi.md#list) | **Get** /apps | 


## Documentation For Models

 - [App](docs/App.md)
 - [AppCreateResponse](docs/AppCreateResponse.md)
 - [AppListResponse](docs/AppListResponse.md)
 - [AppListResponseInner](docs/AppListResponseInner.md)
 - [ErrorMessage](docs/ErrorMessage.md)
 - [Lock](docs/Lock.md)
 - [Plan](docs/Plan.md)
 - [Router](docs/Router.md)
 - [Unit](docs/Unit.md)
 - [Url](docs/Url.md)


## Documentation For Authorization

## Bearer
- **Type**: API key 

Example
```
	auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
		Key: "APIKEY",
		Prefix: "Bearer", // Omit if not necessary.
	})
    r, err := client.Service.Operation(auth, args)
```

## Author



