package utils

import (
	"fmt"
	"strings"
)

// generateQuery generates an insert query based on the table name and parameters.
func generateInsertQuery(table string, params map[string]interface{}) (string, []interface{}) {
	columns, values, par := unpackMapForQueryInsert(params)
	s := fmt.Sprintf("insert into %s (%s) values (%s)", table, columns, values)
	s = strings.ReplaceAll(s, `"`, `'`)
	return s, par
}

// generateQuery generates an insert query based on the table name and parameters.
func generateUpdateQuery(table string, params map[string]interface{}, condition string) (string, []interface{}) {
	par, _ := generateQueryUpdate(params)
	s := fmt.Sprintf("update %s set %s where %s", table, par, condition)
	s = strings.ReplaceAll(s, `"`, `'`)
	return s, nil
}

// unpackMapForQueryInsert unpacks a map into separate column names and values for an insert query.
func unpackMapForQueryInsert(params map[string]interface{}) (string, string, []interface{}) {
	var columns []string
	var values []string
	var par []interface{}

	for key, val := range params {
		columns = append(columns, key)
		values = append(values, fmt.Sprintf(":%s", key))
		par = append(par, val)
	}

	return strings.Join(columns, ", "), strings.Join(values, ", "), par
}

func generateQueryUpdate(params map[string]interface{}) (string, map[string]interface{}) {
	pars := utfDictToString(params)
	setval, par := unpackDictForQueryUpdate(pars)
	return setval, par
}

func utfDictToString(data map[string]interface{}) map[string]interface{} {
	convertedData := make(map[string]interface{})
	for k, v := range data {
		convertedData[k] = fmt.Sprintf("%v", v)
	}
	return convertedData
}

func unpackDictForQueryUpdate(params map[string]interface{}) (string, map[string]interface{}) {
	var v []string
	par := make(map[string]interface{})

	for p := range params {
		if value, ok := params[p].(map[string]interface{}); ok {
			par[p] = value["val"]
			nameFunct := value["name_funct"].(string)
			if nameFunct == "to_date" {
				v = append(v, fmt.Sprintf("%s = to_date(:%s, '%s')", p, p, value["format"]))
			} else {
				panic(fmt.Sprintf("name_funct %s tidak terdaftar", nameFunct))
			}
		} else if value, ok := params[p].([]interface{}); ok {
			v = append(v, fmt.Sprintf("%s = %v", p, value[0]))
		} else {
			par[p] = params[p]
			v = append(v, fmt.Sprintf("%s = :%s", p, p))
		}
	}

	setVal := fmt.Sprintf("%s", strings.Join(v, ", "))
	return setVal, par
}
