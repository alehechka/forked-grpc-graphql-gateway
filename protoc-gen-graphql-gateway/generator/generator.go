package generator

import (
	"errors"
	_ "log"
	"strings"

	"path/filepath"

	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/ysugimoto/grpc-graphql-gateway/graphql"
	"github.com/ysugimoto/grpc-graphql-gateway/protoc-gen-graphql-gateway/builder"
	ext "github.com/ysugimoto/grpc-graphql-gateway/protoc-gen-graphql-gateway/extension"
	"github.com/ysugimoto/grpc-graphql-gateway/protoc-gen-graphql-gateway/format"
	"github.com/ysugimoto/grpc-graphql-gateway/protoc-gen-graphql-gateway/types"
)

type Generator struct {
	req *plugin.CodeGeneratorRequest
	r   *Resolver

	queries   format.Queries
	mutations format.Mutations
	types     format.Types
}

func New(req *plugin.CodeGeneratorRequest) *Generator {
	return &Generator{
		req: req,
		r:   NewResolver(req),

		queries:   format.Queries{},
		mutations: format.Mutations{},
		types:     format.Types{},
	}
}

func (g *Generator) Generate(resp *plugin.CodeGeneratorResponse) {
	var genError error
	defer func() {
		if genError != nil {
			msg := genError.Error()
			resp.Error = &msg
		}
	}()

	args := &types.Params{}
	if g.req.Parameter != nil {
		args, genError = types.NewParams(g.req.GetParameter())
		if genError != nil {
			return
		}
	}

	for _, f := range g.req.GetProtoFile() {
		pkg := f.GetPackage()
		for _, s := range f.GetService() {
			for _, m := range s.GetMethod() {
				if opt := ext.GraphqlQueryOption(m); opt != nil {
					qs, err := g.AnalyzeQuery(m, opt)
					if err != nil {
						genError = err
						return
					}
					g.queries.Add(pkg, qs)
				}
				if opt := ext.GraphqlMutationOption(m); opt != nil {
					ms, err := g.AnalyzeMutation(m, opt)
					if err != nil {
						genError = err
						return
					}
					g.mutations.Add(pkg, ms)
				}
			}
		}
	}

	if len(g.queries) == 0 {
		genError = errors.New("nothing to generate queries")
		return
	}

	queries := g.queries.Concat()
	mutations := g.mutations.Concat()
	types, err := g.r.ResolveTypes(queries, mutations)
	if err != nil {
		genError = err
		return
	}

	if args.QueryOut != "" {
		builders := []builder.Builder{
			builder.NewQuery(queries),
			builder.NewMutation(mutations),
		}
		builders = append(builders, types...)
		schema := format.NewSchema(builders)
		resp.File = append(
			resp.File,
			schema.Format(filepath.Join(args.QueryOut, "./query.graphql")),
		)
	}

	if args.ProgramOut != "" {
		builders := []builder.Builder{
			builder.NewPackage(args.ProgramPackage),
			builder.NewImport(queries, mutations),
		}
		builders = append(builders, types...)
		builders = append(
			builders,
			builder.NewQuery(queries),
			builder.NewMutation(mutations),
			builder.NewHandler(),
		)
		program := format.NewProgram(builders)
		resp.File = append(
			resp.File,
			program.Format(
				filepath.Join(args.ProgramOut, "./app/query.grahpql.go"),
			),
		)
	}
}

func (g *Generator) AnalyzeQuery(
	m *descriptor.MethodDescriptorProto,
	opt *graphql.GraphqlQuery,
) (*types.QuerySpec, error) {
	var req, resp *types.Message
	if req = g.r.FindMessage(strings.TrimPrefix(m.GetInputType(), ".")); req == nil {
		return nil, errors.New("InputType: " + m.GetInputType() + " not exists")
	}
	if resp = g.r.FindMessage(strings.TrimPrefix(m.GetOutputType(), ".")); resp == nil {
		return nil, errors.New("OutputType: " + m.GetOutputType() + " not exists")
	}

	return &types.QuerySpec{
		Input:  req,
		Output: resp,
		Option: opt,
	}, nil
}

func (g *Generator) AnalyzeMutation(
	m *descriptor.MethodDescriptorProto,
	opt *graphql.GraphqlMutation,
) (*types.MutationSpec, error) {
	var req, resp *types.Message
	if req = g.r.FindMessage(strings.TrimPrefix(m.GetInputType(), ".")); req == nil {
		return nil, errors.New("InputType: " + m.GetInputType() + " not exists")
	}
	if resp = g.r.FindMessage(strings.TrimPrefix(m.GetOutputType(), ".")); resp == nil {
		return nil, errors.New("OutputType: " + m.GetOutputType() + " not exists")
	}

	return &types.MutationSpec{
		Input:  req,
		Output: resp,
		Option: opt,
	}, nil
}