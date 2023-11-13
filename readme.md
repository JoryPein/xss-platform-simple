## 服务概述

此服务需要使用到80和8099两个端口，请确保这两个端口没有被占用。该服务提供了一个Web应用，通过Docker容器化部署，方便快捷。

## 安装步骤

1. 克隆仓库到本地：

   ```shell
   git clone https://github.com/your/repository.git
   ```

2. 进入项目目录：

   ```shell
   cd your/repository
   ```

3. 启动服务：

   ```shell
   docker-compose up -d
   ```

   此命令将在后台启动容器，确保端口80和8099没有被占用。

4. 完成安装后，通过以下步骤验证服务是否正常运行。

## 验证服务

向80端口所在的WEB服务发送HTTP请求：

```shell
curl http://<ip>:80?xss=asdfgg
```

通过RESTful接口查看接收到的请求：

```shell
curl http://<ip>:8099/api/cross.json
```

确保替换 `<ip>` 为你的服务器IP地址。

## 注意事项

- 请确保服务器上的80和8099端口未被其他应用占用。
- 如果需要自定义配置，请修改 `docker-compose.yml` 文件。
- 更多配置和使用说明，请参阅相关文档或代码注释。

## 接口文档

详细的接口文档可以在 [docs](docs/) 目录下找到。请查阅该文档以了解更多关于RESTful接口的信息。

## 反馈与贡献

欢迎提出问题、报告bug或贡献代码。请在GitHub上提交issue或pull request。感谢您的参与和支持！
