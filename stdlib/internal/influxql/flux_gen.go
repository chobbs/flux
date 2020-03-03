// DO NOT EDIT: This file is autogenerated via the builtin command.

package influxql

import (
	ast "github.com/influxdata/flux/ast"
	parser "github.com/influxdata/flux/internal/parser"
	runtime "github.com/influxdata/flux/runtime"
)

func init() {
	runtime.RegisterPackage(pkgAST)
}

var pkgAST = &ast.Package{
	BaseNode: ast.BaseNode{
		Errors: nil,
		Loc:    nil,
	},
	Files: []*ast.File{&ast.File{
		BaseNode: ast.BaseNode{
			Errors: nil,
			Loc: &ast.SourceLocation{
				End: ast.Position{
					Column: 41,
					Line:   5,
				},
				File:   "",
				Source: "package influxql\n\nepoch = 1970-01-01T00:00:00Z\nminTime = 1677-09-21T00:12:43.145224194Z\nmaxTime = 2262-04-11T23:47:16.854775806Z",
				Start: ast.Position{
					Column: 1,
					Line:   1,
				},
			},
		},
		Body: []ast.Statement{&ast.VariableAssignment{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 29,
						Line:   3,
					},
					File:   "",
					Source: "epoch = 1970-01-01T00:00:00Z",
					Start: ast.Position{
						Column: 1,
						Line:   3,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 6,
							Line:   3,
						},
						File:   "",
						Source: "epoch",
						Start: ast.Position{
							Column: 1,
							Line:   3,
						},
					},
				},
				Name: "epoch",
			},
			Init: &ast.DateTimeLiteral{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 29,
							Line:   3,
						},
						File:   "",
						Source: "1970-01-01T00:00:00Z",
						Start: ast.Position{
							Column: 9,
							Line:   3,
						},
					},
				},
				Value: parser.MustParseTime("1970-01-01T00:00:00Z"),
			},
		}, &ast.VariableAssignment{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 41,
						Line:   4,
					},
					File:   "",
					Source: "minTime = 1677-09-21T00:12:43.145224194Z",
					Start: ast.Position{
						Column: 1,
						Line:   4,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 8,
							Line:   4,
						},
						File:   "",
						Source: "minTime",
						Start: ast.Position{
							Column: 1,
							Line:   4,
						},
					},
				},
				Name: "minTime",
			},
			Init: &ast.DateTimeLiteral{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 41,
							Line:   4,
						},
						File:   "",
						Source: "1677-09-21T00:12:43.145224194Z",
						Start: ast.Position{
							Column: 11,
							Line:   4,
						},
					},
				},
				Value: parser.MustParseTime("1677-09-21T00:12:43.145224194Z"),
			},
		}, &ast.VariableAssignment{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 41,
						Line:   5,
					},
					File:   "",
					Source: "maxTime = 2262-04-11T23:47:16.854775806Z",
					Start: ast.Position{
						Column: 1,
						Line:   5,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 8,
							Line:   5,
						},
						File:   "",
						Source: "maxTime",
						Start: ast.Position{
							Column: 1,
							Line:   5,
						},
					},
				},
				Name: "maxTime",
			},
			Init: &ast.DateTimeLiteral{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 41,
							Line:   5,
						},
						File:   "",
						Source: "2262-04-11T23:47:16.854775806Z",
						Start: ast.Position{
							Column: 11,
							Line:   5,
						},
					},
				},
				Value: parser.MustParseTime("2262-04-11T23:47:16.854775806Z"),
			},
		}},
		Imports:  nil,
		Metadata: "parser-type=rust",
		Name:     "influxql.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 17,
						Line:   1,
					},
					File:   "",
					Source: "package influxql",
					Start: ast.Position{
						Column: 1,
						Line:   1,
					},
				},
			},
			Name: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 17,
							Line:   1,
						},
						File:   "",
						Source: "influxql",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "influxql",
			},
		},
	}},
	Package: "influxql",
	Path:    "internal/influxql",
}
