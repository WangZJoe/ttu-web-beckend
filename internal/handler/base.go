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
	Base = hBase{}
)

type hBase struct{}

func (h *hBase) DeviceList(ctx context.Context, req *apiv1.DeviceListReq) (res *apiv1.DeviceListRes, err error) {
	log.Println("shit")
	in := model.BaseDeviceListIn{}
	_ = gconv.Struct(req, &in)
	out, err := service.Base().BaseDeviceList(ctx, in)
	res = &apiv1.DeviceListRes{}
	_ = gconv.Struct(out, &res)
	return res, err
}
func (h *hBase) Realtime(ctx context.Context, req *apiv1.RealtimeReq) (res *apiv1.RealtimeRes, err error) {
	in := model.BaseRealtimeIn{}
	_ = gconv.Struct(req, &in)
	out, err := service.Base().BaseRealtime(ctx, in)
	res = &apiv1.RealtimeRes{}
	_ = gconv.Struct(out, &res)
	return res, err
}
func (h *hBase) Record(ctx context.Context, req *apiv1.RecordReq) (res *apiv1.RecordRes, err error) {
	in := model.BaseRecordIn{}
	fmt.Println(in.TimeSpanNumber)
	_ = gconv.Struct(req, &in)
	out, err := service.Base().BaseRecord(ctx, in)
	res = &apiv1.RecordRes{}
	_ = gconv.Struct(out, &res)
	return res, err
}
func (h *hBase) Alarm(ctx context.Context, req *apiv1.AlarmReq) (res *apiv1.AlarmRes, err error) {
	in := model.BaseAlarmIn{}
	_ = gconv.Struct(req, &in)
	out, err := service.Base().BaseAlarm(ctx, in)
	res = &apiv1.AlarmRes{}
	_ = gconv.Struct(out, &res)
	return res, err
}
func (h *hBase) FaultWaveform(ctx context.Context, req *apiv1.FaultWaveformReq) (res *apiv1.FaultWaveformRes, err error) {
	in := model.BaseFaultWaveformIn{}
	_ = gconv.Struct(req, &in)
	out, err := service.Base().BaseFaultWaveform(ctx, in)
	res = &apiv1.FaultWaveformRes{}
	_ = gconv.Struct(out, &res)
	return res, err
}
func (h *hBase) SetConfig(ctx context.Context, req *apiv1.SetConfigReq) (res *apiv1.SetConfigRes, err error) {
	in := model.BaseSetConfigIn{}
	_ = gconv.Struct(req, &in)
	out, err := service.Base().BaseSetConfig(ctx, in)
	res = &apiv1.SetConfigRes{}
	_ = gconv.Struct(out, &res)
	return res, err
}

func (h *hBase) GetConfig(ctx context.Context, req *apiv1.GetConfigReq) (res *apiv1.GetConfigRes, err error) {
	in := model.BaseGetConfigIn{}
	_ = gconv.Struct(req, &in)
	out, err := service.Base().BaseGetConfig(ctx, in)
	res = &apiv1.GetConfigRes{}
	_ = gconv.Struct(out, &res)
	return res, err
}

func (h *hBase) DayAnalysis(ctx context.Context, req *apiv1.DayAnaReq) (res *apiv1.DayAnaRes, err error) {
	in := model.BaseDayAnaIn{}
	_ = gconv.Struct(req, &in)
	out, err := service.Base().BaseDayAnalysis(ctx, in)
	//fmt.Println(out.DayAna)
	res = &apiv1.DayAnaRes{}
	_ = gconv.Struct(out, &res)
	return res, err
}
