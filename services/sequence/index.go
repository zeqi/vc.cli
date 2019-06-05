package sequence

import (
	"context"
	"fmt"

	"vc.cli/services"
	pb "vc.pb/sequence"
)

type Service struct {
	Client pb.SequenceService
}

func NewService() Service {
	return Service{Client: services.GetClient().SequenceClient}
}

func (o *Service) Create(name string, comments string) (*pb.ResDoc, error) {
	rsp, err := o.Client.Create(context.Background(), &pb.ReqCreate{Name: name, Comments: comments})
	if err != nil {
		fmt.Errorf("Sequence Create Error %s", err)
		return rsp, err
	}
	return rsp, err
}

func (o *Service) FInd(skip int64, limit int64) (*pb.ResDocs, error) {
	rsp, err := o.Client.Find(context.Background(), &pb.ReqFind{Skip: skip, Limit: limit, Condition: &pb.Model{}})
	if err != nil {
		fmt.Errorf("Sequence FInd Error %s", err)
		return rsp, err
	}
	return rsp, err
}

func (o *Service) FindDocsAndCount(skip int64, limit int64) (*pb.ResDocsAndCount, error) {
	rsp, err := o.Client.FindDocsAndCount(context.Background(), &pb.ReqFind{Skip: skip, Limit: limit, Condition: &pb.Model{}})
	if err != nil {
		fmt.Errorf("Sequence FindDocsAndCount Error %s", err)
		return rsp, err
	}
	return rsp, err
}

func (o *Service) FIndOne(skip int64, limit int64) (*pb.ResDoc, error) {
	rsp, err := o.Client.FindOne(context.Background(), &pb.ReqFind{Skip: skip, Limit: limit, Condition: &pb.Model{}})
	if err != nil {
		fmt.Errorf("Sequence FIndOne Error %s", err)
		return rsp, err
	}
	return rsp, err
}

func (o *Service) FIndById(id string) (*pb.ResDoc, error) {
	rsp, err := o.Client.FindById(context.Background(), &pb.Model{Id: id})
	if err != nil {
		fmt.Errorf("Sequence FindById Error %s", err)
		return rsp, err
	}
	return rsp, err
}

func (o *Service) IncByName(name string) (*pb.ResDoc, error) {
	rsp, err := o.Client.IncByName(context.Background(), &pb.Model{Name: name})
	if err != nil {
		fmt.Errorf("Sequence IncByName Error %s", err)
		return rsp, err
	}
	return rsp, err
}
