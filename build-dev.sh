imageName=bayungrh/gosseract-http-dev
containerName=gosseract-http-dev

echo ":: Build image"
docker build -t $imageName  .

echo ":: Kill container"
docker kill $containerName || true

echo ":: Delete old container..."
docker rm -f $containerName || true

echo ":: Run new container..."
docker run -dti -p 8080:8080 \
--name $containerName \
-v "$(pwd)":/app \
$imageName

echo ":: Clean all <none> tag image..." 
docker rmi $(docker images --filter "dangling=true" -q --no-trunc)q