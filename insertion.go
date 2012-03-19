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

func insertion_sort() {
	/* add the first element directly */
	e := data.Front()
	sorted_data.PushBack(e.Value.(int))

	/* pick one item from original unsorted list */
	for e = e.Next(); e != nil; e = e.Next() {
		inserted := 0
		s := sorted_data.Front()
		/* search the sorted list for the correct place to insert the element */
		for s = sorted_data.Front(); s != nil; s = s.Next() {
			if s.Value.(int) > e.Value.(int) {
				/* insert the element in the sorted list at the correct place */
				sorted_data.InsertBefore(e.Value.(int), s)
				inserted = 1
				break
			}
		}
		/* insert the element to the end of the list */
		if inserted == 0 && s == nil {
			sorted_data.PushBack(e.Value.(int))
			inserted = 1
		}
		if inserted == 0 {
			fmt.Println("Error inserting item")
			os.Exit(1)
		}
	}
}

func main() {
	fmt.Println("Welcome to Insertion Sort")

	data.Init()
	sorted_data.Init()

	/* Read unsorted data from file and store it into list */
	f, err := os.Open("./data100000.txt")
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

	insertion_sort()

	fmt.Println("Sorted List: ")
	for e := sorted_data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

