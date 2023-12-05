# SRE_GO_Backend

物流後端練習

docker build -t sparkdoaz/sre-web .

[https://hub.docker.com/repository/docker/sparkdoaz/sre-web/general](https://hub.docker.com/repository/docker/sparkdoaz/sre-web/general)

docker build -t sparkdoaz/sre-web .

docker build -t sparkdoaz/sre-web:v1 .

docker build -t sparkdoaz/sre-web:v1 . && docker push sparkdoaz/sre-web:v1

docker build -t sparkdoaz/sre-web:v1 --push .

docker tag sre-web:latest sparkdoaz/sre-web:latest

docker push sparkdoaz/sre-web:latest

docker buildx build . --platform linux/amd64,linux/arm64 --push -t sparkdoaz/sre-web:multiple
