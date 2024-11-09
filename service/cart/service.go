package cart

import (
	"fmt"
	"go-api/types"
)

func getCartItemsIDs(items []types.CartItem) ([]int, error) {
	productIds := make([]int, len(items))

	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %d", item.Quantity)
		}
		productIds[i] = item.ProductID
	}
	return productIds, nil
}

func (h *Handler) createOrder(ps []types.Product, items []types.CartItem, userID int) (int, float64, error) {

	productMap := make(map[int]types.Product)

	for _, product := range ps {
		productMap[product.ID] = product
	}

	// check if all products are actually in stock
	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, nil
	}
	//calculate total price
	totalPrice := calculateTotalPrice(items, productMap)
	//reduce quantity of products in db
	for _, itme := range items {
		product := productMap[itme.ProductID]
		product.Quantity -= itme.Quantity
		h.productStore.UpdateProduct(product)
	}
	//create the order
	orderID, err := h.store.CreateOrder(types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "123 Main St, Anytown USA 12345",
	})
	if err != nil {
		return 0, 0, nil
	}

	//create order items
	for _, item := range items {
		h.store.CreateOrderItem(types.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductID].Price,
		})
	}
	return orderID, totalPrice, nil

}

func checkIfCartIsInStock(cartItems []types.CartItem, products map[int]types.Product) error {

	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}
	for _, item := range cartItems {
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d is not available in the store,please refresh your cart", item.ProductID)
		}
		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in the quantity requested", product.Name)
		}
	}
	return nil
}

func calculateTotalPrice(cartItems []types.CartItem, products map[int]types.Product) float64 {
	var total float64
	for _, item := range cartItems {
		product := products[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}
	return total
}
