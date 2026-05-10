package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)

		go func(tt *Test){
			defer wg.Done()
			fmt.Println(tt.name)
		}(&tt)

	}
	wg.Wait()

}