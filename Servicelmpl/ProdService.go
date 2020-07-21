package ServiceImpl

import (
	Service "LearningNotes-Go/Services"
	"context"
	"strconv"
	"time"
)

type ProdService struct{}

func (*ProdService) GetProdsList(ctx context.Context, in *Service.ProdsRequest, res *Service.ProdListResponse) error {
	time.Sleep(time.Second * 3)

	models := make([]*Service.ProdModel, 0)
	var i int32
	for i = 0; i < in.Size; i++ {
		models = append(models, newProd(100+i, "prodname"+strconv.Itoa(100+int(i))))
	}

	res.Data = models
	return nil
}

func newProd(id int32, pname string) *Service.ProdModel {
	return &Service.ProdModel{ProdID: id, ProdName: pname}
}
