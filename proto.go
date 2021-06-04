package wyLive

type (
	// 创建通道
	LiveCreateRes struct {
		LiveCommon
		Ret CreateRetComm `json:"ret"` // 返回数据结构
	}

	// 批量恢复频道
	ResumeManyLiveRes struct {
		LiveCommon
		Ret ResumeMany `json:"ret"` // 返回数据结构
	}

	// 获取录制文件列表
	ListRecordRes struct {
		LiveCommon
		Ret RecordInfo `json:"ret"` // 返回数据结构
	}

	// 获取直播实时转码相关地址
	ChannelDetailRes struct {
		LiveCommon
		Ret ChannelInfo `json:"ret"` // 返回数据结构
	}

	// 获取一个直播频道的信息
	ChannelListRes struct {
		LiveCommon
		PNum int         `json:"pnum"` // 页数
		Ret  ChannelData `json:"ret"`  // 返回数据结构
	}

	// 直播封面
	LiveCoverRes struct {
		LiveCommon
		RequestId string `json:"requestId"` // 全局唯一请求id
	}

	// 获取直播实时转码相关地址
	LivePushUrlRes struct {
		LiveCommon
		RequestId string      `json:"requestId"` // 全局唯一请求id
		Ret       PushRetComm `json:"ret"`       // 返回数据结构
	}

	LiveCommon struct {
		Code int    `json:"code"` // 错误码
		Msg  string `json:"msg"`  // 错误信息
	}

	LoadInterLiveRes struct {
		Stats       int    `json:"stats"`       // 房间状态1：初始状态。 2：进行中。 3：正常结束。 4：异常结束。
		CreateTime  int64  `json:"createtime"`  // 房间创建时间
		DestroyTime int64  `json:"destroytime"` // 房间结束时间
		Cid         int64  `json:"cid"`         // 间 ID。该 ID 为创建房间接口调用成功后返回的房间 ID
		Uid         int64  `json:"uid"`         // 房间创建者的用户 ID
		Total       int64  `json:"total"`       // 房间内在线用户总数
		CName       string `json:"cname"`       // 房间名称
	}

	LoadInterMembersRes struct {
		Cid   int64  `json:"cid"`   // 间 ID。该 ID 为创建房间接口调用成功后返回的房间 ID
		Uid   int64  `json:"uid"`   // 房间创建者的用户 ID
		Total int64  `json:"total"` // 房间内在线用户总数
		CName string `json:"cname"` // 房间名称
	}

	InterRemoveMembersRes struct {
		Uid []int64 `json:"uids"` // 被移除出房间的成员用户 ID 列表。
	}


	InterLiveComm struct {
		Code int   `json:"code"` // 200 错误码
		Cid  int64 `json:"cid"`  // 房间 ID
	}

	CreateRetComm struct {
		Ctime       int64  `json:"ctime"`       // 创建频道的时间戳
		Cid         string `json:"cid"`         // 频道ID，32位字符串
		Name        string `json:"name"`        // 频道名称
		PushUrl     string `json:"pushUrl"`     // 推流地址
		HttpPullUrl string `json:"httpPullUrl"` // http拉流地址
		HlsPullUrl  string `json:"hlsPullUrl"`  // hls拉流地址
		RtmpPullUrl string `json:"rtmpPullUrl"` // rtmp拉流地址
	}

	ResumeMany struct {
		SuccessList []string `json:"successList"` // 成功禁用cid列表
	}

	RecordInfo struct {
		PNum         int64       `json:"pnum"`         // 当前页
		TotalRecords int64       `json:"totalRecords"` // 总记录数
		TotalPNum    int64       `json:"totalPnum"`    // 总页数
		Records      int64       `json:"records"`      // 单页记录数
		VideoList    []VideoList `json:"videoList"`    // 录制视频列表
	}

	VideoList struct {
		Vid          int    `json:"vid"`            // 视频文件ID
		VideoName    string `json:"video_name"`     // 录制后文件名，格式为filename_YYYYMMDD-HHmmss_YYYYMMDD-HHmmss, 文件名_录制起始时间（年月日时分秒) -录制结束时间（年月日时分秒)
		OrigVideoKey string `json:"orig_video_key"` // 视频文件在点播桶中的存储路径
	}

	ChannelData struct {
		List []ChannelInfo `json:"list"` // 列表
	}

	ChannelInfo struct {
		Status     int    `json:"status"`     // 频道状态（0：空闲； 1：直播； 2：禁用； 3：直播录制）
		Type       int    `json:"type"`       // 频道类型 ( 0 : rtmp, 1 : hls, 2 : http)
		NeedRecord int    `json:"needRecord"` // 1-开启录制； 0-关闭录制
		Format     int    `json:"format"`     // 1-flv； 0-mp4
		Duration   int    `json:"duration"`   // 录制切片时长(分钟)，默认120分钟
		CTime      int64  `json:"ctime"`      // 创建频道的时间戳
		Cid        string `json:"cid"`        // 频道ID，32位字符串
		Name       string `json:"name"`       // 频道名称
		Filename   string `json:"filename"`   // 录制后文件名
	}

	PushRetComm struct {
		Status               int                   `json:"status"`               // 拉流转码状态0->暂未开通,1->已开通
		PushUrl              string                `json:"push_url"`             // 推流地址
		HttpPullUrl          string                `json:"httpPullUrl"`          // http拉流地址
		HlsPullUrl           string                `json:"hlsPullUrl"`           // hls拉流地址
		RtmpPullUrl          string                `json:"rtmpPullUrl"`          // rtmp拉流地址
		TranscodeHttpPullUrl TranscodeHttpPullUrlC `json:"transcodeHttpPullUrl"` // 实时转码http拉流地址,当status=0时该数据结点不存在
		TranscodeRtmpPullUrl TranscodeHttpPullUrlC `json:"transcodeRtmpPullUrl"` // 实时转码rtmp拉流地址,当status=0时该数据结点不存在
		TranscodeHlsPullUrl  TranscodeHttpPullUrlC `json:"transcodeHlsPullUrl"`  // 实时转码hls拉流地址,当status=0时该数据结点不存在
	}

	TranscodeHttpPullUrlC struct {
		F1280 string `json:"1280"` // 16:9,1280x720,1600k格式拉流地址
		F960  string `json:"960"`  // 16:9,960x540,1000k格式拉流地址
		F640  string `json:"640"`  // 16:9,640x360,600k格式拉流地址
		F320  string `json:"320"`  // 16:9,320x180,300k格式拉流地址
		F540  string `json:"540"`  // 9:16,540x960,1000k格式拉流地址
		F360  string `json:"360"`  // 9:16,360x640,600k格式拉流地址
		F180  string `json:"180"`  // 9:16,180x320,300k格式拉流地址
	}
)
