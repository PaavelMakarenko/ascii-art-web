echo "BUILDING IMAGE"
docker image build -f Dockerfile -t ascii-art-web .
echo
echo "IMAGES"
docker images
echo
sleep 2
echo "RUNNING IMAGE IN CONTAINER"
docker container run -p 8080:8080 --detach --name ascii-art-web-container ascii-art-web
echo
echo "CONTAINERS"
docker container ls
echo