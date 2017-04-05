package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(text *string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(*text))
	*text = hex.EncodeToString(md5Ctx.Sum(nil))
	return *text
}
