package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"jobQ/model"
	"jobQ/setting"
	"math/rand"
	"net/http"
	"sync"
)

type SenderClass struct{}

var numberJob = model.JobBuilder[int]{}
var stringJob = model.JobBuilder[string]{}

var numberBuilder = model.NumberProducer{}
var stringBuilder = model.StringProducer{}

var sender = SenderClass{}




func (s SenderClass) sendJob(j any, Type int) {

	if Type%2 == 0 {
		j, _ = j.(model.Job[int])
	}

	if Type%2 != 0 {
		j, _ = j.(model.Job[string])
	}

	jData, err := json.Marshal(j)
	if err != nil {
		fmt.Println("can't convert job to json : ", j, " : ", err)
		return
	}
	
	resp, err := http.Post(setting.ProducerEndPointAddr, setting.ContentType, bytes.NewBuffer(jData))
	
	if err != nil {
		fmt.Println("Error sending job: ",string(jData)," to server. error: ", err)
		return
	}

	defer resp.Body.Close()

}

func StartProduce() {

	var selectorMin = 1
	var selectorMax = 5

	for {

		jType := rand.Intn(selectorMax-selectorMin) + selectorMin

		if jType%2 == 0 {

			setting.Mu.Lock()
			setting.JobId += 1
			job := numberJob.CreateJob(setting.JobId)
			job.SetId(setting.JobId)
			job.SetValue(numberBuilder.GetNumber())
			sender.sendJob(job, jType)
			setting.Mu.Unlock()

		}

		if jType%2 != 0 {

			setting.Mu.Lock()
			setting.JobId += 1
			job := stringJob.CreateJob(setting.JobId)
			job.SetId(setting.JobId)
			job.SetValue(stringBuilder.GetString())
			sender.sendJob(job, jType)
			setting.Mu.Unlock()
		}
	}

}

func main() {

	var wg sync.WaitGroup

	for range setting.ProducerCount {

		wg.Add(1)
		go StartProduce()

	}

}
