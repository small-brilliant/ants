// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Users is the golang structure for table users.
type Users struct {
	Uid         uint64 `json:"uid"         ` // 用户ID 系统生成
	LarkId      string `json:"larkId"      ` // 账户ID 用户设置
	Password    string `json:"password"    ` // 密码
	Udid        string `json:"udid"        ` // 注册设备唯一标识
	Status      uint   `json:"status"      ` // 用户状态
	Nickname    string `json:"nickname"    ` // 昵称
	Firstname   string `json:"firstname"   ` // firstname
	Lastname    string `json:"lastname"    ` // lastname
	Gender      uint   `json:"gender"      ` // 性别
	BirthTs     int64  `json:"birthTs"     ` // 生日
	Email       string `json:"email"       ` // Email
	Mobile      string `json:"mobile"      ` // 手机号
	RegPlatform uint   `json:"regPlatform" ` // 注册平台
	ServerId    int    `json:"serverId"    ` // 分配的ws服务器
	CityId      int    `json:"cityId"      ` // 城市ID
	AvatarKey   string `json:"avatarKey"   ` // 小图 72*72
	CreatedTs   int64  `json:"createdTs"   ` //
	UpdatedTs   int64  `json:"updatedTs"   ` //
	DeletedTs   int64  `json:"deletedTs"   ` //
}
