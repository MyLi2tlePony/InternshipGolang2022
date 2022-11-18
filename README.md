# Задание
[Микросервис для работы с балансом пользователей](docs/TASK.md)

## Запуск проекта
Клонируйте репозиторий с помощью git `git clone https://github.com/MyLi2tlePony/InternshipGolang2022.git`

Создайте и запустите докер контейнер командой `make up` или `cd deployments && docker-compose -f docker-compose.yaml build && docker-compose -f docker-compose.yaml up`

## Запуск интеграционных тестов   
Создайте и запустите докер контейнер с интеграционными тестами командой `make integration-tests` или `cd deployments && docker-compose -f docker-compose.test.yaml build && docker-compose -f docker-compose.test.yaml up`

## Запросы Postman
### Метод начисления средств на баланс - `POST`
Принимает id пользователя и сколько средств зачислить. (если пользователя нет, то он создается)

`http://0.0.0.0:3456/replenish`

Тело запроса:
`
{
    "UserID": 5,
    "Amount": 1000
}
`

### Метод резервирования средств с основного баланса на отдельном счете - `POST`
Принимает id пользователя, id услуги, id заказа, стоимость

`http://0.0.0.0:3456/reserve`

Тело запроса:
`
{
"UserID": 5,
"OrderID": 1,
"ServiceID": 2,
"Amount": 500
}
`

### Метод признания выручки - `POST`
Принимает id пользователя, id услуги, id заказа, сумму. (списывает из резерва деньги, добавляет данные в отчет для бухгалтерии)

`http://0.0.0.0:3456/reserve/confirm`

Тело запроса:
`
{
"UserID": 5,
"OrderID": 1,
"ServiceID": 2,
"Amount": 500
}
`