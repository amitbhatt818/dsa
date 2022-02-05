package main

import "fmt"

func findAllHobbyists(hobby string, hobbies map[string][]string) []string {
	var hobbyists []string
	for h, hobbies := range hobbies {
		for _, ho := range hobbies {
			if ho == hobby {
				hobbyists = append(hobbyists, h)
			}
		}
	}
	return hobbyists

}

func main() {
	hobbies := map[string][]string{
		"Steve": []string{"Fashion", "Piano", "Reading"},
		"Patty": []string{"Drama", "Magic", "Pets"},
		"Chad":  []string{"Puzzles", "Pets", "Yoga"},
	}

	fmt.Println(findAllHobbyists("Yoga", hobbies))
}
