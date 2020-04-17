package v1

// 订单
type OrderMainDAO interface {
	SaveOrderMain()
}

// 订单详情
type OrderDetailDAO interface {
	SavaOrderDetail()
}

type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}
