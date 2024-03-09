# Конфиг файл config.yaml
Структура:
``` yaml
serverPort: int
consensus: 3
reportPeriod: duration
minSolvedCases: int

passSalt: string
signingKey: string

directors: array of
  - username: string
    password: string
```

Пример:
``` yaml
serverPort: 8080
consensus: 3
reportPeriod: 3m
minSolvedCases: 2

passSalt: "salt"
signingKey: "sign"

directors:
  - username: "justnik1"
    password: "justnik1"
  - username: "justnik2"
    password: "justnik2"
```