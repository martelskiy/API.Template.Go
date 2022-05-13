API template project suppose to be used as a starter for a new service development.  
Current template includes:

## Logger
[This](https://github.com/uber-go/zap) package is used for logger implementation

## Open API

[Swag](github.com/swaggo/swag) is used for an Open API in composition with [HTTP Swagger](github.com/swaggo/http-swagger) which helps to integrate with native `net/http` GO package.

In order to generate or update docs for Open API documentation it needs to run command (from the root):

```
swag init
```
Endpoints available in the template:  

* `/status` (it will be visible in _Swagger_)

_Swagger_ url (port may be altered): 
```
http://localhost:8080/swagger/index.html
```
## Web API Router

[Mux](github.com/gorilla/mux) is used as a router for a web API
```