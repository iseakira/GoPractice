package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"sync"
)

func downloadCSV(wg *sync.WaitGroup, urls[]string,ch chan []byte){
	defer wg.Done()
	defer close(ch)

	for _,u := range urls {

		resp,err := http.Get(u)
			if err != nil{
				log.Println("cannot downlaod CSV:",err)
				continue
			}
			b,err := os.ReadAll(resp.Body)
			if err != nil{
				resp.Body.Close()
				log.Println("cannot read content:",err)
				continue
			}
			resp.Body.Close()
			ch <- b
		}

	}


func main(){
	urls := []string {
		"http://my-server.com/data01.csv",
		"http://my-server.com/data02.csv",
		"http://my-server.com/data02.csv",
	}


ch := make(chan []byte)

var wg sync.WaitGroup
wg.Add(1)
go downloadCSV(&wg,urls,ch)

for b := range ch {
	r := csv.NewReader(bytes.NewReader(b))

	for {
		records,err := r.Read()
		if err == io.EOF {
			break
		}
		 if err != nil {
        log.Println("read error:", err)
        break
    }
		insertRecords(records)
	}

}
wg.Wait()
}

select {
case v1 := <- ch1:

case v2 := <- ch2:

default:
}

for {
	select {
	case ev := <- keyCh:

	case ev := <- recvCh:

	deafult:
	}
}



