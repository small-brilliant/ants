// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Avatars is the golang structure of table avatars for DAO operations like Where/Data.
type Avatars struct {
	g.Meta       `orm:"table:avatars, do:true"`
	OwnerId      interface{} // 用户ID/ChatID
	OwnerType    interface{} // 1:用户头像 2:群头像
	AvatarSmall  interface{} // 小图 72*72
	AvatarMedium interface{} // 中图 240*240
	AvatarLarge  interface{} // 大图 480*480
	CreatedTs    interface{} //
	UpdatedTs    interface{} //
	DeletedTs    interface{} //
}
