// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UsersDao is the data access object for table users.
type UsersDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns UsersColumns // columns contains all the column names of Table for convenient usage.
}

// UsersColumns defines and stores column names for table users.
type UsersColumns struct {
	Uid         string // 用户ID 系统生成
	LarkId      string // 账户ID 用户设置
	Password    string // 密码
	Udid        string // 注册设备唯一标识
	Status      string // 用户状态
	Nickname    string // 昵称
	Firstname   string // firstname
	Lastname    string // lastname
	Gender      string // 性别
	BirthTs     string // 生日
	Email       string // Email
	Mobile      string // 手机号
	RegPlatform string // 注册平台
	ServerId    string // 分配的ws服务器
	CityId      string // 城市ID
	AvatarKey   string // 小图 72*72
	CreatedTs   string //
	UpdatedTs   string //
	DeletedTs   string //
}

// usersColumns holds the columns for table users.
var usersColumns = UsersColumns{
	Uid:         "uid",
	LarkId:      "lark_id",
	Password:    "password",
	Udid:        "udid",
	Status:      "status",
	Nickname:    "nickname",
	Firstname:   "firstname",
	Lastname:    "lastname",
	Gender:      "gender",
	BirthTs:     "birth_ts",
	Email:       "email",
	Mobile:      "mobile",
	RegPlatform: "reg_platform",
	ServerId:    "server_id",
	CityId:      "city_id",
	AvatarKey:   "avatar_key",
	CreatedTs:   "created_ts",
	UpdatedTs:   "updated_ts",
	DeletedTs:   "deleted_ts",
}

// NewUsersDao creates and returns a new DAO object for table data access.
func NewUsersDao() *UsersDao {
	return &UsersDao{
		group:   "default",
		table:   "users",
		columns: usersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UsersDao) Columns() UsersColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UsersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
