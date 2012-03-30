// This program implements Comb Sort
// http://en.wikipedia.org/wiki/Comb_sort

package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"strconv"
	"container/list"
)

var data list.List

func comb_sort() {

	const SHRINK_FACTOR int = 3
	gap := data.Len() / SHRINK_FACTOR

	for {
		next := data.Front()
		for c := 0; c < gap; c++ {
			next = next.Next()
		}
		isSwap := false
		for cur := data.Front(); cur != nil; cur = cur.Next() {
			next = next.Next()
			if next == nil {
				break
			}
			if cur.Value.(int) > next.Value.(int) {
				cur.Value, next.Value = next.Value, cur.Value
				isSwap = true
			}
		}
		if isSwap == false {
			break
		}
		if gap > 1 {
			gap -= 1
		}
	}
}

func main() {
	fmt.Println("Welcome to Comb Sort")

	data.Init()

	/* Read unsorted data from file and store it into list */
	f, err := os.Open("./data1000.txt")
	if err == nil {
		r := bufio.NewReader(f)
		for {
			line, _, err := r.ReadLine()
			if err == io.EOF {
				break
			}
			i, err := strconv.Atoi(string(line))
			data.PushBack(i)
		}
	} else {
		fmt.Println("Error reading data file")
		return
	}
	err = f.Close();
	fmt.Println("Read ", data.Len(), " integers")

	comb_sort()

	fmt.Println("Sorted List: ")
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

