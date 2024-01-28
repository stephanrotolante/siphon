package main

func ListContainsValue(list []string, target string) bool {
	for _, value := range list {
		if value == target {
			return true
		}
	}
	return false
}
