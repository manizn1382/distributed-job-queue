package model


type Factory[T AllowedTypes] interface{
	CreateJob(int) Job[T] 
}