package model

import(
	"jobQ/setting"
	"math/rand"
)

type Job[T AllowedTypes] struct{
	id int
	value T
}


func (j Job[T]) GetId() int{
	return j.id
}


func (j Job[T]) GetValue() T{
	return j.value
}

func (j *Job[T]) SetId(id int){
	 j.id = id
}


func (j *Job[T]) SetValue(value T){
	j.value = value
}


func (j *Job[T]) Produce() T {
	return rand.Intn(setting.Max - setting.Min) + setting.Min
}