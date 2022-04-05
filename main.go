package main

import (
	"hacktiv8_assignment_4/routes"
)

func main() {
	PORT := ":8080"
	routes.Serve().Run(PORT)
}
