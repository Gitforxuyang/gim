package ratelimiter

import (
	"errors"
	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/hotspot"
)

func Init() {
	_ = api.InitDefault()
	hotspot.LoadRules([]*hotspot.Rule{
		{
			Resource:        "request_limit_auth",
			MetricType:      hotspot.QPS,
			ControlBehavior: hotspot.Reject,
			ParamIndex:      0,
			Threshold:       30,
			DurationInSec:   1,
		},
		{
			Resource:        "request_limit_wait_auth",
			MetricType:      hotspot.QPS,
			ControlBehavior: hotspot.Reject,
			ParamIndex:      0,
			Threshold:       5,
			DurationInSec:   1,
		},
	})
}

func Token(userId int64, addr string) (*base.SentinelEntry, error) {
	if userId != 0 {
		entry, err := api.Entry("request_limit_auth",
			api.WithTrafficType(base.Inbound),
			api.WithArgs(userId))
		if err != nil {
			return nil, errors.New("rate limit auth")
		}
		return entry, nil
	} else {
		entry, err := api.Entry("request_limit_wait_auth",
			api.WithTrafficType(base.Inbound),
			api.WithArgs(addr))
		if err != nil {
			return nil, errors.New("rate limit wait auth")
		}
		return entry, nil
	}
}
