package handlers

func AboutHandler() (map[string]string, error) {
	return map[string]string{"message": "This is the About page."}, nil
}
