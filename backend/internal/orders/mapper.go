package orders

func EntityToOrder(entity OrderEntity) Order {
	return Order{
		ID:              entity.ID,
		Canal:           entity.Canal,
		Cantidad:        entity.Cantidad,
		Company:         entity.Company,
		CP:              entity.CP,
		CreatedAt:       entity.CreatedAt,
		DaysToDelivery:  entity.DaysToDelivery,
		ErrorCode:       entity.ErrorCode,
		ErrorMessage:    entity.ErrorMessage,
		FechaCompra:     entity.FechaCompra,
		FechaEstimada:   entity.FechaEstimada,
		FulfillmentType: entity.FulfillmentType,
		IsFlash:         entity.IsFlash,
		IsMarketplace:   entity.IsMarketplace,
		NoPedido:        entity.NoPedido,
		Plan:            entity.Plan,
		ProductType:     entity.ProductType,
		SKU:             entity.SKU,
		StoreSelected:   entity.StoreSelected,
		TipoPago:        entity.TipoPago,
		EDD1:            entity.EDD1,
		EDD2:            entity.EDD2,
	}
}

func EntitiesToOrders(entities []OrderEntity) []Order {
	orders := make([]Order, 0, len(entities))

	for _, entity := range entities {
		orders = append(orders, EntityToOrder(entity))
	}

	return orders
}

func OrderToDTO(order Order) OrderDTO {
	return OrderDTO{
		ID:              order.ID,
		NoPedido:        order.NoPedido,
		Canal:           order.Canal,
		SKU:             order.SKU,
		FechaEstimada:   order.FechaEstimada,
		FulfillmentType: order.FulfillmentType,
		ProductType:     order.ProductType,
		Cantidad:        order.Cantidad,
		FechaCompra:     order.FechaCompra,
		Company:         order.Company,
		HasError:        order.ErrorCode != "" || order.ErrorMessage != "",
	}
}

func OrdersToDTOs(orders []Order) []OrderDTO {
	dtos := make([]OrderDTO, 0, len(orders))

	for _, order := range orders {
		dtos = append(dtos, OrderToDTO(order))
	}

	return dtos
}
