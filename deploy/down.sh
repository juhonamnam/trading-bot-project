. ./deploy/app_info.sh

# Stop Running Containers
RUNNING_CONTAINERS=$(docker ps -q --filter ancestor=$APP_NAME:$APP_TAG)

if [ ${#RUNNING_CONTAINERS} -ne 0 ]; then
    echo "Stopping currently running containers..."
    docker stop $RUNNING_CONTAINERS
fi

# Delete Containers
ALL_CONTAINERS=$(docker ps -a -q --filter ancestor=$APP_NAME:$APP_TAG)

if [ ${#ALL_CONTAINERS} -ne 0 ]; then
    echo "Deleting containers..."
    docker rm $ALL_CONTAINERS
fi

# Delete Image
docker rmi -f $APP_NAME:$APP_TAG
