package config

import "forum/main/models"

// List of all models
var Models = []interface{}{
	&models.User{},
	&models.Post{},
}
