# re_ark


## proto to go source

```sh
protoc protos/*.proto --go_out=plugins=grpc:. --go_opt=module=github.com/pistatium/re_ark -I./protos/
```
