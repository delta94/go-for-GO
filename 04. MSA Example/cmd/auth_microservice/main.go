package main

import (
	"MSA.example.com/1/dataservice/userdata"
	"MSA.example.com/1/tool/dbc"
	"MSA.example.com/1/tool/message"
	"MSA.example.com/1/usecase"
	"github.com/go-playground/validator/v10"
	"log"
	"runtime"
)

func main() {
	db, err := dbc.ConnectDB("AuthMicroservice")
	if err != nil {
		log.Fatalf("unable to connect DB server, err: %v\n", err)
	}

	userD := userdata.GetUserDAO(db)
	natsM, err := message.GetDefaultNatsByEnv()
	if err != nil {
		log.Fatalf("unable to connect NATS server, err: %v\n", err)
	}
	validate := validator.New()
	u := usecase.NewAuthDefaultUseCase(userD, natsM, validate)

	_, err = natsM.Subscribe("auth.signup", u.SignUpMsgHandler)
	// handler들을 메서드로 가지고 있는 usecase 구조체 생성 추가 (db, nats 필드 소유)
	if err != nil {
		log.Fatalf("unable to subscribe auth.login from nats message broker, err: %v\n", err)
	}
	log.Println("Auth message pub/sub server is completely started.")
	runtime.Goexit()
}

