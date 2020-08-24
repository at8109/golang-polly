package main

import "github.com/at8109/golang-polly/service"

var (
	kimberly service.PollyService = service.NewKimberlyPollyService()
	joey     service.PollyService = service.NewJoeyPollyService()
)

func main() {
	err := kimberly.Sythesize("Hi, I am Kimberly. How are you?", "kimberly.mp3")
	if err != nil {
		panic(err)
	}

	err = joey.Sythesize("Hi, I am Joey, Nice to meet you", "joey.mp3")
	if err != nil {
		panic(err)
	}

}
