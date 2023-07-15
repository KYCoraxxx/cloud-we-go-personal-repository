package demo

import (
	"context"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"math/rand"
)

func (lb *randomBalancer) GetPicker(result discovery.Result) loadbalance.Picker {
	return &randomPicker{
		instances: result.Instances,
	}
}

func (lb *randomBalancer) Name() string {
	return "Random LoadBalancer"
}

type randomPicker struct {
	instances []discovery.Instance
}

func (rp *randomPicker) Next(ctx context.Context, request interface{}) discovery.Instance {
	return rp.instances[rand.Int()%len(rp.instances)]
}
