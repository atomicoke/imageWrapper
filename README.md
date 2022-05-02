# imageWrapper

用于 图片 resize 的一个小服务

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

1. 准备一个可以访问的图像url

> e.g. https://i0.hdslb.com/bfs/archive/08e42b4078dda8e8ee3a867f61e39317550cd600.jpg@412w_232h_1c.webp

2.用浏览器访问
> http://localhost:8888/100?url=https://i0.hdslb.com/bfs/archive/08e42b4078dda8e8ee3a867f61e39317550cd600.jpg@412w_232h_1c.webp