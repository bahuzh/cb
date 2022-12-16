//编译二进制文件
make build
bin/amd64
cp httpsrv  /xxx/ 
//编写 Dockerfile 将模块二作业编写的 httpserver 容器化
docker build --platform amd64 -t bahuzh/httpsrv:v1.1 .
docker images

//将镜像推送至 docker 官方镜像仓库
docker login
docker push bahuzh/httpsrv:v1.1

//通过 docker 命令本地启动 httpserver

docker run -d -p 80:80  bahuzh/httpsrv:v1.1

//通过 nsenter 进入容器查看 IP 配置
lsns -t net
nsenter -t 182112 -n ip a

