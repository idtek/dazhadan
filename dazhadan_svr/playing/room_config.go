package playing

import (
	"time"
)

type RoomConfig struct {
	NeedPlayerNum			int32        `json:"need_player_num"`
	WinCoin				int32        `json:"win_coin"`
	ShuangjiCoin			int32        `json:"shuangji_coin"`
	DaduCoin			int32        `json:"dadu_coin"`
	PrizeCoin			int32        `json:"prize_coin"`
	InitType			int32        `json:"init_type"`
	MaxPlayGameCnt			int	     `json:"max_play_game_cnt"`	//最大的游戏局数
	RandomDropNum			int          `json:"random_drop_num"`	//随机出牌张数

	WaitPlayerEnterRoomTimeout	int        `json:"wait_player_enter_room_timeout"`
	WaitPlayerOperateTimeout	int        `json:"wait_player_operate_timeout"`
	WaitDaduSec                	time.Duration      `json:"wait_dadu_sec"`	//等待打独时长
	WaitDropSec                	time.Duration      `json:"wait_drop_sec"`	//等待出牌时长
	WaitReadySec              	time.Duration      `json:"wait_ready_sec"`	//等待准备时长
	AfterSwitchPositionSleep        time.Duration      `json:"after_switch_position_sleep"`	//交换位置后sleep时长
}

func NewRoomConfig() *RoomConfig {
	return &RoomConfig{}
}

func (config *RoomConfig) Init(score_type, prize_type, init_type int32) {
	if score_type == 1 {
		config.WinCoin = 10
		config.ShuangjiCoin = 15
		config.DaduCoin = 20
	}else {
		config.WinCoin = 10
		config.ShuangjiCoin = 20
		config.DaduCoin = 30
	}

	if score_type == 1 {
		config.PrizeCoin = 3
	}else {
		config.PrizeCoin = 5
	}

	config.InitType = init_type
	config.NeedPlayerNum = 4
	config.MaxPlayGameCnt = 3
	config.RandomDropNum = 3

	config.WaitPlayerEnterRoomTimeout = 300
	config.WaitPlayerOperateTimeout = 300
	config.WaitDaduSec = 5
	config.WaitDropSec = 3
	config.WaitReadySec = 5
	config.AfterSwitchPositionSleep = 1
}