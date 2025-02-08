# Используем официальный образ PostgreSQL
FROM postgres:15

# Устанавливаем переменные окружения для базы данных
ENV DB_USER=myuser
ENV DB_PASSWORD=mypassword
ENV DB_NAME=containers_db

# Копируем начальные скрипты для базы данных (если есть)
# Например, создаем схему или таблицы при старте
COPY ./init.sql /docker-entrypoint-initdb.d/

# Открываем порт для работы с базой данных
EXPOSE 5432