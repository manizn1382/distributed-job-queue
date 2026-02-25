package client

import (
	"jobQ/model"
	"jobQ/setting"
	"sync"
	"math/rand"
)

var numberJob = model.JobBuilder[int]{}
var stringJob = model.JobBuilder[string]{}

var numberBuilder = model.NumberProducer{} 
var stringBuilder = model.StringProducer{} 

func StartProduce(){

	var selectorMin = 1
	var selectorMax = 5


	for {

		jType := rand.Intn(selectorMax - selectorMin) + selectorMin

		if jType % 2 == 0 {

			setting.Mu.Lock()
			setting.JobId+=1
			job := numberJob.CreateJob(setting.JobId)
			job.SetId(setting.JobId)
			job.SetValue(numberBuilder.GetNumber())
			setting.Mu.Unlock()


		}

		if jType % 2 != 0 {

			setting.Mu.Lock()
			setting.JobId+=1
			job := stringJob.CreateJob(setting.JobId)
			job.SetId(setting.JobId)
			job.SetValue(stringBuilder.GetString())
			setting.Mu.Unlock()

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