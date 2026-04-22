

```bash
# Скачать образ из репозитория в локальное хранилище
# `docker pull <image-name>:<tag>`
$ docker pull nginx:latest
```

***

```bash
# Поиск образов в репозитории Docker Hub
$ docker search <image-name>
```

***

```bash
# Сохранить образ из локального репозитория в файл
$ docker save -o <output-file.tar> <image-name>

# Восстановить образ из файла в локальный репозиторий.
$ docker load -i <in-file.tar>
```

***

```bash
# Просмотр локального репозитория Docker
$ docker images
# или
$ docker image ls

# Запуск контейнера из локального хранилища
$ docker run <image-name>:<tag>

# `-d` - запустить контейнер в фоновом режитме
$ docker run -d <image-name>:<tag>

# `p 8080:80` с пробросом портов
$ docker run -p 8080:80 <image-name>:<tag>

# `--name <container-name> присвоить имя контейнеру
$ docker run --name <container-name> <image-name>:<tag>

# `-it` запуск в интернативном режиме с привязко к терминалу
$ docker run -it <image-name>:<tag> /bin/bash

# `--rm` автоматическое удаление контейнера после его остановки
$ docker run -rm <image-name>:<tag>
```

***

```bash
# Удалить все неиспользуемые образы (висячие или без контейнеров)
$ docker image prune -a

# Принудительное удаление всех образов, даже используемых
$ docker rmi -f $(docker images -q)

# Удаление неиспользуемых образов, контейнеров и сетей
$ docker system prune -a
```