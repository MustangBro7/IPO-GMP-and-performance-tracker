/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import "ipo_tracker/cmd"

func main() {
	cmd.Execute()
}

// func main() {
// 	// command := fmt.Sprint("/Users/abhinavmohan/myenv/bin/python3", "ipos/price_fetcher.py", "INFY")

// 	// Use bash to execute the combined command
// 	cmd := exec.Command("/Users/abhinavmohan/myenv/bin/python3", "ipos/price_fetcher.py", "INFY")
// 	var out bytes.Buffer
// 	cmd.Stdout = &out
// 	err := cmd.Run()
// 	if err != nil {
// 		fmt.Println("Error running Python script:", err)
// 		println(err.Error())
// 	}

// 	println(out.String())

// }
