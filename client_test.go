package xenapi

import (
	"fmt"
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"reflect"
	"testing"
)

func TestAuthentication(t *testing.T) {
	client, err := NewClient("http://10.71.56.91")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	sessionRef, err := client.Session.LoginWithPassword("root", "Lo2RQQS4q2Sf", "1.0", "terraform")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	hostRefs, err := client.Host.GetAll(sessionRef)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	records, err := client.Host.GetAllRecords(sessionRef)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	fmt.Println(records)
	fmt.Println(reflect.TypeOf(records))

	editions, err := client.Host.GetEditions(sessionRef, hostRefs[0])
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	fmt.Println(editions)
	fmt.Println(reflect.TypeOf(editions))
	

	err = client.Session.Logout(sessionRef)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
}

func TestMain(m *testing.M) {
	flag.Parse()
	log.SetOutput(os.Stdout)
	os.Exit(m.Run())
}
