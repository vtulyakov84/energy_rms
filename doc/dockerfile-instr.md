
### Варианты создания и запуска 
### 1. Собрать образ из Dockerfile
Собраный образ будет храниться в локальном хранилище

```bash
# Собрать образ из Dockerfile
# `-t` - <image-name:tag>, например 
# `.`  - путь к директории с Dockerfile
$ docker build -t <image-name:tag> .
```

### 3. Запуск контейнера из локального хранилища
```bash
$ docker run <image-name>:<tag>
```

### Выгрузка Docker-образа в tar-файл
```bash
# Сохранение образа в tar-архив
docker save -o myapp_image.tar myapp:latest

# Сжатие для экономии места (опционально)
docker save myapp:latest | gzip > myapp_image.tar.gz

# Создание архива с несколькими образами
docker save -o multi_images.tar nginx:alpine redis:alpine postgres:13
```

### Проверка целостности tar-файла
```bash
# Проверка, что файл существует и не пустой
if [ -f myapp_image.tar ] && [ -s myapp_image.tar ]; then
    echo "Файл существует и не пуст"
fi

# Проверка целостности tar-архива
tar -tf myapp_image.tar > /dev/null 2>&1
if [ $? -eq 0 ]; then
    echo "Tar-архив не поврежден"
else
    echo "Ошибка: tar-архив поврежден"
fi
```

### Расширенная проерка целостности tar-файла с контрольной суммой
```bash
# Создание контрольной суммы при выгрузке
docker save myapp:latest -o myapp_image.tar
sha256sum myapp_image.tar > myapp_image.tar.sha256

# Проверка контрольной суммы
sha256sum -c myapp_image.tar.sha256
if [ $? -eq 0 ]; then
    echo "Контрольная сумма совпадает"
else
    echo "Ошибка: файл изменен или поврежден"
fi
```

### Полная проверка структуры Docker-образа
```bash
#!/bin/bash
# check_docker_tar.sh

TAR_FILE="$1"

check_docker_tar() {
    # Проверка существования файла
    if [ ! -f "$TAR_FILE" ]; then
        echo "❌ Файл $TAR_FILE не найден"
        return 1
    fi
    
    # Проверка tar-структуры
    echo "📦 Проверка целостности tar-архива..."
    if ! tar -tf "$TAR_FILE" &>/dev/null; then
        echo "❌ Tar-архив поврежден или имеет неверный формат"
        return 1
    fi
    
    # Проверка наличия обязательных файлов Docker-образа
    echo "🔍 Проверка структуры Docker-образа..."
    local required_files=("manifest.json" "repositories")
    
    for file in "${required_files[@]}"; do
        if ! tar -tf "$TAR_FILE" | grep -q "$file"; then
            echo "❌ Отсутствует обязательный файл: $file"
            return 1
        fi
    done
    
    # Проверка размера (должен быть больше 1MB для реальных образов)
    local size=$(stat -f%z "$TAR_FILE" 2>/dev/null || stat -c%s "$TAR_FILE" 2>/dev/null)
    if [ "$size" -lt 1048576 ]; then
        echo "⚠️  Образ подозрительно мал: $(numfmt --to=iec $size)"
    else
        echo "✅ Размер образа: $(numfmt --to=iec $size)"
    fi
    
    echo "✅ Проверка пройдена: архив корректен"
    return 0
}

check_docker_tar "$1"
```

### Загрузка docker-образа из tar-файла
```bash
# Базовая загрузка
docker load -i myapp_image.tar

# Загрузка из сжатого файла
gunzip -c myapp_image.tar.gz | docker load

# Загрузка с проверкой перед импортом
if tar -tf myapp_image.tar &>/dev/null; then
    docker load -i myapp_image.tar
    echo "✅ Образ успешно загружен"
else
    echo "❌ Невозможно загрузить - архив поврежден"
fi
```