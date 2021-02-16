# Plus500 v1 RESTful API

A RESTful API for Plus500 v1 written in Go.

***Notice***: This application is still under development.

## Endpoints

***Notice***:
There is actually one endpoint defined

### Global Methods

* [SCHEME:HOST:PORT]/plus500/v1/home

## Environment

You need to create a .env file in your root folder.

```shell
API_PLUS500_VERSION="0"
API_PLUS500_URL="https://marketools.plus500.com"
API_PLUS500_USERAGENT="GoPlus500Bot/1.0"

#
API_SERVER_HOST="http://localhost"
API_SERVER_PATH_PREFIX="/plus500/v1"
API_SERVER_PORT="8383"
API_SERVER_PATH_SRC="/src"
API_SERVER_CLIENT_LIMIT="10"

```

## Docker

### Build

***Notice***:
Using semantic versioning at the end, which is also tagged at github

```dockerfile
docker build -t deemount/plus500:v0.1.1 .
```

### Run

```dockerfile
docker run --publish 8686:8686 --detach --name pls5 deemount/plus500:v0.1.1  
```

## History

* create architecture, add functionalities, first setup
* first upload & initial commit

### To Do's

* github tags
* more methods & functionality
* more documentation

### More

* id's for later usage are saved in /src directory
