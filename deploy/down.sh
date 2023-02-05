. ./deploy/app_info.sh

# Stop Running Containers
RUNNING_CONTAINERS=$(docker ps -q --filter ancestor=$APP_NAME:$APP_TAG)

if [ ${#RUNNING_CONTAINERS} -ne 0 ]; then
    echo "exec: docker stop $RUNNING_CONTAINERS"
    docker stop $RUNNING_CONTAINERS
fi

# Delete Containers
ALL_CONTAINERS=$(docker ps -a -q --filter ancestor=$APP_NAME:$APP_TAG)

if [ ${#ALL_CONTAINERS} -ne 0 ]; then
    echo "exec: docker rm $ALL_CONTAINERS"
    docker rm $ALL_CONTAINERS
fi

# Delete Image
echo "exec: docker rmi -f $APP_NAME:$APP_TAG"
docker rmi -f $APP_NAME:$APP_TAG
