package model

import(
	"math/rand"
	"jobQ/setting"
)


type StringJob struct{
	id int
	value string
}



func (sj StringJob) GetId() int{
	return sj.id
}


func (sj StringJob) GetValue() string{
	return sj.value
}

func (sj *StringJob) SetId(id int){
	 sj.id = id
}


func (sj *StringJob) SetValue(value string){
	sj.value = value
}


func (j StringJob) Produce() string{


	b := make([]byte, setting.StrLen)

	for i := range b {
		b[i] = setting.Key[rand.Intn(len(setting.Key))]
	}

	return string(b)

}

