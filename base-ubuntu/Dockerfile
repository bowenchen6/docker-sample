FROM registry.cn-hangzhou.aliyuncs.com/bowen/base-ubuntu:test 

ENV HOME /root
RUN /etc/my_init.d/00_regen_ssh_host_keys.sh
CMD ["/sbin/my_init"]

COPY conf/sources.list /etc/apt/sources.list
COPY conf/go1.7.1.linux-amd64.tar.gz $HOME

RUN tar -xzf $HOME/go1.7.1.linux-amd64.tar.gz -C /usr/local
RUN rm $HOME/go1.7.1.linux-amd64.tar.gz
RUN mkdir -p /go/src

RUN export PATH=$PATH:/usr/local/go/bin
RUN export GOPATH=/go/src