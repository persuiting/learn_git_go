package main

import "fmt"

func main()  {
	str := ""
	for i := 0; i <= 255; i++ {
		str = str + "0"
	}
	pwdMap := make(map[string]bool,0)
	//pwdMap[""] = true
	//pwdMap["1"] = true
	fmt.Println(len(pwdMap))
	fmt.Printf(str)
}

