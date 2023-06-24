package main

func main() {
	router := NewRouter()
	router.Logger.Fatal(router.Start(":1324"))
}
