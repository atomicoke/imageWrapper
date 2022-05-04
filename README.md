# imageWrapper

用于 图片 resize 的一个小服务

## Features
1. 图片resize
2. 基于内存的缓存

## 使用方法

step 1

```shell
git clone https://github.com/atomicoke/imageWrapper.git
```

step 2

```shell
cd imageWrapper
```

step 3

```shell
go run main.go
```

step 4

有两种格式使用

- /${width}x${height}?url=${url}
- /${resize}?url=${url}

1. 一个可以访问的图片路径

> e.g. https://pic2.zhimg.com/v2-471f8aa91487ac3c073ab5c5b42361ca_400x224.jpg?source=7e7ef6e2

2. 生成一个宽为100的图片(保留长宽比)

> /100?url=https://pic2.zhimg.com/v2-471f8aa91487ac3c073ab5c5b42361ca_400x224.jpg?source=7e7ef6e2

3. 生成一个200x400的图片

> /200x400?url=https://pic2.zhimg.com/v2-471f8aa91487ac3c073ab5c5b42361ca_400x224.jpg?source=7e7ef6e2