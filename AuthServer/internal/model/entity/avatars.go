// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Avatars is the golang structure for table avatars.
type Avatars struct {
	OwnerId      uint64 `json:"ownerId"      ` // 用户ID/ChatID
	OwnerType    uint   `json:"ownerType"    ` // 1:用户头像 2:群头像
	AvatarSmall  string `json:"avatarSmall"  ` // 小图 72*72
	AvatarMedium string `json:"avatarMedium" ` // 中图 240*240
	AvatarLarge  string `json:"avatarLarge"  ` // 大图 480*480
	CreatedTs    int64  `json:"createdTs"    ` //
	UpdatedTs    int64  `json:"updatedTs"    ` //
	DeletedTs    int64  `json:"deletedTs"    ` //
}
