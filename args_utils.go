package main

import "strings"

func ParseURLS(slice []string) []string {
	for i, v := range slice {
		isAnURL := strings.HasPrefix(v, "http://") || strings.HasPrefix(v, "https://")

		if !isAnURL {
			slice[i] = "http://" + v
		}
	}

	return slice
}
