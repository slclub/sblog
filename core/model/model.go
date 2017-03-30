package model

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"sblog/db"
	"strconv"
	"strings"
	"time"
)

var print = fmt.Println

type Model struct {
	Object map[string]interface{}
	IDV    int
	//changed attributes and values.
	ObjectUpdate map[string]interface{}
}

func (self *Model) GetSource(args ...string) string {
	if len(args) <= 0 {
		return ""
	}

	table := args[0]
	if table == "" {
		return ""
	}
	return table
}

func (self *Model) GetAttr() (result []string) {
	return
}

func (self *Model) Save(data Modeli) {
	if data.ID() <= 0 {
		data.Create(data)
	} else {
		data.Update(data)
	}
}

func (self *Model) Create(data Modeli) (ret int, err error) {
	DB, _ := db.Open()
	defer DB.Close()
	attrs := data.GetAttr()
	values := data.GetObjectValues(data)

	sqlRes, err := DB.Exec("INSERT INTO "+data.GetSource()+" ("+strings.Join(data.GetAttr(), ",")+") values("+GetSqlPlaceholder(len(attrs))+")", values...)
	if err != nil {
		panic(err.Error() + "INsert error")
	}
	lastId, errr := sqlRes.LastInsertId()

	if errr != nil {
		panic(err.Error() + "INsert error")
	}

	return int(lastId), nil
}

func (self *Model) Delete(delStr string, args ...interface{}) {
}

func (self *Model) Update(data Modeli, args ...interface{}) (ret int, err error) {
	DB, _ := db.Open()
	defer DB.Close()
	//attrs := data.GetAttr()
	//values := data.GetObjectValues(data)
	objectUpdate := data.GetObjectUpdate()
	var values = make([]interface{}, len(objectUpdate))

	sql := "UPDATE " + data.GetSource() + " SET "
	if len(objectUpdate) <= 0 {
		return 0, errors.New("ObjectUpdate is empty")
	}

	if errExists := data.Exists(data); errExists == nil {
		panic("the record updated was not existed")
	}

	vali := 0
	for index, val := range objectUpdate {
		if nil == val {
			val = ""
		}
		values[vali] = val
		vali++
		sql += " " + index + "=?,"
	}

	values = append(values, data.ID())
	sql = strings.TrimRight(sql, ",")
	sql += " WHERE " + data.IDField("") + " =?"
	update_stmt, update_err := DB.Prepare(sql)
	if update_err != nil {
		panic(update_err.Error() + "UPDATE error")
	}
	sqlRes, err := update_stmt.Exec(values...)
	if err != nil {
		panic(err.Error() + "UPDATE error")
	}
	effectRows, errr := sqlRes.RowsAffected()

	if errr != nil {
		panic(errr.Error() + "UPDATE error")
	}

	return int(effectRows), nil
}

func (self *Model) Find(args ...interface{}) {
}

func (self *Model) FindOne(args ...interface{}) {
	if len(args) <= 0 {
	}
}

func (self *Model) Exists(data Modeli) map[string]interface{} {
	if data.ID() <= 0 {
		return nil
	}
	DB, _ := db.Open()
	defer DB.Close()
	id := data.ID()
	attrs := data.GetAttr()
	lenAttr := len(attrs)
	sel := strings.Join(attrs, ",")
	sql := "select " + sel + " from " + data.GetSource() + " WHERE " + data.IDField("") + " = ?"

	var res = make(map[string]interface{})
	var values = make([]interface{}, lenAttr)
	for i := 0; i < lenAttr; i++ {
		values[i] = new(interface{})
	}
	select_err := DB.QueryRow(sql, id).Scan(values...)

	if select_err != nil {
		panic("select error exists " + select_err.Error())
	}
	for i := 0; i < lenAttr; i++ {
		res[attrs[i]] = values[i]
	}

	return res
}

func (self *Model) Assgin(param map[string]interface{}) {
	paramStr, err := json.Marshal(param)
	if err != nil {
		panic("core/model Assgin function json marshal error")
	}
	json.Unmarshal(paramStr, &self.ObjectUpdate)
}

func (self *Model) GetObjectValues(m Modeli) (ret []interface{}) {

	attrs := m.GetAttr()
	al := len(attrs)

	for i := 0; i < al; i++ {
		val := self.Object[attrs[i]]
		ret = append(ret, val)
	}

	return
}

func (self *Model) GetObject() map[string]interface{} {
	return self.Object
}

func (self *Model) ID(args ...int) int {
	if len(args) > 0 && args[0] > 0 {
		self.IDV = args[0]
	}

	return (self.IDV)
}

func (self *Model) GetInterValue(val string, typev string, args ...string) interface{} {

	fld := ""
	if len(args) >= 1 {
		fld = args[0]
	}

	switch typev {
	case "int":
		vs, ok := self.Object[val].(string)
		v, _ := strconv.Atoi(vs)
		if !ok {
			v = 0
			var vv int64

			if fld == "created_time" {
				vv = time.Now().Unix()
			}
			if fld == "modified_time" {
				vv = time.Now().Unix()
			}
			v = int(vv)
		}
		self.Object[val] = v
		return v
	case "string":
		v, ok := self.Object[val].(string)
		if ok == false {
			v = ""
		}
		self.Object[val] = v
	}

	return ""
}

func (self *Model) GetObjectUpdate() map[string]interface{} {
	return self.ObjectUpdate
}

func (self *Model) IDField(fld string) string {
	if fld == "" {
		fld = "id"
	}
	return fld
}

//SQL 语句占位符
func GetSqlPlaceholder(n int) (result string) {
	var nw [][]byte
	for i := 0; i < n; i++ {
		nw = append(nw, []byte("?"))
	}

	result = string(bytes.Join(nw, []byte(",")))
	return
}
