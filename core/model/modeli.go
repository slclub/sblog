package model

type Modeli interface {
	GetSource(string)

	GetAttr()

	Save(*Modeli)

	Create(*Modeli)

	Delete(string, ...interface{})

	Update(string, ...interface{})

	Find(...interface{})
}
