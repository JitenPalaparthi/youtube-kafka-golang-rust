package main

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand/v2"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err.Error())
	}
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {

			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("message is not delivered to the topic :%d\n", ev.TopicPartition.Partition)
				} else {
					log.Printf("message successfully delivered to the following topic partition %d\n", ev.TopicPartition.Partition)
				}
			}
		}
	}()

	topic := "user-topic"

	u1 := NewUser(uint(rand.IntN(10000)), "Jiten", "Jitenp@outlook.com")

	bytes, err := u1.ToBytes()
	if err != nil {
		panic(err.Error())
	}

	err = p.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{
		Topic: &topic, Partition: kafka.PartitionAny}, Value: bytes,
	}, nil)

	if err != nil {
		log.Println(err.Error())
	}

	p.Flush(5000)

	log.Println("Message has been successfully sent")

}

type User struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(id uint, name, email string) *User {
	return &User{Id: id, Name: name, Email: email}
}

func (u *User) ToBytes() ([]byte, error) {
	if u == nil {
		return nil, errors.New("nil user")
	}
	return json.Marshal(u)
}
