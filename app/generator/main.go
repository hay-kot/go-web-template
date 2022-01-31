package main

import (
	"github.com/hay-kot/git-web-template/ent"
	"github.com/hay-kot/git-web-template/internal/dtos"
	"github.com/hay-kot/git-web-template/pkgs/automapper"
)

// getMappers serialized the config file into a list of automapper struct
func getMappers() []automapper.AutoMapper {
	return []automapper.AutoMapper{
		{
			Package: "mapper",
			Prefix:  "users",
			Name:    "User Out",
			Schema: automapper.Schema{
				Type:   dtos.UserOut{},
				Prefix: "dtos",
			},
			Model: automapper.Model{
				Type:   ent.User{},
				Prefix: "ent",
			},
			Imports: []string{},
		},
	}
}

func main() {
	automappers := getMappers()
	conf := automapper.DefaultConf()

	automapper.Generate(automappers, conf)
}
