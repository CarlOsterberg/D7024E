package main

import ("bufio"
	"fmt"
	"os"
	"strings"
	)

func main() {
	//x := 0
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("hello")
	for {
		//x++
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		fmt.Println("responding..")
		fmt.Println(text)
	}
}
