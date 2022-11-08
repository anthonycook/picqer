package picqer

type OrdersWebhook struct {
	Idhook           int     `json:"idhook"`
	Name             string  `json:"name"`
	Event            string  `json:"event"`
	EventTriggeredAt string  `json:"event_triggered_at"`
	Data             []Order `json:"data"`
}

type OrdersNotesCreatedWebhook struct {
	Idhook           int    `json:"idhook"`
	Name             string `json:"name"`
	Event            string `json:"event"`
	EventTriggeredAt string `json:"event_triggered_at"`
	Data             struct {
		Order struct {
			Idorder             int     `json:"idorder"`
			Idcustomer          int     `json:"idcustomer"`
			Orderid             string  `json:"orderid"`
			Deliveryname        string  `json:"deliveryname"`
			Deliverycontactname string  `json:"deliverycontactname"`
			Deliveryaddress     string  `json:"deliveryaddress"`
			Deliveryaddress2    string  `json:"deliveryaddress2"`
			Deliveryzipcode     string  `json:"deliveryzipcode"`
			Deliverycity        string  `json:"deliverycity"`
			Deliveryregion      string  `json:"deliveryregion"`
			Deliverycountry     string  `json:"deliverycountry"`
			Invoicename         string  `json:"invoicename"`
			Invoicecontactname  string  `json:"invoicecontactname"`
			Invoiceaddress      string  `json:"invoiceaddress"`
			Invoiceaddress2     string  `json:"invoiceaddress2"`
			Invoicezipcode      string  `json:"invoicezipcode"`
			Invoicecity         string  `json:"invoicecity"`
			Invoiceregion       string  `json:"invoiceregion"`
			Invoicecountry      string  `json:"invoicecountry"`
			Reference           string  `json:"reference"`
			Partialdelivery     bool    `json:"partialdelivery"`
			Discount            float64 `json:"discount"`
			Status              string  `json:"status"`
			PublicStatusPage    string  `json:"public_status_page"`
			Created             string  `json:"created"`
			Updated             string  `json:"updated"`
			Products            []struct {
				Idproduct   int     `json:"idproduct"`
				Idvatgroup  int     `json:"idvatgroup"`
				Productcode string  `json:"productcode"`
				Name        string  `json:"name"`
				Remarks     string  `json:"remarks"`
				Price       float64 `json:"price"`
				Amount      int     `json:"amount"`
				Weight      int     `json:"weight"`
			} `json:"products"`
		} `json:"order"`
		Note struct {
			IdorderLog  int    `json:"idorder_log"`
			Iduser      int    `json:"iduser"`
			Idorder     int    `json:"idorder"`
			Idpicklist  string `json:"idpicklist"`
			Type        string `json:"type"`
			Action      string `json:"action"`
			Description string `json:"description"`
			Details     string `json:"details"`
			CreatedAt   string `json:"created_at"`
		} `json:"note"`
	} `json:"data"`
}

type PicklistsWebhook struct {
	Idhook           int    `json:"idhook"`
	Name             string `json:"name"`
	Event            string `json:"event"`
	EventTriggeredAt string `json:"event_triggered_at"`
	Data             struct {
		Idpicklist       int         `json:"idpicklist"`
		Picklistid       string      `json:"picklistid"`
		Idcustomer       int         `json:"idcustomer"`
		Idorder          int         `json:"idorder"`
		Idwarehouse      int         `json:"idwarehouse"`
		Deliveryname     string      `json:"deliveryname"`
		Deliverycontact  string      `json:"deliverycontact"`
		Deliveryaddress  string      `json:"deliveryaddress"`
		Deliveryaddress2 string      `json:"deliveryaddress2"`
		Deliveryzipcode  string      `json:"deliveryzipcode"`
		Deliverycity     string      `json:"deliverycity"`
		Deliveryregion   interface{} `json:"deliveryregion"`
		Deliverycountry  string      `json:"deliverycountry"`
		Emailaddress     string      `json:"emailaddress"`
		Telephone        interface{} `json:"telephone"`
		Reference        string      `json:"reference"`
		AssignedToIduser interface{} `json:"assigned_to_iduser"`
		Invoiced         bool        `json:"invoiced"`
		Status           string      `json:"status"`
		Totalproducts    int         `json:"totalproducts"`
		Totalpicked      int         `json:"totalpicked"`
		Created          string      `json:"created"`
		Updated          string      `json:"updated"`
		Products         []struct {
			Idproduct     int     `json:"idproduct"`
			Idvatgroup    int     `json:"idvatgroup"`
			Productcode   string  `json:"productcode"`
			Name          string  `json:"name"`
			Remarks       string  `json:"remarks"`
			Amount        int     `json:"amount"`
			AmountPicked  int     `json:"amount_picked"`
			Price         float64 `json:"price"`
			Weight        int     `json:"weight"`
			Stocklocation string  `json:"stocklocation"`
		} `json:"products"`
	} `json:"data"`
}

type PicklistsShipmentsCreatedWebhook struct {
	Idhook           int    `json:"idhook"`
	Name             string `json:"name"`
	Event            string `json:"event"`
	EventTriggeredAt string `json:"event_triggered_at"`
	Data             struct {
		Idshipment                int    `json:"idshipment"`
		Idpicklist                int    `json:"idpicklist"`
		Idorder                   int    `json:"idorder"`
		Idshippingprovider        int    `json:"idshippingprovider"`
		IdcompanyShippingprovider int    `json:"idcompany_shippingprovider"`
		Labelurl                  string `json:"labelurl"`
		Provider                  string `json:"provider"`
		Providername              string `json:"providername"`
		PublicProvidername        string `json:"public_providername"`
		Trackingcode              string `json:"trackingcode"`
		Trackingurl               string `json:"trackingurl"`
		Tracktraceurl             string `json:"tracktraceurl"`
		Created                   string `json:"created"`
		Updated                   string `json:"updated"`
	} `json:"data"`
}

type ProductsWebhook struct {
	Idhook           int    `json:"idhook"`
	Name             string `json:"name"`
	Event            string `json:"event"`
	EventTriggeredAt string `json:"event_triggered_at"`
	Data             struct {
		Idproduct               int         `json:"idproduct"`
		Idvatgroup              int         `json:"idvatgroup"`
		Idsupplier              interface{} `json:"idsupplier"`
		Productcode             string      `json:"productcode"`
		Name                    string      `json:"name"`
		Price                   float64     `json:"price"`
		Fixedstockprice         int         `json:"fixedstockprice"`
		ProductcodeSupplier     string      `json:"productcode_supplier"`
		Deliverytime            interface{} `json:"deliverytime"`
		Description             string      `json:"description"`
		Barcode                 string      `json:"barcode"`
		Type                    string      `json:"type"`
		Unlimitedstock          bool        `json:"unlimitedstock"`
		Weight                  int         `json:"weight"`
		MinimumPurchaseQuantity int         `json:"minimum_purchase_quantity"`
		PurchaseInQuantitiesOf  int         `json:"purchase_in_quantities_of"`
		HsCode                  interface{} `json:"hs_code"`
		CountryOfOrigin         interface{} `json:"country_of_origin"`
		Active                  bool        `json:"active"`
		Productfields           []struct {
			Idproductfield int    `json:"idproductfield"`
			Title          string `json:"title"`
			Value          string `json:"value"`
		} `json:"productfields"`
		Images []string `json:"images"`
		Stock  []struct {
			Idwarehouse         int `json:"idwarehouse"`
			Stock               int `json:"stock"`
			Reserved            int `json:"reserved"`
			Reservedbackorders  int `json:"reservedbackorders"`
			Reservedpicklists   int `json:"reservedpicklists"`
			Reservedallocations int `json:"reservedallocations"`
			Freestock           int `json:"freestock"`
			Freepickablestock   int `json:"freepickablestock"`
		} `json:"stock"`
		Tags map[string]Tag `json:"tags"`
	} `json:"data"`
}

type ProductsStockChangedWebhook struct {
	Idhook           int    `json:"idhook"`
	Name             string `json:"name"`
	Event            string `json:"event"`
	EventTriggeredAt string `json:"event_triggered_at"`
	Data             struct {
		Product struct {
			Idproduct               int         `json:"idproduct"`
			Idvatgroup              int         `json:"idvatgroup"`
			Idsupplier              interface{} `json:"idsupplier"`
			Productcode             string      `json:"productcode"`
			Name                    string      `json:"name"`
			Price                   float64     `json:"price"`
			Fixedstockprice         int         `json:"fixedstockprice"`
			ProductcodeSupplier     string      `json:"productcode_supplier"`
			Deliverytime            interface{} `json:"deliverytime"`
			Description             string      `json:"description"`
			Barcode                 string      `json:"barcode"`
			Type                    string      `json:"type"`
			Unlimitedstock          bool        `json:"unlimitedstock"`
			Weight                  int         `json:"weight"`
			MinimumPurchaseQuantity int         `json:"minimum_purchase_quantity"`
			PurchaseInQuantitiesOf  int         `json:"purchase_in_quantities_of"`
			HsCode                  interface{} `json:"hs_code"`
			CountryOfOrigin         interface{} `json:"country_of_origin"`
			Active                  bool        `json:"active"`
			Productfields           []struct {
				Idproductfield int    `json:"idproductfield"`
				Title          string `json:"title"`
				Value          string `json:"value"`
			} `json:"productfields"`
			Images []string `json:"images"`
			Stock  []struct {
				Idwarehouse         int `json:"idwarehouse"`
				Stock               int `json:"stock"`
				Reserved            int `json:"reserved"`
				Reservedbackorders  int `json:"reservedbackorders"`
				Reservedpicklists   int `json:"reservedpicklists"`
				Reservedallocations int `json:"reservedallocations"`
				Freestock           int `json:"freestock"`
				Freepickablestock   int `json:"freepickablestock"`
			} `json:"stock"`
		} `json:"product"`
		ProductStockHistory struct {
			IdproductStockHistory int    `json:"idproduct_stock_history"`
			Idproduct             int    `json:"idproduct"`
			Idwarehouse           int    `json:"idwarehouse"`
			Iduser                int    `json:"iduser"`
			OldStock              int    `json:"old_stock"`
			StockChange           int    `json:"stock_change"`
			NewStock              int    `json:"new_stock"`
			Reason                string `json:"reason"`
			ChangeType            string `json:"change_type"`
			ChangedAt             string `json:"changed_at"`
		} `json:"product_stock_history"`
	} `json:"data"`
}

type PurchaseOrdersWebhook struct {
	Idhook           int    `json:"idhook"`
	Name             string `json:"name"`
	Event            string `json:"event"`
	EventTriggeredAt string `json:"event_triggered_at"`
	Data             struct {
		Idpurchaseorder   int         `json:"idpurchaseorder"`
		Idsupplier        int         `json:"idsupplier"`
		Idwarehouse       int         `json:"idwarehouse"`
		Idtemplate        int         `json:"idtemplate"`
		Purchaseorderid   string      `json:"purchaseorderid"`
		SupplierOrderid   interface{} `json:"supplier_orderid"`
		SupplierName      interface{} `json:"supplier_name"`
		Status            string      `json:"status"`
		Remarks           interface{} `json:"remarks"`
		DeliveryDate      string      `json:"delivery_date"`
		Language          string      `json:"language"`
		PurchasedByIduser interface{} `json:"purchased_by_iduser"`
		PurchasedAt       interface{} `json:"purchased_at"`
		CompletedByIduser interface{} `json:"completed_by_iduser"`
		CompletedAt       interface{} `json:"completed_at"`
		CreatedByIduser   int         `json:"created_by_iduser"`
		Created           string      `json:"created"`
		Updated           string      `json:"updated"`
		Products          []struct {
			IdpurchaseorderProduct int         `json:"idpurchaseorder_product"`
			Idproduct              int         `json:"idproduct"`
			Idvatgroup             int         `json:"idvatgroup"`
			Productcode            string      `json:"productcode"`
			ProductcodeSupplier    string      `json:"productcode_supplier"`
			Name                   string      `json:"name"`
			Price                  float64     `json:"price"`
			Amount                 int         `json:"amount"`
			Amountreceived         int         `json:"amountreceived"`
			DeliveryDate           interface{} `json:"delivery_date"`
			Weight                 int         `json:"weight"`
		} `json:"products"`
		IdfulfilmentCustomer interface{} `json:"idfulfilment_customer"`
	} `json:"data"`
}

type ReceiptsWebhook struct {
	Idreceipt   int `json:"idreceipt"`
	Idwarehouse int `json:"idwarehouse"`
	Supplier    struct {
		Idsupplier int    `json:"idsupplier"`
		Name       string `json:"name"`
	} `json:"supplier"`
	Purchaseorder struct {
		Idpurchaseorder int    `json:"idpurchaseorder"`
		Purchaseorderid string `json:"purchaseorderid"`
	} `json:"purchaseorder"`
	Receiptid               string      `json:"receiptid"`
	Status                  string      `json:"status"`
	Remarks                 interface{} `json:"remarks"`
	CompletedBy             interface{} `json:"completed_by"`
	AmountReceived          int         `json:"amount_received"`
	AmountReceivedExcessive int         `json:"amount_received_excessive"`
	CompletedAt             interface{} `json:"completed_at"`
	Created                 string      `json:"created"`
	Updated                 string      `json:"updated"`
	Products                []struct {
		IdreceiptProduct         int         `json:"idreceipt_product"`
		IdpurchaseorderProduct   int         `json:"idpurchaseorder_product"`
		Idproduct                int         `json:"idproduct"`
		Idpurchaseorder          int         `json:"idpurchaseorder"`
		Productcode              string      `json:"productcode"`
		ProductcodeSupplier      string      `json:"productcode_supplier"`
		Name                     string      `json:"name"`
		Barcode                  string      `json:"barcode"`
		Amount                   int         `json:"amount"`
		AmountOrdered            int         `json:"amount_ordered"`
		AmountReceiving          int         `json:"amount_receiving"`
		AddedByReceipt           bool        `json:"added_by_receipt"`
		AbcClassification        interface{} `json:"abc_classification"`
		AmountPreviouslyReceived int         `json:"amount_previously_received"`
		AmountMoreThanOrdered    int         `json:"amount_more_than_ordered"`
		AmountBackorders         int         `json:"amount_backorders"`
		Image                    interface{} `json:"image"`
		Stock                    int         `json:"stock"`
		StockLocation            struct {
			Idlocation       int         `json:"idlocation"`
			Idwarehouse      int         `json:"idwarehouse"`
			ParentIdlocation int         `json:"parent_idlocation"`
			Name             string      `json:"name"`
			Remarks          interface{} `json:"remarks"`
			UnlinkOnEmpty    bool        `json:"unlink_on_empty"`
			LocationType     struct {
				IdlocationType int    `json:"idlocation_type"`
				Name           string `json:"name"`
				Default        bool   `json:"default"`
				Color          string `json:"color"`
			} `json:"location_type"`
			IsBulkLocation bool `json:"is_bulk_location"`
		} `json:"stock_location"`
	} `json:"products"`
}

type ReturnsWebhook struct {
	Idhook           int    `json:"idhook"`
	Name             string `json:"name"`
	Event            string `json:"event"`
	EventTriggeredAt string `json:"event_triggered_at"`
	Data             struct {
		Idreturn             int         `json:"idreturn"`
		IdreturnStatus       int         `json:"idreturn_status"`
		Idcustomer           int         `json:"idcustomer"`
		Idorder              int         `json:"idorder"`
		Idtemplate           int         `json:"idtemplate"`
		Returnid             string      `json:"returnid"`
		Name                 string      `json:"name"`
		Contactname          interface{} `json:"contactname"`
		Address              string      `json:"address"`
		Address2             interface{} `json:"address2"`
		Zipcode              string      `json:"zipcode"`
		City                 string      `json:"city"`
		Region               string      `json:"region"`
		Country              string      `json:"country"`
		FullAddress          string      `json:"full_address"`
		Telephone            interface{} `json:"telephone"`
		Emailaddress         interface{} `json:"emailaddress"`
		Language             string      `json:"language"`
		Reference            string      `json:"reference"`
		TrackingCode         interface{} `json:"tracking_code"`
		ReceivedAt           interface{} `json:"received_at"`
		CompletedAt          interface{} `json:"completed_at"`
		CreatedAt            string      `json:"created_at"`
		UpdatedAt            string      `json:"updated_at"`
		IdfulfilmentCustomer interface{} `json:"idfulfilment_customer"`
		ReturnStatus         struct {
			IdreturnStatus int    `json:"idreturn_status"`
			Name           string `json:"name"`
			Default        bool   `json:"default"`
			Completed      bool   `json:"completed"`
			Color          string `json:"color"`
			CreatedAt      string `json:"created_at"`
			UpdatedAt      string `json:"updated_at"`
		} `json:"return_status"`
		Customer struct {
			Idcustomer  int         `json:"idcustomer"`
			Customerid  string      `json:"customerid"`
			Name        string      `json:"name"`
			Contactname interface{} `json:"contactname"`
		} `json:"customer"`
		Order struct {
			Idorder   int    `json:"idorder"`
			Orderid   string `json:"orderid"`
			Reference string `json:"reference"`
			CreatedAt string `json:"created_at"`
		} `json:"order"`
		ReturnedProducts []struct {
			IdreturnProduct int         `json:"idreturn_product"`
			Idreturn        int         `json:"idreturn"`
			IdreturnReason  int         `json:"idreturn_reason"`
			Idproduct       int         `json:"idproduct"`
			Idwarehouse     interface{} `json:"idwarehouse"`
			Price           float64     `json:"price"`
			Amount          int         `json:"amount"`
			Status          string      `json:"status"`
			Changeable      bool        `json:"changeable"`
			Product         struct {
				Idproduct   int    `json:"idproduct"`
				ImageURL    string `json:"image_url"`
				Productcode string `json:"productcode"`
				Name        string `json:"name"`
			} `json:"product"`
			ReturnReason struct {
				IdreturnReason int    `json:"idreturn_reason"`
				Name           string `json:"name"`
			} `json:"return_reason"`
		} `json:"returned_products"`
		ReplacementProducts []struct {
			IdreturnProductReplacement int         `json:"idreturn_product_replacement"`
			Idreturn                   int         `json:"idreturn"`
			Idproduct                  int         `json:"idproduct"`
			Idpicklist                 interface{} `json:"idpicklist"`
			Price                      float64     `json:"price"`
			Amount                     int         `json:"amount"`
			Status                     string      `json:"status"`
			Changeable                 bool        `json:"changeable"`
			Product                    struct {
				Idproduct   int    `json:"idproduct"`
				ImageURL    string `json:"image_url"`
				Productcode string `json:"productcode"`
				Name        string `json:"name"`
			} `json:"product"`
		} `json:"replacement_products"`
		Totals struct {
			ReturnedProductsCount    int         `json:"returned_products_count"`
			ReturnedProductsValue    float64     `json:"returned_products_value"`
			ReplacementProductsCount int         `json:"replacement_products_count"`
			ReplacementProductsValue float64     `json:"replacement_products_value"`
			DaysAfterOrderShipped    interface{} `json:"days_after_order_shipped"`
		} `json:"totals"`
	} `json:"data"`
}
