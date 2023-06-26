package main

import "fmt"

func interface2String(inter interface{}) string {

	switch inter.(type) {

	case string:
		return fmt.Sprint(inter.(string))
	case int:
		return string(inter.(int))
	case float64:
		return fmt.Sprint(inter.(float64))
	}
	return ""
}
