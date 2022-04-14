package product

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/luoboding/mall/db"
	"github.com/luoboding/mall/models/catalogue"
	"gorm.io/datatypes"
)

type Product struct {
	ID          uint64         `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title"`                     // 产品标题
	SubTitle    string         `json:"sub_title"`                 // 产品简介
	Thumbnail   string         `json:"thumbnail"`                 // 标题缩略图
	Pictures    datatypes.JSON `json:"pictures" gorm:"type:JSON"` // 产品图片
	CatalogueId int            `json:"catalogue_id"`              // 分类id
	Description string         `json:"description"`               // 产品详细描述
	Sort        uint8          `json:"sort"`                      // 顺序
	Status      uint           `json:"status" gorm:"index"`       // 状态 0 开启 1 下架 2 售罄
	CreatedAt   time.Time      `json:"created_at"`
	// forign key
	// catalogue catalogue.Catalogue `gorm:"foreignKey:CatalogueId"`
}

func (p *Product) validate() bool {
	return p.Title != "" && p.SubTitle != "" && p.Thumbnail != "" && p.CatalogueId != 0
}

func (p *Product) Create() error {
	if !p.validate() {
		return errors.New("参数错误")
	}
	cata, e := catalogue.One(p.CatalogueId)
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
	Size    int // 每页多少条
	Current int // 当前多少页
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
	Pagination Pagination `json:"pagination"`
	Result     []Product  `json:"result"`
}

func List(query *ProductSearchQuery) (*ProductListResponse, error) {
	response := &ProductListResponse{
		Pagination: query.Pagination,
	}
	fmt.Println("pagination", query.Pagination.Current)
	paramters := []interface{}{}
	where := []string{}
	if query.Title != "" {
		paramters = append(paramters, fmt.Sprintf("%s%%", query.Title))
		where = append(where, "title like ?")
	}
	if query.Catalogue_ID != 0 {
		paramters = append(paramters, query.Catalogue_ID)
		where = append(where, "catalogue_id = ?")
	}

	if query.Price_gt != 0 {
		paramters = append(paramters, query.Price_gt)
		where = append(where, "price >= ?")
	}
	if query.Price_lt != 0 {
		paramters = append(paramters, query.Price_lt)
		where = append(where, "price < ?")
	}
	conn := db.Get_DB()
	q := conn.Table("products").Where(strings.Join(where, " and "), paramters...)
	offset := query.Size * (query.Current - 1)
	q = q.Offset(offset).Limit(query.Size)
	for _, v := range query.By {
		q = q.Order(v)
	}
	r := q.Find(&response.Result)
	fmt.Println("r", response.Pagination, response.Result)
	return response, r.Error
}
