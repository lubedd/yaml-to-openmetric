![Github Repository Size](https://img.shields.io/github/languages/code-size/lubedd/yaml-to-openmetric)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/lubedd/yaml-to-openmetric)
![Lines of code](https://img.shields.io/tokei/lines/github/lubedd/yaml-to-openmetric)
![Go Report](https://img.shields.io/badge/go%20report-A%2B-brightgreen)
![Repository Top Language](https://img.shields.io/github/languages/top/lubedd/yaml-to-openmetric)
![License](https://img.shields.io/conda/l/conda-forge/setuptools)
![Github Open Issues](https://img.shields.io/github/issues/lubedd/yaml-to-openmetric)
![GitHub last commit](https://img.shields.io/github/last-commit/lubedd/yaml-to-openmetric)
![GitHub contributors](https://img.shields.io/github/contributors/lubedd/yaml-to-openmetric)
![Simply the best ;)](https://img.shields.io/badge/simply-the%20best%20%3B%29-orange)

<img align="right" width="50%" src="./images/gopher.png">

# Converting YAML to OpenMetrics service

## Описание задания

Создайте микросервис, который конвертирует данные из YAML файла в формат OpenMetrics и возвращает их при обращении к
endpoint-у http://localhost:8080/metrics

YAML – это входные данные, они не меняются. Файл лежит в директории
проекта. Пример YAML-данных:
```
currencies:
- name: usd
  value: 70
- name: eur
  value: 80
```

## Примечания к решению

- :trident: Чистая архитектура
- :book: Стандартный макет проекта Go (ну более-менее :blush:)
- :wrench: Настройка адреса сервера и пути к yaml файлу происходит через ./configs/config.json
- :truck: Маршрутизация сделана без фреймворков для максимального быстродействия системы
## Как запустить?
```
cd ./cmd/openmetric
go run main.go
```
или
```
docker-compose up
```