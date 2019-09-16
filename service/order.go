package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/triardn/inventory/common"
	"github.com/triardn/inventory/model"
	"github.com/triardn/inventory/repository"
)

type OrderService struct {
	repository            repository.IOrderRepository
	orderDetailRepository repository.IOrderDetailRepository
	restockRepository     repository.IRestockRepository
	ServiceOption
}

func NewOrderService(orderRepository repository.IOrderRepository, orderDetailRepository repository.IOrderDetailRepository, restockRepository repository.IRestockRepository, logger *common.APILogger) *OrderService {
	orderService := &OrderService{}
	orderService.Logger = logger
	orderService.repository = orderRepository
	orderService.orderDetailRepository = orderDetailRepository
	orderService.restockRepository = restockRepository
	return orderService
}

func (os *OrderService) GetAllOrder() (orders []model.Order, err error) {
	return os.repository.GetAllOrder()
}

func (os *OrderService) GetOrderByID(id uint64) (order model.Order, err error) {
	return os.repository.GetOrderByID(id)
}

func (os *OrderService) GetOrderIDByInvoice(invoice string) (orderID uint64, err error) {
	return os.repository.GetOrderIDByInvoice(invoice)
}

func (os *OrderService) PopulateExportData(start string, end string) (data [][]string, grandTotalProfit int64, err error) {
	layout := "2006-01-02 15:04:05"
	data = [][]string{}

	if end == "" {
		start = fmt.Sprintf("%s 00:00:01", start)
		end = fmt.Sprintf("%s 23:59:59", start)
	} else {
		start = fmt.Sprintf("%s 00:00:01", start)
		end = fmt.Sprintf("%s 23:59:59", end)
	}

	timeStart, _ := time.Parse(layout, start)
	timeEnd, _ := time.Parse(layout, end)

	orders, err := os.repository.GetAllOrderWithTimeFrame(timeStart.Unix(), timeEnd.Unix())
	if err != nil {
		return
	}

	for _, order := range orders {
		details, err := os.orderDetailRepository.GetDetailByOrderID(order.ID)
		if err != nil {
			continue
		}

		for i, detail := range details {
			var tempData []string
			averagePrice := os.restockRepository.GetAveragePriceByProductID(detail.ProductID)
			profit := detail.Total - (int64(averagePrice) * detail.Quantity)
			if i == 0 {
				tempData = []string{order.Invoice, strconv.FormatInt(order.Created, 10), detail.Product.Sku, detail.Product.Name, strconv.FormatInt(detail.Quantity, 10), common.FormatCurrency("id_ID", detail.Price, true), common.FormatCurrency("id_ID", detail.Total, true), common.FormatCurrency("id_ID", int64(averagePrice), true), common.FormatCurrency("id_ID", profit, true)}
			} else {
				tempData = []string{"", "", detail.Product.Sku, detail.Product.Name, strconv.FormatInt(detail.Quantity, 10), common.FormatCurrency("id_ID", detail.Price, true), common.FormatCurrency("id_ID", detail.Total, true), common.FormatCurrency("id_ID", int64(averagePrice), true), common.FormatCurrency("id_ID", profit, true)}
			}

			data = append(data, tempData)

			grandTotalProfit += profit
		}
	}

	return
}

func (os *OrderService) ReportStatistics(start string, end string) (statistics map[string]int64) {
	layout := "2006-01-02 15:04:05"

	if end == "" {
		start = fmt.Sprintf("%s 00:00:01", start)
		end = fmt.Sprintf("%s 23:59:59", start)
	} else {
		start = fmt.Sprintf("%s 00:00:01", start)
		end = fmt.Sprintf("%s 23:59:59", end)
	}

	timeStart, _ := time.Parse(layout, start)
	timeEnd, _ := time.Parse(layout, end)

	turnover := os.repository.GetTotalTurnover(timeStart.Unix(), timeEnd.Unix())
	orderCount := os.repository.GetCountOrder(timeStart.Unix(), timeEnd.Unix())
	productCount := os.orderDetailRepository.GetTotalOrderedProduct(timeStart.Unix(), timeEnd.Unix())

	statistics = make(map[string]int64)
	statistics["turnover"] = turnover
	statistics["orderCount"] = orderCount
	statistics["productCount"] = productCount

	return
}
