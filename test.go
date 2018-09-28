package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"runtime"
	"strconv"
	"sync/atomic"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type Rank struct {
	Type7     string `json:"Type7"`
	Rank      int    `json:"Rank"`
	Type5Ch   string `json:"Type5Ch"`
	Type5En   string `json:"Type5En"`
	CardPoint []int  `json:"CardPoint"`
}

var rankTable7CF map[string]*Rank
var rankTable7CNF map[string]*Rank
var config map[string]interface{}

func InitJSONToMap() {
	jsonFileConfig, err1 := ioutil.ReadFile("./config.json")
	if err1 != nil {
		fmt.Println(err1)
	}
	jsonFileRankTable7CF, err2 := ioutil.ReadFile("./RankTable7CF.json")
	if err2 != nil {
		fmt.Println(err2)
	}
	jsonFileRankTable7CNF, err3 := ioutil.ReadFile("./RankTable7CNF.json")
	if err3 != nil {
		fmt.Println(err3)
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	error1 := json.Unmarshal(jsonFileConfig, &config)
	if error1 != nil {
		fmt.Println("Unmarshal failed, ", error1)
		return
	}
	error2 := json.Unmarshal(jsonFileRankTable7CF, &rankTable7CF)
	if error2 != nil {
		fmt.Println("Unmarshal failed, ", error2)
		return
	}
	error3 := json.Unmarshal(jsonFileRankTable7CNF, &rankTable7CNF)
	if error3 != nil {
		fmt.Println("Unmarshal failed, ", error3)
		return
	}
}

func GoPokerCalculator(in []string) {
	var cardMap = [5][15]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	var point int
	var suit int
	suitMap := make(map[int][]int, 7)

	for i := 0; i < 7; i++ {
		suitInt, _ := strconv.Atoi(in[i])
		suit = suitInt % 10
		point = (suitInt - suit) / 10

		cardMap[suit][point]++
		cardMap[suit][0]++
		cardMap[0][point]++

		if suitMap[point] != nil {
			suitMap[point] = append(suitMap[point], suit)
		} else {
			suitMap[point] = []int{suit}
		}

	}

	isFlush := false
	selectSuit := 0
	for i := 0; i < 5; i++ {
		if cardMap[i][0] >= 5 {
			isFlush = true
			selectSuit = i
			break
		}
	}
	var buffer bytes.Buffer
	for i := 14; i > 1; i-- {
		buffer.WriteString(strconv.Itoa(cardMap[selectSuit][i]))
	}
	keyOfRank := buffer.String()
	var rankInfo *Rank
	selectCards := make([]string, 5)
	if isFlush {
		rankInfo = rankTable7CF[keyOfRank]
		for i := 0; i < 5; i++ {
			point := rankInfo.CardPoint[i]
			selectCards[i] = strconv.Itoa(point*10 + selectSuit)
		}
	} else {
		rankInfo = rankTable7CNF[keyOfRank]
		for i := 0; i < 5; i++ {
			point := rankInfo.CardPoint[i]
			suit, suitMap[point] = suitMap[point][len(suitMap[point])-1], suitMap[point][:len(suitMap[point])-1]
			selectCards[i] = strconv.Itoa(point*10 + suit)
		}
	}
}

func PokerCalculator(in []string) {
	var cardMap = [5][15]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	var point int
	var suit int
	suitMap := make(map[int][]int, 7)

	for i := 0; i < 7; i++ {
		suitInt, _ := strconv.Atoi(in[i])
		suit = suitInt % 10
		point = (suitInt - suit) / 10

		cardMap[suit][point]++
		cardMap[suit][0]++
		cardMap[0][point]++

		if suitMap[point] != nil {
			suitMap[point] = append(suitMap[point], suit)
		} else {
			suitMap[point] = []int{suit}
		}

	}

	isFlush := false
	selectSuit := 0
	for i := 0; i < 5; i++ {
		if cardMap[i][0] >= 5 {
			isFlush = true
			selectSuit = i
			break
		}
	}
	var buffer bytes.Buffer
	for i := 14; i > 1; i-- {
		buffer.WriteString(strconv.Itoa(cardMap[selectSuit][i]))
	}
	keyOfRank := buffer.String()
	var rankInfo *Rank
	selectCards := make([]string, 5)
	if isFlush {
		rankInfo = rankTable7CF[keyOfRank]
		for i := 0; i < 5; i++ {
			point := rankInfo.CardPoint[i]
			selectCards[i] = strconv.Itoa(point*10 + selectSuit)
		}
	} else {
		rankInfo = rankTable7CNF[keyOfRank]
		for i := 0; i < 5; i++ {
			point := rankInfo.CardPoint[i]
			suit, suitMap[point] = suitMap[point][len(suitMap[point])-1], suitMap[point][:len(suitMap[point])-1]
			selectCards[i] = strconv.Itoa(point*10 + suit)
		}
	}
}

// 這是單純自己call自己測試
func main() {
	InitJSONToMap()
	loopTimes := int(config["testloop"].(float64))
	configGOMAXPROCS := int(config["runtimeGOMAXPROCS"].(float64))
	runtime.GOMAXPROCS(configGOMAXPROCS)
	c := make(chan bool)
	a := make([]string, 7)
	a[0] = "144"
	a[1] = "131"
	a[2] = "121"
	a[3] = "122"
	a[4] = "91"
	a[5] = "24"
	a[6] = "132"

	remain := int64(loopTimes)
	poolSize := 300
	pool := make(chan struct{}, poolSize)
	for range make([]struct{}, poolSize) {
		pool <- struct{}{}
	}
	t1 := time.Now()
	for i := 0; i < loopTimes; i++ {
		<-pool
		go func() {
			GoPokerCalculator(a)
			pool <- struct{}{}
			if atomic.AddInt64(&remain, -1) == 0 {
				c <- true
			}
		}()
	}
	<-c
	elapsed1 := time.Since(t1)
	fmt.Println("Golang - Benchmark times: ", loopTimes)
	fmt.Println("Golang - Benchmark Multi-core poolSize: ", poolSize)
	fmt.Println("Golang - Benchmark Multi-core processor took: ", elapsed1)
	t2 := time.Now()
	for i := 0; i < loopTimes; i++ {
		PokerCalculator(a)
	}
	elapsed2 := time.Since(t2)
	fmt.Println("Golang - Benchmark Single-core processor took: ", elapsed2)
}
