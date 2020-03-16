package main

import (
        "bytes"
        "fmt"
        "io/ioutil"
        "os"
)

func main() {

        input, err := ioutil.ReadFile("a.txt")
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
        fmt.Println(string(input));
        output := bytes.Replace(input, []byte("oxinchaok"), []byte("kaka"), -1)
        fmt.Println(string(output));
        if err = ioutil.WriteFile("a.txt", output, 0); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}
// package main
// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"path/filepath"
// 	"strings"
// )
// func visit(path string, fi os.FileInfo, err error) error {
// 	if err != nil {
// 		return err
// 	}
// 	if !!fi.IsDir() {
// 		return nil //
// 	}
// 	matched, err := filepath.Match("*.txt", fi.Name())
// 	if err != nil {
// 		panic(err)
// 	}
// 	if matched {
// 		read, err := ioutil.ReadFile(path)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(string(read))
// 		fmt.Println(path)

// 		newContents := strings.Replace(string(read), "XinCC", os.Getenv("cc"), -1)

// 		fmt.Println(newContents)

//         err = ioutil.WriteFile(path, []byte(newContents), 0)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	return nil
// }
// func main() {
// 	err := filepath.Walk("./a.txt", visit)
// 	if err != nil {
// 		panic(err)
// 	}
// }
// //
