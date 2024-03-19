# Backend Traffic Police
Данный проект - это тестовое задание в RTUITLab

# Стек проекта
Общий стек проекта:
1. Go 1.22
2. Postgres 16
3. RabbitMQ
4. Docker и Docker Compose
5. Swagger

Используемые библиотеки:
1. pgx - драйвер для PostgreSQL
2. swaggo - преобразует аннотации Go в документацию Swagger
3. amqp091-go - для работы с RabbitMQ  из Go
4. migrate - для работы с миграциями БД

Я не использовал веб фреймворки, так как Go имеет сильную стандартную библиотеку и использование веб фреймворка не было необходимостью

# Описание 
## Описание базы данных
![](images/db.png)
Таблицы:
cases - основная таблица, которая хранит айди транспортного средства, замеченного в правонарушении. Айди камеры, которая засняла случай, айди правонарушения, и остальную необходимую информацию.

users - хранит данные для аутентификации и авторизации пользователей

directors, experts, cameras - виды ролей в системе. 

cameras - хранит основную информацию о камерах.

camera_types - хранит информацию о типах камер.

expert_cases - хранит информацию об оценках экспертов по каждому случаю.

rating - хранит текущую информацию о рейтинге экспертов по количествую правильно и неправильно решенных случаев.

violations - хранит информацию о правонарушениях

transports - хранит информацию о различных транспортах, который были зафиксированы в различных случаях.

persons - информация о владельце каждого транспорта.

## Описание общей архитектуры проекта
![](images/arch.png)

В проекте используется чистая архитекрута, где 3 слоя: транспортный, сервисный (бизнес логика), репозиторий. Каждый выше стоящий слой зависит от ниже стоящего через интерфейс, что дает гибкость. Также благодаря этому реализовано юнит тестирование. Для тестирования конкретного слоя, мокается интерфейс нижестоящего слоя. 

Для транспортного слоя используются модели dto, а для сервисного слоя и репозитория используются модели domain. Использование отдельных моделей для репозитория я посчитал излишним. Такое разделение моделей позволит легко изменять взаимодействие с бэкендом (например, добавить взаимодействие по gRPC или GraphQL).

Прроект состоит из 2 сервисов:
1. service - основной сервис, который занимается всей логикой приложения, принимает запросы от клиентов, обрабатывает и возвращает ответ.
2. fine_notification - сервис, который занимается отправкой уведомлений по доступным каналам свзяи (В проекте реализовано только уведомление по почте).

Эти 2 сервиса связаны через очередь сообщений RabbitMQ. Принцип работы прост: service отправляет сообщение в очередь fine_notification читает сообщение и отправляет уведомление.

Для авторизации используется Middleware, который парсит JWT токен. В этом токене зашит айди пользователя и его роль. По роли пользователя проверяется возможность доступа к ресурсу, а по айди пользователя проверяется, что эксперт подтвержден директором.

# Запуск проекта
### 1. Установить Docker
Для запуска проекта необходим Docker.

### 2. Создать конфиг файлы
 Перед запускам нужно ОБЯЗАТЕЛЬНО создать 2 конфиг файла: 
1. Конфиг файл `service_ config.yaml` в директории service (cd service). 
2. Конфиг файл `notification_config.yaml` в директории fine_notification (cd fine_notification).

Конфиг файл `service_ config.yaml` для service имеет следующую структуру:
``` yaml
serverPort: <int: Порт, на котором работает сервис>
consensus: <int: Необходимое количество проверок специалистов для оценки случая> 

passSalt: <string: Соль для хеширования паролей>
signingKey: <string: Ключ подписи JWT токенов>

rating: <Информация для рейтинга>
  reportPeriod: <duration: Время отчетного периода. Формат hms>
  minSolvedCases: <int: Минимальное количество решенных кейсов экспертом для его оценки в отчетный период>
  minExperts: <int: Минимальное количество экспретов для оценки рейтинга. Минимально - 3>
  
postgres: <Информация о БД>
  user: <string: Имя пользователя БД>
  password: <string: Пароль пользователя БД>
  host: <string: Хост БД>
  port: <int: Порт БД>
  database: <string: Наименование БД>

rabbitmq: <Информация о RabbitMQ>
  user: <string: Имя пользователя RabbitMQ>
  password: <string: Пароль пользователя RabbitMQ>
  host: <string: Хост RabbitMQ>
  port: <int: Порт RabbitMQ>

directors: <array: Массив директоров>
  - username: <string: Имя директора>
    password: <string: Пароль директора>
```

Пример для service_ config.yaml (Можно просто копипастить):
``` yaml
serverPort: 8080
consensus: 2

passSalt: "salt"
signingKey: "sign"

rating:
  reportPeriod: 8h
  minSolvedCases: 1
  minExperts: 3

postgres:
  user: "user"
  password: "user"
  host: "postgres"
  port: 5432
  database: "traffic_police_db"

rabbitmq:
  user: "guest"
  password: "guest"
  host: "rabbitmq"
  port: 5672

directors:
  - username: "director1"
    password: "director1"
  - username: "director2"
    password: "director2"
```

Конфиг файл `notification_config.yaml` для fine_notification имеет следующую структуру:
``` yaml
emailSender: <Информация об отправителе сообщений по почте>
  host: <string: Хост отправителя сообщений>
  port: <string: Порт отправителя сообщений>
  username: <string: Имя пользователя отправителя сообщений>
  password: <string: Пароль пользователя>
  subject: <string: Заголовок сообщения о правонарушении>


rabbitmq: <Информация о RabbitMQ>
  user: <string: Имя пользователя RabbitMQ>
  password: <string: Пароль пользователя RabbitMQ>
  host: <string: Хост RabbitMQ>
  port: <string: Порт RabbitMQ>
```

Пример для notification_config.yaml. Для отправителя сообщений проще всего использовать smtp сервер gmail и пароль приложения в gmail.
``` yaml
emailSender:
  host: "smtp.gmail.com"
  port: 587
  username: "emailsender@gmail.com"
  password: "secret"
  subject: "Информация о правонарушении"


rabbitmq:
  user: "guest"
  password: "guest"
  host: "rabbitmq"
  port: 5672
```

### 3. Запустить контейнеры
```
docker compose up -d
```
