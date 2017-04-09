package app

import (
	"github.com/gin-gonic/gin"
	//"html/template"
	"net/http"
	sbstring "sblog/provider/string"
	sbtext "sblog/provider/text"
	"sblog/source"
	"strconv"
)

var ConfTitleNumber = 50
var ConfContentNumber = 200

var Index gin.HandlerFunc = func(c *gin.Context) {
	//c.JSON(200, gin.H{"iwas": "fine thank you"})
	post := source.NewPost()

	page, _ := strconv.Atoi(c.Query("page"))
	search := c.Query("search")

	fwhere, fbind := BindFind("search", search)

	if page <= 0 {
		page = 1
	}
	post.Limit(0, 20)

	post.Page(uint(page))
	ret := post.Find(post, fwhere, fbind)

	for posti, postv := range ret {
		postvv := postv.(map[string]interface{})
		ptitle := postvv["title"].(string)
		pcontent := postvv["content"].(string)
		postvv["title"] = sbstring.Substr(ptitle, 0, ConfTitleNumber)
		postvv["content"] = (sbtext.DeScript(sbstring.Substr(sbtext.DeHtml(pcontent), 0, ConfContentNumber)))
		ret[posti] = postvv
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"content":  "Hello world",
		"postList": ret,
		"prepage":  page - 1,
		"nextpage": page + 1,
		"setting":  Setting("all"),
	})
}

func init() {
	//r.Any("/", diFront.Di(index))
}

var Page = func() {
}

var BindFind = func(fld string, val interface{}) (ret string, bindArr []interface{}) {
	bindArr = make([]interface{}, 2)
	if fld == "search" && val != nil && val != "" {
		ret = " (title like ? or tags like ?) "
		valStr := val.(string)
		bindArr[0] = "%" + valStr + "%"
		bindArr[1] = "%" + valStr + "%"
		return ret, bindArr
	}
	return ret, bindArr[:0]
}
