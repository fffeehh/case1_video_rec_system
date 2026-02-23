package main

import (
	"sort"
	"time"
)

func CalculateSimilarity(u1, u2 User) float64 {
	if len(u1.ViewedVideoIDs) == 0 || len(u2.ViewedVideoIDs) == 0 {
		return 0
	}
	commonCount := 0

	// ускоряем наш алгоритм, используя хэш таблицы вместо вложенных списков. Таким образом мы вместо алгоритма O(n*m) используем алгоритм O(n+m)
	hash := make(map[int]bool)
	for _, id := range u1.ViewedVideoIDs {
		hash[id] = true
	}
	for _, id := range u2.ViewedVideoIDs {
		if hash[id] {
			commonCount++
		}
	}
	
	unionCount := len(u1.ViewedVideoIDs) + len(u2.ViewedVideoIDs) - commonCount
	return float64(commonCount) / float64(unionCount)

}


// -----------------------------------------------------------------------
// дописать дописать дописать дописать дописать дописать дописать дописать
// -----------------------------------------------------------------------
func Recommend(user User, allVideos []Video) []Video {
	favoriteCategory := ""
	if len(user.ViewedVideoIDs) > 0 {
		lastVideoID := user.ViewedVideoIDs[len(user.ViewedVideoIDs)-1]
		for _, v := range allVideos {
			if v.ID == lastVideoID {
				favoriteCategory = v.Category
				break
			}
		}
	}

	type ratedVideo struct {
		video Video
		score int
	}

	var rankedVideos []ratedVideo

	for _, v := range allVideos {
		if contains(user.ViewedVideoIDs, v.ID) {
			continue
		}
		score := 0
		if v.Category == favoriteCategory {
			score += 10
		}
		if time.Since(v.CreatedAt).Hours() < 48 {
			score += 5
		}
		score += v.Likes / 10
		rankedVideos = append(rankedVideos, ratedVideo{v, score}) 
	}
	sort.Slice(rankedVideos, func(i, j int) bool {
		return rankedVideos[i].score > rankedVideos[j].score
	})
	var result []Video
	for _, v := range rankedVideos {
		result = append(result, v.video)
	}

	return result
}

func contains(slice []int, id int) bool {
	for _, v := range slice {
		if v == id {
			return true
		}
	}
	return false
}
