package wyLive

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"unsafe"
)

// 网易云 直播创建
type wYLvChannel struct {
	appKey     string // app key
	appSecret  string // app secret
	liveBaseU  string // 直播api url
	interBaseU string // 互动api url
}

/* 创建一个直播频道 name 必须唯一 创建直播名时需要加入唯一值
* param name string	频道名（最大长度64个字符）	是
* param type int	频道类型 ( 0 : rtmp)	否
 */
func (w *wYLvChannel) CreateChannel(name string) (res LiveCreateRes, err error) {
	err = w.sendHttp("/channel/create", map[string]interface{}{"name": name, "type": 0}, &res)
	return
}

/* 修改频道
* param name string	频道名（最大长度64个字符）是
* param cid  string	频道ID，32位字符串 是
* param type int	频道类型 ( 0 : rtmp)	否
 */
func (w *wYLvChannel) UpdateChannel(name, cid string) (res LiveCommon, err error) {
	err = w.sendHttp("/channel/update", map[string]interface{}{"name": name, "cid": cid, "type": 0}, &res)
	return
}

/* 删除频道
* param cid string	频道ID	是
 */
func (w *wYLvChannel) DelChannel(cid string) (res LiveCommon, err error) {
	err = w.sendHttp("/channel/delete", map[string]interface{}{"cid": cid}, &res)
	return
}

/* 禁用频道
* param cid string	频道ID	是
 */
func (w *wYLvChannel) PauseChannel(cid string) (res LiveCommon, err error) {
	err = w.sendHttp("/channel/pause", map[string]interface{}{"cid": cid}, &res)
	return
}

/* 批量禁用频道
* param cidList []string	频道ID	是
 */
func (w *wYLvChannel) PauseManyChannel(cidList []string) (res ResumeManyLiveRes, err error) {
	err = w.sendHttp("/channellist/pause", map[string]interface{}{"cidList": cidList}, &res)
	return
}

/* 恢复用户被禁用的频道
* param cid string	频道ID	是
 */
func (w *wYLvChannel) ResumeChannel(cid string) (res LiveCommon, err error) {
	err = w.sendHttp("/channel/resume", map[string]interface{}{"cid": cid}, &res)
	return
}

/* 批量恢复频道
* param cidList	[]string	频道ID列表	是
 */
func (w *wYLvChannel) ResumeManyChannel(cidList []string) (res ResumeManyLiveRes, err error) {
	err = w.sendHttp("/channellist/resume", map[string]interface{}{"cidList": cidList}, &res)
	return
}

/* 获取录制文件列表 获取某频道录制视频文件列表，按生成时间由近至远提供分页。
* param cid	    string	频道ID，32位字符串	是
* param records	int	单页记录数，默认值为10，最多1000条	否
* param pnum	int	要取第几页，默认值为1	否
 */
func (w *wYLvChannel) LoadRecordList(cid string, pnum, records int) (res ListRecordRes, err error) {
	err = w.sendHttp("/videolist", map[string]interface{}{"cid": cid, "records": records, "pnum": pnum}, &res)
	return
}

/* 获取一个直播频道的信息
* param cid	string	频道ID，32位字符串
 */
func (w *wYLvChannel) LoadChannelInfo(cid string) (res ChannelDetailRes, err error) {
	err = w.sendHttp("/channelstats", map[string]interface{}{"cid": cid}, &res)
	return
}

/*
* 获取频道列表
* records	int	单页记录数，默认值为10，最多1000条	否
* pnum	    int	要取第几页，默认值为1	否
* ofield	string	排序的域，支持的排序域为：ctime（默认）	否
* sort	    int	升序还是降序，1升序，0降序，默认为desc	否
* status	int	筛选频道状态，status取值：（0：空闲,1：直播，2：禁用，3：录制中） 否
 */
func (w *wYLvChannel) LoadChannelList(records, pnum, sort, status int, ofield string) (res ChannelListRes, err error) {
	err = w.sendHttp("/channellist", map[string]interface{}{"records": records, "pnum": pnum, "ofield": ofield, "sort": sort, "status": status}, &res)
	return
}

/*
* 重新获取推拉流地址
* param cid	string	频道ID，32位字符串 是
 */
func (w *wYLvChannel) ResetPushAddr(cid string) (res LiveCreateRes, err error) {
	err = w.sendHttp("/address", map[string]interface{}{"cid": cid}, &res)
	return
}

/*
* 设置录制配置
* param cid	String	频道ID	是
* param duration int	录制周期, 取值范围[5,120], 单位分钟, 默认120分钟	是
* param format	 int	录制切片视频容器格式, 取值范围[0-2], 默认1, 含义：0-mp4, 1-flv, 2-mp3	是
* param filename string	录制切片文件名前缀, 要求非空, 字符长度不大于64, 默认为「创建频道」时传入的频道名称(name参数)	是
 */
func (w *wYLvChannel) LiveSaveRec(cid, name string) (res LiveCommon, err error) {
	err = w.sendHttp("/record/channel/updateConfig", map[string]interface{}{"cid": cid, "duration": 300, "format": 1, "filename": name}, &res)
	return
}

/*
* 直播截图
* param cid	         string	频道ID	是
* param timeInterval int 截图周期，取值范围[5,3600]，单位秒，默认5	否
* param imageFormat	 string	图片格式：jpg、png，默认jpg	否
* param imageWidth	 int 图片宽度，默认0	否
* param imageHeight	 int 图片高度，默认0	否
 */
func (w *wYLvChannel) GetLiveCover(cid string) (res LiveCoverRes, err error) {
	err = w.sendHttp("/channel/snapshot/updateconfig", map[string]interface{}{"cid": cid}, &res)
	return
}

/*
* 获取直播实时转码相关地址
* param cid	string	频道ID，32位字符串	是
 */
func (w *wYLvChannel) LivePushUrl(cid string) (res LivePushUrlRes, err error) {
	err = w.sendHttp("/transcodeAddress", map[string]interface{}{"cid": cid}, &res)
	return
}

// 发送请求POST
func (w *wYLvChannel) sendHttp(url2 string, params map[string]interface{}, res interface{}) (err error) {
	bytesData, err1 := json.Marshal(params)
	if err1 != nil {
		err = err1
		return
	}
	request, err1 := http.NewRequest("POST", w.liveBaseU+url2, bytes.NewReader(bytesData))
	if err1 != nil {
		err = err1
		return
	}
	var nt = time.Now()
	var tm = strconv.FormatInt(nt.Unix(), 10)
	var rand = strconv.FormatInt(nt.UnixNano(), 10)

	request.Header.Set("AppKey", w.appKey)
	request.Header.Set("Nonce", rand)
	request.Header.Set("CurTime", tm)
	request.Header.Set("CheckSum", sHA1(w.appSecret+rand+tm)) // 参数加密
	request.Header.Set("Content-Type", "application/json;charset=utf-8")

	var client = &http.Client{Timeout: 5 * time.Second}
	resp, err1 := client.Do(request)
	if err1 != nil {
		err = err1
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("not found 404")
		return
	}

	bs, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		err = err1
		return
	}
	err = json.Unmarshal(bs, res)
	return
}

// 发送请求Get
func (w *wYLvChannel) sendDOHttp(method,url2 string, res interface{}) (err error) {
	request, err1 := http.NewRequest(method, w.liveBaseU+url2, nil)
	if err1 != nil {
		err = err1
		return
	}

	var nt = time.Now()
	var tm = strconv.FormatInt(nt.Unix(), 10)
	var rand = strconv.FormatInt(nt.UnixNano(), 10)

	request.Header.Set("AppKey", w.appKey)
	request.Header.Set("Nonce", rand)
	request.Header.Set("CurTime", tm)
	request.Header.Set("CheckSum", sHA1(w.appSecret+rand+tm)) // 参数加密
	request.Header.Set("Content-Type", "application/json;charset=utf-8")

	var client = &http.Client{Timeout: 5 * time.Second}
	resp, err1 := client.Do(request)
	if err1 != nil {
		err = err1
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("not found 404")
		return
	}

	bs, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		err = err1
		return
	}
	err = json.Unmarshal(bs, res)
	return
}

// sha1 加密 转为16进制小写
func sHA1(s string) string {
	hash := sha1.New()
	hash.Write(strToBytes(s))
	return hex.EncodeToString(hash.Sum(nil))
}

//StrToBytes string 转为byte 高效
func strToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
