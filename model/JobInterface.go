package model

type JobInterface[T AllowedTypes] interface{
	GetId() int
	GetValue() T
	SetId(int)
	SetValue(T)
	Produce() T
}

