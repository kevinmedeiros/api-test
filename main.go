// main.go
package main

func main() {
	a := App{}
	// Mysql Conf
	a.Initialize("DB_USERNAME", "DB_PASSWORD", "api_test")

	a.Run(":8080")
}
