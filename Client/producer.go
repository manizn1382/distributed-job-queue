package client

import (
	"jobQ/model"
	"jobQ/setting"
	"sync"
	"math/rand"
)





func StartProduce(){

	var selectorMin = 1
	var selectorMax = 5


	for {

		jType := (rand.Intn(selectorMax - selectorMin) + selectorMin)

		if jType % 2 == 0 {

			job := model.IntegerJob{}
			job.SetId(setting.JobId)
			job.SetValue(job.Produce())

		}

		if jType % 2 != 0 {

			job := model.StringJob{}
			job.SetId(setting.JobId)
			job.SetValue(job.Produce())

		}
	}


}


func main(){

	var wg sync.WaitGroup

	for range setting.ProducerCount{

		wg.Add(1)
		go StartProduce()

	}



}