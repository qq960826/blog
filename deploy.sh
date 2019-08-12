docker stop $(docker ps -q --filter ancestor=qq960826/blog)
docker rm $(docker ps -all -q --filter ancestor=qq960826/blog)
docker pull qq960826/blog
docker run -d qq960826/blog