package my

import (
	"regexp"
	"os"
	"fmt"
)

//すべての正規表現殿完全一致判定 (文字列,正規表現配列) 
//f=true すべて含む　f=false 一つも含まない
func MatchAll(str string, arr []*regexp.Regexp, f bool) bool {
	for _, v := range arr {
		flg := v.MatchString(str)
		if f{
			flg = !flg
		}
		if flg{
			return false
		}
	}
	return true
}

func CompileAll(s []string) []*regexp.Regexp{
	regs := []*regexp.Regexp{}
	for _, v := range s{
		if v == ""{
			continue
		}
		reg, err := regexp.Compile(v)
		if err != nil{
			fmt.Println("正規表現エラー:", err)
			os.Exit(1)
		}
		regs = append(regs, reg)
	}
	return regs
}