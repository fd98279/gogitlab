package util

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"time"
)

func Average(xs []float64) float64 {
  total := float64(0)
  for _, x := range xs {
    total += x
  }
  return total / float64(len(xs))
}

func CallURL(_done chan<- bool) {
	url := "https://gitlab.com"
	httpClient := http.Client{}			
	startTime := time.Now()
	elapsed := time.Since(time.Now()).Seconds()
	counter := 0.0
	
	for (int(time.Now().Sub(startTime).Seconds()) < 5 ) {
		counter += 1
		start := time.Now()
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}
		
		res, getErr := httpClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}
			
		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}
		
		_ = body
		
		elapsed += time.Since(start).Seconds()
	}
			
	fmt.Println("Elapsed Time: ", elapsed/counter, " seconds")
	_done <- true	
}
