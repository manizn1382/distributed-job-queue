package setting

import(
	"sync"
)
var Min = 10
var Max = 100

var Key = "laklakzardzibaast"

var StrLen = 9

var ProducerCount = 10

var WorkerCount = 4

var BrokerChannel = make(chan any , ProducerCount)

var JobId = 0

var Mu sync.Mutex

var ProducerEndPoint string = "api/job"

var ProducerEndPointAddr = "http://localhost:8080/" + (ProducerEndPoint) 

var ContentType = "application/json"