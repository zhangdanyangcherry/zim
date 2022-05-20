package main

import (
	"github.com/smallnest/rpcx/server"
	"github.com/zmicro-team/zim/app/chat/internal/model"
	"github.com/zmicro-team/zim/app/chat/internal/service"
	"github.com/zmicro-team/zim/pkg/runtime"
	"github.com/zmicro-team/zmicro"
	"github.com/zmicro-team/zmicro/core/log"
)

func main() {
	app := zmicro.New(
		zmicro.InitRpcServer(InitRpcServer),
		zmicro.Before(before),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

func InitRpcServer(s *server.Server) error {
	if err := s.RegisterName("Chat", service.GetChatService(), ""); err != nil {
		log.Fatal(err.Error())
	}
	if err := s.RegisterName("Conv", service.GetConvService(), ""); err != nil {
		log.Fatal(err.Error())
	}
	return nil
}

func before() error {
	runtime.Setup()
	db := runtime.GetDB()
	if err := db.AutoMigrate(
		&model.Msg{},
		&model.User{},
		&model.Group{},
		&model.GroupMember{},
	); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
