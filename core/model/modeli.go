package model

type Modeli interface {
	GetSource(...string) string

	GetAttr() []string

	Save(Modeli)

	Create(Modeli) (int, error)

	Delete(string, ...interface{})

	Update(Modeli, ...interface{}) (int, error)

	Find(Modeli, string, []interface{}) []interface{}

	ID(...int) int
	IDField(string) string

	GetObjectValues(Modeli) []interface{}
	GetObject() map[string]interface{}
	GetObjectUpdate() map[string]interface{}
	Exists(Modeli) map[string]interface{}
	DataDecode(interface{}) error
}
