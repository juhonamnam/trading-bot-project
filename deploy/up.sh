. ./deploy/app_info.sh

mkdir -p logs

docker build --tag $APP_NAME:$APP_TAG .

docker run --name $APP_NAME -v $(pwd)/logs:/trading-bot-project/logs -d $APP_NAME:$APP_TAG
