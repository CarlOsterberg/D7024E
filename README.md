# D7024E

build image: docker build . -t kadlab <br>
compose(windows): docker-compose up <br>
compose(ubuntu): sudo docker-compose --compatibility up -d <br>
connect to container cli: sudo docker exec -it d7024e_kademliaNodes_1 /bin/sh <br>
test coverage: go test -race -covermode=atomic -coverprofile=coverage <br>
Inspect test coverge: go tool cover -html=coverage <br>

