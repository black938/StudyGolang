package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	_ = 1
	_ = 2
	a := []int{1, 3, 5, 7, 8}
	rr := twoSum(a, 8)
	fmt.Println(rr, len(rr), cap(rr))

	s1 := []int{1, 2}
	s2 := []int{3, 4, 5}
	s3 := append(s1, s2...) //解构
	fmt.Println(s3)
	ss := []int{1, 2, 3}
	sss := make([]int, 3)
	copy(sss, ss)
	sss[0] = 99
	fmt.Println(ss, sss)
	sort.Ints(sss)
	fmt.Println(sss)

	//dict
	scoreMap := make(map[string]int, 8)
	scoreMap2 := map[string]int{"xiaoming": 99, "xiaohong": 98}
	fmt.Println(scoreMap, scoreMap2)
	delete(scoreMap2, "xiaoming")
	//字典构成的切片
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	fmt.Println(countWord("how do you do"))

	fmt.Println(strconv.Atoi("319"))
	fmt.Println(strconv.Itoa(319))

	//var (
	//	name    string
	//	age     int
	//	married bool
	//)
	//fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	//fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
	fmt.Println(time.Now().Unix())
	start := time.Now().UnixNano()
	time.Sleep(1 * time.Second)
	end := time.Now().UnixNano()
	fmt.Println(end - start)

	t := time.Now()
	fmt.Print(t.Format("2006 01 02 15 04"))

}

//var twoPointList [][]int
//var stringList []string

func twoSum(nums []int, target int) [][]int {
	var result [][]int
	for index1, _ := range nums {
		for index2, _ := range nums {
			if nums[index1]+nums[index2] == target && index1 != index2 && index1 < index2 {
				result = append(result, []int{index1, index2})
			}
		}
	}
	return result
}

func countWord(sentence string) map[string]int {
	resultMap := make(map[string]int, 8)
	wordList := strings.Split(sentence, " ")
	for _, word := range wordList {
		resultMap[word] += 1
	}
	return resultMap
}
