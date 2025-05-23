# Интернет-магазин футбольной атрибутики
## Кучер Артем ЭФМО-02-24

## Запуск проекта
### 1. Клонирование репозитория
```
git clone https://github.com/KucherArtyom/Online-store-of-football-products.git
cd Online-store-of-football-products
```
### 2. Настройка базы данных PostgreSQL
```
createdb footballstore
psql footballstore < footballstore.sql
```
### 3. Настройка бэкенда (Go)
```
cd backend
```
Создайте файл .env:
```
APP_ENV=development
SERVER_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=ваш_пароль
DB_NAME=footballstore
```

```
go mod download
go run main.go
```
### 4. Настройка фронтенда (Vue.js)
```
cd frontend
npm install
npm run dev
```

## Схема базы данных
![Схема базы данных](https://github.com/user-attachments/assets/148d24bc-2810-45ba-ab24-5553367a4bfe)

## Диаграмма последовательности "Процесс оформления заказа"
![Диаграмма последовательности Процесс оформления заказа](https://github.com/user-attachments/assets/43ed3674-76c7-42f7-ab69-e1ed7c8f04b8)

## Главная страница сайта
![Главная страница](https://github.com/user-attachments/assets/09facf58-cd42-4c8d-bfa1-3b94cc3fc11f)

## Страница регистрации
![Страница регистрации](https://github.com/user-attachments/assets/f7a91569-105b-43f2-adf0-76da09e8cdd9)

## Страница авторизации
![Страница авторизации](https://github.com/user-attachments/assets/a950d8b8-c760-48bc-acde-a6fb6a301cfc)

## Страница авторизованного пользователя
![Страница авторизованного пользователя](https://github.com/user-attachments/assets/7f29a81b-9717-463c-a02a-b79f4bd74a27)

## Страница с футболками
![Страница с футболками](https://github.com/user-attachments/assets/ca459285-c311-4afe-8d0b-40e5fdc0eed8)

## Страница с шарфами
![Страница с шарфами](https://github.com/user-attachments/assets/613e8ee0-9bf7-496b-ab3b-2bf0ac01e4da)

## Страница с мячами
![Страница с мячами](https://github.com/user-attachments/assets/58e224ea-7c0d-4441-8277-25980c4ccc8e)

## Страница с бутсами
![Страница с бутсами](https://github.com/user-attachments/assets/0d8d8aed-c309-48ff-85dc-9d33d63af8be)

## Избранное
![Избранное](https://github.com/user-attachments/assets/929339cf-d467-4df9-8b6e-65917a555a28)

## Корзина
![Корзина](https://github.com/user-attachments/assets/d8cd7970-75e4-4cb1-8082-1beab10f7e3d)

## Страница оформления заказа
![Страница оформления заказа](https://github.com/user-attachments/assets/30689959-1adf-4899-b760-6b8c9b016a50)
