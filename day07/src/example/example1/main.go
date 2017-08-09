package main

import (
	"fmt"
	"os"
)

func main() {

	//fmt.Fprint(os.Stdout,"hahahhah")

	file, err := os.OpenFile("/Users/victor/test.log",os.O_CREATE|os.O_WRONLY,777)

	if err != nil {
		fmt.Printf("not found file %s",err);
	}



	fmt.Fprintf(file,"hello %s world","-")

}
