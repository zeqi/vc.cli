package services

import (
	"sync"
	"time"

	// "github.com/micro/go-grpc"
	// "github.com/micro/go-plugins/registry/kubernetes"
	// k8s "github.com/micro/kubernetes/go/micro"
	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
	"vc.cli/models"
	pb_sequence "vc.pb/sequence"
)

type Client struct {
	MicroApi       micro.Service
	SequenceClient pb_sequence.SequenceService
}

var instance *Client
var once sync.Once

func Init(config models.MicroServices) *Client {
	once.Do(func() {
		instance = &Client{}
		instance.InitService(config.MicroApi)
		instance.InitClient(config.MicroMongo)
	})
	return instance
}

func (o *Client) InitService(config models.MicroConfig) {
	// Create a new service. Optionally include some options here.
	// adds := strings.Split(config.Etcd.Addrs, ",")
	// ros := registry.Addrs(adds...)
	// r := etcdv3.NewRegistry(func(op *registry.Options) {
	// 	op.Addrs = strings.Split(config.Etcd.Addrs, ",")
	// })
	service := k8s.NewService(
		micro.Name(config.Name),
		micro.Version(config.Version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	// micro.Registry(r),
	// micro.Selector(static.NewSelector()),
	)
	// Init will parse the command line flags.
	// service.Init(micro.Flags(
	// 	cli.StringFlag{
	// 		Name:        config.Name,
	// 		Value:       config.Name,
	// 		Destination: &sequenceService,
	// 	}))
	o.MicroApi = service

}

func (o *Client) InitClient(config models.MicroConfig) {
	o.SequenceClient = pb_sequence.NewSequenceService(config.Name, o.MicroApi.Client())
}

func GetClient() *Client {
	return instance
}
