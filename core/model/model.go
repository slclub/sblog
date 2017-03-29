package model

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

var print = fmt.Println

type Model struct {
	Object map[string]interface{}
	IDV    int
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
	return ret, nil
}

func (self *Model) Delete(delStr string, args ...interface{}) {
}

func (self *Model) Update(data Modeli, args ...interface{}) {
}

func (self *Model) Find(args ...interface{}) {
}

func (self *Model) Assgin(param map[string]interface{}) {
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

//SQL 语句占位符
func GetSqlPlaceholder(n int) (result string) {
	var nw [][]byte
	for i := 0; i < n; i++ {
		nw = append(nw, []byte("?"))
	}

	result = string(bytes.Join(nw, []byte(",")))
	return
}
