package generator

import (
	"fmt"
	"strconv"
	"strings"
)

func GenerateRoot(text string, spec *Spec) {
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
	} else if strings.HasPrefix(text, "@tag") {
		tagData := strings.SplitN(strings.TrimSpace(strings.TrimPrefix(text, "@tag")), " ", 2)
		if len(spec.Tags) == 0 {
			spec.Tags = make([]map[string]interface{}, 1)
			spec.Tags[0] = map[string]interface{}{
				"name": tagData[0],
			}
			if len(tagData) > 1 {
				spec.Tags[0]["description"] = tagData[1]
			}
		} else {
			spec.Tags = append(spec.Tags, map[string]interface{}{
				"name": tagData[0],
			})
			if len(tagData) > 1 {
				spec.Tags[len(spec.Tags)-1]["description"] = tagData[1]
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
}
