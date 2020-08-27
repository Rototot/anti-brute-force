# Анти-брутфорс

Сервис предназначен для борьбы с подбором паролей при авторизации в какой-либо системе.

Сервис вызывается перед авторизацией пользователя и может либо разрешить, либо заблокировать попытку.

Предполагается, что сервис используется только для server-server, т.е. скрыт от конечного пользователя.


## Dependencies 
* Docker
* Docker-compose 3.7+
* Make (for run Makefile)

## Install
```bash
cp .env.dist .env

make run
```


## Docs
### API doc
Add To Blacklist 
```
POST /blacklist
{
  "subnet" : "192.168.1.1/12"
}
```

Remove From Blacklist 
```
DELETE /blacklist
{
  "subnet" : "192.168.1.1/12"
}
```

Add To Whitelist 
```
POST /whitelist
{
  "subnet" : "192.168.1.1/12"
}
```

Remove From Whitelist 
```
DELETE /whitelist
{
  "subnet" : "192.168.1.1/12"
}
```

Login Attempt
```
POST /login/attempt
{
  "login" : "test_login",
  "ip"    : "192.168.1.1",
  "password" : "hash_password"
}
```

Reset Login Attempt
```
DELETE /login/attempt
{
  "login" : "test_login",
  "ip"    : "192.168.1.1"
}
```

### Cli doc
```
Usage:
  anti-brute-force [command]

Available Commands:
  cli         cli app control
  help        Help about any command
  server      Start web API server

Flags:
  -h, --help   help for anti-brute-force

```
### ТЗ

* [Техническое задание](./docs/tz)

### Project structure
В основе лежит гексогональная архитектура + принципы DDD. 

#### Слои
1. Domain - ядро бизнес-логики системы с портами-интерфейсами для реализации их в последующих словях
2. Application - слой ответачает за координацию потоков данных и содержит сценарии использования приложения 
3. Infrastrucutre - слой содержащий непосредственную реализацию портов из предыдущих слове, а также здесь происходит непосредственное взаимодействие с инфраструктурными компонентами сервиса
4. Presentation - слой содержит способы взаимодействия приложения с внешним миров (http/cli и т.д); задача слоя обрабатывать ввод и отдавать данные по запросам; данные отдаются, запрашивая домен

```
pkg

├── domain
│   ├── constants
│   ├── entities
│   ├── factories
│   ├── repositories
│   ├── services
│   └── valueobjects
├── application
│   └── usecases
│       └── mocks
├── infrastructure
│   ├── configurators
│   └── persistence
│       ├── postgres
│       │   └── repositories
│       └── redis
│           └── repositories
└── presentation
    ├── cli
    │   └── commands
    └── rest
        ├── controllers
        ├── httputils
        ├── queries
        └── routers

```
