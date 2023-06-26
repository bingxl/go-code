package main

import "flag"

// FlagT 单个参数类型
type FlagT struct {
	Name, DefVal, Usage string
}

// Parse 获取参数
func Parse(f []FlagT) map[string]string {
	tmp := make(map[string]*string)

	for i := 0; i < len(f); i++ {
		name, defVal, usage := f[i].Name, f[i].DefVal, f[i].Usage
		tmp[name] = flag.String(name, defVal, usage)
	}

	flag.Parse()

	result := make(map[string]string)
	for k := range tmp {
		result[k] = *tmp[k]
		delete(tmp, k)
	}
	return result
}
