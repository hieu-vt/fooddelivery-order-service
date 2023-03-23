package common

type Restaurant struct {
	Name      string     `json:"name" gorm:"-"`
	Addr      string     `json:"addr" gorm:"-"`
	Logo      *Image     `json:"logo" gorm:"-"`
	Cover     []string   `json:"cover" gorm:"-"`
	LikeCount int        `json:"like_count" gorm:"-"`
	Owner     SimpleUser `json:"owner" gorm:"-"`
}
