package wyLive

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// 初始化
func InitWYLvChannel(appKey, appSecret string) WYLv {
	return &wYLvChannel{appKey: appKey, appSecret: appSecret, liveBaseU: "https://vcloud.163.com/app", interBaseU: "https://logic-dev.netease.im"}
}

type WYLv interface {
	// 直播
	CreateChannel(name string) (res LiveCreateRes, err error)                                       // 创建一个直播频道
	UpdateChannel(name, cid string) (res LiveCommon, err error)                                     // 修改频道
	DelChannel(cid string) (res LiveCommon, err error)                                              // 删除频道
	PauseChannel(cid string) (res LiveCommon, err error)                                            // 禁用频道
	PauseManyChannel(cidList []string) (res ResumeManyLiveRes, err error)                           // 批量禁用频道
	ResumeChannel(cid string) (res LiveCommon, err error)                                           // 恢复用户被禁用的频道
	ResumeManyChannel(cidList []string) (res ResumeManyLiveRes, err error)                          // 批量恢复频道
	LoadRecordList(cid string, pnum, records int) (res ListRecordRes, err error)                    // 获取录制文件列表
	LoadChannelInfo(cid string) (res ChannelDetailRes, err error)                                   // 获取一个直播频道的信息
	LoadChannelList(records, pnum, sort, status int, ofield string) (res ChannelListRes, err error) // 获取频道列表
	ResetPushAddr(cid string) (res LiveCreateRes, err error)                                        // 重新获取推拉流地址
	LiveSaveRec(cid, name string) (bs LiveCommon, err error)                                        // 设置录制配置
	GetLiveCover(cid string) (bs LiveCoverRes, err error)                                           // 直播截图
	LivePushUrl(cid string) (res LivePushUrlRes, err error)                                         // 获取直播实时转码相关地址
	// 互动直播
	CreateRoom(name string, uid int64) (res InterLiveComm, err error)                // 创建房间
	LoadRoom(cid int64, name string) (res LoadInterLiveRes, err error)               // 查看房间信息
	LoadRoomMembers(cid int64, name string) (res LoadInterMembersRes, err error)     // 查看房间内成员信息
	RemoveRoomMembers(cid, uid int64, name string) (res InterLiveComm, err error)    // 移除成员
	LoadRemoveMembers(cid int64, name string) (res InterRemoveMembersRes, err error) // 查看被移除成员列表
	ReturnRemoveMembers(cid, uid int64, name string) (res InterLiveComm, err error)  // 撤销移除
	RemoveRoom(cid int64, name string) (res InterLiveComm, err error)                // 删除房间
}
