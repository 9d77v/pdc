package utils

import (
	"context"
	"strings"

	"github.com/99designs/gqlgen/graphql"
)

//GetFieldData 获取字段map和数组
func GetFieldData(ctx context.Context, prefix string) (map[string]bool, []string) {
	fieldMap := make(map[string]bool)
	if prefix == "" {
		fields := graphql.CollectAllFields(ctx)
		for _, v := range fields {
			fieldMap[v] = true
		}
	} else {
		fields := getPreloads(ctx)
		for _, v := range fields {
			if strings.HasPrefix(v, prefix) {
				trimStr := strings.TrimPrefix(v, prefix)
				trimArr := strings.Split(trimStr, ".")
				fieldMap[trimArr[0]] = true
			}
		}
	}
	fields := make([]string, 0)
	for k := range fieldMap {
		fields = append(fields, k)
	}
	return fieldMap, fields
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
