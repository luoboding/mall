package product

import (
	"errors"
	"fmt"
	"time"

	"github.com/luoboding/mall/db"
	"github.com/luoboding/mall/models/catalogue"
	"gorm.io/datatypes"
)

type Product struct {
	ID          uint64              `json:"id" gorm:"primaryKey"`
	Title       string              `json:"title" gorm:"index;comment:标题"`             // 产品标题
	SubTitle    string              `json:"sub_title" gorm:"comment:副标题"`              // 产品简介
	Thumbnail   string              `json:"thumbnail" gorm:"comment:缩略图"`              // 标题缩略图
	Pictures    datatypes.JSON      `json:"pictures" gorm:"type:JSON;comment:产品图片"`    // 产品图片
	CatalogueID int                 `json:"catalogue_id" gorm:"comment:分类ID"`          // 分类id
	Description string              `json:"description" gorm:"comment:产品描述"`           // 产品详细描述
	Sort        uint8               `json:"sort" gorm:"comment:产品排序"`                  // 顺序
	Status      uint                `json:"status" gorm:"index;comment:产品状态0正常1下架2售罄"` // 状态 0 开启 1 下架 2 售罄
	CreatedAt   time.Time           `json:"created_at" gorm:"comment:创建时间"`
	Catalogue   catalogue.Catalogue `json:"catalogue"`
}

func (p *Product) validate() bool {
	return p.Title != "" && p.SubTitle != "" && p.Thumbnail != "" && p.CatalogueID != 0
}

func (p *Product) Create() error {
	if !p.validate() {
		return errors.New("参数错误")
	}
	cata, e := catalogue.One(p.CatalogueID)
	if e != nil || cata.ID == 0 {
		return errors.New("分类不存在")
	}
	connection := db.Get_DB()
	re := connection.Create(p)
	return re.Error
}

func (p *Product) Update() error {
	if !p.validate() {
		return errors.New("参数错误")
	}
	var first Product
	connection := db.Get_DB()
	r := connection.Table("products").Where("id = ?", p.ID).First(&first)
	if r.Error != nil {
		return r.Error
	}
	action := connection.Model(first).Updates(p)
	return action.Error
}

// 包方法

type Pagination struct {
	Size    int `json:"size"`    // 每页多少条
	Current int `json:"current"` // 当前多少页
}

type Order struct {
	By []string // 排序
}

type ProductSearchQuery struct {
	Pagination
	Order
	Title        string
	Price_gt     int
	Price_lt     int
	Catalogue_ID uint64
}

type ProductListResponse struct {
	Pagination
	Result []Product `json:"result"`
}

func List(query *ProductSearchQuery) (*ProductListResponse, error) {
	response := &ProductListResponse{
		Pagination: query.Pagination,
	}
	conn := db.Get_DB()
	request := conn.Table("products")
	if query.Title != "" {
		request = request.Where("title like ?", fmt.Sprintf("%s%%", query.Title))
	}
	if query.Catalogue_ID != 0 {
		request = request.Where("catalogue_id = ?", query.Catalogue_ID)
	}

	if query.Price_gt != 0 {
		request = request.Where("price >= ?", query.Price_gt)
	}

	if query.Price_lt != 0 {
		request = request.Where("price < ?", query.Price_lt)
	}

	offset := query.Size * (query.Current - 1)
	request = request.Offset(offset).Limit(query.Size)
	for _, v := range query.By {
		request = request.Order(v)
	}
	r := request.Preload("Catalogue").Find(&response.Result)
	return response, r.Error
}
