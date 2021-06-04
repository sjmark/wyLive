package wyLive

import (
	"fmt"
)

/* 创建房间
* param channelName	string	必选 abc 房间名称。
* param mode	    int	    必选 2 固定为 2。
* param uid	        int64	必选 163 房间创建者的用户 ID，是您的业务系统中的实际用户 ID。
 */
func (w *wYLvChannel) CreateRoom(name string, uid int64) (res InterLiveComm, err error) {
	err = w.sendHttp("/v2/api/room", map[string]interface{}{"channelName": name, "mode": 2, "uid": uid}, &res)
	return
}

/* 查看房间信息
* param cid	  int64  必选 6207760637435905	房间 ID。该 ID 为创建房间接口调用成功后返回的房间 ID。
* param cname string 必选 abc 房间名称。
 */
func (w *wYLvChannel) LoadRoom(cid int64, name string) (res LoadInterLiveRes, err error) {
	if cid > 0 {
		err = w.sendDOHttp("GET", fmt.Sprintf("/v2/api/rooms/%d", cid), &res)
		return
	}

	if name != "" {
		err = w.sendDOHttp("GET", fmt.Sprintf("/v3/api/rooms?cname=%s", name), &res)
	}
	return
}

/* 查看房间内成员信息
* param cid	  int64  必选 6207760637435905	房间 ID。该 ID 为创建房间接口调用成功后返回的房间 ID。
* param cname string 必选 abc 房间名称。
 */
func (w *wYLvChannel) LoadRoomMembers(cid int64, name string) (res LoadInterMembersRes, err error) {
	if cid > 0 {
		err = w.sendDOHttp("GET", fmt.Sprintf("/v2/api/rooms/%d/members", cid), &res)
		return
	}

	if name != "" {
		err = w.sendDOHttp("GET", fmt.Sprintf("/v3/api/rooms/members?cname=%s", name), &res)
	}
	return
}

/* 移除成员
* param cid	  int64  必选 6207760637435905	房间 ID。该 ID 为创建房间接口调用成功后返回的房间 ID。
* param uid	  int64	是 193992653091841 需要被移除出房间的用户 ID。
* param cname string 必选 abc 房间名称。
 */
func (w *wYLvChannel) RemoveRoomMembers(cid, uid int64, name string) (res InterLiveComm, err error) {
	if cid > 0 && uid > 0 {
		err = w.sendDOHttp("POST", fmt.Sprintf("/v2/api/kicklist/%d/members/%d", cid, uid), &res)
		return
	}

	if name != "" && uid > 0 {
		err = w.sendDOHttp("POST", fmt.Sprintf("/v3/api/kicklist/members?cname=%s&uid=%d", name, uid), &res)
	}
	return
}

/* 查看被移除成员列表
* param cid	  int64  必选 6207760637435905	房间 ID。该 ID 为创建房间接口调用成功后返回的房间 ID。
* param cname string 必选 abc 房间名称。
 */
func (w *wYLvChannel) LoadRemoveMembers(cid int64, name string) (res InterRemoveMembersRes, err error) {
	if cid > 0 {
		err = w.sendDOHttp("GET", fmt.Sprintf("/v2/api/kicklist/%d", cid), &res)
		return
	}

	if name != "" {
		err = w.sendDOHttp("GET", fmt.Sprintf("/v3/api/kicklist?cname=%s", name), &res)
	}
	return
}

/* 撤销移除
* param cid	  int64  必选 6207760637435905	房间 ID。该 ID 为创建房间接口调用成功后返回的房间 ID。
* param uid	  int64	是 193992653091841 需要被移除出房间的用户 ID。
* param cname string 必选 abc 房间名称。
 */
func (w *wYLvChannel) ReturnRemoveMembers(cid, uid int64, name string) (res InterLiveComm, err error) {
	if cid > 0 && uid > 0 {
		err = w.sendDOHttp("DELETE", fmt.Sprintf("/v2/api/kicklist/%d/members/%d", cid, uid), &res)
		return
	}

	if name != "" && uid > 0 {
		err = w.sendDOHttp("DELETE", fmt.Sprintf("/v3/api/kicklist/members?cname=%s&uid=%d", name, uid), &res)
	}
	return
}

/* 删除房间
* param cid	  int64  必选 6207760637435905	房间 ID。该 ID 为创建房间接口调用成功后返回的房间 ID。
* param cname string 必选 abc 房间名称。
 */
func (w *wYLvChannel) RemoveRoom(cid int64, name string) (res InterLiveComm, err error) {
	if cid > 0 {
		err = w.sendDOHttp("DELETE", fmt.Sprintf("/v2/api/rooms/%d", cid), &res)
		return
	}

	if name != "" {
		err = w.sendDOHttp("DELETE", fmt.Sprintf("v3/api/rooms?cname=%s", name), &res)
	}
	return
}
