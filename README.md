# Golang Docker

## Docker command
### Build Image 
```sh
docker build -t go-example .
```

### Run container
```sh
docker run \
--name go-example \
-dit \
-p 8000:8000 \
-e DB_USER=root \
-e DB_HOST=mysql \
-e DB_NAME=docker_test \
-e DB_PASSWORD=secret \
-v "$(pwd)/storage:/app/storage" \
go-example
```

## Docker compose
### Start docker compose
```sh
docker-compose up
```

### Stop docker compose
```sh
docker-compose down
```

## Docker with network
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
--name go-example \
--network test-docker \
-dit \
-p 8000:8000 \
-e DB_USER=root \
-e DB_HOST=mysql \
-e DB_NAME=docker_test \
-e DB_PASSWORD=secret \
-v "$(pwd)/storage:/app/storage" \
go-example
```