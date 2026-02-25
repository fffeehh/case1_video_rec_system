package main

import "fmt"

type RecService struct {
	Videos []Video
	Users []User
}

func NewRecService(v []Video, u []User) *RecService {
	return &RecService{
		Videos: v,
		Users: u,
	}
}


func(s *RecService) GetRecommendationsForUser(userName string) ([]Video, error) {
	var targetUser User
	found := false

	for _, u := range s.Users {
		if u.Name == userName {
			targetUser = u
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("user not found")
	}

	return Recommend(targetUser, s.Users, s.Videos), nil
}
