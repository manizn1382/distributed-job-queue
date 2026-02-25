package model


type Factory[T JobTypes] interface{
	CreateJob(int) Job[T] 
}


type JobBuilder[T JobTypes] struct{

}


func (b JobBuilder[T]) CreateJob(id int) Job[T] {
	return Job[T]{
		
	}
}