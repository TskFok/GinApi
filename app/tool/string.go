package tool

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//func Title(s string) {
//	for _, c := range []cases.Caser{
//		cases.Lower(language.Und),
//		cases.Upper(language.Turkish),
//		cases.Title(language.Dutch),
//		cases.Title(language.Und, cases.NoLower),
//	} {
//		fmt.Println(c.String(s))
//	}
//}

// Title /首字母转大写
func Title(s string) (res string) {
	for _, c := range []cases.Caser{
		cases.Title(language.Und, cases.NoLower),
	} {
		res = c.String(s)
	}

	return
}
