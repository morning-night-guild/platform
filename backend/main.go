package main

//go:generate wire

func main() {
	InitializeHTTPServer().Run()
}
