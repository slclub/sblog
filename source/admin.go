package source

import (
	"sblog/core/model"
)

type Admin struct {
	*model.Model
	uid          int
	username     string
	password     string
	name         string
	email        string
	created_time int
}

func NewAdmin() *Admin {
	return &Admin{model.NewModel(), 0, "", "", "", "", 0}
}
func (admin *Admin) GetAttr() (result []string) {

	result = admin.Model.GetAttr()
	ret := []string{"uid", "username", "password", "name", "email", "created_time"}
	result = append(result, ret...)
	return
}

func (admin *Admin) GetSource(args ...string) string {
	return "s_admin"
}

func (object *Admin) Assign(param map[string]interface{}) {
	object.Model.Assgin(param)
	object.Object = param
	//var okk interface{}
	object.uid, _ = object.GetInterValue("uid", "int").(int)
	object.username, _ = object.GetInterValue("username", "string").(string)
	object.password, _ = object.GetInterValue("password", "string").(string)
	object.email, _ = object.GetInterValue("email", "string").(string)
	object.name, _ = object.GetInterValue("name", "int").(string)
	object.created_time, _ = object.GetInterValue("created_time", "int", "created_time").(int)

	object.ID(int(object.uid))
}

func (object *Admin) Save(value model.Modeli) {

	id := value.ID()
	if id <= 0 {
		value.ID(object.uid)
	}

	object.Model.Save(value)
}

func (object *Admin) Create(value model.Modeli) (int, error) {
	return object.Model.Create(value)

}

func (object *Admin) Update(value model.Modeli, args ...interface{}) (ret int, err error) {

	return object.Model.Update(value, args...)
}

func (object *Admin) IDField(fld string) string {
	return object.Model.IDField("uid")
}

func (object *Admin) Order() model.Modeli {
	object.OrderSql = " order by uid desc "
	return object
}

func (object *Admin) DataDecode(convertData interface{}) error {

	object.Model.DataDecode(convertData)

	dest, ok := convertData.(map[string]interface{})
	if ok == false {
		return nil
	}

	for index, val := range dest {
		if val == nil {
			continue
		}
		if index == "username" || index == "password" || index == "name" || index == "email" {
			val = string(val.([]uint8))
		}
		dest[index] = val
	}
	dest["ID"] = dest["uid"]
	return nil
}
