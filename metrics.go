package main


// Реализация метрики precision. 
//  Подаем на вход список рекомендованных видео и айди просмотренных видео, которые мы скроем от алгоритма. 
//  Проверяем, сколько в списке рекомендаций скрытых видео.
func CalculatePrecision(recommend []Video, hiddenIDs []int) float64 {
	if len(recommend) == 0 || len(hiddenIDs) == 0 {
		return 0
	}

	hits := 0
	topN := 3
	if len(recommend) < topN {
		topN = len(recommend)
	}

	for i := 0; i < topN; i++ {
		for _, hiddenID := range hiddenIDs {
			if recommend[i].ID == hiddenID {
				hits++
			}
		}
	}
	return float64(hits) / float64(topN)
}
