package model

import(
	"math/rand"
	"jobQ/setting"
)

type Producer interface{
	GetNumber() int
	GetString() string
}



type NumberProducer struct{
	
}

type StringProducer struct{

}


func (np NumberProducer) GetNumber() int{
	return rand.Intn(setting.Max - setting.Min) + setting.Min
}

func (sp StringProducer) GetString() string{
	
	b := make([]byte, setting.StrLen)

	for i := range b {
		b[i] = setting.Key[rand.Intn(len(setting.Key))]
	}

	return string(b)
}

