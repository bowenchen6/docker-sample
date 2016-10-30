FROM registry.cn-hangzhou.aliyuncs.com/bowen/base-ubuntu:test 

RUN         DEBIAN_FRONTEND=noninteractive apt-get update && DEBIAN_FRONTEND=noninteractive  apt-get install -y redis-server
EXPOSE      6379
ENTRYPOINT  ["/usr/bin/redis-server"]