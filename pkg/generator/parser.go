package generator

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
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
		var method, path string
		pathSpec := &PathSpec{}

		for _, comment := range commentGroup.List {
			text := strings.TrimSpace(strings.TrimPrefix(comment.Text, "//"))

			if isMain {
				GenerateRoot(text, spec)
			} else {
				GeneratePath(text, &method, &path, pathSpec)
			}
		}

		if method != "" && path != "" {
			if _, exists := spec.Paths[path]; !exists {
				spec.Paths[path] = map[string]interface{}{}
			}
			spec.Paths[path].(map[string]interface{})[strings.ToLower(method)] = pathSpec
		}
	}
}
