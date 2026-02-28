package main

import "fmt"

type RecService struct {
	videos []Video
	users []User
	engine Recommender
}

func NewRecService(v []Video, u []User, e Recommender) *RecService {
	return &RecService{
		videos: v,
		users: u,
		engine: e,
	}
}


func(s *RecService) GetRecommendationsForUser(userName string) ([]Video, error) {
	var targetUser User
	found := false

	for _, u := range s.users {
		if u.Name == userName {
			targetUser = u
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("user not found")
	}
		return s.engine.Build(targetUser, s.users, s.videos), nil
}
