package process

import "fmt"

// 因为UserMgr实例在服务器端只有一个
// 很多地方会用到，因此定义成一个全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// 完成对UserMgr进行初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 完成对onlineUsers添加的功能
func (this *UserMgr) AddOnlineUser(u *UserProcess) {
	this.onlineUsers[u.UserID] = u
}

// 完成对onlineUsers删除的功能
func (this *UserMgr) DeleteOnlineUser(user int) {
	delete(this.onlineUsers, user)
}

// 完成查询功能
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

// 根据id返回对应的值
func (this *UserMgr) GetOnlineUserById(userID int) (u *UserProcess, err error) {
	//  从map中取出值，带检测
	u, ok := this.onlineUsers[userID]
	if !ok { //要查找的用户当前不在线
		err = fmt.Errorf("用户%d不在线", userID)
		return
	}
	return
}
