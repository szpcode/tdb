package tdb

import(
	"strings"
	"strconv"
)

var varMapStr = map[string][]string{}
var varStr = make(map[string]string)
var varMapInt = map[string][]int{}
var varInt = make(map[string]int)

func escapeString(str string) string {
	str = strings.Replace(str , "\"", "\\\"", -1)
	str = strings.Replace(str, "'", "\\'", -1)
	return str
}

func clearData() {
    if len(varStr) > 0 {
        for k := range varStr {
            delete(varStr, k)
        }
    }
  
    if len(varInt) > 0 {
        for k := range varInt {
            delete(varInt, k)
        }
    }
  
    if len(varMapInt) > 0 {
        for k := range varMapInt {
            delete(varMapInt, k)
        }
    }

    if len(varMapStr) > 0 {
        for k := range varMapStr {
            delete(varMapStr, k)
        }
    }
}

func Prepare(sql string) string {

	for k, v := range varStr { 
		sql = strings.Replace(sql, k, "\""+ escapeString(v) +"\"", -1)
	}

	for k, v := range varMapStr { 
		params := make([]string, 0, len(v))
		for _, value := range v {
			params = append(params, escapeString(value))
		}
		sql = strings.Replace(sql, k, "\""+ strings.Join(params, "\", \"") +"\"", -1)
	}
	
	for k, v := range varInt { 
		sql = strings.Replace(sql, k, strconv.Itoa(v), -1)
	}

	for k, v := range varMapInt {
		params := make([]string, 0, len(v))
		for _, value := range v {
			params = append(params, strconv.Itoa(value))
		}	
		sql = strings.Replace(sql, k, strings.Join(params, ", "), -1) 
	}
  clearData()
	return sql;
}

func BindStrValue(name string, value string) {
	varStr[name] = value
}

func BindStrValues(name string, value []string) {
	varMapStr[name] = value
}


func BindIntValue(name string, value int) {
	varInt[name] = value
}

func BindIntValues(name string, value []int) {
	varMapInt[name] = value
}