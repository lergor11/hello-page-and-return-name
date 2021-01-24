# Hello-page-and-return-name - example webserver


### ENV

| ENV               |    Default               |
|-------------------|--------------------------|
| HTTP_ADDR         | `:8080`                  |
| USER_LOG_FILE     | `current dir + user.log` |
| SQLITE_FILE       | nil                      |

`
If env SQLITE_FILE variable is set, then the log will be written to the database
`

### URL

GET `/hello`:  response "Hello page"

GET `/user?name=*name*`:  response all data when the method POST was requested for this users

POST `/user?name=*name*`: response 200 and log this users

GET `/metrics`: response metrics

GET `/status`: response 200 and "ok"

### Starting service

```shell
$ go run cmd/hello-page-and-return-name/main.go
```


### Starting service in containers

```sh
$ docker build -t hello-page-and-return-name .
$ docker run --rm --name hello-page-and-return-name -it -p 8080:8080 hello-page-and-return-name
```

### Starting service in k8s

All variables are change in ```helm/values.yaml```
```sh
$ helm upgrade --install --wait --namespace <*namespace*> hello-page-and-return-name ./helm
```