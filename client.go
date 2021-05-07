package picqer

// Client is a single Picqer client
type Client struct {
	BaseURL string
	Token   string

	Orders         OrderService
	Picklists      PicklistService
	Products       ProductService
	PurchaseOrders PurchaseOrderService
	Backorders     BackorderService
	Stats          StatsService
}

// NewClient creates a new conntection to Piqcer
// using it's endpoint and API key
func NewClient(baseURL, token string) *Client {
	c := &Client{
		BaseURL: baseURL,
		Token:   token,
	}

	c.Products = &ProductServiceOp{client: c}
	c.Orders = &OrderServiceOp{client: c}
	c.Picklists = &PicklistServiceOp{client: c}
	c.PurchaseOrders = &PurchaseOrderServiceOp{client: c}
	c.Backorders = &BackorderServiceOp{client: c}
	c.Stats = &StatsServiceOp{client: c}

	return c
}
