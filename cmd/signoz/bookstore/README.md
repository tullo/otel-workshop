# Bookstore REST API using Gin and Gorm

Read the [article](https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/).


## Build and configure app to send data to SigNoz:

```sh
go build
SERVICE_NAME=Bookstore INSECURE_MODE=true OTEL_EXPORTER_OTLP_ENDPOINT=<IP of SigNoz backend:4317> ./gin-bookstore
```

This runs the gin application at port `8090`. Try accessing API at `http://localhost:8090/books`

Below are the apis available to play around. The API calls will generate telemetry data which will be sent to SigNoz which can be viewed at `http://<IP of SigNoz backend:3000`

## Endpoint and http verbs available:
```
GET    /books                    
GET    /books/:id               
POST   /books                    
PATCH  /books/:id                
DELETE /books/:id                
```

## Gin server boot output when run in debug mode

```sh
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /books       --> github.com/tullo/otel-workshop/cmd/signoz/bookstore/controllers.FindBooks (4 handlers)
[GIN-debug] GET    /books/:id   --> github.com/tullo/otel-workshop/cmd/signoz/bookstore/controllers.FindBook (4 handlers)
[GIN-debug] POST   /books       --> github.com/tullo/otel-workshop/cmd/signoz/bookstore/controllers.CreateBook (4 handlers)
[GIN-debug] PATCH  /books/:id   --> github.com/tullo/otel-workshop/cmd/signoz/bookstore/controllers.UpdateBook (4 handlers)
[GIN-debug] DELETE /books/:id   --> github.com/tullo/otel-workshop/cmd/signoz/bookstore/controllers.DeleteBook (4 handlers)
[GIN-debug] Listening and serving HTTP on :8090
```
