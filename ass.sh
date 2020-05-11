docker build -t consumer -f cons.dockerfile .
docker image tag consumer controllerbase/consumer
docker push controllerbase/consumer

docker build -t producer -f prod.dockerfile .
docker image tag producer controllerbase/producer
docker push controllerbase/producer