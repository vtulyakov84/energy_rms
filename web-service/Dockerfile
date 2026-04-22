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

# Создание директорий для конфигураций и веб-файлов
RUN mkdir -p /var/www/html /etc/apache2/sites-available /etc/php8

# Копирование локальных конфигураций Apache
COPY ./apache-config/000-default.conf /etc/apache2/sites-available/000-default.conf

# Копирование локальной конфигурации PHP
COPY ./php-config/php.ini /etc/php8/

# Копирование файлов веб-страниц
COPY ./www/ /var/www/html/

# Настройка прав доступа
RUN chown -R www-data:www-data /var/www/html && \
    chmod -R 755 /var/www/html

# Включение необходимых модулей Apache
RUN a2enmod php8 && \
    a2enmod rewrite && \
    a2enmod headers && \
    a2enmod ssl

# Настройка PHP для работы с PostgreSQL
RUN echo "extension=pgsql.so" > /etc/php8/conf.d/20-pgsql.ini && \
    echo "extension=pdo_pgsql.so" > /etc/php8/conf.d/21-pdo_pgsql.ini

# Открытие порта 80
EXPOSE 80

# Запуск Apache в foreground режиме
CMD ["/usr/sbin/apache2", "-D", "FOREGROUND"]