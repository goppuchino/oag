# OpenAPI генератор для Go
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/goppuchino/oag)
![GitHub contributors](https://img.shields.io/github/contributors/goppuchino/oag)
![GitHub last commit](https://img.shields.io/github/last-commit/goppuchino/oag)
![GitHub License](https://img.shields.io/github/license/goppuchino/oag)


🌏 [English](README.md) | [Russian](README-ru.md)

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

## Описание

### main.go

Основное описание и общие настройки спецификации OpenAPI читаются из аннотаций основного файла проекта.
Перечисленные далее объекты описываются в аннотациях к функции `main()`.

#### [OpenAPI Object][4]

| Имя поля               | Аннотация             | Описание                                                                                   | Пример                                                                                                              |
|------------------------|-----------------------|--------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------|
| [openapi][4]           |                       | Выставляется стандартное значение `3.1.0`                                                  |                                                                                                                     |
| [info][1]              |                       | Применяются аннотации описанные в Info Object                                              |                                                                                                                     |
| [jsonSchemaDialect][5] | @jsonSchemaDialect    | Применяемый диалект с схеме данных (лучше не указывать)                                    | `// @jsonSchemaDialect https://spec.openapis.org/oas/3.1/dialect/base`                                              |
| [servers][6]           | @server.N.url         | Адрес сервера, где N — порядковый номер, начинается с 0                                    | `// @server.0.url https://development.gigantic-server.com/v1`                                                       |
|                        | @server.N.description | Описание сервера, где N — порядковый номер, начинается с 0                                 | `// @server.0.description Development server`                                                                       |
|                        | @server.N.url         | Адрес сервера, где N — порядковый номер, начинается с 0                                    | `// @server.1.url https://{username}.gigantic-server.com:{port}/{basePath}`                                         |
|                        | @server.N.description | Описание сервера, где N — порядковый номер, начинается с 0                                 | `// @server.1.description The production API server`                                                                |
|                        | @server.N.variables   | Переменные сервера, где N — порядковый номер, начинается с 0                               | `// @server.1.variables username.default demo`                                                                      |
|                        |                       |                                                                                            | `// @server.1.variables username.description A user-specific subdomain. Use `demo` for a free sandbox environment.` |
|                        |                       |                                                                                            | `// @server.1.variables port.default 8433`                                                                          |
|                        |                       |                                                                                            | `// @server.1.variables port.enum 8433,433`                                                                         |
|                        |                       |                                                                                            | `// @server.1.variables basePath.default v1`                                                                        |
| paths                  |                       | Применяются аннотации из функций описывающих конкретные API методы                         |                                                                                                                     |
| _webhooks_             | ![TODO][7]            | ==Пока не реализовано==                                                                    |                                                                                                                     |
| _components_           | ![TODO][7]            | _Пока не реализовано_                                                                      |                                                                                                                     |
| _security_             | ![TODO][7]            | _Пока не реализовано_                                                                      |                                                                                                                     |
| [tags][8]              | @tag                  | Через теги в основной схеме можно задавать порядок и описание тегов в спецификации OpenAPI | `// @tag Auth Авторизация/Регистрация`                                                                              |
|                        |                       |                                                                                            | `// @tag Users Пользователи`                                                                                        |
|                        |                       |                                                                                            | `// @tag Pets Питомцы`                                                                                              |
|                        |                       |                                                                                            | `// @tag Other`                                                                                                     |
**Переменные сервера**

Для использования переменных сервера, необходимо в адресе сервера указать сначала именованные переменные, в примере это:

- `username` — Имя пользователя
- `port` — Порт
- `basePath` — Базовый путь

Далее после необходимо описать данные переменные, иначе запрос будет формироваться без изменения указанного адреса. В примере мы указываем для:

- `username` — `demo` через аннотацию `// @server.1.variables username.default demo`
- `port` — `8433` через аннотацию `// @server.1.variables port.default 8433`
- `basePath` — `v1` через аннотацию `// @server.1.variables basePath.default v1`

Стоит заметить, что переменные принимают следующие модификаторы:

| Модификатор   | Тип      | Описание                                                                                                  | Пример аннотации                                     |
|---------------|----------|-----------------------------------------------------------------------------------------------------------|------------------------------------------------------|
| `enum`        | []string |                                                                                                           | `// @server.1.variables port.enum 8433,433`          |
| `default`     | string   | **Обязательное**. Стандартное значение, из него подставляется в случае отсутствия данных от пользователя. | `// @server.1.variables username.default demo`       |
| `description` | string   |                                                                                                           | `// @server.1.description The production API server` |

#### [Info Object][1]

| Имя поля            | Аннотация           | Описание                         | Пример                                                             |
|---------------------|---------------------|----------------------------------|--------------------------------------------------------------------|
| [title][1]          | @title              | **ОБЯЗАТЕЛЬНОЕ**. Имя вашего API | `// @title Example Pet Store App`                                  |
| [summary][1]        | @summary            | Краткое описание API             | `// @summary A pet store manager.`                                 |
| [description][1]    | @description        | Подробно описание вашего API     | `// @description This is an example server for a pet store.`       |
| [termsOfService][1] | @termsOfService     |                                  | `// @termsOfService https://example.com/terms/`                    |
| [contact][2]        | @contact.name       |                                  | `// @contact.name API Support`                                     |
|                     | @contact.url        |                                  | `// @contact.email support@example.com`                            |
| [license][3]        | @license.name       |                                  | `// @license.name Apache 2.0`                                      |
|                     | @license.url        |                                  | `// @license.url https://www.apache.org/licenses/LICENSE-2.0.html` |
|                     | @license.identifier |                                  | `// @license.identifier Apache-2.0`                                |
| [version][1]        | @version            |                                  | `// @version 1.0.1`                                                |

[1]: https://swagger.io/specification/#info-object
[2]: https://swagger.io/specification/#contact-object
[3]: https://swagger.io/specification/#license-object
[4]: https://swagger.io/specification/#openapi-object
[5]: https://swagger.io/specification/#schema-object
[6]: https://swagger.io/specification/#server-object
[7]: https://img.shields.io/badge/TODO-1?color=orange
[8]: https://swagger.io/specification/#tag-object