# Используем базовый образ с nginx для сервировки статических файлов
FROM nginx:alpine

# Копируем HTML файл в директорию для сервировки
COPY index.html /usr/share/nginx/html/index.html
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

# Открываем порт 80
EXPOSE 80
