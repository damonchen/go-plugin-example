# echo example

实现了基本的 grpc 下的处理


## 运行

```
go build
cd plugin
go build
cd ..

export ECHO_PLUGIN=./plugin/plugin
./echo-example 'test echo server'
```

