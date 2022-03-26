package services

import "github.com/hay-kot/git-web-template/backend/internal/repo"

type AllServices struct {
	Auth *AuthService
}

func NewServices(repos *repo.AllRepos) *AllServices {
	return &AllServices{
		Auth: &AuthService{repos},
	}
}
