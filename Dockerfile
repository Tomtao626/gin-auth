FROM golang:1.18.5 as builder

WORKDIR ./gin-auth
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="tp320670258@gmail.com"

WORKDIR ./gin-blog

COPY --from=0 ./gin-auth/server ./
COPY --from=0 ./gin-auth/conf/app.ini ./

EXPOSE 8888
ENTRYPOINT ./server


#作用
#golang:latest 镜像为基础镜像，将工作目录设置为 $GOPATH/src/go-gin-example，并将当前上下文目录的内容复制到 $GOPATH/src/go-gin-example 中
#
#在进行 go build 编译完毕后，将容器启动程序设置为 ./go-gin-example，也就是我们所编译的可执行文件
#
#注意 go-gin-example 在 docker 容器里编译，并没有在宿主机现场编译
#
#说明
#Dockerfile 文件是用于定义 Docker 镜像生成流程的配置文件，文件内容是一条条指令，每一条指令构建一层，因此每一条指令的内容，就是描述该层应当如何构建；这些指令应用于基础镜像并最终创建一个新的镜像
#
#你可以认为用于快速创建自定义的 Docker 镜像
#
#1、 FROM
#
#指定基础镜像（必须有的指令，并且必须是第一条指令）
#
#2、 WORKDIR
#
#格式为 WORKDIR <工作目录路径>
#
#使用 WORKDIR 指令可以来指定工作目录（或者称为当前目录），以后各层的当前目录就被改为指定的目录，如果目录不存在，WORKDIR 会帮你建立目录
#
#3、COPY
#
#格式：
#
#COPY <源路径>... <目标路径>
#COPY ["<源路径1>",... "<目标路径>"]
#COPY 指令将从构建上下文目录中 <源路径> 的文件/目录复制到新的一层的镜像内的 <目标路径> 位置
#
#4、RUN
#
#用于执行命令行命令
#
#格式：RUN <命令>
#
#5、EXPOSE
#
#格式为 EXPOSE <端口1> [<端口2>...]
#
#EXPOSE 指令是声明运行时容器提供服务端口，这只是一个声明，在运行时并不会因为这个声明应用就会开启这个端口的服务
#
#在 Dockerfile 中写入这样的声明有两个好处
#
#帮助镜像使用者理解这个镜像服务的守护端口，以方便配置映射
#运行时使用随机端口映射时，也就是 docker run -P 时，会自动随机映射 EXPOSE 的端口
#6、ENTRYPOINT
#
#ENTRYPOINT 的格式和 RUN 指令格式一样，分为两种格式
#
#exec 格式：
#<ENTRYPOINT> "<CMD>"
#shell 格式：
#ENTRYPOINT [ "curl", "-s", "http://ip.cn" ]
#ENTRYPOINT 指令是指定容器启动程序及参数