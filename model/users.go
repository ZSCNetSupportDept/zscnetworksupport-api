package model

// User是报修系统用户的对象模型，其中:
// id为用户学号，
// name为用户名字，
// ISP为用户运营商：0代表其它，1代表电信，2联通，3移动，
// Block为用户的楼栋编号，23-26分别代表了香晖的ABCD,27是朝晖，
// Room是用户的房间编号，
// Phone是用户的手机号
type User struct {
	ID    int
	Name  string
	ISP   int
	Block int
	Room  int
	Phone int
}

// Store将一个用户的信息储存到数据库里
func (self *User) Store() error

// InfoUpdate将更新用户的信息，需要传入最新的信息
func (self *User) InfoUpdate(source *User) error

// GetByID通过用户的ID返回用户的信息
func (self *User) GetByID(id int) error

// GeyByName通过用户的名字返回用户的信息
func (self *User) GetByName(name string) error

// ChangeISP将修改用户的运营商，需要传入最新的运营商代码
func (self *User) ChangeISP(isp int) error
