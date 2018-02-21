package main


import (
	"fmt"
	"lib/util"
)



func main() {
	
	done := make(chan bool)
	
	fmt.Println("Starting http get...")
	go util.CallURL(done)
	<- done
	
	fmt.Println("Finished http get...")
}


