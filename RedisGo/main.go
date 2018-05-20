package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client

func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func ExampleClient() {
	err := client.Set("key", "Michael Jordan", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func PubSubDemo() {
	pubsub := client.Subscribe("mychannel1")
	defer pubsub.Close()

	// Wait for subscription to be created before publishing message.
	subscr, err := pubsub.ReceiveTimeout(time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println(subscr)

	err = client.Publish("mychannel1", "hello").Err()
	if err != nil {
		panic(err)
	}

	msg, err := pubsub.ReceiveMessage()
	if err != nil {
		panic(err)
	}

	fmt.Println(msg.Channel, msg.Payload)
}

func PubSubGoRoutine() {
	pubsub := client.Subscribe("mychannel2")
	defer pubsub.Close()

	// Wait for subscription to be created before publishing message.
	subscr, err := pubsub.ReceiveTimeout(time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println(subscr)

	go func() {
		for i := 0; i < 10; i++ {
			msg := "msg" + strconv.Itoa(i)
			err = client.Publish("mychannel2", msg).Err()
			fmt.Println("Publish " + msg)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			duration, _ := time.ParseDuration("500ms")
			time.Sleep(duration)
		}
	}()

	go func() {
		for {
			msg, err := pubsub.ReceiveMessage()
			if err != nil {
				panic(err)
			}

			fmt.Println(msg.Channel, msg.Payload)
		}
	}()
	select {}

}
func main() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//ExampleNewClient()
	//ExampleClient()
	//PubSubDemo()
	PubSubGoRoutine()

}
