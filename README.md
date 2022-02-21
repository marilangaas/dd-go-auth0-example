Repo for demonstrating how `Transport` is overwritten when using datadog [dd-trace-go.v1/contrib/net/http tracing](https://pkg.go.dev/gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http) 
in custom http client given as argument to WithClient when using the go-auth0 package to initialize a management client. 

To use this repo ...

1. fill in config values in `config/local.yaml`:
```
auth0_domain: ""
auth0_client_id: ""
auth0_client_secret: ""
```

2. send in a valid `id` to `GetUserEmail` in `server.go`


