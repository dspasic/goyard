# Go Yard

Personal Go Playground

## Todo's

* Get the basic
* * Continue here: https://go.dev/doc/tutorial/database-access
* [] Read [Effective Go](https://go.dev/doc/effective_go)
* [] Read [How to write Go code](https://go.dev/doc/code)

## Set up Database 

Use can set up a database by using podman

```bash
podman run -dt -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=recordings --name gotut-mysql -p 3306:3306 docker.io/library/mysql
```

Then etner the running container with podman exec

```
podman exec -ti gotut-mysql /bin/bash
```

and follow the instruction of the tutorial.
