package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"go-callback-request", // name
		true,                  // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := bodyFrom(os.Args)
	//doc := `[{"message" : "true"}]`
	doc := `[{"callback":{"request_id":"20210909-live-recording-file-61398de53a573","created_at":1643180970,"retry_created_at":-1,"callback_type":"kollus-channel","content_provider_key":"","media_content_id":-1,"callback_url":"https://ryu-cb.requestcatcher.com/complete","callback_method":"post","callback_data":{"recording_file_id":"907715251","service_account_key":"sarabose","target_filename":"meaqfjq9uclaeuuc-sarabose_ch-2021-09-09.mp4","recording_file_transfer_result":"1","recording_file_filename":"/sarabose/meaqfjq9uclaeuuc/20210909-ygpecl4n/segment/20210909043011068_2021-09-09-04.30.11.069-UTC_0.mp4","broadcast_key":"20210909-ygpecl4n","version":"2018-11-01","recording_file_kind":"segment","stream_key":"7dwtg3uw","channel_key":"meaqfjq9uclaeuuc"},"callback_max_retry":3,"callback_retry":0,"callback_connect_timeout":2,"callback_response_timeout":3,"callback_response":200,"callback_response_body":"request caught"},"request_id":"20210909-live-recording-file-61398de53a573","message":"SUCCESS"}]`
	//err = ch.Publish(
	//	"",     // exchange
	//	q.Name, // routing key
	//	false,  // mandatory
	//	false,
	//	amqp.Publishing{
	//		DeliveryMode: amqp.Persistent,
	//		ContentType:  "text/plain",
	//		Body:         []byte(body),
	//	})
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         []byte(doc),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

type Article struct {
	Message string
}

func jsonStrAdd() string {
	doc := `
	[{
		"message": "true",
	}]
	`

	var data []Article // JSON 문서의 데이터를 저장할 구조체 슬라이스 선언

	json.Unmarshal([]byte(doc), &data) // doc의 내용을 변환하여 data에 저장

	fmt.Println(data) // [{1 Hello, world! {Maria maria@exa... (생략)
	return doc
}
