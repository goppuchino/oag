# OpenAPI Generator for Go
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/goppuchino/oag)
![GitHub contributors](https://img.shields.io/github/contributors/goppuchino/oag)
![GitHub last commit](https://img.shields.io/github/last-commit/goppuchino/oag)
![GitHub License](https://img.shields.io/github/license/goppuchino/oag)

🌏 [English](README.md) | [Russian](README-ru.md)

<img align="right" width="180px" src="https://raw.githubusercontent.com/goppuchino/oag/master/assets/oag.png">

A fast and lightweight tool to **generate clean, standards-compliant OpenAPI 3 specifications** directly from your Go code. Perfect for:  
- Auto-documenting REST APIs 🏗️  
- Eliminating manual spec maintenance ✨  
- Ensuring compatibility with Swagger UI, Postman & more 🔌  

Features:  
✅ Struct-to-Schema auto-mapping  
✅ Built-in validation & linting  
✅ Custom template support

```go  
// Just annotate & generate!
// @method get
// @path /users/{id}
// @summary Get user
// @description Get user information
// @param id path isRequired User ID for get user information
func GetUser(w http.ResponseWriter, r *http.Request) { ... }  
```  

**Get started:**  
```shell
$ go install github.com/goppuchino/oag
$ cd <Your Project>
$ oag
```

*Because nobody loves writing YAML by hand.* 🐹

## Description

### main.go

The main description and general settings of the OpenAPI specification are read from annotations in the project's main file.
The objects listed below are described in annotations to the `main()` function.

#### [OpenAPI Object][4]

| Field Name           | Annotation           | Description                                                               | Example                                                                                           |
|----------------------|----------------------|---------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------|
| [openapi][4]         |                      | Standard value is set to `3.1.0`                                          |                                                                                                   |
| [info][1]            |                      | Annotations described in Info Object are applied                          |                                                                                                   |
| [jsonSchemaDialect][5]| @jsonSchemaDialect   | The dialect applied to the data schema (better not to specify)            | `// @jsonSchemaDialect https://spec.openapis.org/oas/3.1/dialect/base`                            |
| [servers][6]         | @server.N.url        | Server address, where N is the ordinal number, starting from 0            | `// @server.0.url https://development.gigantic-server.com/v1`                                     |
|                      | @server.N.description| Server description, where N is the ordinal number, starting from 0         | `// @server.0.description Development server`                                                     |
|                      | @server.N.url        | Server address, where N is the ordinal number, starting from 0            | `// @server.1.url https://{username}.gigantic-server.com:{port}/{basePath}`                       |
|                      | @server.N.description| Server description, where N is the ordinal number, starting from 0         | `// @server.1.description The production API server`                                              |
|                      | @server.N.variables  | Server variables, where N is the ordinal number, starting from 0          | `// @server.1.variables username.default demo`                                                    |
|                      |                      |                                                                           | `// @server.1.variables username.description A user-specific subdomain. Use `demo` for a free sandbox environment.` |
|                      |                      |                                                                           | `// @server.1.variables port.default 8433`                                                        |
|                      |                      |                                                                           | `// @server.1.variables port.enum 8433,433`                                                       |
|                      |                      |                                                                           | `// @server.1.variables basePath.default v1`                                                      |
| paths                |                      | Annotations from functions describing specific API methods are applied    |                                                                                                   |
| _webhooks_           | ![TODO][7]           | ==Not yet implemented==                                                   |                                                                                                   |
| _components_         | ![TODO][7]           | _Not yet implemented_                                                     |                                                                                                   |
| _security_           | ![TODO][7]           | _Not yet implemented_                                                     |                                                                                                   |
| [tags][8]            | @tag                 | Tags in the main schema can set the order and description of tags in the OpenAPI specification | `// @tag Auth Authorization/Registration`                                    |
|                      |                      |                                                                           | `// @tag Users Users`                                                                             |
|                      |                      |                                                                           | `// @tag Pets Pets`                                                                               |
|                      |                      |                                                                           | `// @tag Other`                                                                                   |

**Server Variables**

To use server variables, you must first specify named variables in the server address, in the example these are:

- `username` — Username
- `port` — Port
- `basePath` — Base path

Then you need to describe these variables, otherwise the request will be formed without changing the specified address. In the example, we specify for:

- `username` — `demo` through annotation `// @server.1.variables username.default demo`
- `port` — `8433` through annotation `// @server.1.variables port.default 8433`
- `basePath` — `v1` through annotation `// @server.1.variables basePath.default v1`

Note that variables accept the following modifiers:

| Modifier      | Type     | Description                                                                                  | Annotation Example                                   |
|---------------|----------|----------------------------------------------------------------------------------------------|------------------------------------------------------|
| `enum`        | []string |                                                                                              | `// @server.1.variables port.enum 8433,433`          |
| `default`     | string   | **Required**. Default value, used when no data is provided by the user.                      | `// @server.1.variables username.default demo`       |
| `description` | string   |                                                                                              | `// @server.1.description The production API server` |

#### [Info Object][1]

| Field Name          | Annotation         | Description                      | Example                                                           |
|--------------------|--------------------|----------------------------------|-------------------------------------------------------------------|
| [title][1]         | @title             | **REQUIRED**. Name of your API   | `// @title Example Pet Store App`                                 |
| [summary][1]       | @summary           | Brief API description            | `// @summary A pet store manager.`                                |
| [description][1]   | @description       | Detailed description of your API | `// @description This is an example server for a pet store.`      |
| [termsOfService][1]| @termsOfService    |                                  | `// @termsOfService https://example.com/terms/`                   |
| [contact][2]       | @contact.name      |                                  | `// @contact.name API Support`                                    |
|                    | @contact.url       |                                  | `// @contact.email support@example.com`                           |
| [license][3]       | @license.name      |                                  | `// @license.name Apache 2.0`                                     |
|                    | @license.url       |                                  | `// @license.url https://www.apache.org/licenses/LICENSE-2.0.html`|
|                    | @license.identifier|                                  | `// @license.identifier Apache-2.0`                               |
| [version][1]       | @version           |                                  | `// @version 1.0.1`                                               |

[1]: https://swagger.io/specification/#info-object
[2]: https://swagger.io/specification/#contact-object
[3]: https://swagger.io/specification/#license-object
[4]: https://swagger.io/specification/#openapi-object
[5]: https://swagger.io/specification/#schema-object
[6]: https://swagger.io/specification/#server-object
[7]: https://img.shields.io/badge/TODO-1?color=orange
[8]: https://swagger.io/specification/#tag-object

