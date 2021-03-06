package filters

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentFilter struct {
	SortField *string `form:"sort_field"`
	SortDest  *string `form:"sort_dir"`
	Search    *string `form:"search"`
	Name      *string `form:"name"`
	Email     *string `form:"email"`
	Phone     *string `form:"phone"`
	Time      *string `form:"time"`
	Date      *string `form:"date"`
}

func (sf *StudentFilter) GetSortField() string {
	if sf.SortField == nil {
		return "id"
	}

	return *sf.SortField
}

func (sf *StudentFilter) GetSortDest() string {
	if sf.SortDest == nil {
		return "asc"
	}

	return *sf.SortDest
}

func (sf *StudentFilter) GetSearchValue() string {
	if sf.Search == nil {
		return ""
	}

	return *sf.Search
}

func (sf *StudentFilter) GetDate() string {
	if sf.Date == nil {
		return ""
	}

	return *sf.Date
}

func (sf *StudentFilter) Scope(db *gorm.DB) *gorm.DB {
	if sf.SortField != nil || sf.SortDest != nil {
		db = db.Debug().Order(fmt.Sprintf("%s %s", sf.GetSortField(), sf.GetSortDest()))
	}

	if sf.Name != nil {
		db = db.Where("name LIKE ?", "%"+*sf.Name+"%")
	}
	if sf.Email != nil {
		db = db.Where("email LIKE ?", "%"+*sf.Email+"%")
	}
	if sf.Phone != nil {
		db = db.Where("name LIKE ?", "%"+*sf.Phone+"%")
	}
	if sf.Search != nil {
		db = db.Debug().Where(fmt.Sprintf("name LIKE %s OR email LIKE %s OR phone LIKE %s", "'%"+*sf.Search+"%'", "'%"+*sf.Search+"%'", "'%"+*sf.Search+"%'"))
	}
	if sf.Date != nil {
		db = db.Debug().Where(fmt.Sprintf("created_at LIKE %s OR updated_at LIKE %s", "'"+*sf.Date+"%'", "'"+*sf.Date+"%'"))
	}

	if sf.Time != nil {
		now := time.Now()
		if *sf.Time == "weekly" {
			db = db.Where("created_at > ?", now.AddDate(0, 0, -7))
		}
		if *sf.Time == "daily" {
			db = db.Where("created_at > ?", now.AddDate(0, 0, -1))
		}
		if *sf.Time == "monthly" {
			db = db.Where("created_at > ?", now.AddDate(0, -1, 0))
		}
		if *sf.Time == "yearly" {
			db = db.Where("created_at > ?", now.AddDate(-1, 0, 0))
		}
	}

	return db
}

func NewStudentFilterFromCtx(c *gin.Context) (*StudentFilter, error) {
	var sf StudentFilter

	if err := c.ShouldBindQuery(&sf); err != nil {
		return nil, err
	}
	return &sf, nil
}
