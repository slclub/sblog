package source

import (
	//"errors"
	"fmt"
	"sblog/core/model"
	//"sblog/db"
	//"strings"
	"time"
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

func NewPost() *Post {
	var Object = make(map[string]interface{})
	var ObjectUpdate = make(map[string]interface{})
	var limit = make(map[string]uint)
	return &Post{&model.Model{Object, 0, ObjectUpdate, limit, ""}, 0, 0, 0, "", "", 0, "", 0, 0}
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
	post.Model.Assgin(param)
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
	return post.Model.Create(value)

}

func (post *Post) Update(value model.Modeli, args ...interface{}) (ret int, err error) {
	post.ObjectUpdate["modified_time"] = int(time.Now().Unix())

	if val, _ := post.ObjectUpdate["sort"]; val == 0 || val == "" || val == nil {
		delete(post.ObjectUpdate, "sort")
	}

	return post.Model.Update(value, args...)
}

func (post *Post) IDField(fld string) string {
	return post.Model.IDField("p_id")
}

func (post *Post) Order() model.Modeli {
	post.OrderSql = " order by sort desc,p_id desc "
	return post
}

func (post *Post) DataDecode(convertData interface{}) error {

	post.Model.DataDecode(convertData)

	dest, ok := convertData.(map[string]interface{})
	if ok == false {
		return nil
	}

	for index, val := range dest {
		if val == nil {
			continue
		}
		if index == "content" || index == "title" || index == "tags" {
			val = string(val.([]uint8))
		}
		dest[index] = val
	}
	dest["ID"] = dest["p_id"]
	return nil
}
