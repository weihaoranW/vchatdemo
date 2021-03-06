package model

import (
	"fmt"
	"log"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/vhaoran/vchat/lib"
	"github.com/vhaoran/vchat/lib/ypg"
)

type Abc struct {
	ypg.BaseModel
	ID    int `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	CName string
}

func (r *Abc) TableName() string {
	return "abc"
}

func prepare() {
	// load config
	opt := &lib.LoadOption{
		LoadMicroService: false,
		LoadEtcd:         false,

		//-----------attention here------------
		LoadPg: true,
		//-----------attention here------------

		LoadRedis: false,

		LoadMongo: false,

		LoadMq:  false,
		LoadJwt: false,
	}
	_, err := lib.InitModulesOfOptions(opt)
	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
}

func Test_pg_insert(t *testing.T) {
	var err error
	prepare()

	if ypg.X.HasTable(new(Abc)) {
		err := ypg.X.DropTable(new(Abc)).Error
		if err != nil {
			log.Println(err)
			return
		}
	}

	if !ypg.X.HasTable(new(Abc)) {
		er := ypg.X.CreateTable(new(Abc)).Error
		if er != nil {
			fmt.Println("---create table err---", err, "-----------")
			return
		}
	}

	//
	for i := 0; i < 10; i++ {
		bean := &Abc{
			ID: i,
			CName: "whr_test" +
				"",
		}

		err = ypg.X.Save(bean).Error
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("------", "ok", "-----------")
	}

	fmt.Println("------", "demo find", "-----------")
	l := make([]*Abc, 0)
	err = ypg.X.Find(&l).Error
	if err != nil {
		log.Println(err)
		return
	}

	spew.Dump(l)
}

func Test_pg_update(t *testing.T) {
	prepare()
	var err error

	for i := 0; i < 10; i++ {
		bean := &Abc{
			ID: i,
		}

		//err = ypg.X.Model(bean).Update("CName", "hello").Error
		err = ypg.X.Model(bean).Updates(Abc{CName: "test "}).Error
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("------", "ok", "-----------")
	}

	fmt.Println("------", "demo find", "-----------")
	l := make([]*Abc, 0)
	err = ypg.X.Find(&l).Error
	if err != nil {
		log.Println(err)
		return
	}

	spew.Dump(l)
}
