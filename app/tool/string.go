package tool

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
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

func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

func UUID() string {
	uuidWithHyphen := uuid.New()

	return strings.Replace(uuidWithHyphen.String(), "-", "", -1)
}

func Password(password string, salt string) string {
	builder := strings.Builder{}
	builder.WriteString(password)
	builder.WriteString(salt)

	after := builder.String()
	after = MD5(after)

	after = string([]rune(after)[0:30])

	return after
}
