// Every operation in the specification must be reachable on the client.
//
// A naming scheme that shortens method names can collapse two operations onto one silently, and a
// namespace nothing reaches would pass a count that only added methods up. So this walks the
// generated client from the root, the way a caller does.
package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"sort"
	"strings"
)

const (
	specPath = "openapi/company-v3.json"
	apiPath  = "api.gen.go"
	root     = "Client"
)

var methods = map[string]bool{"get": true, "post": true, "put": true, "patch": true, "delete": true}

// A group of operations: the methods it holds, and the groups under it.
type group struct {
	operations []string
	children   map[string]string
}

func main() {
	expected := operations()
	groups, held := parseAPI()

	var reached []string

	walk(groups, held, root, "clockster", &reached)

	sort.Strings(reached)

	if len(reached) != expected {
		fmt.Fprintf(os.Stderr,
			"%d operations in the specification, %d reachable on the client. Two operations whose "+
				"names collide are silently merged; check `overrides` in "+
				"internal/cmd/generate/names.go.\n", expected, len(reached))

		os.Exit(1)
	}

	fmt.Printf("%d operations, all reachable.\n", len(reached))
}

func operations() int {
	raw, err := os.ReadFile(specPath)
	if err != nil {
		fail("%s cannot be read: %v", specPath, err)
	}

	var document struct {
		Paths map[string]map[string]json.RawMessage `json:"paths"`
	}

	if err := json.Unmarshal(raw, &document); err != nil {
		fail("%s is not a document this can read: %v", specPath, err)
	}

	count := 0

	for _, item := range document.Paths {
		for method := range item {
			if methods[method] {
				count++
			}
		}
	}

	return count
}

// parseAPI answers the namespaces the generated file declares, and the operations on each. Read
// from the source rather than by reflection: a method is reachable because a field leads to it,
// which is what a caller has.
func parseAPI() (map[string]*group, map[string]bool) {
	file, err := parser.ParseFile(token.NewFileSet(), apiPath, nil, 0)
	if err != nil {
		fail("%s cannot be read: %v", apiPath, err)
	}

	groups := map[string]*group{}
	iterators := map[string]bool{}

	for _, declaration := range file.Decls {
		switch held := declaration.(type) {
		case *ast.GenDecl:
			readTypes(held, groups)
		case *ast.FuncDecl:
			readMethod(held, groups, iterators)
		}
	}

	return groups, iterators
}

func readTypes(declaration *ast.GenDecl, groups map[string]*group) {
	for _, spec := range declaration.Specs {
		held, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}

		structure, ok := held.Type.(*ast.StructType)
		if !ok {
			continue
		}

		found := &group{children: map[string]string{}}

		for _, field := range structure.Fields.List {
			pointer, ok := field.Type.(*ast.StarExpr)
			if !ok || len(field.Names) != 1 {
				continue
			}

			name, ok := pointer.X.(*ast.Ident)
			if !ok || name.Name == "Client" {
				continue
			}

			found.children[field.Names[0].Name] = name.Name
		}

		groups[held.Name.Name] = found
	}
}

func readMethod(declaration *ast.FuncDecl, groups map[string]*group, iterators map[string]bool) {
	if declaration.Recv == nil || len(declaration.Recv.List) != 1 || !declaration.Name.IsExported() {
		return
	}

	pointer, ok := declaration.Recv.List[0].Type.(*ast.StarExpr)
	if !ok {
		return
	}

	// A method on the client itself is an operation of the root, which is what `me` is.
	name, ok := pointer.X.(*ast.Ident)
	if !ok {
		return
	}

	receiver := name.Name

	found, ok := groups[receiver]
	if !ok {
		found = &group{children: map[string]string{}}
		groups[receiver] = found
	}

	if strings.HasSuffix(declaration.Name.Name, "All") {
		iterators[receiver+"."+declaration.Name.Name] = true
	}

	found.operations = append(found.operations, declaration.Name.Name)
}

func walk(groups map[string]*group, iterators map[string]bool, name, prefix string, reached *[]string) {
	held, ok := groups[name]
	if !ok {
		return
	}

	for _, operation := range held.operations {
		// ListAll walks the pages of List; it is one operation between them rather than two.
		if iterators[name+"."+operation] && contains(held.operations, strings.TrimSuffix(operation, "All")) {
			continue
		}

		*reached = append(*reached, prefix+"."+operation+"()")
	}

	for field, child := range held.children {
		walk(groups, iterators, child, prefix+"."+field, reached)
	}
}

func contains(held []string, name string) bool {
	for _, one := range held {
		if one == name {
			return true
		}
	}

	return false
}

func fail(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
