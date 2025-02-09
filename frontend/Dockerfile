# Используем официальный Node.js образ
FROM node:16

# Рабочая директория
WORKDIR /app

# Копируем package.json и package-lock.json (если он есть)
COPY package*.json ./

# Устанавливаем зависимости
RUN npm install

# Копируем весь проект
COPY . .

# Собираем проект
RUN npm run build

# Указываем порт, на котором приложение будет доступно
EXPOSE 3000

# Запускаем приложение
CMD ["npm", "start"]
