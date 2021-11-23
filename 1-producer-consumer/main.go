//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, tweets chan<- *Tweet) {
	defer close(tweets)
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			return
		}
		fmt.Println(1)
		tweets <- tweet
	}
}

func consumer(tweets <-chan *Tweet) {
	for tweet := range tweets {
		if tweet.IsTalkingAboutGo() {
			fmt.Println(tweet.Username, "\ttweets about golang")
		} else {
			fmt.Println(tweet.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	tweets := make(chan *Tweet, 5)

	// Producer
	go producer(stream, tweets)

	// Consumer
	consumer(tweets)

	fmt.Printf("Process took %s\n", time.Since(start))
}
