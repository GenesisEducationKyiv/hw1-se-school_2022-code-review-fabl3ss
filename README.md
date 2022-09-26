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

### Architecture diagram
![architecture_diagram](https://raw.github.com/GenesisEducationKyiv/hw1-se-school_2022-code-review-fabl3ss/hw5/static/architecture_diagram.png)
### Folders structure

``` bash
.
│   .dockerignore
│   .env
│   .gitignore
│   .golangci.toml
│   docker-compose.yml
│   Dockerfile
│   go.mod
│   go.sum
│   Makefile
│   README.md
│
├───.github
│       .keep
│
<<<<<<< HEAD
=======
├───.idea
│       hw1-se-school_2022-code-review-fabl3ss.iml
│       modules.xml
│       workspace.xml
│
├───.vscode
│       launch.json
│
>>>>>>> 6ba15fbcdd4ce98ad7459106697281edfc366b7b
├───bin
│       golangci-lint
│
├───credentials
│   └───gmail
│           client_secret.json
│           token.json
│
├───logs
├───src
│   ├───cmd
│   │       main.go
│   │       start_http_server.go
│   │
│   ├───config
│   │       config_integration_test.go
│   │       env_names.go
│   │       exchangers_names.go
│   │       fiber_config.go
│   │
│   ├───loggers
│   │       logger.go
│   │       zap_logger.go
│   │
│   ├───pkg
│   │   ├───application
│   │   │   │   application_contracts.go
│   │   │   │
│   │   │   ├───exchange
│   │   │   │       crypto_exchange_chain.go
│   │   │   │       crypto_exchange_ucase.go
│   │   │   │       crypto_exchange_ucase_test.go
│   │   │   │
│   │   │   ├───mailing
│   │   │   │       crypto_mailing_ucase.go
│   │   │   │       crypto_mailing_ucase_test.go
│   │   │   │
│   │   │   ├───mocks
│   │   │   │       persistence_mocks.go
│   │   │   │
│   │   │   └───subscription
│   │   │           subscription_ucase.go
│   │   │           subscription_ucase_test.go
│   │   │
│   │   ├───delivery
│   │   │   └───http
│   │   │       │   config_rate_handler.go
│   │   │       │   handlers.go
│   │   │       │   http_contracts.go
│   │   │       │   mailing_handler.go
│   │   │       │
│   │   │       ├───middleware
│   │   │       │       fiber_middleware.go
│   │   │       │       validator.go
│   │   │       │       validator_test.go
│   │   │       │
│   │   │       ├───presenters
│   │   │       │       json_presenter.go
│   │   │       │
│   │   │       ├───responses
│   │   │       │       http_responses.go
│   │   │       │
│   │   │       └───routes
│   │   │               init.go
│   │   │               routes.go
│   │   │               routes_contracts.go
│   │   │
│   │   ├───domain
│   │   │   ├───models
│   │   │   │       currency.go
│   │   │   │       email.go
│   │   │   │       rate.go
│   │   │   │
│   │   │   └───usecases
│   │   │           crypto_usecase_contract.go
│   │   │           mailing_usecase_contract.go
│   │   │           subscription_usecase_contract.go
│   │   │           usecases.go
│   │   │
│   │   ├───persistence
│   │   │   ├───crypto
│   │   │   │   │   crypto_cache.go
│   │   │   │   │   crypto_logger.go
│   │   │   │   │
│   │   │   │   ├───banners
│   │   │   │   │       crypto_bannerbear.go
│   │   │   │   │       crypto_bannerbear_integration_test.go
│   │   │   │   │
│   │   │   │   ├───charts
│   │   │   │   │       coinbase_chart_integration_test.go
│   │   │   │   │       coinbase_chart_provider.go
│   │   │   │   │
│   │   │   │   └───exchangers
│   │   │   │           coinapi_exchange_provider.go
│   │   │   │           coinapi_integration_test.go
│   │   │   │           coinbase_exchange_provider.go
│   │   │   │           coinbase_integration_test.go
│   │   │   │           crypto_exchanger_node.go
│   │   │   │           crypto_logging_exchanger.go
│   │   │   │           exchangers_test.go
│   │   │   │           nomics_exchange_provider.go
│   │   │   │           nomics_integration_test.go
│   │   │   │
│   │   │   ├───mailing
│   │   │   │       mailing_gmail.go
│   │   │   │
│   │   │   └───storage
│   │   │       ├───csv
│   │   │       │       csv_email_storage.go
│   │   │       │
│   │   │       └───redis
│   │   │               redis_cache.go
│   │   │
│   │   ├───types
│   │   │   ├───errors
│   │   │   │       errors.go
│   │   │   │
│   │   │   └───filemodes
│   │   │           filemodes.go
│   │   │
│   │   └───utils
│   │           files.go
│   │           files_test.go
│   │           http.go
│   │           slices.go
│   │           slices_test.go
    │       utils_arch_test.go
    │
    └───postman
            test_postman.json
```
