---
layout: home
title: Introduction
nav_order: 1
---

# OpenAPI Generator for Go
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/goppuchino/oag)
![GitHub contributors](https://img.shields.io/github/contributors/goppuchino/oag)
![GitHub last commit](https://img.shields.io/github/last-commit/goppuchino/oag)
![GitHub License](https://img.shields.io/github/license/goppuchino/oag)

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