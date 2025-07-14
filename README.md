# Go Yard

Personal Go Playground

## Todo's

Get the basic
* [x] Complete: https://go.dev/doc/tutorial/database-access
* [x] Complete: [Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)
* [x] Complete: [Getting started with generics](https://go.dev/doc/tutorial/generics)
* [x] Complete: [Getting started with fuzzing](https://go.dev/doc/tutorial/fuzz)
* [x] Complete: [Writting Web Applications](https://go.dev/doc/articles/wiki/)
* [] Complete: [Effective Go](https://go.dev/doc/effective_go)
* [] Complete: [How to write Go code](https://go.dev/doc/code)
* [x] Complete: [Go Slices: usage and internals](https://go.dev/blog/slices-intro)

## Set up Database

Use can set up a database by using podman

```bash
podman run -dt -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=recordings --name gotut-mysql -p 3306:3306 docker.io/library/mysql
```

Then etner the running container with podman exec

```
podman exec -ti gotut-mysql /bin/bash
```

Read [web-service-gin](https://go.dev/doc/tutorial/web-service-gin) for furhter
instructions.
