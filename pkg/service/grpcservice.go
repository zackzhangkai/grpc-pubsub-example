package service

import (
	"context"
	"github.com/zackzhangkai/grpc-pubsub-example/api/proto"
	"github.com/zackzhangkai/grpc-pubsub-example/pkg/pubsub"
	"google.golang.org/protobuf/types/known/emptypb"
	klog "k8s.io/klog/v2"
	"time"

	"strings"
)

type Service struct {
	bus *pubsub.Publisher
	proto.UnimplementedPubSubServiceServer
}

func (s *Service) Publish(ctx context.Context, in *proto.PublishRequest) (*emptypb.Empty, error) {
	data := in.GetTopic() + ":// " + string(in.GetPayload())
	s.bus.Publish(data)
	return &emptypb.Empty{}, nil
}

func (service *Service) Subscribe(in *proto.SubscribeRequest, stream proto.PubSubService_SubscribeServer) error {

	fn := func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.HasPrefix(s, in.Topic+":// ")
		}
		return false
	}

	for v := range service.bus.SubscribeTopic(fn) {
		err := stream.Send(&proto.SubscribeResponse{Payload: []byte(v.(string)[len(in.Topic+":// "):])})
		if err != nil {
			return err
		}
	}

	klog.Info("stream closed")
	return nil
}

func NewService() *Service {
	publisher := pubsub.NewPublisher(1*time.Second, 1)
	return &Service{bus: publisher}
}
