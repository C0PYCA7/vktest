# Проект VkPinger
***
## Описание
VkPinger - это проект, который предоставляет возможность следить за состоянием контейнеров 
и добавлять новые. Проект состоит из нескольких компонентов: база данных PostgreSQL, Backend сервер
, сервис Pinger, Frontend, Kafka

***

## В разработке использовались следующие технологии: 
- **Go**
- **React**
- **Docker**
- **PostgreSQL**
- **Kafka**
- **Nginx**
- **Cron**

***
## Структура проекта:
- **Backend** - серверная часть приложения
- **Pinger** - сервис для проверки состояния контейнеров
- **DB** - база данных
- **Frontend** - фронтенд-приложение


### Backend
Серверное приложение, которое обрабатывает входящие запросы по двум эндпоинтам 
- /
- /create

Для обработки / сервер обращается к базе данных и забирает информацию обо всех контейнерах
которые были добавлены
Для обработки /create сервер получает необходимые json данные, обрабатывает их
и добавляет в базу данных

При запуске приложение создает consumerGroup, подключается к топику, слушает входящие
сообщения и обрабатывает их

### Pinger
Приложение, которое получает данные о контейнерах добавленных в базу данных, пингует их по ip
и отдает данные в Backend

При запуске приложение создает топик если его еще не существует, создается asyncProducer
который отправляет данные в топик после обработки

***
## Запуск приложения
- Клонировать репозиторий <pre>git clone https://github.com/C0PYCA7/vktest/tree/main</pre>
- Сбилдить docker-compose файл <pre>docker-compose build</pre>
- Запустить docker-compose файл <pre>docker-compose up</pre>
- Перейти на localhost:80
- Ввести ip необходимого контейнера из одинаковой сети с проектом
- Перезагрузить страницу и через минуту вы увидите состояние контейнера