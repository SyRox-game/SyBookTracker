package main

import "fmt"

var currentWordCount int64
const FILENAME_WORDCOUNT = "wordcount.json"

func getWrittenWords(currentWordCountVar int64) int64 {
	loadWordcount()
	writtenwords := currentWordCountVar - currentWordCount
	currentWordCount = currentWordCountVar
	writeWordcount()
	
	return writtenwords
}

func loadWordcount () {
	data, err := readData[int64](FILENAME_WORDCOUNT)
	
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	
	currentWordCount = data
		
}

func writeWordcount () {
	err := writeData[int64](currentWordCount, FILENAME_WORDCOUNT)
	if err != nil {
		fmt.Printf(err.Error())
	}
}