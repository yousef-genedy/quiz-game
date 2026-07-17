package main

import "math/rand"

func Shuffle(p [][]string) [][]string {
	rand.Shuffle(len(p), func(i, j int) {
		p[i], p[j] = p[j], p[i]
	})
	return p
}
