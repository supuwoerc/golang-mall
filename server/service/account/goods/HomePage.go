package goods

import (
	"github.com/gin-gonic/gin"
	"server/helper"
	"server/logic/orm/dal"
	"server/logic/orm/model"
	"server/service/h"
)

type HomePageFeed struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Price         string `json:"price"`
	OriginalPrice string `json:"originalPrice"`
	Img           string `json:"img"`
	Desc          string `json:"desc"`
}
type HomePageResponse struct {
	Blocks map[string][]*HomePageFeed `json:"blocks"`
}

func HomePage(c *gin.Context) {
	blocks, err := dal.Block.Find()
	if err != nil {
		h.Fail(c, err)
		return
	}
	response := HomePageResponse{Blocks: make(map[string][]*HomePageFeed)}
	var ids []string
	for _, block := range blocks {
		ids = append(ids, block.GoodsID)
	}
	ids = helper.RemoveDuplicateStrings(ids)
	goods, err := dal.Good.Where(dal.Good.ID.In(ids...)).Find()
	if err != nil {
		h.Fail(c, err)
		return
	}
	var goodMap = make(map[string]*model.Good)
	for _, good := range goods {
		goodMap[good.ID] = good
	}
	for _, block := range blocks {
		if good, ok := goodMap[block.GoodsID]; ok {
			response.Blocks[block.Key] = append(response.Blocks[block.Key], &HomePageFeed{
				ID:            good.ID,
				Name:          good.Name,
				Price:         good.Price.String(),
				OriginalPrice: good.OriginalPrice.String(),
				Img:           good.Img,
				Desc:          good.Desc,
			})
		}

	}
	h.OK(c, response)
}
