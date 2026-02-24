package model


type JobBuilder[T AllowedTypes] struct{

}


func (b JobBuilder[T]) CreateJob(id int) Job[T] {
	return Job[T]{
		id:    id,
	}
}