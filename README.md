# Welcome 
## [Documentation](https://maxym.gitbook.io/crypto-mailer/)
``` ENG ```
### **Hi!** This project was written to send the current rate of cryptocurrencies to email
* Backend was written entirely in **Go**, and thoroughly ent-to-end tested in **Postman**
* By convention, this project does not use a database, so email data is stored in a csv file
* The data in this file is stored in order, which allows you to use **binary search** to check email existance
* The project is written to work out of the box, so all the credentials are public, and some of them are **trial** :( \
Therefore, if the API stopped working for some reason, first try to set up the [config](https://maxym.gitbook.io/crypto-mailer/reference/setup-config)

\
``` УКР ```
### **Привіт!** Це проєкт для отримання поточного курсу криптовалют і розсилання його по email
* Бекенд повністю написаний на **Go**, і протестований end-to-end у **Postman**
* За умовою завдання, не потрібно було підключати бд, тому емейли зберігаються у csv файлі
* Дані в цьому csv файлі зберігаються впорядковано, тому використовується **бінарний пошук** для перевірки чи емейл вже записано
* Проєкт написано так, щоб він працював "з коробки", тому деякі облікові дані з **пробним періодом** :( \
Якщо API перестав працювати як належне, можливо варто переналаштувати [конфіг](https://maxym.gitbook.io/crypto-mailer/reference/setup-config)


## Deployment
``` Docker ```
```bash 
  docker-compose up
```

## Routes

```bash 
  localhost:8000/rate       --> Get current cryptocurrency rate
  localhost:8000/subscribe  --> Subscribe email to the newsletter
  localhost:8080/sendEmails --> Send the newsletter to all subscribed emails
```

## Linter
[golangci-lint](https://github.com/golangci/golangci-lint) was used as a linter
```bash
  make lint      --> to execute all configured linters
  make lint-fast --> to execute only fast linters
```

## Tests
``` Module and Integration ```
```bash
  make test      --> to execute tests once
  make test100   --> to execute tests 100 times
  make cover     --> to see tests coverage in html
```

``` End-to-end ```
```bash
  1. Import tests/postman/test_postman.json into postman
  2. Run collection
```

## Project structure
``` bash
.
├── Dockerfile
├── Makefile
├── README.md
├── bin
│   └── golangci-lint
├── credentials
│   └── gmail
│       ├── client_secret.json
│       └── token.json
├── docker-compose.yml
├── go.mod
├── go.sum
├── src
│   ├── cmd
│   │   └── main.go
│   ├── config
│   │   ├── config_integration_test.go
│   │   ├── env_names.go
│   │   └── fiber_config.go
│   ├── pkg
│   │   ├── delivery
│   │   │   └── http
│   │   │       ├── contracts.go
│   │   │       ├── mailing_handler.go
│   │   │       ├── middleware
│   │   │       │   └── fiber_middleware.go
│   │   │       └── responses
│   │   │           └── http_responses.go
│   │   ├── domain
│   │   │   ├── crypto.go
│   │   │   ├── emails.go
│   │   │   └── mocks
│   │   │       └── mockRepositories.go
│   │   ├── repository
│   │   │   ├── crypto
│   │   │   │   ├── banners
│   │   │   │   │   ├── crypto_bannerbear.go
│   │   │   │   │   └── crypto_bannerbear_integration_test.go
│   │   │   │   └── exchangers
│   │   │   │       ├── crypto_coinbase.go
│   │   │   │       └── crypto_coinbase_integration_test.go
│   │   │   ├── mailing
│   │   │   │   └── mailing_gmail.go
│   │   │   └── storage
│   │   │       └── csv_email_storage.go
│   │   ├── routes
│   │   │   ├── init.go
│   │   │   └── routes.go
│   │   ├── types
│   │   │   ├── errors
│   │   │   │   └── errors.go
│   │   │   └── filemodes
│   │   │       └── filemodes.go
│   │   ├── usecase
│   │   │   ├── contracts.go
│   │   │   ├── crypto_exchange_ucase.go
│   │   │   ├── crypto_exchange_ucase_test.go
│   │   │   ├── crypto_mailing_ucase.go
│   │   │   ├── crypto_mailing_ucase_test.go
│   │   │   ├── subscription_ucase.go
│   │   │   └── subscription_ucase_test.go
│   │   └── utils
│   │       ├── files.go
│   │       ├── files_test.go
│   │       ├── http.go
│   │       ├── slices.go
│   │       ├── slices_test.go
│   │       ├── start-server.go
│   │       ├── validator.go
│   │       └── validator_test.go
│   └── platform
│       ├── csv
│       │   ├── data.csv
│       │   └── test.csv
│       ├── gmail_api
│       │   └── gmail_api.go
│       └── storage
├── static
│   └── crypto-message.html
└── tests
    └── postman
        └── test_postman.json
```
