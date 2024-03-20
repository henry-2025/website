package main

import "henry2025/website/internal"

func main() {
	s := internal.NewSite(false)
	s.SetupAndServe()
}
