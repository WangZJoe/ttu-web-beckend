package consts

import "time"

const (
	uid        = "app"
	MQTTBroker = "mqtt://sinpower.3322.org:11883"
	//MQTTLAN    = "mqtt://192.168.2.33:1883"
	MQTTLAN = "mqtt://47.110.134.175:8183"
	MQTTQos = 0x0

	Publish_register_get      = uid + "/get/request/database/register"
	Publish_realtime_data_get = uid + "/get/request/database/realtime"
	Publish_history_data_get  = uid + "/get/request/database/history"
	Publish_alarm_data_get    = uid + "/get/request/database/SOE"
	Publish_getParams         = uid + "/get/request/TestApp/params"
	Publish_setParams         = uid + "/set/request/TestApp/params"

	Received_device_get        = "database/get/response/" + uid + "/vT"
	Received_modelschema_get   = "database/get/response/" + uid + "/modelschema"
	Received_model_get         = "database/get/response/" + uid + "/model"
	Received_guid_get          = "database/get/response/" + uid + "/guid"
	Received_initData          = "database/action/response/" + uid + "/InitData"
	Received_register_get      = "database/get/response/" + uid + "/register"
	Received_realtime_data_get = "database/get/response/" + uid + "/realtime"
	Received_history_data_get  = "database/get/response/" + uid + "/history"
	Received_alarm_data_get    = "database/get/response/" + uid + "/SOE"
	Received_getParams         = "TestApp/get/response/" + uid + "/params"
	Received_setParams         = "TestApp/set/response/" + uid + "/params"
)
const (
	ContextKey = "ContextKey"
)

const (
	JwtPrivateKey = "^:P}*39a3LWxzXq="
	JwtExpTime    = time.Minute * 1
	JwtIssuerName = "niubx"
)

const (
	JwtHeaderKey    = "Authorization"
	JwtHeaderPrefix = "Bearer "
)

var ModelConfig = map[string]struct {
	Name    string
	Mapping map[string]string
}{
	"LTU": {
		Name: "线路终端",
		Mapping: map[string]string{
			"AMax_phsA": "InMax",
			"BMax_phsA": "InAvg",
			"A_phsA":    "Ia",
			"B_phsA":    "Ib",
			"C_phsA":    "Ic",
			"PhV_phsA":  "Ua",
			"PhV_phsB":  "Ub",
			"PhV_phsC":  "Uc",
			"Tmp":       "T",
			"EnvHum":    "H",
		},
	},
	"Switch": {
		Name: "智能开关",
		Mapping: map[string]string{
			"AMax_phsA": "InMax",
			"BMax_phsA": "InAvg",
			"A_phsA":    "Ia",
			"B_phsA":    "Ib",
			"C_phsA":    "Ic",
			"PhV_phsA":  "Ua",
			"PhV_phsB":  "Ub",
			"PhV_phsC":  "Uc",
			"Tmp":       "T",
			"EnvHum":    "H",
		},
	},
}

var AlarmDict = map[string]string{
	"IndN":              "开入量N",
	"PTOV_Open_Op_phsA": "A相断相",
	"PTUV_Loss_Op_phsA": "A相失压",
	"PTUV_Op_phsA":      "A相欠压",
	"PTOV_Op_phsA":      "A相过压",
	"PTOC_Op_phsA":      "A相过流",
	"PTUC_Op_phsA":      "A相失流",
	"PTOC_Ovld_Op_phsA": "A相过载",
	"PTOC_Hvld_Op_phsA": "A相重载",
	"PTUC_Open_Op_phsA": "A相断流",
	"RDIR_dirPhsA":      "A相潮流反向",
	"PTOC_Dly_Op_phsA":  "A相短路短延时",
	"PIOC_Op_phsA":      "A相短路瞬时保护",
	"PTLK_Op_phsA":      "A相剩余电流保护",
	"PTUV_lim_Op_phsB":  "B相电压越下限",
	"PTOV_lim_Op_phsB":  "B相电压越上限",
	"PTOV_Open_Op_phsB": "B相断相",
	"PTUV_Loss_Op_phsB": "B相失压",
	"PTUV_Op_phsB":      "B相欠压",
	"PTOV_Op_phsB":      "B相过压",
	"PTOC_Op_phsB":      "B相过流",
	"PTUC_Op_phsB":      "B相失流",
	"PTOC_Ovld_Op_phsB": "B相过载",
	"PTOC_Hvld_Op_phsB": "B相重载",
	"PTUC_Open_Op_phsB": "B相断流",
	"RDIR_dirPhsB":      "B相潮流反向",
	"PTOC_Dly_Op_phsB":  "B相短路短延时保护",
	"PIOC_Op_phsB":      "B相短路瞬时保护",
	"PTLK_Op_phsB":      "B相剩余电流保护",
	"PTUV_lim_Op_phsC":  "C相电压越下限",
	"PTOV_lim_Op_phsC":  "C相电压越上限",
	"PTOV_Open_Op_phsC": "C相断相",
	"PTUV_Loss_Op_phsC": "C相失压",
	"PTUV_Op_phsC":      "C相欠压",
	"PTOV_Op_phsC":      "C相过压",
	"PTOC_Op_phsC":      "C相过流",
	"PTUC_Op_phsC":      "C相失流",
	"PTOC_Ovld_Op_phsC": "C相过载",
	"PTOC_Hvld_Op_phsC": "C相重载",
	"PTUC_Open_Op_phsC": "C相断流",
	"RDIR_dirPhsC":      "C相潮流反向",
	"PTOC_Dly_Op_phsC":  "C相短路短延时保护",
	"PIOC_Op_phsC":      "C相短路瞬时保护",
	"PTLK_Op_phsC":      "C相剩余电流保护",
	"SeqAAlm":           "电压逆相序",
	"SeqVAlm":           "电流逆相序",
	"ImbAAlm":           "电压不平衡",
	"ImbVAlm":           "电流不平衡",
	"PwrSupAlm":         "辅助电源失电",
	"PwrOffAlm":         "停电",
	"PwrOnAlm":          "上电",
	"DmdUpAlm":          "需量超限",
	"DoorAlm":           "门磁报警",
	"FloodAlm":          "水浸报警",
	"SmokAlm":           "烟感报警",
	"ShockAlm":          "振动报警",
	"Oilchrom":          "油色谱",
	"CSWI_ pos":         "开关机械状态",
	"TEST_pos":          "断路器自检状态",
	"ImbLAlm":           "负荷不平衡",
}
