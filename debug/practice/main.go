package practice

func greet(user string) string {
	if user == "" {
		user = "world"
	}
	return "Hello " + user
}
