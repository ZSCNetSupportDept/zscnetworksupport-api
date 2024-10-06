package model

// Member是网维成员的对象模型，其中：
// ID是成员的工号，
// Name是成员的姓名，
// Access是成员的权限：0是开发组，1是科长,2是API,3是值班组长，4是正式成员，5是实习成员，6是退休成员，
// Phone是成员的手机号
type Member struct {
	ID     int
	Name   string
	Access int
	Phone  int
}

// GetByID通过一个成员的ID来返回这个成员的所有信息，还会返回一个错误类型，你需要判断操作有没有出错
func (self *Member) GetByID(id int) error {
	return nil
	// 这里在施工中。。。
}

// Store将一个成员的信息存储到数据库中，你需要提供一个指向Member数据的指针，该成员的ID不能和已有的ID相同，否则会返回错误，修改信息不要用这个，这个一般用作录入信息
func (self *Member) Store(source *Member) error {
	return nil
}

// AccessChange修改一个成员的权限，你需要提供新的权限代码，出错了会返回一些错误
func (self *Member) AccessChange(accessNew int) error {
	return nil
}
