package shaders_test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"path"
	"reflect"
	"testing"

	"qlova.tech/gpu"
	"qlova.tech/gpu/driver/opengl/glsl"
	"qlova.tech/gpu/internal/shaders"
	"qlova.tech/gpu/mat4"
	"qlova.tech/gpu/rgba"
	"qlova.tech/gpu/vec4"
)

var fset = token.NewFileSet()

type Importer struct {
	imported map[string]*types.Package
}

const embedSource = `package embed`

func (imp *Importer) Import(p string) (*types.Package, error) {
	if imp.imported[p] != nil {
		return imp.imported[p], nil
	}

	base := path.Base(p)

	var source string

	switch p {
	case "embed":
		source = embedSource
	case "qlova.tech/gpu":
		source = gpu.Source
	case "qlova.tech/gpu/mat4":
		source = mat4.Source
	case "qlova.tech/gpu/rgba":
		source = rgba.Source
	case "qlova.tech/gpu/vec4":
		source = vec4.Source
	default:
		return nil, fmt.Errorf("unknown package %q", p)
	}

	f, err := parser.ParseFile(fset, base+".go", source, 0)
	if err != nil {
		return nil, err
	}

	// A Config controls various options of the type checker.
	// The defaults work fine except for one setting:
	// we must specify how to deal with imports.
	conf := types.Config{Importer: imp}

	pkg, err := conf.Check(p, fset, []*ast.File{f}, nil)
	if err != nil {
		return nil, err
	}

	if imp.imported == nil {
		imp.imported = make(map[string]*types.Package)
	}
	imp.imported[p] = pkg

	return pkg, nil
}

func Parse[Frag, Vert any](renderer gpu.Renderer[Frag, Vert]) error {

	name := reflect.TypeOf(renderer).Name()

	f, err := parser.ParseFile(fset, name+".go", renderer.Source(), 0)
	if err != nil {
		return err
	}

	//Type check pass.
	conf := types.Config{Importer: new(Importer)}
	pkg, err := conf.Check("shaders", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err)
	}

	b, err := glsl.Compile(fset, f, pkg, name)
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}

func TestShaders(t *testing.T) {
	err := Parse(shaders.TexturedMeshRenderer{})
	if err != nil {
		t.Fatal(err)
	}
}
