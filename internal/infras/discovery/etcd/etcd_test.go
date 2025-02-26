package etcd

import (
	"log"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/daheige/athena/internal/infras/discovery"
)

func TestRegister(t *testing.T) {
	r, err := New([]string{
		"192.168.10.121:2379",
	})
	if err != nil {
		log.Fatal("failed to init registry,error: ", err)
	}

	s := discovery.Service{
		Name:       "test_etcd_service",
		Address:    "127.0.0.1:3000",
		InstanceID: uuid.New().String(),
	}

	err = r.Register(s, 30)
	if err != nil {
		log.Fatal("failed to register service,error: ", err)
	}

	log.Println("register success")

	for {
		time.Sleep(1 * time.Second)
	}
}

// 获取服务列表
func TestEtcdServices(t *testing.T) {
	r, err := New([]string{
		"192.168.10.121:2379",
	})
	if err != nil {
		log.Fatal(err)
	}

	services, err := r.GetServices("test_etcd_service")
	if err != nil {
		log.Fatal("failed to get services,error: ", err)
	}

	for _, service := range services {
		log.Printf("service:%v\n", service)
	}

	// 剔除某个服务，一般来说，在应用程序退出之前，需要将其剔除该服务
	// r.Deregister("test_etcd_service", "c199fc34-05cf-4f27-bd63-bbcc2b6b9433")
}
