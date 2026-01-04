package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) AddTrack(date string, status string, wordstatus int64) Tracker{
	return addTrack(date, status, wordstatus)
}

func (a *App) GetAllTracks() []Tracker{
	loadTracker()
	return TrackerList
}

func (a *App) DeleteTrack(id int16) {
	deleteTrack(id)
}

func (a *App) IsHighestID(id int16) bool {
	return IsHighestID(id)
}