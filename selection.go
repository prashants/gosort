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
var sorted_data list.List

func selection_sort() {
	for {
		/* find the minimum element and move it to the sorted list */
		min := data.Front()
		for e := data.Front(); e != nil; e = e.Next() {
			if e.Value.(int) < min.Value.(int) {
				min = e
			}
		}
		/* move the minimum element to sorted list */
		sorted_data.PushBack(min.Value.(int))
		data.Remove(min)
		/* if unsorted list is empty break out of the main loop */
		if data.Len() <= 0 {
			return
		}
	}
}

func main() {
	fmt.Println("Welcome to Selection Sort")

	data.Init()
	sorted_data.Init()

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

	selection_sort()

	fmt.Println("Sorted List: ")
	for e := sorted_data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

