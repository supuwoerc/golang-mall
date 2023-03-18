package goods

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/logic/orm/dal"
	"server/service/h"
)

type GetRequest struct {
	ID string `form:"id"`
}

func Get(c *gin.Context) {
	var getRequest GetRequest
	if err := c.ShouldBindQuery(&getRequest); err != nil {
		h.Validator(c, err)
		return
	}
	if good, err := dal.Good.Where(dal.Good.ID.Eq(getRequest.ID)).First(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.FailMessage(c, "不存在该商品，请检查查询参数")
		} else {
			h.Fail(c, err)
		}
	} else {
		h.OK(c, &HomePageFeed{
			ID:            good.ID,
			Name:          good.Name,
			Price:         good.Price.String(),
			OriginalPrice: good.OriginalPrice.String(),
			Img:           good.Img,
			Desc:          good.Desc,
		})
	}
}
