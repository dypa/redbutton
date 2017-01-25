CGO_ENABLED=0 go build

docker-compose up

http://localhost:6001/

docker build -t demo . && docker run --publish 8888:8888 demo
