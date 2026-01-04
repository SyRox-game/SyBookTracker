package main

type Tracker struct {
	ID int16 `json:"id"`
	Date string `json:"date"`
	Status string `json:"status"`
	Wordstatus int64 `json:"wordstatus"`
	Writtenwords int64 `json:"writtenwords"`
}

func createNewTracker(date string, status string, wordstatus int64) Tracker {
	
	
	newTracker := Tracker{GetNextID(), date, status, wordstatus, getWrittenWords(wordstatus)}
	
	return newTracker
}