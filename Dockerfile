FROM ubuntu:18.04
RUN apt update && apt-get install -y wget
RUN apt install -y gcc npm
RUN npm install -g n
RUN n stable
RUN npm install npm@latest -g
RUN wget https://golang.google.cn/dl/go1.14.10.linux-amd64.tar.gz
RUN tar -zxvf go1.14.10.linux-amd64.tar.gz -C /usr/local/
RUN rm -rf ./go1.14.10.linux-amd64.tar.gz
RUN ["/bin/bash", "-c", "ln -s /usr/local/go/bin/go /bin/go"]
RUN ["/bin/bash", "-c", "/usr/local/go/bin/go env -w GO111MODULE=on"]
RUN ["/bin/bash", "-c", "/usr/local/go/bin/go env -w GOPROXY=https://goproxy.cn,direct"]

