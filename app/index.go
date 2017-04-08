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
	if page <= 0 {
		page = 1
	}
	Print(uint(page))
	post.Limit(0, 20)
	post.Page(uint(page))
	ret := post.Find(post, "", []interface{}{})

	for posti, postv := range ret {
		postvv := postv.(map[string]interface{})
		ptitle := postvv["title"].(string)
		pcontent := postvv["content"].(string)
		postvv["title"] = sbstring.Substr(ptitle, 0, ConfTitleNumber)
		postvv["content"] = (sbtext.DeScript(sbstring.Substr(sbtext.DeHtml(pcontent), 0, ConfContentNumber)))
		ret[posti] = postvv
	}

	//Print(ret)

	Print(Setting("all"))
	c.HTML(http.StatusOK, "index.html", gin.H{
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
