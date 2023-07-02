# Mongo Example

# Quick Start

## Prerequisite
运行一个mongodb 实例

```shell
docker run  --name mongo -p 27017:27017 --restart always \
-e MONGO_INITDB_ROOT_USERNAME=admin \
-e MONGO_INITDB_ROOT_PASSWORD=admin \
-d mongo
```