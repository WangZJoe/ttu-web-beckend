package handler

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"log"
	"ttu-backend/apiv1"
	"ttu-backend/internal/model"
	"ttu-backend/internal/service"
)

var (
	Data = hData{}
)

type hData struct{}

// DataInit 加载初始数据，设备列表和对应模型
//func (h *hData) DataInit(ctx context.Context, req *apiv1.DataInitReq) (res *apiv1.DataRes, err error) {
//	res = &apiv1.DataRes{}
//	in := model.EmptyIn{
//		Token:     "123",
//		Timestamp: "2019-12-25T00:00:00.727+0800",
//		Body:      "",
//	}
//_:
//	out, err := service.Mqtt().MqttInit(ctx, in)
//	gconv.Struct(out, &res)
//	return res, nil
//}

// DataHistory 按时间段访问历史数据
func (h *hData) DataHistory(ctx context.Context, req *apiv1.DataHistoryReq) (res *apiv1.DataHistoryRes, err error) {
	res = &apiv1.DataHistoryRes{}
	print(req.FrozenType)
	in := model.MqttDatabaseGetHistoryIn{
		Token:      "633",
		TimeType:   req.TimeType,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		TimeSpan:   req.TimeSpan,
		FrozenType: req.FrozenType,
		Body:       req.Body,
	}
_:
	gconv.Struct(req, &in)
	out, err := service.Mqtt().MqttDatabaseGetHistory(ctx, in)
	if err != nil {
		return nil, err
	}
_:
	gconv.Struct(out.Body, &res)
	return res, nil
}

// DataHistoryN 按上N条访问历史数据
func (h *hData) DataHistoryN(ctx context.Context, req *apiv1.DataHistoryNReq) (res *apiv1.DataHistoryRes, err error) {
	res = &apiv1.DataHistoryRes{}
	print(req.FrozenType)
	in := model.MqttDatabaseGetHistoryInN{
		Token:      "633",
		TimeStamp:  "2022-03-29T23:32:49.223+0800",
		Dev:        req.Dev,
		UpperN:     req.UpperN,
		FrozenType: req.FrozenType,
		Body:       req.Body,
	}
_:
	gconv.Struct(req, &in)
	out, err := service.Mqtt().MqttDatabaseGetHistoryN(ctx, in)
	if err != nil {
		return nil, err
	}
_:
	gconv.Struct(out.Body, &res)
	return res, nil
}

func (h *hData) DataRealtime(ctx context.Context, req *apiv1.DataRealtimeReq) (res *apiv1.DataRealtimeRes, err error) {
	res = &apiv1.DataRealtimeRes{}
	in := "{ \"token\":\"123\",\"timestamp\":\"2022-03-21T09:30:08.230+0800\",\"body\":[{\"dev\":\"" + req.Dev + "\", \"totalcall\":\"" + req.TotalCall + "\", \"body\":[\"" + req.Attribute + "\"]}]}"

_:
	// 传入 请求消息体和 请求主题
	fmt.Println(in)
	out, err := service.Mqtt().MqttDatabaseGetRealtime(ctx, in)
	if err != nil {
		return nil, err
	}
_:
	//gconv.Struct(out.Body[0], &res)
	if len(out.Body) != 0 {
		gconv.Struct(out.Body[0], &res)
		return res, nil
	} else {
		return res, fmt.Errorf("no realtime data")
	}
}

func (h *hData) DataTopo(ctx context.Context, req *apiv1.DataTopoReq) (res *apiv1.DataTopoRes, err error) {
	res = &apiv1.DataTopoRes{}
	log.Println("shit")
	in := "{ \"token\":\"123\",\"timestamp\":\"2022-03-21T09:30:08.230+0800\",\"body\":[]}]}"
	modelName := []string{"LTU", "Switch"}
_:
	out, err := service.Mqtt().MqttDatabaseGetTopo(ctx, in, modelName)
	if err != nil {
		return nil, err
	}
_:
	gconv.Struct(out, &res)
	fmt.Println(res)
	return res, nil
}
