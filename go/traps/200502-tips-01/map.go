package main

type Param map[string]interface{}

type Show struct {
	Param
}

// https://zhuanlan.zhihu.com/p/35058068

func main() {
	s := new(Show)
	//s.Param = map[string]interface{}{}
	s.Param["RMB"] = 10000
}