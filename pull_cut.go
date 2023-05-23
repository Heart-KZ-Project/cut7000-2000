package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

func listOfFileNames(files []fs.FileInfo) []string {
	var ans = make([]string, 0)
	for _, file := range files {
		f := strings.Split(file.Name(), "_")
		// fmt.Println(f)
		v, _ := strconv.Atoi(f[0][len(f[0])-3:])
		if v > 400 && v%3 == 0 && f[1] == "fake" && f[2] == "B.png" {
			ans = append(ans, file.Name())
		}
	}
	return ans
}

func copyToPull(files []string) {
	for _, v := range files {
		in, err := ioutil.ReadFile(filepath.Join("images", v))
		if err != nil {
			fmt.Println(err)
			return
		}
		err = ioutil.WriteFile(filepath.Join(".", "cut7000-pull", v), in, 0644)
		if err != nil {
			fmt.Println("Error creating", filepath.Join(".", "cut7000-pull", v))
			fmt.Println(err)
			return
		}
	}

}

func main() {
	files, err := ioutil.ReadDir("images")
	if err != nil {
		log.Fatal(err)
	}

	copyToPull(listOfFileNames(files))
	// fmt.Println(listOfFileNames(files), len(listOfFileNames(files)))
}
