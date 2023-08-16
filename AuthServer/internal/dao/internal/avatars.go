// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AvatarsDao is the data access object for table avatars.
type AvatarsDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns AvatarsColumns // columns contains all the column names of Table for convenient usage.
}

// AvatarsColumns defines and stores column names for table avatars.
type AvatarsColumns struct {
	OwnerId      string // 用户ID/ChatID
	OwnerType    string // 1:用户头像 2:群头像
	AvatarSmall  string // 小图 72*72
	AvatarMedium string // 中图 240*240
	AvatarLarge  string // 大图 480*480
	CreatedTs    string //
	UpdatedTs    string //
	DeletedTs    string //
}

// avatarsColumns holds the columns for table avatars.
var avatarsColumns = AvatarsColumns{
	OwnerId:      "owner_id",
	OwnerType:    "owner_type",
	AvatarSmall:  "avatar_small",
	AvatarMedium: "avatar_medium",
	AvatarLarge:  "avatar_large",
	CreatedTs:    "created_ts",
	UpdatedTs:    "updated_ts",
	DeletedTs:    "deleted_ts",
}

// NewAvatarsDao creates and returns a new DAO object for table data access.
func NewAvatarsDao() *AvatarsDao {
	return &AvatarsDao{
		group:   "default",
		table:   "avatars",
		columns: avatarsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AvatarsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AvatarsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AvatarsDao) Columns() AvatarsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AvatarsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AvatarsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AvatarsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
