package model

// Volunteer 是新生报名招新时的一个对象模型，其中：
// ID是新生的学号，
// Name是新生的名字，
// Phone是新生的手机号
// Major是新生的专业
// College是新生的学院
type Volunteer struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Phone   int    `json:"Phone"`
	Major   int    `json:"Major"`
	College int    `json:"College"`
}

// Create创建一个报名信息
func (self *Volunteer) Create() error

// Cancel取消一个报名的信息
func (self *Volunteer) Cancel() error

// Change修改报名信息
func (self *Volunteer) Change() error

// GetByID通过学号查询报名的信息
func (self *Volunteer) GetByID(id int) error
