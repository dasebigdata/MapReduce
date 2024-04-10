package main

import (
	"bufio"
	"fmt"
	"mr/mr"
	"os"
	"plugin"
	"sort"
)

func main() {
	inputPath := os.Args[1]
	outputPath := os.Args[2]
	mapPath := os.Args[3]
	reducePath := os.Args[4]

	mapSo, err := plugin.Open(mapPath)
	if err != nil {
		panic(err)
	}
	mapp, err := mapSo.Lookup("Map")
	if err != nil {
		panic(err)
	}
	mapFunc := *(mapp.(*mr.Map))

	reduceSo, err := plugin.Open(reducePath)
	if err != nil {
		panic(err)
	}
	reduce, err := reduceSo.Lookup("Reduce")
	if err != nil {
		panic(err)
	}
	reduceFunc := *(reduce.(*mr.Reduce))

	// inputs
	lines := make([]string, 0)
	inputFile, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// map
	pairs := make([]*mr.Pair, 0)
	for _, line := range lines {
		pairs = append(pairs, mapFunc(line)...)
	}

	// shuffle
	intermediate := make(map[any][]any)
	for _, pair := range pairs {
		if _, ok := intermediate[pair.Key]; !ok {
			intermediate[pair.Key] = make([]any, 0)
		}
		intermediate[pair.Key] = append(intermediate[pair.Key], pair.Value)
	}
	intermediateList := make([]*mr.Pair, 0, len(intermediate))
	for key, value := range intermediate {
		intermediateList = append(intermediateList, &mr.Pair{Key: key, Value: value})
	}
	sort.Slice(intermediateList, func(i, j int) bool {
		return fmt.Sprintf("%v", intermediateList[i].Key) < fmt.Sprintf("%v", intermediateList[j].Key)
	})

	// reduce
	result := make([]*mr.Pair, 0, len(intermediateList))
	for _, pair := range intermediateList {
		result = append(result, reduceFunc(pair.Key, pair.Value.([]any)))
	}

	// output
	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	for _, pair := range result {
		fmt.Fprintf(outputFile, "%v\t%v", pair.Key, pair.Value)
		fmt.Fprintln(outputFile)
	}
}
