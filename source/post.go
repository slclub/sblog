package source

import (
	"fmt"
	"sblog/core/model"
	"sblog/db"
	"strings"
)

var print = fmt.Println

type Post struct {
	*model.Model
	p_id          int
	c_id          int
	uid           int
	tags          string
	title         string
	sort          int
	content       string
	created_time  int
	modified_time int
}

func (post *Post) GetAttr() (result []string) {

	result = post.Model.GetAttr()
	ret := []string{"p_id", "c_id", "uid", "tags", "title", "sort", "content", "created_time", "modified_time"}
	result = append(result, ret...)
	return
}

func (post *Post) GetSource(args ...string) string {
	return "s_post"
}

func (post *Post) Assign(param map[string]interface{}) {
	post.Object = param
	//var okk interface{}
	post.p_id, _ = post.GetInterValue("p_id", "int").(int)
	post.c_id, _ = post.GetInterValue("c_id", "int").(int)
	post.uid, _ = post.GetInterValue("uid", "int").(int)
	post.tags, _ = post.GetInterValue("tags", "string").(string)
	post.title, _ = post.GetInterValue("title", "string").(string)
	post.content, _ = post.GetInterValue("content", "string").(string)
	post.sort, _ = post.GetInterValue("sort", "int").(int)
	post.created_time, _ = post.GetInterValue("created_time", "int", "created_time").(int)
	post.modified_time, _ = post.GetInterValue("modified_time", "int", "modified_time").(int)

	post.ID(int(post.p_id))
}

func (post *Post) Save(value model.Modeli) {

	id := value.ID()
	if id <= 0 {
		value.ID(post.p_id)
	}

	post.Model.Save(value)
}

func (post *Post) Create(value model.Modeli) (int, error) {
	DB, _ := db.Open()
	defer DB.Close()
	attrs := post.GetAttr()
	values := post.GetObjectValues(post)
	print(post.Object, values)
	_, err := DB.Exec("INSERT INTO "+post.GetSource()+" ("+strings.Join(post.GetAttr(), ",")+") values("+model.GetSqlPlaceholder(len(attrs))+")", values...)
	if err != nil {
		panic(err.Error() + "INsert error")
	}
	return 1, nil
}

func (post *Post) Update(value model.Modeli, args ...interface{}) {
}
