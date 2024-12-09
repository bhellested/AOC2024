package main

import (
	"fmt"
	"os"
	"strconv"
	"math"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calculateChecksum(arr []uint64) uint64 {
	var ret uint64=0;
	for i,val := range arr{
		if val == math.MaxUint64 {
			continue
		}
		ret+=val*uint64(i)
	}
	return ret
}

func rearrageStoragePart1(arr *[]uint64){
	left :=0
	right := len(*arr)-1

	for left < right {
		//move left forward until the next empty
		for (*arr)[left] != math.MaxUint64{
			left++
		}
		//move right backward until the next nonEmpty
		for (*arr)[right]==math.MaxUint64 {
			right--
		}
		if(left>=right){
			break
		}
		//swap left and right, then increment one
		temp := (*arr)[left]
		(*arr)[left] = (*arr)[right]
		(*arr)[right] = temp
		left++;
		right--
	}
}

func rearrageStoragePart2(arr *[]uint64){
	right := len(*arr)-1
	rightSeen := make(map[uint64]struct{})
	rightSeen[math.MaxUint64]=struct{}{}
	for {
		count:=0
		for{
			_,exists :=rightSeen[(*arr)[right]]
			if exists {
				right--
				if right<=0 {
					return
				}
			} else{
				break
			}
			
		}
		//right is now at the last position of a file, count how many occurances we have
		val := (*arr)[right]
		for {
			if(*arr)[right]==val{
				count++
				right--
				if(right<=0){
					return
				}
			}else{
				right++//make sure that right is positioned on the last occurance
				break
			}
		}
		rightSeen[(*arr)[right]]=struct{}{}
		left:=-1
		//now we need to find a line of open spaces at least count in length, that are to the left of right.
		for i:=0;i<right;{
			if((*arr)[i]==math.MaxUint64){
				//start the count
				candidateLeft:=i
				openingSize:=0

				for i<right && (*arr)[i] == math.MaxUint64 {
					openingSize++
					i++
					if(openingSize>=count){
						break
					}
				}
				if(openingSize>=count){
					left = candidateLeft
					break
				}
			}else{
				i++
			}
		}
		//perform the move if we found a good position, otherwise we just move on
		if left > 0 {
			rightVal := (*arr)[right]
			for i:=0;i<count;i++{
				(*arr)[left+i]=rightVal
				(*arr)[right+i] = math.MaxUint64
			}
			
		}
	}
}

func main(){
	dat, err := os.ReadFile("../../inputs/day9.txt")
	check(err)
	var totalLength uint64 =0
	for _,char := range string(dat) {
		val,err := strconv.ParseUint(string(char),10,64)
		check(err)
		totalLength+=val
	}
	fileVec := make([]uint64,totalLength)
	isFile := true;
	totalLength=0//reusing this for current position
	for ID,char := range string(dat) {
		val,err := strconv.ParseUint(string(char),10,64)
		check(err)

		for i:=uint64(0);i<val;i++{
			if(isFile){
				fileVec[totalLength+i]=uint64(ID/2)
			}else{
				fileVec[totalLength+i] = math.MaxUint64
			}
		}
		totalLength+=val
		isFile = !isFile
	}
	cloneForPart2 := make([]uint64, len(fileVec))
	copy(cloneForPart2, fileVec)
	rearrageStoragePart1(&fileVec)
	rearrageStoragePart2(&cloneForPart2)

	fmt.Println("CheckSum for Part 1: ", calculateChecksum(fileVec))
	fmt.Println("CheckSum for Part 2: ", calculateChecksum(cloneForPart2))
}