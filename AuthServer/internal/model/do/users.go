// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Users is the golang structure of table users for DAO operations like Where/Data.
type Users struct {
	g.Meta      `orm:"table:users, do:true"`
	Uid         interface{} // 用户ID 系统生成
	LarkId      interface{} // 账户ID 用户设置
	Password    interface{} // 密码
	Udid        interface{} // 注册设备唯一标识
	Status      interface{} // 用户状态
	Nickname    interface{} // 昵称
	Firstname   interface{} // firstname
	Lastname    interface{} // lastname
	Gender      interface{} // 性别
	BirthTs     interface{} // 生日
	Email       interface{} // Email
	Mobile      interface{} // 手机号
	RegPlatform interface{} // 注册平台
	ServerId    interface{} // 分配的ws服务器
	CityId      interface{} // 城市ID
	AvatarKey   interface{} // 小图 72*72
	CreatedTs   interface{} //
	UpdatedTs   interface{} //
	DeletedTs   interface{} //
}
