// This program implements Gnome Sort
// http://en.wikipedia.org/wiki/Gnome_sort

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

func gnome_sort() {
	for cur := data.Front().Next(); cur != nil; cur = cur.Next() {
		prev := cur.Prev()
		if cur.Value.(int) < prev.Value.(int) {
			cur.Value, prev.Value = prev.Value, cur.Value
			/* run algorithm backwards */
			for curBack := cur.Prev(); ; curBack = curBack.Prev() {
				prevBack := curBack.Prev()
				if prevBack == nil {
					break
				}
				if curBack.Value.(int) < prevBack.Value.(int) {
					curBack.Value, prevBack.Value = prevBack.Value, curBack.Value
				} else {
					break
				}
			}
		}
	}
}

func main() {
	fmt.Println("Welcome to Gnome Sort")

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

	gnome_sort()

	fmt.Println("Sorted List: ")
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

