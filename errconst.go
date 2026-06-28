// Package errconst provides a go/analysis analyzer enforcing the gomatic Go
// standard that errors are sentinel constants (errs.Const), never created with
// errors.New or fmt.Errorf — fmt.Errorf is permitted only to wrap a cause with %w.
package errconst

import (
	"go/ast"
	"strings"

	goyze "github.com/gomatic/go-yze"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/types/typeutil"
)

const (
	messageErrorsNew = "use a sentinel error constant (errs.Const) instead of errors.New"
	messageErrorf    = "use a sentinel error constant, or wrap a cause with %%w, instead of fmt.Errorf"
)

// Analyzer reports disallowed error construction.
var Analyzer = &analysis.Analyzer{
	Name:     "errconst",
	Doc:      "reports errors.New and non-wrapping fmt.Errorf, which the gomatic Go standard forbids",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

// Registration declares this analyzer to the yze framework.
var Registration = goyze.Registration{
	Name:       "errconst",
	Group:      "go",
	Categories: []goyze.Category{"errors"},
	URL:        "https://docs.gomatic.dev/yze/go/errconst",
	Analyzer:   Analyzer,
}

// run reports each disallowed error-construction call.
func run(pass *analysis.Pass) (any, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	insp.Preorder([]ast.Node{(*ast.CallExpr)(nil)}, func(n ast.Node) {
		check(pass, n.(*ast.CallExpr))
	})
	return nil, nil
}

// check reports the call when it is a disallowed error constructor.
func check(pass *analysis.Pass, call *ast.CallExpr) {
	switch calleeName(pass, call) {
	case "errors.New":
		pass.Reportf(call.Pos(), messageErrorsNew)
	case "fmt.Errorf":
		reportErrorf(pass, call)
	}
}

// reportErrorf flags an fmt.Errorf call that does not wrap a cause with %w.
func reportErrorf(pass *analysis.Pass, call *ast.CallExpr) {
	if !wrapsCause(call) {
		pass.Reportf(call.Pos(), messageErrorf)
	}
}

// wrapsCause reports whether an fmt.Errorf call wraps a cause with %w. A
// non-literal format string cannot be judged statically and is assumed to wrap.
func wrapsCause(call *ast.CallExpr) bool {
	lit, ok := call.Args[0].(*ast.BasicLit)
	if !ok {
		return true
	}
	return strings.Contains(lit.Value, "%w")
}

// calleeName returns the fully-qualified name of the statically-called function
// (e.g. "errors.New"), or "" when the callee is not a static function.
func calleeName(pass *analysis.Pass, call *ast.CallExpr) string {
	if fn := typeutil.StaticCallee(pass.TypesInfo, call); fn != nil {
		return fn.FullName()
	}
	return ""
}
