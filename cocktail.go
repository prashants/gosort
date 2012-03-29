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

func cocktail_sort() {
	ptrStart := data.Front()
	ptrEnd := data.Back()
	cur := ptrStart

	for {
		if ptrStart == ptrEnd {
			break
		}

		/* forward scan */
		cur = ptrStart
		for {
			if cur == ptrEnd || cur == nil {
				break
			}
			next := cur.Next()
			if cur.Value.(int) >= next.Value.(int) {
				/* swap current and next elements */
				cur.Value, next.Value = next.Value, cur.Value
			}
			cur = cur.Next()
		}
		ptrEnd = ptrEnd.Prev()

		if ptrStart == ptrEnd {
			break
		}

		/* reverse scan */
		cur = ptrEnd
		for {
			if cur == ptrStart || cur == nil {
				break
			}
			prev := cur.Prev()
			if cur.Value.(int) < prev.Value.(int) {
				/* swap current and previous elements */
				cur.Value, prev.Value = prev.Value, cur.Value
			}
			cur = cur.Prev()
		}
		ptrStart = ptrStart.Next()
	}
}

func main() {
	fmt.Println("Welcome to Bubble Sort")

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

	cocktail_sort()

	fmt.Println("Sorted List: ")
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

