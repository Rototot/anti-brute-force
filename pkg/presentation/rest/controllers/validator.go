package controllers

type StructValidator interface {
	Struct(s interface{}) error
}
