package main

import (
	"github.com/gaku3601/clean-blog/src/infrastructure/router"
	_ "github.com/lib/pq"
)

func main() {
	router.Start()
}
