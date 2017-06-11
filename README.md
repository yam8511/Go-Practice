# Initial Set Up
```shell=
docker run -tid --name=go_web -p 8000:8000 -v `pwd`:/app -w /app golang:latest
docker exec -ti go_web bash
go get github.com/yam8511/ZuZuGo
go build
./app
(Ctrl + C  to stop server)
(Ctrl + D  to leave terminal)
```

# Start docker container & Start Server
```shell=
docker start go_web && docker exec -ti go_web bash
go build
./app
(Ctrl + C  to stop server)
(Ctrl + D  to leave terminal)
```

# Stop docker container
```shell=
docker stop go_web
```

# Remove docker container
```shell=
docker rm -f go_web
```
