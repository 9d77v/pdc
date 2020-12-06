package base

import (
	"context"
	"strings"

	"github.com/99designs/gqlgen/graphql"
)

//GraphQLField ..
type GraphQLField struct {
	FieldMap map[string]bool
	Fields   []string
}

//NewGraphQLField get graphql field from context
func NewGraphQLField(ctx context.Context, prefix string) GraphQLField {
	field := GraphQLField{}
	if prefix == "" {
		field.FieldMap = getFieldMap(ctx)
	} else {
		field.FieldMap = getFieldMapWithPrefix(ctx, prefix)
	}
	field.Fields = getFields(field.FieldMap)
	return field
}

func getFieldMap(ctx context.Context) map[string]bool {
	fieldMap := make(map[string]bool)
	fields := graphql.CollectAllFields(ctx)
	for _, v := range fields {
		fieldMap[v] = true
	}
	return fieldMap
}

func getFieldMapWithPrefix(ctx context.Context, prefix string) map[string]bool {
	fieldMap := make(map[string]bool)
	fields := getPreloads(ctx)
	for _, v := range fields {
		if strings.HasPrefix(v, prefix) {
			trimStr := strings.TrimPrefix(v, prefix)
			trimArr := strings.Split(trimStr, ".")
			fieldMap[trimArr[0]] = true
		}
	}
	return fieldMap
}

func getFields(fieldMap map[string]bool) []string {
	fields := make([]string, 0)
	for k := range fieldMap {
		fields = append(fields, k)
	}
	return fields
}

//getPreloads ..
func getPreloads(ctx context.Context) []string {
	return getNestedPreloads(
		graphql.GetRequestContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}

func getNestedPreloads(ctx *graphql.RequestContext, fields []graphql.CollectedField, prefix string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := getPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, getNestedPreloads(ctx, graphql.CollectFields(ctx, column.SelectionSet, nil), prefixColumn)...)
		preloads = append(preloads, getNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)

	}
	return
}

func getPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}
	return name
}
