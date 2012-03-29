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

func quick_sort() {
	l := data.Front()
	r := data.Back()
	p := l
	for {
		for {
			fmt.Println("R")
			if p.Value.(int) > r.Value.(int) {
				fmt.Println("SR", r.Value)
				fmt.Println(r.Value, p.Value)
				temp := r.Value
				r.Value = p.Value
				p.Value = temp
				fmt.Println(r.Value, p.Value)
				p = r
				l = l.Next()
				break
			} else {
				fmt.Println("MR")
				r = r.Prev()
			}
		}
		fmt.Println("Pivot", p.Value)	
		for {
			fmt.Println("L")
			if p.Value.(int) > l.Value.(int) {
				fmt.Println("SL")
				temp := l.Value
				l.Value = p.Value
				p.Value = temp
				p = l
				r = r.Prev()
				break
			} else {
				fmt.Println("ML")
				l = l.Next()
			}
		}
	}
}

func main() {
	fmt.Println("Welcome to Quick Sort")

	data.Init()

	/* Read unsorted data from file and store it into list */
	f, err := os.Open("./data10.txt")
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

	quick_sort()

	fmt.Println("Sorted List: ")
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

