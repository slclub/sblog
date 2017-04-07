package model

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sblog/db"
	"strconv"
	"strings"
	"time"
)

const (
	PAGE_SIZE = 10
)

var Print = fmt.Println

type Model struct {
	Object map[string]interface{}
	IDV    int
	//changed attributes and values.
	ObjectUpdate map[string]interface{}
	LimitSql     map[string]uint
	OrderSql     string
	SelectAttr   []string
}

func NewModel() *Model {

	var Object = make(map[string]interface{})
	var ObjectUpdate = make(map[string]interface{})
	var limit = make(map[string]uint)

	return &Model{Object, 0, ObjectUpdate, limit, "", []string{}}
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

func (self *Model) Delete(data Modeli, args ...interface{}) (int, error) {
	if len(args) == 0 {
		//panic(errors.New("Incorrect args In model.Delete"))
		return 0, errors.New("Incorrect args In model.Delete")
	}

	where := ""
	var values = make([]interface{}, 0)
	switch len(args) {
	case 1:
		arg1, _ := args[0].(int)
		if id := (arg1); id > 0 {
			where = data.IDField("") + "=? "
			values = append(values, id)
		} else {
			return 0, errors.New("Incorrect id In Model.Delete")
		}

		break
	case 2:
		if vali, ok := args[1].([]interface{}); ok {
			values = vali
			where, _ = args[0].(string)
		} else {
			//panic(errors.New("Incorrect condition for delete operation"))
			return 0, errors.New("Incorrect condition for delete operation")
		}

	}

	sql := " DELETE FROM " + data.GetSource() + " WHERE " + where

	DB, _ := db.Open()
	defer DB.Close()

	res, err := DB.Exec(sql, values...)
	if err != nil {
		return 0, err
	}
	rowCnt, err := res.RowsAffected()

	return int(rowCnt), err
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
		//panic("the record updated was not existed")
		return 0, errors.New("the record updated was not existed")
	}

	vali := 0
	for index, val := range objectUpdate {

		if index == "created_time" {
			continue
		}

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
		return 0, errors.New(update_err.Error() + "UPDATE error")
	}
	sqlRes, err := update_stmt.Exec(values...)
	if err != nil {
		return 0, errors.New(update_err.Error() + "UPDATE error")
	}
	effectRows, err := sqlRes.RowsAffected()

	if err != nil {
		return 0, errors.New(update_err.Error() + "UPDATE error")
	}

	return int(effectRows), nil
}

func (self *Model) Fields(args ...string) (ret []string) {
	if len(args) > 0 {
		self.SelectAttr = args
	}
	ret = self.SelectAttr
	return
}

func (self *Model) Find(data Modeli, where string, bindArr []interface{}) []interface{} {
	DB, _ := db.Open()
	defer DB.Close()

	id := data.ID()

	attrs := data.Fields()
	if len(attrs) == 0 {
		attrs = data.GetAttr()
	}
	lenAttr := len(attrs)

	where = self.Where(where)

	if len(bindArr) == 0 {
		bindArr = append(bindArr, id)
	}
	if where == "" || bindArr == nil {
		where = " " + data.IDField("") + ">?"
	} else {
		where = where
	}
	sel := strings.Join(attrs, ",")

	if (self.LimitSql["limit"]) == 0 {
		self.LimitSql["limit"] = PAGE_SIZE
	}
	limit := " limit " + strconv.Itoa(int(self.LimitSql["offset"])) + "," + strconv.Itoa(int(self.LimitSql["limit"]))
	data.Order()
	order := self.OrderSql

	sql := "select " + sel + " from " + data.GetSource() + " WHERE " + where + " " + order + limit

	rows, err := DB.Query(sql, bindArr...)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	ret := make([]interface{}, 0)
	cols, err := rows.Columns()
	getVals := make([]interface{}, len(cols))
	getValsAddr := make([]interface{}, len(cols))

	for i := 0; i < len(getVals); i++ {
		addr := new(interface{})
		getVals[i] = addr
		getValsAddr[i] = &getVals[i]
	}
	for rows.Next() {
		err := rows.Scan(getValsAddr...)
		dest := make(map[string]interface{})
		for i := 0; i < lenAttr; i++ {
			dest[attrs[i]] = getVals[i]
		}
		data.DataDecode(dest)
		if err != nil {
			log.Fatal(err)
			continue
		}

		ret = append(ret, dest)
	}

	return ret
}

func (self *Model) FindOne(data Modeli, args ...interface{}) (ret map[string]interface{}) {

	where := ""
	bindArr := []interface{}{}
	switch len(args) {
	case 0:
		if data.ID() == 0 {
			return
		} else {
			where = data.IDField("") + "=? "
			bindArr = append(bindArr, data.ID())
		}
	case 1:
		where = data.IDField("") + "=? "
		bindArr[0] = append(bindArr, args[0].(int))
	case 2:
		where = args[0].(string)
		bindArr = args[1].([]interface{})
	}

	self.Limit(0, 1)
	s := self.Find(data, where, bindArr)
	if len(s) > 0 {
		ret = s[0].(map[string]interface{})
	}
	return
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
		var vall *interface{}
		vall = values[i].(*interface{})
		res[attrs[i]] = *vall
	}

	data.DataDecode(res)
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
		vint, ok := self.Object[val].(int)
		if ok {
			self.Object[val] = vint
			return vint
		}

		vint64, ok := self.Object[val].(int64)
		if ok {
			self.Object[val] = vint64
			return vint64
		}

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

func (self *Model) Where(where string) string {
	var ret string

	ret = where
	return ret
}

func (self *Model) Limit(offset, limit uint) Modeli {
	if limit == 0 {
		limit = PAGE_SIZE
	}
	self.LimitSql["offset"] = offset
	self.LimitSql["limit"] = limit
	return self
}

func (self *Model) Page(page uint, args ...uint) {
	var limit uint = 0
	if len(args) == 0 {
		if self.LimitSql["limit"] == 0 {
			limit = PAGE_SIZE
			self.LimitSql["limit"] = PAGE_SIZE
		} else {
			limit = self.LimitSql["limit"]
		}
	} else {
		limit = args[0]
	}

	if page <= 0 {
		page = 1
	}

	self.LimitSql["offset"] = (page - 1) * limit
}

func (self *Model) Order() Modeli {
	self.OrderSql = ""
	return self
}

func (self *Model) DataDecode(convertData interface{}) error {
	dest, ok := convertData.(map[string]interface{})

	if ok == false {
		return nil
	}

	for index, val := range dest {
		//vall := val.(**interface{})
		if index == "created_time" || index == "modified_time" {
			vint64 := (val).(int64)
			dest[index] = time.Unix(vint64, 0).Format("2006-01-02 03:04:05 PM")
			continue
		}
	}

	return nil
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
