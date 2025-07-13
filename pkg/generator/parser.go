package generator

import (
	"fmt"
	"github.com/goppuchino/oag/pkg/utils"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"
)

func GenerateOpenAPISpec(root string) (*Spec, error) {
	spec := &Spec{
		OpenAPI: "3.1.0",
		Info: map[string]interface{}{
			"title":   "API Documentation",
			"version": "1.0.0",
		},
		Paths: map[string]interface{}{},
	}

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		fileSet := token.NewFileSet()
		node, err := parser.ParseFile(fileSet, path, nil, parser.ParseComments)
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, "main.go") {
			parseAnnotations(node.Comments, spec, true)
		} else {
			parseAnnotations(node.Comments, spec, false)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return spec, nil
}

func parseAnnotations(comments []*ast.CommentGroup, spec *Spec, isMain bool) {
	for _, commentGroup := range comments {
		var method, path, summary, description string
		var tags []string
		var parameters []map[string]interface{}

		for _, comment := range commentGroup.List {
			text := strings.TrimSpace(strings.TrimPrefix(comment.Text, "//"))

			if isMain {
				// OpenAPI root objects
				if strings.HasPrefix(text, "@jsonSchemaDialect") {
					spec.JsonSchemaDialect = strings.TrimSpace(strings.TrimPrefix(text, "@jsonSchemaDialect"))
				} else if strings.HasPrefix(text, "@server.") {
					parts := strings.SplitN(strings.TrimPrefix(text, "@server."), " ", 2)
					if len(parts) == 2 {
						serverParts := strings.Split(parts[0], ".")
						value := strings.TrimSpace(parts[1])

						if len(spec.Servers) == 0 {
							spec.Servers = make([]map[string]interface{}, 1)
							spec.Servers[0] = make(map[string]interface{})
						}

						serverIndex := 0
						if len(serverParts) > 1 {
							if idx, err := strconv.Atoi(serverParts[0]); err == nil {
								serverIndex = idx
								if serverIndex >= len(spec.Servers) {
									newServers := make([]map[string]interface{}, serverIndex+1)
									copy(newServers, spec.Servers)
									spec.Servers = newServers
								}
								if spec.Servers[serverIndex] == nil {
									spec.Servers[serverIndex] = make(map[string]interface{})
								}
							}
						}

						server := spec.Servers[serverIndex]
						lastPart := serverParts[len(serverParts)-1]

						switch lastPart {
						case "url":
							server["url"] = value
						case "description":
							server["description"] = value
						default:
							if strings.HasPrefix(lastPart, "variables") {
								varParts := strings.SplitN(value, " ", 2)
								if len(varParts) < 2 {
									fmt.Println("Некорректная структура аннотации для variables")
									return
								}
								keyPath := varParts[0]
								actualValue := varParts[1]
								keySegments := strings.Split(keyPath, ".")
								if len(keySegments) < 2 {
									fmt.Println("Некорректный путь ключа для variables")
									return
								}

								varName := keySegments[0]
								varAttr := keySegments[1]

								if server["variables"] == nil {
									server["variables"] = make(map[string]interface{})
								}
								variables := server["variables"].(map[string]interface{})

								if variables[varName] == nil {
									variables[varName] = make(map[string]interface{})
								}
								variable := variables[varName].(map[string]interface{})

								switch varAttr {
								case "default":
									variable["default"] = actualValue
								case "description":
									variable["description"] = actualValue
								case "enum":
									enumValues := strings.Split(actualValue, ",")
									for i, v := range enumValues {
										enumValues[i] = strings.TrimSpace(v)
									}
									variable["enum"] = enumValues
								}
							}
						}
					}
				}

				// Info schema
				if strings.HasPrefix(text, "@title") {
					spec.Info["title"] = strings.TrimSpace(strings.TrimPrefix(text, "@title"))
				} else if strings.HasPrefix(text, "@version") {
					spec.Info["version"] = strings.TrimSpace(strings.TrimPrefix(text, "@version"))
				} else if strings.HasPrefix(text, "@summary") {
					spec.Info["summary"] = strings.TrimSpace(strings.TrimPrefix(text, "@summary"))
				} else if strings.HasPrefix(text, "@description") {
					spec.Info["description"] = strings.TrimSpace(strings.TrimPrefix(text, "@description"))
				} else if strings.HasPrefix(text, "@termsOfService") {
					spec.Info["termsOfService"] = strings.TrimSpace(strings.TrimPrefix(text, "@termsOfService"))
				} else if strings.HasPrefix(text, "@contact.") {
					if spec.Info["contact"] == nil {
						spec.Info["contact"] = make(map[string]interface{})
					}
					contact := spec.Info["contact"].(map[string]interface{})

					if strings.HasPrefix(text, "@contact.name") {
						contact["name"] = strings.TrimSpace(strings.TrimPrefix(text, "@contact.name"))
					} else if strings.HasPrefix(text, "@contact.url") {
						contact["url"] = strings.TrimSpace(strings.TrimPrefix(text, "@contact.url"))
					} else if strings.HasPrefix(text, "@contact.email") {
						contact["email"] = strings.TrimSpace(strings.TrimPrefix(text, "@contact.email"))
					}

				} else if strings.HasPrefix(text, "@license.") {
					if spec.Info["license"] == nil {
						spec.Info["license"] = make(map[string]interface{})
					}
					license := spec.Info["license"].(map[string]interface{})

					if strings.HasPrefix(text, "@license.name") {
						license["name"] = strings.TrimSpace(strings.TrimPrefix(text, "@license.name"))
					} else if strings.HasPrefix(text, "@license.identifier") {
						license["identifier"] = strings.TrimSpace(strings.TrimPrefix(text, "@license.identifier"))
					} else if strings.HasPrefix(text, "@license.url") {
						license["url"] = strings.TrimSpace(strings.TrimPrefix(text, "@license.url"))
					}
				}

				// Other schemas not specified in file main.go
			} else {
				if strings.HasPrefix(text, "@method") {
					method = strings.TrimSpace(strings.TrimPrefix(text, "@method"))
				} else if strings.HasPrefix(text, "@path") {
					path = strings.TrimSpace(strings.TrimPrefix(text, "@path"))
				} else if strings.HasPrefix(text, "@summary") {
					summary = strings.TrimSpace(strings.TrimPrefix(text, "@summary"))
				} else if strings.HasPrefix(text, "@description") {
					description = strings.TrimSpace(strings.TrimPrefix(text, "@description"))
				} else if strings.HasPrefix(text, "@tags") {
					tagList := strings.TrimSpace(strings.TrimPrefix(text, "@tags"))
					tags = append(tags, strings.Split(tagList, ",")...)
				} else if strings.HasPrefix(text, "@param") {
					parapmContent := strings.TrimSpace(strings.TrimPrefix(text, "@param"))
					parts := strings.SplitN(parapmContent, " ", 4) // format: name* in* required description
					fmt.Println("parts:", parts)
					fmt.Println("len(parts):", len(parts))
					if len(parts) < 2 {
						fmt.Println("Параметр @param имеет некорректный формат:", text)
						continue
					}

					param := map[string]interface{}{
						"name":     parts[0],
						"in":       parts[1],
						"required": utils.StringToBool(parts[2]),
						"schema": map[string]interface{}{
							"type": "string",
						},
					}

					if len(parts) > 3 {
						param["description"] = strings.TrimSpace(parts[3])
					}

					parameters = append(parameters, param)
				}
			}
		}

		if method != "" && path != "" {
			if _, exists := spec.Paths[path]; !exists {
				spec.Paths[path] = map[string]interface{}{}
			}

			methodSpec := PathSpec{
				Summary:     summary,
				Description: description,
				Tags:        utils.Unique(tags),
			}

			if len(parameters) > 0 {
				methodSpec.Parameters = parameters
			}

			spec.Paths[path].(map[string]interface{})[strings.ToLower(method)] = methodSpec
		}
	}
}
