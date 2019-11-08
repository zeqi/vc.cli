package services

import (
	"context"
	"fmt"
	"sync"
	"time"

	// "github.com/micro/go-grpc"
	// "github.com/micro/go-plugins/registry/kubernetes"
	// k8s "github.com/micro/kubernetes/go/micro"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"

	// k8s "github.com/micro/kubernetes/go/micro"
	"vc.cli/models"
	pb_sequence "vc.pb/sequence"
)

type Registrar struct {
	client.Client
	MicroApi       micro.Service
	SequenceClient pb_sequence.SequenceService
}

var instance *Registrar
var once sync.Once

func Init(config models.MicroServices) *Registrar {
	once.Do(func() {
		instance = &Registrar{}
		instance.initService(config.MicroApi)
		instance.initClient(config.MicroMongo)
	})
	return instance
}

func (o *Registrar) initService(config models.MicroConfig) {
	// Create a new service. Optionally include some options here.
	// adds := strings.Split(config.Etcd.Addrs, ",")
	// ros := registry.Addrs(adds...)
	// r := etcdv3.NewRegistry(func(op *registry.Options) {
	// 	op.Addrs = strings.Split(config.Etcd.Addrs, ",")
	// })
	service := grpc.NewService(
		micro.Name(config.Name),
		micro.Version(config.Version),
		micro.RegisterTTL(time.Minute),
		micro.RegisterInterval(time.Second*30),
		// micro.WrapClient(logWrap),
		// micro.WrapClient()
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
	// service.Init()
	o.MicroApi = service

}

func (o *Registrar) initClient(config models.MicroConfig) {
	o.SequenceClient = pb_sequence.NewSequenceService(config.Name, o.MicroApi.Client())
}

type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Printf("[wrapper] client request to service: %s endpoint: %s\n", req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

// implements client.Wrapper as logWrapper
func logWrap(c client.Client) client.Client {
	return &logWrapper{c}
}

// Return Registrar instance
func GetClient() *Registrar {
	return instance
}
