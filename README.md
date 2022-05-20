# favorites service
## Сервис для хранения избранных аптек и товаро пользователей.

From docs [grpc-gateway install](https://github.com/grpc-ecosystem/grpc-gateway/blob/master/README.md#installation):

>Run go mod tidy to resolve the versions. <br>
Install by running
```bash
$ go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
>This will place four binaries in your $GOBIN;

- protoc-gen-grpc-gateway
- protoc-gen-openapiv2
- protoc-gen-go
- protoc-gen-go-grpc

>Make sure that your $GOBIN is in your $PATH.

tools добавлен.

Параметры берутся из файла configs/config.local.yml по умолчанию. Папку конфига и сам конфиг можно изменить с помощью флагов "config_file" и "config_dir". Пример можно посмотреть в makefile, run.

В папке docker лежит [docker-compose](./docker/docker-compose.yml), с помощью которого можно поднять postgres. С помощью [goose](https://github.com/pressly/goose) можно залить [миграцию](./database/schema).

Для работы с сервисом есть [makefile](Makefile).

Для тестирования добавлен [postman_collection](./postman/favorites.postman_collection.json), который эмулирует запрос фронта.

## Description

Сервис имеет два эндпоинта:
- /api/favorite/v1/pharmacy - для аптек
- /api/favorite/v1/product - для товар

Они обрабатывают только метод POST. Действия задаются с помощью query-параметров:
- ACTION = "ADD" / "DELETE" (на данный момент сделано добавление и удаление)
- ID = id аптеки или товара

Id пользователя передается в JWT-токене в заголовке. Есть 2 возможности передать его:
- Bearer Token
- Cookie "APP.token"
- Metadata "grpcgateway-cookie: APP.token=[token]" для grpc клиента