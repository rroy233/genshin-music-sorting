package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//读取两个目录
	lessDict, err := os.Open("./oldVersion")
	if err != nil {
		panic(err)
	}
	lde, err := lessDict.ReadDir(-1)
	if err != nil {
		panic(err)
	}
	moreDict, err := os.Open("./newVersion")
	if err != nil {
		panic(err)
	}
	mde, err := moreDict.ReadDir(-1)
	if err != nil {
		panic(err)
	}

	lessMap := make(map[string]string, len(lde))

	//遍历less
	for _, entry := range lde {
		data, err := ioutil.ReadFile("./oldVersion/" + entry.Name())
		if err != nil {
			panic(err)
		}
		lessMap[fmt.Sprintf("%x", md5.Sum(data))] = "./oldVersion/" + entry.Name()
		fmt.Printf("lessMap[%s] = %s\n", fmt.Sprintf("%x", md5.Sum(data)), "./oldVersion/"+entry.Name())
	}

	//遍历more
	for _, entry := range mde {
		data, err := ioutil.ReadFile("./newVersion/" + entry.Name())
		if err != nil {
			panic(err)
		}
		fMd5 := fmt.Sprintf("%x", md5.Sum(data))
		fName := "./newVersion/" + entry.Name()
		if lessMap[fMd5] == "" {
			log.Printf("New file found:%s !!\n", fName)
			copyfile(fName, "./out/"+entry.Name())
		} else {
			log.Printf("%s skipped.\n", fName)
		}
	}

}

func copyfile(from, to string) {
	originalFile, err := os.Open(from)
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()
	newFile, err := os.Create(to)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
	_, err = io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
