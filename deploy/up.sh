. ./deploy/app_info.sh

mkdir -p logs

# Build Image
echo "docker build --tag $APP_NAME:$APP_TAG ."
docker build --tag $APP_NAME:$APP_TAG .

# Run Container
echo "exec: docker run --name $APP_NAME -v $(pwd)/logs:/trading-bot-project/logs -d $APP_NAME:$APP_TAG"
docker run --name $APP_NAME -v $(pwd)/logs:/trading-bot-project/logs -d $APP_NAME:$APP_TAG
