FROM registry.cn-hangzhou.aliyuncs.com/acs/ubuntu

RUN apt-get -y update \
    && apt-get install -y software-properties-common \
    && add-apt-repository ppa:ondrej/php -y \
    && apt-get -y update \
    && apt-get install php5.6 -y \
    && apt-get install -my \
    curl \
    wget \
    git \
    php5.6-curl \
    php5.6-cli \
    php5.6-cgi \
    php5.6-fpm \
    php5.6-gd \
    php5.6-memcached \
    php5.6-mcrypt \
    php5.6-mbstring

#install redis extension
RUN cd ~/ \
    && git clone https://github.com/phpredis/phpredis.git \
    && cd phpredis \
    && phpize \
    && ./configure \
    && make \
    && make install

RUN cd /etc/php/5.6/mods-available/ \
    && echo "extension=redis.so" >> redis.ini \
    && ln -s /etc/php/5.6/mods-available/redis.ini /etc/php/5.6/fpm/conf.d/20-redis.ini