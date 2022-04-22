package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Task: Given a list of 4 integers, find all possible valid 24 hour times (eg: 12:34) that the given integers can be assembled into and return the total number of valid times.
You can not use the same number twice.
Times such as 34:12 and 12:60 are not valid.
Provided integers can not be negative.
Notes: Input integers can not be negative.
Input integers can yeald 0 possible valid combinations.
Example:
	Input: [1, 2, 3, 4]
	Valid times: ["12:34", "12:43", "13:24", "13:42", "14:23", "14:32", "23:14", "23:41", "21:34", "21:43"]
	Return: 10
*/


func filterData (req []string) (resp []string){

	for _, v := range req{
		s := strings.Split(v, ":")

		isValid := false

		if v1, e :=  strconv.Atoi(s[0]); e == nil{
			if v1 < 24{
				isValid = true
			}
		}
		
		//error occured
		if isValid == false{
			continue
		}

		isValid = false
		if v2, e :=  strconv.Atoi(s[1]); e == nil{
			if v2 < 60{
				isValid = true
			}
		}

		//error occured
		if isValid == false{
			continue
		}

		if(isValid){
			resp = append(resp, v)
		}
	}

	return
}

func calNextPos(asc bool, input ...int) (r int){
	
	arr := []int {0,1,2,3}
	if len(input) == 2 && asc == false{
		arr = []int {3,2,1,0}
	}

	for _, v := range arr{
		exists := false
		for _, vv:= range input{
				if v == vv{
					exists = true
					break
				}
		}

		if exists == false{
			r = v
			break 
		}
	}

	// fmt.Println("====")
	// fmt.Println(input, r)
	return
}

func possibleTimes(digits []int) int{
	if len(digits) != 4{
		fmt.Println("INVALID INPUT COUNT :", len(digits))
		return 0
	}

	res := make([]string, 0)
	l:= len(digits)
	totalPosLen:= l
	
	for i:= l -1 ; i > 0; i--{
		totalPosLen*= i
	}

	for i := 0; i <  l; i++{
		for j := i+1; j < i + 4 ; j++ {
		   for  k := j+1; k < j + 3; k++{
			for  m := k+1; m < k+2; m++{
				tmpJ := j
				if tmpJ >= l{
					tmpJ = tmpJ-l
				}
				tmpK := k
				if tmpK >= l{
					tmpK = calNextPos(tmpK == l, i,tmpJ)
				}
				tmpM := m
				if tmpM >= l{
					tmpM = calNextPos(false, i,tmpJ,tmpK)
				}
				//appending result..
				res = append(res, fmt.Sprintf("%d%d:%d%d", digits[i], digits[tmpJ],digits[tmpK],digits[tmpM]))
			 }
		   }
		}
	 }

	 if totalPosLen != len(res){
		fmt.Println("EXP Possible Count not Matching :", len(res))
		return 0
	 }
	 
	 expTimerCount := len(filterData(res))
	 return expTimerCount
}

func main() {
	
	// Example test cases.
	fmt.Println(possibleTimes([]int{1, 2, 3, 4}))
	fmt.Println(possibleTimes([]int{9, 1, 2, 0}))
	fmt.Println(possibleTimes([]int{2, 2, 1, 9}))
}