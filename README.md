# GO-DOCKER

## description

- gin
- gorm
- gin-swagger
- mysql
- air

## install

touch src/.env

```
HTTP_HOST=""
HTTP_PORT=8080
DB_HOST="mysql"
DB_PORT=3306
DB_USERNAME=root
DB_PASSWORD=password
DB_DATABASE=gin_db
```

docker

```
docker-compose build
docker-compose up
```

## usage

api base url

```
http://localhost:3000/
```

swagger url
```
http://localhost:3000/swagger/index.html
```

update swagger

```
docker exec -it go-docker_app_1 /bin/sh

/go/src # go get -u github.com/swaggo/swag/cmd/swag

/go/src # swag init
```

update go.mod
```
docker-compose run --rm app go mod tidy
```