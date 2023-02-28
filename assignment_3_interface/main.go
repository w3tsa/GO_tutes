package main

import (
	"fmt"
	"os"
)

func main()  {
	file, err := os.Open(os.Args[1])
	if err!= nil {
        fmt.Println("Error:", err)
		os.Exit(1)
    }

	data := make([]byte, 100)
	count, err := file.Read(data)

	if err!= nil {
        fmt.Println("Error:", err)
		os.Exit(1)
    }

	fmt.Println(string(data[:count]))
}

/* 

// func main()  {
// 	file, err := os.Open(os.Args[1])
// 	if err!= nil {
//         fmt.Println("Error:", err)
// 		os.Exit(1)
//     }

// 	io.Copy(os.Stdout, file)
// }
*/