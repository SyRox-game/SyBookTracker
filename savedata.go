package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

func GetAppDataDir() string {

	home, err := os.UserHomeDir()
	if err != nil {
		println("err:", err.Error())
	}

	var base string

	switch runtime.GOOS {
		
	case "windows":
		base = os.Getenv("APPDATA")
		if base == "" {
			base = filepath.Join(home, "AppData", "Roaming")
		}

	case "darwin":
		base = filepath.Join(home, "Library", "Application Support")
		
	default:
		base = filepath.Join(home, ".local", "share")
	}

	path := filepath.Join(base, "SyBookTracker")
	err1 := os.MkdirAll(path, 0o755)

	if err1 != nil {
		println("err", err1.Error())
		return ""
	}

	return path

}

func writeData[T any] (data T, filename string) error {
	path := filepath.Join(GetAppDataDir(), filename)
	
	writeData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, writeData, 0644)

}

func readData[T any] (filename string) (T, error) {
	var zero T
	
	path := filepath.Join(GetAppDataDir(), filename)

	readData, err := os.ReadFile(path)

	if err != nil {
		println("err:", err.Error())
		return zero, err
	}

	var data T
	err = json.Unmarshal(readData, &data)

	return data, err
}