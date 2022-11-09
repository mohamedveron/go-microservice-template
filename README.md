# go-microservice-template

This is a standard template to follow for any new microservice in golang


## Setup of the component:

Must have golang installed version >= 12.0.0

make file consists of 4 steps: generate, test, build, run you can run all of them

```make all```

Run test cases run

```make test```
or

```go test -v ./tests```

Start the http server on port 9090:

```make run```

## Run By Docker

```
 docker build -t go_graphql .
 
 docker run -p 9090:9090 go_graphql
 
```
