# OpenAPI Generator for Go

🚀 *Effortless OpenAPI Spec Generation in Go*  

A fast and lightweight tool to **generate clean, standards-compliant OpenAPI specifications** directly from your Go code. Perfect for:  
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
func GetUser(w http.ResponseWriter, r *http.Request) { ... }  
```  

**Get started:**  
`go get github.com/your-repo/go-openapi-gen`  

*Because nobody loves writing YAML by hand.* 🐹
