package main

import "fmt"

var TrackerList []Tracker
const FILENAME_TRACKER = "tracker.json"

func GetNextID() int16 {

	if len(TrackerList) == 0 {
		return 0
	}

	maxID := TrackerList[0].ID
	for _, track := range TrackerList[1:] {
		if track.ID > maxID {
			maxID = track.ID
		}
	}
	return (maxID + 1)

}

func loadTracker () {
	data, err := readData[[]Tracker](FILENAME_TRACKER)
	
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	
	TrackerList = data
		
}

func writeTracker () {
	err := writeData[[]Tracker](TrackerList, FILENAME_TRACKER)
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func addTrack (date string, status string, wordstatus int64) Tracker{
	loadTracker()
	newTrack := createNewTracker(date, status, wordstatus)
	TrackerList = append(TrackerList, newTrack)
	writeTracker()
	
	return newTrack
}

func GetTrackByID(id int16) *Tracker {
	for i := range TrackerList {
		if TrackerList[i].ID == id {
			return &TrackerList[i]
		}
	}
	return nil
}

func deleteTrack(id int16) {
	trackToDelete := GetTrackByID(id)
	writtenwords := trackToDelete.Writtenwords
	currentWordCount = currentWordCount - writtenwords
	writeWordcount()
	
	n := 0
	for _, track := range TrackerList {
		if track.ID != id {
			TrackerList[n] = track
			n++
		}
	}

	TrackerList = TrackerList[:n]
	writeTracker()
}

func IsHighestID(id int16) bool {
	for i := range TrackerList {
		if TrackerList[i].ID > id {
			return false
		}
	}
	return  true
}



