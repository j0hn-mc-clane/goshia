package main

import (
	handler "goshia/listener/webhooks"
)

func main() {
	// Serve Gin as separate goroutine
	handler.Serve()
}
