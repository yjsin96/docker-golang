# Golang Docker

## Docker command
### Build Image 
```sh
docker build -t gogin .
```
### Run container
```sh
docker run \
--name gogin \
-dit \
-p 8000:8000 \
-e DB_USER=root \
-e DB_HOST=host.docker.internal:3306 \
-e DB_NAME=docker_test \
-e DB_PASSWORD= \
-v "$(pwd)/storage:/app/storage" \
gogin
```


## Docker with network
### Create network
```sh
docker network create test-docker
```
### Start MySQL container
```sh
docker run -d \
--name mysql \
--network test-docker \
-v mysql-data:/var/lib/mysql \
-e MYSQL_ROOT_PASSWORD=secret \
-e MYSQL_DATABASE=docker_test \
mysql
```
### Start golang container
```sh
docker run \
--name gogin \
--network test-docker \
-dit \
-p 8000:8000 \
-e DB_USER=root \
-e DB_HOST=mysql:3306 \
-e DB_NAME=docker_test \
-e DB_PASSWORD=secret \
-v "$(pwd)/storage:/app/storage" \
gogin
```


## Docker compose
### Start docker compose
```sh
docker-compose up -d
```
### Stop docker compose
```sh
docker-compose down
```
### Start docker-compose-volume docker compose
```sh
docker-compose -f docker-compose-volume.yml up -d
```
### Start docker-compose-network docker compose
```sh
docker-compose -f docker-compose-network.yml up -d
```
