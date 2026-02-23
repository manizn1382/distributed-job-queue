package model

import (
	"math/rand"
	"jobQ/setting"
)


type IntegerJob struct{
	id int
	value int
}


func (ij IntegerJob) GetId() int{
	return ij.id
}


func (ij IntegerJob) GetValue() int{
	return ij.value
}

func (ij *IntegerJob) SetId(id int){
	 ij.id = id
}


func (ij *IntegerJob) SetValue(value int){
	ij.value = value
}


func (j IntegerJob) Produce() int{
	return rand.Intn(setting.Max - setting.Min) + setting.Min
}