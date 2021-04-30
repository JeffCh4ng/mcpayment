package main

import (
	"fmt"
	"strings"
)

func main(){
	//1. We have array of integers called nums, write a function to
	//return all numbers (in a form of array of integers)
	//that when subtracted by any of integers in nums doesn't
	//return number that is < 0.
	//For example : nums = [3,1,4,2], your output should be : [4],
	//because when 4 is subtracted by 3 or 1 or 4 or 2 doesn't
	//return number that is < 0

	maxNumber := getMaxNumber([]int{3,1,4,2})
	fmt.Println("Maximum number is : ",maxNumber)

	//2. We have array of integers called nums and an integer called x,
	//write a function to return all numbers (in a form of array of integers)
	//that when divided by any of integers in nums doesn't return x.
	//For example : nums = [1,2,3,4], x = 4,
	//your output should be : [1,2,3], because only 4 divided by 1 is 4 (x)
	exceptionNumber := 4
	listNumber := []int{1,2,3,4}
	res := getListNumberExceptionX(listNumber, exceptionNumber)
	fmt.Printf("list number from %v without %d is %v \n",listNumber,exceptionNumber,res)

	//3. We have a string called word and an integer called x,
	//write a function to return an array of strings containing all strings that has length x.
	//For example : word = "souvenir loud four lost", x = 4,
	//your output should be ["loud", "four", "lost"] because those strings has length of 4
	word := "souvenir loud four lost"
	indexFilter := 4
	resultWord := getWordFilterByIndex(word,indexFilter)
	fmt.Printf("result word from %s to %v filter by index %d \n",word,resultWord, indexFilter)

}

func getMaxNumber(arrNumber []int)int{
	var tempNumber int
	if len(arrNumber) > 0{
		tempNumber = arrNumber[0]
	}
	for _,data := range arrNumber{
		if tempNumber < data {
			tempNumber = data
		}
	}
	return tempNumber
}

func getListNumberExceptionX(arrNumber []int, x int)(result []int){
	for _, data := range arrNumber{
		if x != data{
			result = append(result, data)
		}
	}
	return result
}

func getWordFilterByIndex(word string, index int) (resultWord []string){
	arrWord := strings.Split(word," ")
	for i,_ := range arrWord{
		if arrWord[i] != arrWord[index-1]{
			resultWord = append(resultWord, arrWord[i])
		}
	}
	return resultWord
}

//COMMAND TO RUN this code --> go run main.go