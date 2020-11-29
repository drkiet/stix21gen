package main

import (
	"fmt"
	stix21 "github.com/drkiet/stix21"
	"os"
)


func main() {
	fmt.Println("Hello, world!")

	artifact := GenerateArtifact()
	fmt.Printf("%v\n", stix21.MarshalArtifact(artifact))
	bundle := GenerateBundle([]byte(stix21.MarshalArtifact(artifact)))
	fmt.Println("bundle: main{}")
	if len(os.Args) < 2 {
		fmt.Println("*** error: no input file ***")
		os.Exit(-1)
	}
	jsonData := stix21.MarshalBundle(bundle)
	fmt.Printf("Bundle:%v\n", jsonData)
	//data := readAll(os.Args[1])
	//bundle = stix21.Unmarshal(data)
	//
	//jsonData := stix21.MarshalBundle(bundle)
	//fmt.Printf("Bundle:%v\n", jsonData)
}
