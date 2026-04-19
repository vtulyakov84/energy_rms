FROM php:8.2-apache

# Копируем конфиги Apache
COPY apache-config/000-default.conf /etc/apache2/sites-available/000-default.conf

# Копируем конфиги PHP
COPY php-config/php.ini /usr/local/etc/php/php.ini

# Копируем файлы сайта
COPY www/ /var/www/html/

# Включаем mod_rewrite (опционально)
RUN a2enmod rewrite

# Даем права на запись (если нужно)
RUN chown -R www-data:www-data /var/www/html