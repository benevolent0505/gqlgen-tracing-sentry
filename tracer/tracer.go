package tracer

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/getsentry/sentry-go"
)

type Tracer struct{}

var _ interface {
	graphql.HandlerExtension
	graphql.ResponseInterceptor
	graphql.FieldInterceptor
} = &Tracer{}

func (t Tracer) ExtensionName() string {
	return "SentryTracing"
}

func (t Tracer) Validate(graphql.ExecutableSchema) error {
	return nil
}

func (t Tracer) InterceptResponse(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	rc := graphql.GetOperationContext(ctx)

	span := sentry.StartTransaction(
		ctx,
		operatioName(rc),
		sentry.OpName("gql"),
	)
	defer span.Finish()

	span.SetData("request.query", rc.RawQuery)

	sctx := span.Context()

	return next(sctx)
}

func (t Tracer) InterceptField(ctx context.Context, next graphql.Resolver) (interface{}, error) {
	fc := graphql.GetFieldContext(ctx)

	span := sentry.StartSpan(
		ctx,
		"resolver",
	)
	defer span.Finish()

	span.Description = fc.Field.ObjectDefinition.Name + "." + fc.Field.Name

	span.SetData("resolver.path", fc.Path().String())
	span.SetData("resolver.object", fc.Field.ObjectDefinition.Name)
	span.SetData("resolver.field", fc.Field.Name)
	span.SetData("resolver.alias", fc.Field.Alias)

	sctx := span.Context()

	return next(sctx)
}

func operatioName(rc *graphql.OperationContext) string {
	requestName := "nameless-operation"
	if rc.Doc != nil && len(rc.Doc.Operations) != 0 {
		op := rc.Doc.Operations[0]
		if op.Name != "" {
			requestName = op.Name
		}
	}

	return requestName
}
