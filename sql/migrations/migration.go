package migrations

import (
	"github.com/luoboding/mall/db"
	"github.com/luoboding/mall/models/catalogue"
	"github.com/luoboding/mall/models/product"
	"github.com/luoboding/mall/models/user"
)

func Migrate() {
	db := db.Get_DB()
	migrations := []interface{}{}
	migrations = append(migrations, user.User{})
	migrations = append(migrations, product.Product{})
	migrations = append(migrations, catalogue.Catalogue{})
	db.AutoMigrate(migrations...)
}
