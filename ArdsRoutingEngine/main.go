// ArdsRoutingEngine project main.go
package main

import (
	"fmt"
	"time"
)

type Configuration struct {
	RedisIp   string
	RedisPort string
	RedisDb   int
	Port      string
}

type EnvConfiguration struct {
	RedisIp   string
	RedisPort string
	RedisDb   string
	Port      string
}

type Request struct {
	Company          int
	Tenant           int
	Class            string
	Type             string
	Category         string
	SessionId        string
	ArriveTime       string
	Priority         string
	QueueId          string
	ReqHandlingAlgo  string
	ReqSelectionAlgo string
	ServingAlgo      string
	HandlingAlgo     string
	SelectionAlgo    string
	RequestServerUrl string
	HandlingResource string
	ResourceCount    int
	OtherInfo        string
	LbIp             string
	LbPort           string
}

const layout = "2006-01-02T15:04:05Z07:00"

func main() {
	fmt.Println("Starting Ards Route Engine")
	InitiateRedis()
	for {
		//fmt.Println("Searching...")
		availablePHashes := GetAllProcessingHashes()
		for _, h := range availablePHashes {
			go ExecuteRequestHash(h)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func AppendIfMissing(dataList []Request, i Request) []Request {
	for _, ele := range dataList {
		if ele == i {
			return dataList
		}
	}
	return append(dataList, i)
}

type timeSlice []Request

func (p timeSlice) Len() int {
	return len(p)
}
func (p timeSlice) Less(i, j int) bool {
	t1, _ := time.Parse(layout, p[i].ArriveTime)
	t2, _ := time.Parse(layout, p[j].ArriveTime)
	return t1.Before(t2)
}
func (p timeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
