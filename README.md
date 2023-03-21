# Test-BE-IMP

Backend Test Code IMP ([PT Informatika Media Pratama](http://impstudio.id))

# Folder Structure in Project
```bash
├── api
│   └── swagger.yml                 # swagger spec file
├── build                           # executable file build
├── cmd
│   └── server                      # main application folder
├── config.yaml                     # main application config
├── gen                             # go-swagger generated folders
│   ├── models
│   └── restapi
├── generate.go
├── go.mod
├── go.sum
├── internal
│   ├── app                         # private application code
│   │   ├── handlers                # handler layer
│   │   ├── rest                    # route layer
│   │   └── server.go               # main server code
│   ├── app.go                      # # main application code
│   └── pkg
│   │   ├── constant.go             # internal pkg constant
│   │   ├── errors                  # internal pkg error
│   │   └── migrations              # internal pkg migration
│   └── utils                       # internal utils
├── Makefile
├── migrations                      # migrations file
├── pkg                             # shared pkg application
│   ├── config                      # application config
│   ├── db                          # db config
│   └── log                         # application log
```

# Built With
- [Go v1.20](https://go.dev/)
- [Gorm](https://gorm.io/)
- [Go-swagger](https://github.com/go-swagger/go-swagger)
- [Golang-Migrate](https://github.com/golang-migrate/migrate)
- [MySQL](https://www.mysql.com/)
- [zerolog](https://github.com/rs/zerolog)
- [Godoc](https://go.dev/blog/godoc)
- [cli](https://github.com/urfave/cli)

## Prerequisite
- [go-swagger](https://pkg.go.dev/github.com/go-swagger/go-swagger)
- [golang-migrate](https://pkg.go.dev/github.com/golang-migrate/migrate/v4)

## Getting Started
1. Install Go on your device
2. Clone repository
3. Setup application config 
4. Running Project

## Adding API
1. Edit swagger.yml File in api/swagger.yml. If you need help with swagger docs use the [Editor](https://swagger.io/docs/open-source-tools/swagger-editor/) and read the [Documentation](https://swagger.io/docs/specification/about/).
2. To Generate Routes and Validation from Swagger API run `make generate`.

## Need Help to Running App? 
```bash  
$ make help
```
it will show all option to running app

## Running the app

```bash
$ make all
$ make run          # for running binary file
$ make run-local    # for running main.go file
```

## Want to see all routes?
- run app `make run` || `make run-local`
- go to `http://localhost:8080/docs`

## Migration

- run command `migrate create -ext sql -dir migrations/schema create_tablename_table`
- write sql code in migration file
- run migration `make migrate-up` it will run migration automatically

## Using Godoc

- install godoc 
    ```shell 
    $ go get golang.org/x/tools/cmd/godoc
    ```
- add comment on the package, reference : https://tip.golang.org/doc/comment
- run godoc
    ```shel
    godoc -http:=6060
    ```
- open `localhost:6060` or `127.0.0.1:6060`. it will show go documentation
- access on `http://localhost:6060/pkg/test-be-IMP/internal/app/usecases/?m=all`

## How to run docker
1. Install docker, reference : https://docs.docker.com/engine/install/ubuntu/

2. Run this command
```bash
sudo docker build -t server .
sudo docker run --network='host' server
```

3. access on 127.0.0.1:3000

    *note server = image name
