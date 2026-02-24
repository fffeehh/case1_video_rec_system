package main

import (
	"fmt"
)

func main() {
	videos, err := LoadVideos("data/videos.json")
	if err != nil {
		panic(err)
	}

	users, err := LoadUsers("data/users.json")
	if err != nil {
		panic(err)
	}	
	currentUser := users[1]

		// Получаем рекомендации
	recommendations := Recommend(currentUser, users, videos)

	fmt.Printf("Привет, %s! Вот что мы подобрали для тебя:\n", currentUser.Name)
	for i, v := range recommendations {
		fmt.Printf("%d. %s (Категория: %s, Лайков: %d)\n", i+1, v.Title, v.Category, v.Likes)
	}
}
