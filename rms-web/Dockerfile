# Базовый образ ALT Linux p10
FROM alt:p10

# Установка необходимых пакетов
RUN apt-get update && apt-get install -y \
    apache2 \
    apache2-mod_php8.2 \
    php8.2 \
    php8.2-pgsql \
    php8.2-curl \
    php8.2-libs \
    php8.2-mbstring \
    php8.2-xmlreader \
    php8.2-gd \
    php8.2-zip \
    php8.2-fpm-fcgi \
    && apt-get clean \
    && rm -rf /var/cache/apt/* /var/lib/apt/lists/*

# Копирование локальных конфигураций Apache
COPY ./httpd2/default.conf /etc/httpd2/conf/sites-available/default.conf

# Копирование локальной конфигурации PHP
#COPY ./php-config/php.ini /etc/php/8.2/cli/

# Копирование файлов веб-страниц
COPY ./www/ /var/www/html/

RUN chown -R apache:apache /var/www/html && \
    chmod -R 755 /var/www/html

# Просмотр доступных модулей PHP (опционально, для отладки)
#RUN ls -la /etc/httpd2/conf/mods-available/ | grep php

# Открытие порта 80
EXPOSE 80

# Запуск Apache в foreground режиме
CMD ["/usr/sbin/httpd2", "-D", "FOREGROUND"]