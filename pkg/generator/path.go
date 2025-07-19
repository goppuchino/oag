package generator

import (
	"fmt"
	"github.com/goppuchino/oag/pkg/utils"
	"strings"
)

func GeneratePath(text string, method, path *string, spec *PathSpec) {
	var tags []string

	if strings.HasPrefix(text, "@method") {
		*method = strings.TrimSpace(strings.TrimPrefix(text, "@method"))
	} else if strings.HasPrefix(text, "@path") {
		*path = strings.TrimSpace(strings.TrimPrefix(text, "@path"))
	} else if strings.HasPrefix(text, "@summary") {
		spec.Summary = strings.TrimSpace(strings.TrimPrefix(text, "@summary"))
	} else if strings.HasPrefix(text, "@description") {
		descriptionLine := strings.TrimSpace(strings.TrimPrefix(text, "@description"))
		if spec.Description != "" {
			spec.Description += "\n"
		}
		spec.Description += descriptionLine
	} else if strings.HasPrefix(text, "@tags") {
		tagList := strings.TrimSpace(strings.TrimPrefix(text, "@tags"))
		tags = append(tags, strings.Split(tagList, ",")...)
		spec.Tags = utils.Unique(tags)
	} else if strings.HasPrefix(text, "@param") {
		paramContent := strings.TrimSpace(strings.TrimPrefix(text, "@param"))
		isRequired := strings.Contains(paramContent, "isRequired")
		isDeprecated := strings.Contains(paramContent, "isDeprecated")
		allowEmptyValue := strings.Contains(paramContent, "allowEmptyValue")
		style := ParamStyleSimple
		if isRequired {
			paramContent = strings.ReplaceAll(paramContent, "isRequired", "")
		}
		if isDeprecated {
			paramContent = strings.ReplaceAll(paramContent, "isDeprecated", "")
		}
		if allowEmptyValue {
			paramContent = strings.ReplaceAll(paramContent, "allowEmptyValue", "")
		}
		// format: name* in* isRequired (xor) isDeprecated (xor) allowEmptyValue description
		parts := strings.SplitN(paramContent, " ", 4)
		if len(parts) >= 2 {
			param := map[string]interface{}{
				"name":     parts[0],
				"in":       parts[1],
				"required": isRequired,
				"schema": map[string]interface{}{
					"type": "string",
				},
				"style": ParamStyleName[style],
			}

			if parts[1] == "path" {
				param["required"] = true
			} else if parts[1] == "query" {
				param["style"] = ParamStyleName[ParamStyleForm]
			}

			if len(parts) > 3 {
				param["description"] = strings.TrimSpace(parts[3])
			}

			spec.Parameters = append(spec.Parameters, param)
		} else {
			fmt.Println("Параметр @param имеет некорректный формат:", text)
		}
	}
}
