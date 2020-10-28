package main

import (
	_ "trygin/databases"
	"trygin/routers"
)

func main() {
	routers.Engine.Run(":8080")
}
