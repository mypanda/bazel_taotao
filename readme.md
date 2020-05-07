

https://segmentfault.com/a/1190000020301131

https://goproxy.io/zh/



```
1. 根据go.mod下载包
go mod download

2. 第二种方式
go get packagename@v1.2.3

3. 
go run、go build 也会自动下载依赖
```

```
将依赖包下载到vendor目录
go mod vendor
```

```
bazel run //:gazelle update-repos -from_file=go.mod
bazel run //:gazelle update-repos github.com/json-iterator/go@v1.1.7


bazel run //:gazelle
bazel run //:taotao
bazel clean
```


```

gazell的prefix和项目的module名称一致，不然找不到对应的依赖
例如：用taotao或者用github.com/mypanda/taotao都可

gazell 是收集依赖
BUILD.bazel 可以只写load gazell，不定义go_binary,执行gazell会生成go_binary
```
```BUILD.bazel
多文件
go_binary(
    name = "main",
    srcs = [
        "calc.go",
        "main.go",
    ],
    visibility = ["//visibility:public"],
)

多模块，模块下面也需要定义BUILD.bazel
go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "taotao",
    visibility = ["//visibility:private"],
    deps = [
        "//handler:go_default_library",
        "@com_github_json_iterator_go//:go_default_library",
    ],
)
handle模块
load("@io_bazel_rules_go//go:def.bzl", "go_library")
go_library(
    name = "go_default_library",
    srcs = ["handler.go"],
    importpath = "taotao/handler",
    visibility = ["//visibility:public"],
)
```
```
WORKSPACE是固定形式
```


```
项目中有src/github.com/mypanda/xxx
需要把BUILD.bazel WORKSPACE放到src文件夹中
可以在项目根目录添加Makefile
```
```
怎么在WORSPACE中添加新的包
```