package utils

import (
	"context"
	"fmt"
	"strings"
	"unicode"

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

//CamelToSnack 驼峰转蛇形
func CamelToSnack(s string) string {
	newStr := ""
	for i := 0; i < len(s); i++ {
		if unicode.IsUpper(rune(s[i])) {
			newStr += "_" + strings.ToLower(string(s[i]))
		} else {
			newStr += string(s[i])
		}
	}
	newStr = strings.ReplaceAll(newStr, "_i_d", "_id")
	return strings.ReplaceAll(newStr, "_u_r_l", "_url")
}

//ToDBFields ..
func ToDBFields(fields []string, omitFields ...string) []string {
	dbFields := make([]string, 0)
	omitFieldMap := make(map[string]bool)
	for _, v := range omitFields {
		omitFieldMap[v] = true
	}
	for _, v := range fields {
		if !omitFieldMap[v] {
			value := CamelToSnack(v)
			if strings.Contains(value, "price") {
				dbFields = append(dbFields, fmt.Sprintf("\"%s\"::money::numeric::float8", CamelToSnack(v)))
			} else if strings.Contains(value, ".") {
				dbFields = append(dbFields, CamelToSnack(v))
			} else {
				dbFields = append(dbFields, fmt.Sprintf("\"%s\"", CamelToSnack(v)))
			}
		}
	}
	return dbFields
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
