# re_arc


## proto to go source

```sh
protoc protos/*.proto --go_out=plugins=grpc:. --go_opt=module=github.com/pistatium/re_ark -I./protos/
```

## Build Services

__user_service__

```sh
docker build -f user_service/Dockerfile -t user_service .
```
