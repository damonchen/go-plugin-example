# 操作步骤

```
brew install protoc
```

```
protoc -I service service.proto --go_out=plugins=grpc:service
```

