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

	//qName := "go-callback-request"
	qName := "go-email-request"

	q, err := ch.QueueDeclare(
		qName, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	failOnError(err, "Failed to declare a queue")

	body := bodyFrom(os.Args)
	//doc := `{"callback":{"request_id":"20210909-live-recording-file-61398de53a573","created_at":1643180970,"retry_created_at":-1,"callback_type":"kollus-channel","content_provider_key":"","media_content_id":-1,"callback_url":"https://ryu-cb.requestcatcher.com/complete","callback_method":"post","callback_data":{"recording_file_id":"907715251","service_account_key":"sarabose","target_filename":"meaqfjq9uclaeuuc-sarabose_ch-2021-09-09.mp4","recording_file_transfer_result":"1","recording_file_filename":"/sarabose/meaqfjq9uclaeuuc/20210909-ygpecl4n/segment/20210909043011068_2021-09-09-04.30.11.069-UTC_0.mp4","broadcast_key":"20210909-ygpecl4n","version":"2018-11-01","recording_file_kind":"segment","stream_key":"7dwtg3uw","channel_key":"meaqfjq9uclaeuuc"},"callback_max_retry":3,"callback_retry":0,"callback_connect_timeout":2,"callback_response_timeout":3,"callback_response":200,"callback_response_body":"request caught"},"request_id":"20210909-live-recording-file-61398de53a573","message":"SUCCESS"}`
	//doc := `[{"message" : "true"}]`
	//doc := `{"message" : "true"}`
	//doc := `{"callback":{"request_id":"20210909-live-recording-file-61398de53a573","created_at":1643180970,"retry_created_at":-1,"callback_type":"kollus-channel","content_provider_key":"","media_content_id":-1,"callback_url":"http://192.168.10.108:3000/ping","callback_method":"get","callback_data":{"recording_file_id":"907715251","service_account_key":"sarabose","target_filename":"meaqfjq9uclaeuuc-sarabose_ch-2021-09-09.mp4","recording_file_transfer_result":"1","recording_file_filename":"/sarabose/meaqfjq9uclaeuuc/20210909-ygpecl4n/segment/20210909043011068_2021-09-09-04.30.11.069-UTC_0.mp4","broadcast_key":"20210909-ygpecl4n","version":"2018-11-01","recording_file_kind":"segment","stream_key":"7dwtg3uw","channel_key":"meaqfjq9uclaeuuc"},"callback_max_retry":3,"callback_retry":0,"callback_connect_timeout":2,"callback_response_timeout":3,"callback_response":200,"callback_response_body":"request caught"},"request_id":"20210909-live-recording-file-61398de53a573","message":"SUCCESS"}`
	//doc := `{"from":"KOLLUS <kollus_noreply@kollus.com>","to":"arneg0shua@gmail.com","bcc":"","subject":"Kollus DEV KR - \ud68c\uc6d0\uac00\uc785 \uc778\uc99d \uba54\uc77c","body":"<table cellpadding=\"0\" cellspacing=\"0\" width=\"700px\" style=\"color:#888; font-size: 14px; font-family:Lucida Grande,Apple SD Gothic Neo,Dotum,Arial,Helvetica,Geneva,Verdana,sans-serif\">\n\t<tr>\n\t\t<td style=\"border-bottom:3px solid #3498db; padding:10px 0;\"><a href=\"\" style=\"color:#3498db; font-size:14px; font-weight:bold; text-decoration:none;\">Kollus DEV KR<\/a><\/td>\n    <\/tr>\n\t<tr>\n\t\t<td style=\"padding:50px 0 80px; text-align:center;\"><a href=\"\" style=\"color:#3498db; letter-spacing:-1px; font-size:30px; text-decoration:none;\">\ud68c\uc6d0\uac00\uc785 \uc778\uc99d\uba54\uc77c<\/a><\/td>\n    <\/tr>\n\t<tr>\n        <td style=\"padding:0 50px;line-height:1.5;\">Kollus DEV KR \uac00\uc785\uc744 \ud658\uc601\ud569\ub2c8\ub2e4.<br \/>\r\n\uc544\ub798 \ub9c1\ud06c\ub97c \ud074\ub9ad\ud558\uc2dc\uba74 Kollus DEV KR \uc11c\ube44\uc2a4 \ud398\uc774\uc9c0\ub85c \ub2e4\uc2dc \uc5f0\uacb0\ub429\ub2c8\ub2e4.<br \/>\r\n\ub610\ub294 \uc778\uc99d \ud1a0\ud070\uc744 \ubcf5\uc0ac\ud574\uc11c Kollus DEV KR \uacc4\uc815\uc0dd\uc131 \ucc3d \ucf54\ub4dc\uc5d0 \ubd99\uc5ec\ub123\uae30 \ud574 \uc8fc\uc2ed\uc2dc\uc694.<\/td>     \n    <\/tr>\n    <tr>\n\t\t<td style=\"padding:20px 50px 0;\">\n\t\t\t<table cellpadding=\"0\" cellspacing=\"0\" width=\"100%\" align=\"center\" style=\"border-top:2px solid #3498db; color:#888; font-size:14px\">\n\t\t\t\t<tr>\n              \t\t<td style=\"background:#f8f8f8; border-bottom:1px solid #ddd; border-right:1px solid #ddd; color:#666; font-weight:bold; padding:12px; text-align:center; width:150px;\">\ub9c1\ud06c<\/td>\n                \t<td style=\"border-bottom:1px solid #ddd; padding:12px;\"><a href=\"https:\/\/dev-kr.kollus.com\/account\/signin_by_authentication\/1986?token=04e7aa76\" style=\"color:#3498db;\" target=\"_blank\">https:\/\/dev-kr.kollus.com\/account\/signin_by_authentication\/1986?token=04e7aa76<\/a><\/td>\n            \t<\/tr>\n                <tr>\n                    <td style=\"background:#f8f8f8; border-bottom:1px solid #ddd; border-right:1px solid #ddd; color:#666; font-weight:bold; padding:12px; text-align:center; width:150px;\">\uc778\uc99d \ud1a0\ud070<\/td>\n                    <td style=\"border-bottom:1px solid #ddd; padding:12px;\">04e7aa76<\/td>\n                <\/tr>\n              \t<tr>\n              \t\t<td style=\"background:#f8f8f8; border-bottom:1px solid #ddd; border-right:1px solid #ddd; color:#666; font-weight:bold; padding:12px; text-align:center; width:150px;\">\uc778\uc99d\uc720\ud6a8\uae30\uac04<\/td>\n                    <td style=\"border-bottom:1px solid #ddd; padding:12px;\">2022\/04\/27 02:47<\/td>\n           \t<\/tr>\n    \t\t<\/table>\n\t\t<\/td>      \n    <\/tr>\n    <tr>\n        <td align=\"center\" style=\"color:#bbb; font-size:12px; padding:10px 0 80px;\"># \uac1c\uc778\uc815\ubcf4 \ubcf4\ud638\ub97c \uc704\ud574 \uc778\uc99d \uc720\ud6a8\uc77c\uc2dc\uae4c\uc9c0(\uc778\uc99d\uba54\uc77c \ubc1c\uc1a1 \ud6c4 24\uc2dc\uac04) \uc778\uc99d\uc774 \uc644\ub8cc\ub418\uc9c0 \uc54a\uc73c\uba74 \uacc4\uc815\uc774 \uc0ad\uc81c\ub429\ub2c8\ub2e4.<\/td>    \n\t<\/tr>\n    <tr>\n        <td style=\"padding:0 50px;\">\n            <table cellpadding=\"0\" cellspacing=\"0\" width=\"100%\" align=\"center\" style=\"background:#f9f9f9; color:#888; padding:15px;\">\n                <tr>       \n        \t\t\t<td style=\"font-size:13px; line-height:1.5; text-align:justify;\">Kollus DEV KR\ub294 \uace0\uac1d \uc5ec\ub7ec\ubd84\uaed8\uc11c \ub9cc\uc871\ud558\uc2e4 \uc218 \uc788\ub294 \uc591\uc9c8\uc758 \uc11c\ube44\uc2a4\ub97c \uc81c\uacf5\ud558\ub294 \uac83\uc740 \ubb3c\ub860, <br\/> \uace0\uac1d\uc758 \ubb38\uc758\uc0ac\ud56d\uacfc \ubd88\ud3b8\uc0ac\ud56d\uc744 \ud56d\uc2dc \uc811\uc218\ud558\uc5ec \uc11c\ube44\uc2a4 \uc720\uc9c0\uc5d0 \ub9cc\uc804\uc744 \uae30\ud558\uaca0\uc2b5\ub2c8\ub2e4. <br\/> \uc11c\ube44\uc2a4\uc5d0 \uad00\ud55c \uc0ac\ud56d\uc740 \ub2f9\uc0ac FAQ \ud398\uc774\uc9c0\ub97c \ud1b5\ud574 \uc0c1\uc138\ud788 \ud655\uc778\ud558\uc2e4 \uc218 \uc788\uc2b5\ub2c8\ub2e4.<br \/><br \/>\uac10\uc0ac\ud569\ub2c8\ub2e4.<\/td>\n\t\t\t\t<\/tr>\n\t\t\t<\/table>\n    \t<\/td>        \n    <\/tr>\n    \n        <tr>\n        <td style=\"color:#888; font-size:13px; line-height:1.5; padding:20px 0; text-align:center;\">\ubcf8 \uba54\uc77c\uc740 \ubc1c\uc2e0 \uc804\uc6a9\uc785\ub2c8\ub2e4. \uad81\uae08\ud558\uc2e0 \ub0b4\uc6a9\uc740 \uc774\uba54\uc77c\ub85c \ubb38\uc758\ud574\uc8fc\uc2dc\uae30 \ubc14\ub78d\ub2c8\ub2e4.<br \/>\uc774\uba54\uc77c : <a href=\"mailto:kollus_sales@catenoid.net\" style=\"color:#c00;\" target=\"_blank\">kollus_sales@catenoid.net<\/a><\/td>\n    <\/tr>\n    <\/table>\n<table cellpadding=\"0\" cellspacing=\"0\" width=\"700px\" style=\"border-top:1px solid #ddd; color:#888; font-size: 13px; font-family:Lucida Grande,Apple SD Gothic Neo,Dotum,Arial,Helvetica,Geneva,Verdana,sans-serif; padding:20px; 0\">\n    <tr>\n        <td align=\"center\" valign=\"middle\">\n            Copyright 2013 Kollus.com, Inc. All Right Reserved. \n            <a href=\"http:\/\/kr.dev.kollus.com\" style=\"color:#3498db;\" target=\"_blank\">http:\/\/kr.dev.kollus.com<\/a>\n        <\/td>\n    <\/tr>\n<\/table> \n"}`
	doc := `{"request_id" : "ddd", "message":"success", "callback" : {"from":"KOLLUS <kollus_noreply@kollus.com>","to":"jungin.kim@catenoid.net","cc":"arneg0shua@gmail.com","bcc":"copying_y@naver.com","subject":"Kollus DEV KR - \ud68c\uc6d0\uac00\uc785 \uc778\uc99d \uba54\uc77c","body":"<table cellpadding=\"0\" cellspacing=\"0\" width=\"700px\" style=\"color:#888; font-size: 14px; font-family:Lucida Grande,Apple SD Gothic Neo,Dotum,Arial,Helvetica,Geneva,Verdana,sans-serif\">\n\t<tr>\n\t\t<td style=\"border-bottom:3px solid #3498db; padding:10px 0;\"><a href=\"\" style=\"color:#3498db; font-size:14px; font-weight:bold; text-decoration:none;\">Kollus DEV KR<\/a><\/td>\n    <\/tr>\n\t<tr>\n\t\t<td style=\"padding:50px 0 80px; text-align:center;\"><a href=\"\" style=\"color:#3498db; letter-spacing:-1px; font-size:30px; text-decoration:none;\">\ud68c\uc6d0\uac00\uc785 \uc778\uc99d\uba54\uc77c<\/a><\/td>\n    <\/tr>\n\t<tr>\n        <td style=\"padding:0 50px;line-height:1.5;\">Kollus DEV KR \uac00\uc785\uc744 \ud658\uc601\ud569\ub2c8\ub2e4.<br \/>\r\n\uc544\ub798 \ub9c1\ud06c\ub97c \ud074\ub9ad\ud558\uc2dc\uba74 Kollus DEV KR \uc11c\ube44\uc2a4 \ud398\uc774\uc9c0\ub85c \ub2e4\uc2dc \uc5f0\uacb0\ub429\ub2c8\ub2e4.<br \/>\r\n\ub610\ub294 \uc778\uc99d \ud1a0\ud070\uc744 \ubcf5\uc0ac\ud574\uc11c Kollus DEV KR \uacc4\uc815\uc0dd\uc131 \ucc3d \ucf54\ub4dc\uc5d0 \ubd99\uc5ec\ub123\uae30 \ud574 \uc8fc\uc2ed\uc2dc\uc694.<\/td>     \n    <\/tr>\n    <tr>\n\t\t<td style=\"padding:20px 50px 0;\">\n\t\t\t<table cellpadding=\"0\" cellspacing=\"0\" width=\"100%\" align=\"center\" style=\"border-top:2px solid #3498db; color:#888; font-size:14px\">\n\t\t\t\t<tr>\n              \t\t<td style=\"background:#f8f8f8; border-bottom:1px solid #ddd; border-right:1px solid #ddd; color:#666; font-weight:bold; padding:12px; text-align:center; width:150px;\">\ub9c1\ud06c<\/td>\n                \t<td style=\"border-bottom:1px solid #ddd; padding:12px;\"><a href=\"https:\/\/dev-kr.kollus.com\/account\/signin_by_authentication\/1986?token=04e7aa76\" style=\"color:#3498db;\" target=\"_blank\">https:\/\/dev-kr.kollus.com\/account\/signin_by_authentication\/1986?token=04e7aa76<\/a><\/td>\n            \t<\/tr>\n                <tr>\n                    <td style=\"background:#f8f8f8; border-bottom:1px solid #ddd; border-right:1px solid #ddd; color:#666; font-weight:bold; padding:12px; text-align:center; width:150px;\">\uc778\uc99d \ud1a0\ud070<\/td>\n                    <td style=\"border-bottom:1px solid #ddd; padding:12px;\">04e7aa76<\/td>\n                <\/tr>\n              \t<tr>\n              \t\t<td style=\"background:#f8f8f8; border-bottom:1px solid #ddd; border-right:1px solid #ddd; color:#666; font-weight:bold; padding:12px; text-align:center; width:150px;\">\uc778\uc99d\uc720\ud6a8\uae30\uac04<\/td>\n                    <td style=\"border-bottom:1px solid #ddd; padding:12px;\">2022\/04\/27 02:47<\/td>\n           \t<\/tr>\n    \t\t<\/table>\n\t\t<\/td>      \n    <\/tr>\n    <tr>\n        <td align=\"center\" style=\"color:#bbb; font-size:12px; padding:10px 0 80px;\"># \uac1c\uc778\uc815\ubcf4 \ubcf4\ud638\ub97c \uc704\ud574 \uc778\uc99d \uc720\ud6a8\uc77c\uc2dc\uae4c\uc9c0(\uc778\uc99d\uba54\uc77c \ubc1c\uc1a1 \ud6c4 24\uc2dc\uac04) \uc778\uc99d\uc774 \uc644\ub8cc\ub418\uc9c0 \uc54a\uc73c\uba74 \uacc4\uc815\uc774 \uc0ad\uc81c\ub429\ub2c8\ub2e4.<\/td>    \n\t<\/tr>\n    <tr>\n        <td style=\"padding:0 50px;\">\n            <table cellpadding=\"0\" cellspacing=\"0\" width=\"100%\" align=\"center\" style=\"background:#f9f9f9; color:#888; padding:15px;\">\n                <tr>       \n        \t\t\t<td style=\"font-size:13px; line-height:1.5; text-align:justify;\">Kollus DEV KR\ub294 \uace0\uac1d \uc5ec\ub7ec\ubd84\uaed8\uc11c \ub9cc\uc871\ud558\uc2e4 \uc218 \uc788\ub294 \uc591\uc9c8\uc758 \uc11c\ube44\uc2a4\ub97c \uc81c\uacf5\ud558\ub294 \uac83\uc740 \ubb3c\ub860, <br\/> \uace0\uac1d\uc758 \ubb38\uc758\uc0ac\ud56d\uacfc \ubd88\ud3b8\uc0ac\ud56d\uc744 \ud56d\uc2dc \uc811\uc218\ud558\uc5ec \uc11c\ube44\uc2a4 \uc720\uc9c0\uc5d0 \ub9cc\uc804\uc744 \uae30\ud558\uaca0\uc2b5\ub2c8\ub2e4. <br\/> \uc11c\ube44\uc2a4\uc5d0 \uad00\ud55c \uc0ac\ud56d\uc740 \ub2f9\uc0ac FAQ \ud398\uc774\uc9c0\ub97c \ud1b5\ud574 \uc0c1\uc138\ud788 \ud655\uc778\ud558\uc2e4 \uc218 \uc788\uc2b5\ub2c8\ub2e4.<br \/><br \/>\uac10\uc0ac\ud569\ub2c8\ub2e4.<\/td>\n\t\t\t\t<\/tr>\n\t\t\t<\/table>\n    \t<\/td>        \n    <\/tr>\n    \n        <tr>\n        <td style=\"color:#888; font-size:13px; line-height:1.5; padding:20px 0; text-align:center;\">\ubcf8 \uba54\uc77c\uc740 \ubc1c\uc2e0 \uc804\uc6a9\uc785\ub2c8\ub2e4. \uad81\uae08\ud558\uc2e0 \ub0b4\uc6a9\uc740 \uc774\uba54\uc77c\ub85c \ubb38\uc758\ud574\uc8fc\uc2dc\uae30 \ubc14\ub78d\ub2c8\ub2e4.<br \/>\uc774\uba54\uc77c : <a href=\"mailto:kollus_sales@catenoid.net\" style=\"color:#c00;\" target=\"_blank\">kollus_sales@catenoid.net<\/a><\/td>\n    <\/tr>\n    <\/table>\n<table cellpadding=\"0\" cellspacing=\"0\" width=\"700px\" style=\"border-top:1px solid #ddd; color:#888; font-size: 13px; font-family:Lucida Grande,Apple SD Gothic Neo,Dotum,Arial,Helvetica,Geneva,Verdana,sans-serif; padding:20px; 0\">\n    <tr>\n        <td align=\"center\" valign=\"middle\">\n            Copyright 2013 Kollus.com, Inc. All Right Reserved. \n            <a href=\"http:\/\/kr.dev.kollus.com\" style=\"color:#3498db;\" target=\"_blank\">http:\/\/kr.dev.kollus.com<\/a>\n        <\/td>\n    <\/tr>\n<\/table> \n"}}`
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
