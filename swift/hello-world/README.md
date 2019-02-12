# Step 1 - CLI
Run our first docker container
```BASH
docker run hello-world
```

Run a swift container
```BASH
docker run --privileged -it --rm --name swiftfun swiftdocker/swift:latest swift
```

List running containers
```BASH
docker ps -a | grep swift
```
The container is running, if we quit the container and retry this command, we'll see that the container is down.

# Step 2 - Custom Image
Build the custom image
```BASH
docker build . -t btor/swift-webserver
```

Launch the web serer
```BASH
docker run -p=8181:8181 -it btor/swift-webserver
```

The same with GO ?
```
docker build . -t btor/go-webserver
```

```
docker run -p=8181:80 -it btor/go-webserver
```

With docker-compose
```
docker-compose up --build
docker exec -it go-webserver sh
go run database.go
```

