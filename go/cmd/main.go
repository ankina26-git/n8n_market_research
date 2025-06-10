package main

import (
	v1 "n8n_project_go/app/user/v1"
)

func main() {
	app := v1.New()
	app.Listen(":3000")
}
