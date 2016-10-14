FROM registry.cn-hangzhou.aliyuncs.com/bowen/base-ubuntu:test 

ENV HOME /root
# Docker Build Arguments
ARG RESTY_VERSION="1.11.2.1"
ARG RESTY_LUAROCKS_VERSION="2.4.1"
ARG RESTY_OPENSSL_VERSION="1.0.2j"
ARG RESTY_PCRE_VERSION="8.39"
ARG RESTY_J="1"
ARG RESTY_CONFIG_OPTIONS="\
    --with-file-aio \
    --with-http_addition_module \
    --with-http_auth_request_module \
    --with-http_dav_module \
    --with-http_flv_module \
    --with-http_geoip_module=dynamic \
    --with-http_gunzip_module \
    --with-http_gzip_static_module \
    --with-http_image_filter_module=dynamic \
    --with-http_mp4_module \
    --with-http_random_index_module \
    --with-http_realip_module \
    --with-http_secure_link_module \
    --with-http_slice_module \
    --with-http_ssl_module \
    --with-http_stub_status_module \
    --with-http_sub_module \
    --with-http_v2_module \
    --with-http_xslt_module=dynamic \
    --with-ipv6 \
    --with-mail \
    --with-mail_ssl_module \
    --with-md5-asm \
    --with-pcre-jit \
    --with-sha1-asm \
    --with-stream \
    --with-stream_ssl_module \
    --with-threads \
    "

# These are not intended to be user-specified
ARG _RESTY_CONFIG_DEPS="--with-openssl=$HOME/openssl-${RESTY_OPENSSL_VERSION} --with-pcre=$HOME/pcre-${RESTY_PCRE_VERSION}"

RUN /etc/my_init.d/00_regen_ssh_host_keys.sh
COPY conf/sources.list /etc/apt/sources.list
COPY conf/nginx_signing.key $HOME
COPY conf/openssl-${RESTY_OPENSSL_VERSION}.tar.gz $HOME
COPY conf/pcre-${RESTY_PCRE_VERSION}.tar.gz $HOME
COPY conf/openresty-${RESTY_VERSION}.tar.gz $HOME
COPY conf/luarocks-${RESTY_LUAROCKS_VERSION}.tar.gz $HOME
RUN DEBIAN_FRONTEND=noninteractive apt-get update && apt-key add $HOME/nginx_signing.key
RUN echo "deb http://nginx.org/packages/ubuntu/ xenial nginx \
          deb-src http://nginx.org/packages/ubuntu/ xenial nginx" | tee -a /etc/apt/sources.list

# 1) Install apt dependencies
# 2) Download and untar OpenSSL, PCRE, and OpenResty
# 3) Build OpenResty
# 4) Cleanup

RUN \
    DEBIAN_FRONTEND=noninteractive apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
        nginx \
        build-essential \
        ca-certificates \
        curl \
        libgd-dev \
        libgeoip-dev \
        libncurses5-dev \
        libperl-dev \
        libreadline-dev \
        libxslt1-dev \
        make \
        perl \
        unzip \
        zlib1g-dev \
    && cd $HOME \
    && tar xzf openssl-${RESTY_OPENSSL_VERSION}.tar.gz \
    && tar xzf pcre-${RESTY_PCRE_VERSION}.tar.gz \
    && tar xzf openresty-${RESTY_VERSION}.tar.gz \
    && cd $HOME/openresty-${RESTY_VERSION} \
    && ./configure -j${RESTY_J} ${_RESTY_CONFIG_DEPS} ${RESTY_CONFIG_OPTIONS} \
    && make -j${RESTY_J} \
    && make -j${RESTY_J} install \
    && cd $HOME \
    && rm -rf \
        nginx_signing.key \
        openssl-${RESTY_OPENSSL_VERSION} \
        openssl-${RESTY_OPENSSL_VERSION}.tar.gz \
        openresty-${RESTY_VERSION}.tar.gz \
        openresty-${RESTY_VERSION} \
        pcre-${RESTY_PCRE_VERSION}.tar.gz \
        pcre-${RESTY_PCRE_VERSION} \
    && tar xzf luarocks-${RESTY_LUAROCKS_VERSION}.tar.gz \
    && cd $HOME/luarocks-${RESTY_LUAROCKS_VERSION} \
    && ./configure \
        --prefix=/usr/local/openresty/luajit \
        --with-lua=/usr/local/openresty/luajit \
        --lua-suffix=jit-2.1.0-beta2 \
        --with-lua-include=/usr/local/openresty/luajit/include/luajit-2.1 \
    && make build \
    && make install \
    && cd $HOME \
    && rm -rf luarocks-${RESTY_LUAROCKS_VERSION} luarocks-${RESTY_LUAROCKS_VERSION}.tar.gz \
    && DEBIAN_FRONTEND=noninteractive apt-get autoremove -y \
    && ln -sf /dev/stdout /usr/local/openresty/nginx/logs/access.log \
    && ln -sf /dev/stderr /usr/local/openresty/nginx/logs/error.log

ENTRYPOINT ["/usr/local/openresty/bin/openresty", "-g", "daemon off;"]