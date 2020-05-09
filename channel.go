package main

import "log"

var (
	DefalutChannelName       = "CSO2-Channel"
	MAXCHANNELNUM      uint8 = 16
)

//频道信息，隶属于Server频道,用于请求服务器和请求频道
type channelInfo struct {
	channelID   uint8
	channelName []byte
	unk00       uint16
	unk01       uint16
	unk02       uint8
	unk03       uint8
	unk04       uint8
	nextRoomID  uint8
	roomNum     uint16
	rooms       []roomInfo
}

// //频道内容,包括房间,用于请求频道
// type channel struct {
// 	channelID uint8
// 	roomNum   uint16
// 	rooms     []room
// }

func BuildChannelList(num uint8, channels []channelInfo) []byte {
	var list []byte
	for i := 0; i < int(num); i++ {
		temp := make([]byte, 9+len(channels[i].channelName))
		offset := 0
		WriteUint8(&temp, channels[i].channelID, &offset)
		WriteString(&temp, channels[i].channelName, &offset)
		WriteUint16(&temp, channels[i].unk00, &offset)
		WriteUint16(&temp, channels[i].unk01, &offset)
		WriteUint8(&temp, channels[i].unk02, &offset)
		WriteUint8(&temp, channels[i].unk03, &offset)
		WriteUint8(&temp, channels[i].unk04, &offset)
		for j := 0; j < len(temp); j++ {
			list = append(list, temp[j])
		}
	}
	return list
}

func newChannelInfo(name []byte, chlsrv channelServer) channelInfo {
	return channelInfo{
		getNewChannelID(chlsrv),
		name,
		4,
		0x1F4,
		1,
		0,
		1,
		1,
		0,
		[]roomInfo{},
	}
}

//getNewChannelID() 暂定
func getNewChannelID(chlsrv channelServer) uint8 {
	if chlsrv.channelCount > MAXCHANNELNUM {
		log.Fatalln("Channel is too much ! Unable to create more !")
		//ID=0 是非法的
		return 0
	}
	var intbuf [20]uint8
	//哈希思想
	for i := 0; i < int(chlsrv.channelCount); i++ {
		intbuf[chlsrv.channels[i].channelID] = 1
	}
	//找到空闲的ID
	for i := 1; i < 20; i++ {
		if intbuf[i] == 0 {
			//找到了空闲ID
			return uint8(i)
		}
	}
	return 0
}

//通过ID获取频道
func getChannelWithID(id uint8, chlsrv channelServer) *channelInfo {
	count := chlsrv.channelCount
	for i := 0; i < int(count); i++ {
		//log.Println("ChannelIndex:", strconv.Itoa(int(GameServer.channels[i].channelID)))
		if chlsrv.channels[i].channelID == id {
			return &(chlsrv.channels[i])
		}
	}
	return nil
}

//新建房间
func addChannelRoom(host uint32, id uint8, chlsrv channelServer) {
	chlptr := getChannelWithID(id, chlsrv)
	// room := roomInfo{
	// 	(*chlptr).nextRoomID,

	// }
	//(*chlptr).rooms = append((*chlptr).rooms, room)
	(*chlptr).roomNum++

}

//删除频道房间
func delChannelRoom() {

}
