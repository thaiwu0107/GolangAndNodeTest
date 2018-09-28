package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var rankTable7CF map[string]interface{}
var rankTable7CNF map[string]interface{}
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
	fmt.Println("Successfully ioutil.ReadFile")
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
	fmt.Println("Successfully All Json To Map")
}

func GoPokerCalculator(c chan bool, in []string) {
	var cardMap = [5][15]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	var point int
	var suit int
	suitMap := make(map[int][]int)

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
			var osa = make([]int, 0)
			sa := &osa
			*sa = append(*sa, suit)
			suitMap[point] = osa
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
	var keyOfRankString []string
	var keyOfRank string
	for i := 14; i > 1; i-- {
		keyOfRankString = append(keyOfRankString, strconv.Itoa(cardMap[selectSuit][i]))
	}
	keyOfRank = strings.Join(keyOfRankString, "")
	var rankInfo interface{}
	selectCards := make([]string, 5)
	if isFlush {
		rankInfo = rankTable7CF[keyOfRank]
		for i := 0; i < 5; i++ {
			cardPoint := rankInfo.(map[string]interface{})["CardPoint"].([]interface{})[i]
			point := int(cardPoint.(float64))
			selectCards[i] = strconv.Itoa(point*10 + selectSuit)
		}
	} else {
		rankInfo = rankTable7CNF[keyOfRank]
		for i := 0; i < 5; i++ {
			cardPoint := rankInfo.(map[string]interface{})["CardPoint"].([]interface{})[i]
			point := int(cardPoint.(float64))
			suit, suitMap[point] = suitMap[point][len(suitMap[point])-1], suitMap[point][:len(suitMap[point])-1]
			selectCards[i] = strconv.Itoa(point*10 + suit)
		}
	}
	c <- true
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
	suitMap := make(map[int][]int)

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
			var osa = make([]int, 0)
			sa := &osa
			*sa = append(*sa, suit)
			suitMap[point] = osa
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
	var keyOfRankString []string
	var keyOfRank string
	for i := 14; i > 1; i-- {
		keyOfRankString = append(keyOfRankString, strconv.Itoa(cardMap[selectSuit][i]))
	}
	keyOfRank = strings.Join(keyOfRankString, "")
	var rankInfo interface{}
	selectCards := make([]string, 5)
	if isFlush {
		rankInfo = rankTable7CF[keyOfRank]
		for i := 0; i < 5; i++ {
			cardPoint := rankInfo.(map[string]interface{})["CardPoint"].([]interface{})[i]
			point := int(cardPoint.(float64))
			selectCards[i] = strconv.Itoa(point*10 + selectSuit)
		}
	} else {
		rankInfo = rankTable7CNF[keyOfRank]
		for i := 0; i < 5; i++ {
			cardPoint := rankInfo.(map[string]interface{})["CardPoint"].([]interface{})[i]
			point := int(cardPoint.(float64))
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

	t1 := time.Now()
	for i := 0; i < loopTimes; i++ {
		go GoPokerCalculator(c, a)
	}
	<-c
	elapsed1 := time.Since(t1)
	fmt.Println("Golang - Benchmark times: ", loopTimes)
	fmt.Println("Golang - Benchmark Multi-core processor took: ", elapsed1)
	t2 := time.Now()
	for i := 0; i < loopTimes; i++ {
		PokerCalculator(a)
	}
	elapsed2 := time.Since(t2)
	fmt.Println("Golang - Benchmark times: ", loopTimes)
	fmt.Println("Golang - Benchmark Single-core processor took: ", elapsed2)
}
