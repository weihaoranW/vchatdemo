package ctl

import (
	"fmt"
	"log"
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/vhaoran/vchat/lib"
	"github.com/vhaoran/vchat/lib/ymq"
)

func Test_emq_send(t *testing.T) {
	// load config
	opt := &lib.LoadOption{
		LoadMicroService: false,
		LoadEtcd:         false,
		LoadPg:           false,
		LoadRedis:        false,
		LoadMongo:        false,

		//-----------attention here------------
		LoadMq: true,
		//-----------attention here------------

		LoadJwt: false,
	}
	_, err := lib.InitModulesOfOptions(opt)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = ymq.X.Subscribe("a", func(c mqtt.Client, m mqtt.Message) {
		log.Println("on received:", string(m.Payload()))
	})
	if err != nil {
		fmt.Println("subscribe err :", err)
	}

	for i := 0; i < 10; i++ {
		err := ymq.X.PublishQos("a", 0, fmt.Sprint("hello,__,", i))
		if err != nil {
			log.Println(err)
			return
		} else {
			fmt.Println("------", fmt.Sprint(i), "-- ok---------")
		}
		time.Sleep(time.Second)
	}
}
