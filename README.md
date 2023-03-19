## Go template from (hvt)

## Run image order api
```dockerfile
docker run -d --name fd-order --network my-main-net -e GINPORT=3005 -e MYSQL_GORM_DB_TYPE="mysql" -e MYSQL_GORM_DB_URI="root:ead8686ba57479778a76e@tcp(mysql:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local" -e GRPC_AUTH_CLIENT_URL="fd-auth-user:50051" -e NATS_URL="nats://nats:4222" -e VIRTUAL_HOST="local.200lab.io" -e VIRTUAL_PORT=3005 -e VIRTUAL_PATH="/v1/orders"  --expose 3005 -p 3005:3005 food-delivery-order-service:1.0
```
## Networks
```
docker network connect my-main-net fd-order
docker network connect order-net fd-order
docker network connect order-net nats
docker network connect order-net redis
```

## Run image order pubsub create update order
```dockerfile
docker run -d --name fd-order-create-detail-tracking --network order-net -e MYSQL_GORM_DB_TYPE="mysql" -e MYSQL_GORM_DB_URI="root:ead8686ba57479778a76e@tcp(mysql:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local" -e NATS_URL="nats://nats:4222" food-delivery-order-service /app/app create-order-detail-tracking
```