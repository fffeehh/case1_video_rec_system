package main

import (
	"fmt"
)

func main() {
	// Загрузка данных (слой Data)
	videos, err := LoadVideos("data/videos.json")
	if err != nil {
		panic(err)
	}
	users, err := LoadUsers("data/users.json")
	if err != nil {
		panic(err)
	}
	
	// Инициализация сервиса (слой Logic)
	service := NewRecService(videos, users)
	
	userName := "Миша"
	recs, err := service.GetRecommendationsForUser(userName)
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}

	fmt.Printf("=== Рекомендации для %s ===\n", userName)
	for i, v := range recs {
		if i >= 3 {
			break
		}
		fmt.Printf("- %s [%s]\n", v.Title, v.Category)
	}
	
	// Оценка качества (метрики)
	hiddenHistory := []int{3}
	precision := CalculatePrecision(recs, hiddenHistory)
	fmt.Printf("\nТочность алгоритма (Precision@3): %.2f\n", precision)
}
