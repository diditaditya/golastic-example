# Golang - ElasticSearch Example

This is simply taken from [this amazing post](https://outcrawl.com/go-elastic-search-service) on getting started with dockerized Golang and ElasticSearch development tutorial.

There are few changes as the following:
1. Instead of using `go-dep`, `go mod` is used instead
2. ElasticSearch v6.8.2 instead of v6.2.3
3. The content of `main.go` is separated based on the concern as separate packages

Note that this was tested on windows 10 wsl ubuntu 18.04, however the path of `volumes` in `docker-compose.yml` has been changed to `${PWD}` prior to pushing the files to the repo.

## Usage

Clone the repo
```sh
$ git clone https://github.com/diditaditya/golastic-example.git
```

Run the `docker-compose`
```sh
$ cd golastic-example
$ docker-compose up
```