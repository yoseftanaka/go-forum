package config

import "forum/models"

// List of all models
var Models = []interface{}{
	&models.User{},
	&models.Post{},
}
