---
layout: home
title: Описание
nav_order: 1
permalink: /ru/
---

# OpenAPI генератор для Go
{: .fs-9 }
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/goppuchino/oag)
![GitHub contributors](https://img.shields.io/github/contributors/goppuchino/oag)
![GitHub last commit](https://img.shields.io/github/last-commit/goppuchino/oag)
![GitHub License](https://img.shields.io/github/license/goppuchino/oag)

<img align="right" width="180px" src="https://raw.githubusercontent.com/goppuchino/oag/master/assets/oag.png">

Быстрый и легкий инструмент для **генерации чистых спецификаций OpenAPI 3, соответствующих стандартам** напрямую из
вашего Go кода. Идеально подходит для:

- Автоматического документирования REST API 🏗️
- Устранения ручной поддержки спецификаций ✨
- Обеспечения совместимости со Swagger UI, Postman и другими инструментами 🔌

Возможности:  
✅ Автоматическое сопоставление структур и схем  
✅ Встроенная валидация и проверка  
✅ Поддержка пользовательских шаблонов

```go  
// Just annotate & generate!
// @method get
// @path /users/{id}
// @summary Get user
// @description Get user information
// @param id path isRequired User ID for get user information
func GetUser(w http.ResponseWriter, r *http.Request) { ... }  
```  

**Начало работы:**
```shell
$ go install github.com/goppuchino/oag
$ cd <ваш проект>
$ oag
```

*Потому что никто не любит писать YAML вручную.* 🐹