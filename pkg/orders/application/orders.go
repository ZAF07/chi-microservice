package application

type productsService interface {}

type paymentsinterface interface {}

type OrdersService struct {

}

func NewOrderService() {

}

func (s OrdersService) PlaceOrder(cmd placeOrderCommand) error {

}

type MarkOrderAsPaidCommand struct {

}

func (s OrdersService) MarkOrderAsPaid(cmd MarkOrderAsPaid) error {

}

func (s OrdersService) OrderByID(id orders.ID) (orders.Orders, error) {

}