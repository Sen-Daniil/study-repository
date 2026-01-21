package main

import (
	"fmt"
	"study/feature1"
	"study/feature2"
	"study/feature_postgres/simle_connection"
)

func main() {
	fmt.Println("Hello, Git!")
	feature1.Feature1()
	feature2.Feature2()
	simle_connection.CheckConnection()
}