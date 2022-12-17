package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

type method struct {
	Name string
	In   string
	Out  string
}

func getInterface(node *ast.File, name string) *ast.InterfaceType {
	for _, dec := range node.Decls {
		if gen, ok := dec.(*ast.GenDecl); ok {
			if gen.Tok != token.TYPE {
				continue
			}
			for _, specs := range gen.Specs {
				if ts, ok := specs.(*ast.TypeSpec); ok {
					if ts.Name.String() != name {
						continue
					}
					if iface, ok := ts.Type.(*ast.InterfaceType); ok {
						return iface
					}
				}
			}
		}
	}
	return nil
}

func getMethods(iface *ast.InterfaceType) []method {
	var methods []method
	for _, m := range iface.Methods.List {
		if ft, ok := m.Type.(*ast.FuncType); ok {
			methods = append(methods, method{
				Name: m.Names[0].Name,
				In:   getExprString(ft.Params.List[1].Type),
				Out:  getExprString(ft.Results.List[0].Type),
			})
		}
	}
	return methods
}

func getExprString(expr ast.Expr) string {
	var buf bytes.Buffer
	processExpr(expr, &buf)
	return buf.String()
}

func processExpr(expr ast.Expr, buf *bytes.Buffer) {
	switch t := expr.(type) {
	case *ast.SelectorExpr:
		processSelectorExpr(t, buf)
	case *ast.StarExpr:
		processStarExpr(t, buf)
	case *ast.Ident:
		buf.WriteString(t.Name)
	}
}

func processSelectorExpr(t *ast.SelectorExpr, buf *bytes.Buffer) {
	if ident, ok := t.X.(*ast.Ident); ok {
		buf.WriteString(ident.Name)
	}
	buf.WriteString(".")
	buf.WriteString(t.Sel.Name)
}

func processStarExpr(t *ast.StarExpr, buf *bytes.Buffer) {
	buf.WriteString("*")
	processExpr(t.X, buf)
}

func processApiType(t string) string {
	if strings.ContainsRune(t, '.') {
		return t
	}
	return strings.Replace(t, "*", "*api.", 1)
}

func genMethod(f *os.File, m method) {
	f.WriteString(fmt.Sprintf(`
func (g *GrpcClient) %s(ctx context.Context, in %s, opts ...grpc.CallOption) (%s, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.%s(ctx, in, opts...)
}
`, m.Name, processApiType(m.In), processApiType(m.Out), m.Name))
}

func main() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "api/api_grpc.pb.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	iface := getInterface(node, "WalletClient")
	methods := getMethods(iface)

	f, err := os.Create("generated.go")
	if err != nil {
		log.Fatalln(err)
	}

	f.WriteString(`package tron_grpc

import (
	"context"
	"github.com/fullstackwang/tron-grpc/api"
	"github.com/fullstackwang/tron-grpc/core"
	"google.golang.org/grpc"
)
`)

	for _, m := range methods {
		genMethod(f, m)
	}
}
