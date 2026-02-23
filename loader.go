package main

import (
	"encoding/json"
	"os"
)

func LoadVideos(filename string) ([]Video, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var videos []Video

	err = json.Unmarshal(data, &videos)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func LoadUsers(filename string) ([]User, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
