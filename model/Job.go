package model


type StrJobInterface interface{
	GetId() int
	GetValue() string
	SetId(int)
	SetValue(string)
	Produce()
}


type IntJobInterface interface{
	GetId() int
	GetValue() int
	SetId(int)
	SetValue(int)
	Produce()
}