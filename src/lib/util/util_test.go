package util


import (
	"testing"
)


func TestAverage(t *testing.T) {
  var v float64
  input := []float64{1, 2, 3} 
  v = Average(input)
  if v != 2 {
    t.Error("Expected 2, got ", v)
  }  
}

func TestCallURL(t *testing.T) {
  done := make(chan bool)
  go CallURL(done)
  v := <- done
  if !v {
    t.Error("Expected true, got ", v)
  }  
}
