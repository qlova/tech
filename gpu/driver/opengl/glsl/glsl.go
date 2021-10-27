package glsl

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"io"
	"strings"
)

//Compiler compiles a Go AST to glsl.
type Compiler struct {
	pkg   *types.Package
	Error error
	fset  *token.FileSet

	vertexType   string
	fragmentType string
}

func (c *Compiler) compileUniforms(w io.Writer, spec *ast.TypeSpec) error {
	if structType, ok := spec.Type.(*ast.StructType); ok {
		for i, field := range structType.Fields.List {
			if len(field.Names) != 1 {
				return fmt.Errorf("unexpected number of names for field %s", field.Names)
			}

			var name = field.Names[0].Name

			var goType = types.ExprString(field.Type)

			var glslType string

			//Use shader storage buffer instead of ordinary uniform
			//for slice fields.
			var shaderStorage bool
			if strings.HasPrefix(goType, "[]") {
				shaderStorage = true
				goType = goType[2:]
			}

			switch goType {
			case "gpu.Mat4":
				glslType = "mat4"
			case "gpu.Texture":
				glslType = "sampler2D"
			default:
				return fmt.Errorf("unsupported uniform type %s", types.ExprString(field.Type))
			}

			fmt.Fprintln(w)

			if shaderStorage {
				fmt.Fprintf(w, "layout(binding=%v) readonly buffer %vBuffer {\n", i, name)
				fmt.Fprintf(w, "\t%v %v[];\n", glslType, name)
				fmt.Fprintf(w, "};\n")
			} else {
				fmt.Fprintf(w, "uniform %v %v;\n", glslType, name)
			}
		}
	}
	return nil
}

func (c *Compiler) compileExpr(expr ast.Expr) string {
	var b bytes.Buffer
	switch expr := expr.(type) {
	case *ast.Ident:
		fmt.Println(types.Eval(c.fset, c.pkg, expr.NamePos, expr.Name))
		return fmt.Sprintf("%v", expr.Name)
	case *ast.BasicLit:
		return fmt.Sprintf("%v", expr.Value)
	case *ast.CallExpr:
		if len(expr.Args) != 1 {
			c.Error = fmt.Errorf("unexpected number of arguments for %v", expr)
			return ""
		}
		fmt.Fprintf(&b, "%v", expr.Fun)
		fmt.Fprintf(&b, "(")
		fmt.Fprintf(&b, "%v", c.compileExpr(expr.Args[0]))
		fmt.Fprintf(&b, ")")
		return b.String()
	case *ast.UnaryExpr:
		fmt.Fprintf(&b, "%v", expr.Op)
		fmt.Fprintf(&b, "%v", c.compileExpr(expr.X))
		return b.String()
	case *ast.BinaryExpr:
		fmt.Fprintf(&b, "(")
		fmt.Fprintf(&b, "%v", c.compileExpr(expr.X))
		fmt.Fprintf(&b, " %v ", expr.Op)
		fmt.Fprintf(&b, "%v", c.compileExpr(expr.Y))
		fmt.Fprintf(&b, ")")
		return b.String()
	case *ast.ParenExpr:
		fmt.Fprintf(&b, "(")
		fmt.Fprintf(&b, "%v", c.compileExpr(expr.X))
		fmt.Fprintf(&b, ")")
		return b.String()
	case *ast.SelectorExpr:
		fmt.Fprintf(&b, "%v.%v", c.compileExpr(expr.X), expr.Sel.Name)
		return b.String()
	default:
		c.Error = fmt.Errorf("unsupported expression type %T", expr)
	}
	return ""
}

func (c *Compiler) compileStmt(w io.Writer, stmt ast.Stmt) error {
	switch stmt := stmt.(type) {
	case *ast.BlockStmt:
		for _, s := range stmt.List {
			if err := c.compileStmt(w, s); err != nil {
				return err
			}
		}
	case *ast.ReturnStmt:
		fmt.Fprintf(w, "return;\n")
	/*case *ast.BinaryExpr:
	if stmt.Op == token.ADD {
		fmt.Fprintf(w, "%v += %v;\n", compileExpr(w, stmt.X), compileExpr(w, stmt.Y))
	} else {
		return fmt.Errorf("unsupported binary expression %v", stmt.Op)
	}*/
	case *ast.AssignStmt:
		if len(stmt.Lhs) != 1 {
			return fmt.Errorf("unsupported number of lhs %v", len(stmt.Lhs))
		}
		if len(stmt.Rhs) != 1 {
			return fmt.Errorf("unsupported number of rhs %v", len(stmt.Rhs))
		}
		fmt.Fprintf(w, "%v = %v;\n", c.compileExpr(stmt.Lhs[0]), c.compileExpr(stmt.Rhs[0]))
	case *ast.ExprStmt:
		fmt.Fprintf(w, "%v;\n", c.compileExpr(stmt.X))
	case *ast.IfStmt:
		fmt.Fprintf(w, "if (%v) {\n", c.compileExpr(stmt.Cond))
		if err := c.compileStmt(w, stmt.Body); err != nil {
			return err
		}
		fmt.Fprintf(w, "}\n")
		if stmt.Else != nil {
			fmt.Fprintf(w, "else {\n")
			if err := c.compileStmt(w, stmt.Else); err != nil {
				return err
			}
			fmt.Fprintf(w, "}\n")
		}
	case *ast.ForStmt:
		fmt.Fprintf(w, "for (%v; %v; %v) {\n", c.compileStmt(w, stmt.Init), c.compileExpr(stmt.Cond), c.compileStmt(w, stmt.Post))
		if err := c.compileStmt(w, stmt.Body); err != nil {
			return err
		}
		fmt.Fprintf(w, "}\n")
	case *ast.SwitchStmt:
		fmt.Fprintf(w, "switch (%v) {\n", c.compileExpr(stmt.Tag))
		for _, s := range stmt.Body.List {
			if s, ok := s.(*ast.CaseClause); ok {
				fmt.Fprintf(w, "case %v:\n", c.compileExpr(s.List[0]))
				for _, s := range s.Body {
					if err := c.compileStmt(w, s); err != nil {
						return err
					}
				}
			}
		}
		fmt.Fprintf(w, "}\n")
	/*case *ast.TypeAssertExpr:
	return fmt.Errorf("unsupported type assertion %v", stmt)*/
	case *ast.BranchStmt:
		if stmt.Tok == token.BREAK {
			fmt.Fprintf(w, "break;\n")
		} else {
			return fmt.Errorf("unsupported branch statement %v", stmt)
		}
	case *ast.IncDecStmt:
		if stmt.Tok == token.INC {
			fmt.Fprintf(w, "%v++;\n", c.compileExpr(stmt.X))
		} else {
			return fmt.Errorf("unsupported inc statement %v", stmt)
		}
	case *ast.DeferStmt:
		return fmt.Errorf("unsupported defer statement %v", stmt)
	case *ast.GoStmt:
		return fmt.Errorf("unsupported go statement %v", stmt)
	case *ast.EmptyStmt:
		//Do nothing
	default:
		return fmt.Errorf("unsupported statement %v", stmt)
	}
	return c.Error
}

func (c *Compiler) compileVertex(w io.Writer, body *ast.BlockStmt) error {
	fmt.Fprintln(w, "void main() {")
	if err := c.compileStmt(w, body); err != nil {
		return err
	}
	fmt.Fprintln(w, "}")
	return nil
}

func (c *Compiler) extractType(ident *ast.Ident) string {
	t, _ := types.Eval(c.fset, c.pkg, ident.NamePos, ident.Name)
	return t.Type.String()
}

func Compile(fs *token.FileSet, f *ast.File, pkg *types.Package, shader string) ([]byte, error) {
	var c = new(Compiler)
	c.pkg = pkg
	c.fset = fs

	var buffer bytes.Buffer
	buffer.WriteString("#version 460 core\n")
	buffer.WriteString("#extension GL_ARB_bindless_texture : require\n")

	//Look for the shader type with name 'shader'
	//this shader contains the uniforms.
	for _, decl := range f.Decls {
		if gen, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range gen.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if typeSpec.Name.Name == shader {
						if err := c.compileUniforms(&buffer, typeSpec); err != nil {
							return nil, err
						}
					}
				}
			}
		}
	}

	fmt.Fprintln(&buffer)

	var vertexBody *ast.BlockStmt

	//Look for the Vertex method on the shader type 'shader'.
	//We need to extract the vertex and fragment types of the
	//method ie.
	//
	//		func (Shader) Vertex(Vertex,*Fragment) gpu.Vec4
	//
	//These are used to compile the in and out attributes.
	for _, decl := range f.Decls {
		if gen, ok := decl.(*ast.FuncDecl); ok {
			if gen.Recv != nil && gen.Recv.List[0].Type != nil {
				if gen.Recv.List[0].Type.(*ast.Ident).Name == shader {
					if gen.Name.Name == "Vertex" {

						//We need to extract the vertex and fragment types
						//from the method signature.
						c.vertexType = c.extractType(gen.Type.Params.List[0].Type.(*ast.Ident))
						c.fragmentType = "*" + c.extractType(gen.Type.Params.List[1].Type.(*ast.StarExpr).X.(*ast.Ident))

						vertexBody = gen.Body
					}
				}
			}
		}
	}

	if err := c.compileVertex(&buffer, vertexBody); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
