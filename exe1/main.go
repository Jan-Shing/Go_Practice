package main

import(
	"encoding/csv"
	"fmt"
	"flag"
	"io"
	"os"
	"log"
	"time"
)



func main(){
	var(
	 h bool
	 file_name string
	 timeout int
	)

	flag.StringVar(&file_name, "f", "program.csv", "file name for quiz question")
	flag.BoolVar(&h, "h", false, "help message")
	flag.IntVar(&timeout, "t", -1, "timer limited mode")
	flag.Parse()

	if h {
		flag.Usage()
	}else{
		fmt.Printf("file name for quiz is %s\n", file_name)
	}

	file, err := os.Open(file_name)
	if err != nil{
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	
	var correct, total int
	answerCH := make(chan string)


	timer1 := time.NewTimer(time.Duration(timeout) * time.Second)
	if timeout < 0{
		timer1.Stop()
	}
	

loop:
	for {

		record, err := reader.Read()
		if err == io.EOF{
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		total++

		go func(){
			var ans string
			fmt.Printf("Quiz %d: %s",total, record[0])
			fmt.Scanf("%s\n", &ans)
			answerCH <- ans
		}()

		select {
		case <- timer1.C:
			fmt.Printf("time out.....")
			break loop
		case answer := <-answerCH:
			if answer == record[1]{
			correct++
			}
		}
		

	}

fmt.Printf(" You scored %d out of %d\n", correct, total)


}