# readme

此服务需要使用到80和8099两个端口，请确保这两个端口没有被占用

```shell
docker-compose up -d
```

安装完成后，向80端口所在的WEB服务发送HTTP请求，如下

```shell
curl http://<ip:80>?xss=asdfgg
```

通过RESTful接口，可以查看接收到的请求

```shell
curl http://<ip:8099>/api/cross.json
```

