package model

type Sender interface{
	sendJob(any,int)
}