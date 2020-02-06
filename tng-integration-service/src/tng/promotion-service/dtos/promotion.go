package dtos

type MCInfo struct {
	MCCode string `json:"mc_code"`
	AppID  string `json:"appid"`
}

type Customers struct {
	UCode    string `json:"ucode"`
	FullName string `json:"fullname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
type Order struct {
	OrderNo string `json:"orderno"`
	Amt     string `json:"amt"`
	Created string `json:"created"`
}

type OrderResponse struct {
	OrderNo string `json:"orderno"`
	Amt     string `json:"amt"`
	Created string `json:"created"`
	OldAmt  string `json:"old_amt"`
}

type OrderDetail struct {
	UCode     string `json:"ucode"`
	Type      string `json:"type"`
	Price     int64  `json:"price"`
	Quantity  int    `json:"quantity"`
	ItemTotal int64  `json:"item_total"`
}

type OrderDetailResponse struct {
	UCode     string `json:"ucode"`
	Type      string `json:"type"`
	Price     int64  `json:"price"`
	OldPrice  int64  `json:"old_price"`
	Quantity  int    `json:"quantity"`
	ItemTotal int64  `json:"item_total"`
}

type TixPromotionDetail struct {
	UCode          string `json:"ucode"`
	PaymentDate    string `json:"payment_date"`
	DeliveryDate   string `json:"delivery_date"`
	PCinemaCode    string `json:"p_cinema_code"`
	CinemaCode     string `json:"cinema_code"`
	RoomCode       string `json:"room_code"`
	FilmCode       string `json:"film_code"`
	SessionCode    string `json:"session_code"`
	SeatType       string `json:"seat_type"`
	SeatList       string `json:"seat_list"`
	EVN            string `json:"evn"`
	Channel        string `json:"channel"`
	PaymentMethod  string `json:"payment_method"`
	PaymentChannel string `json:"payment_channel"`
	BankCode       string `json:"bank_code"`
}

type TixPromotionData struct {
	MCInfo      *MCInfo             `json:"mc_info"`
	Promotion   *TixPromotionDetail `json:"promotion"`
	Customers   *Customers          `json:"customers"`
	Order       *Order              `json:"order"`
	OrderDetail *OrderDetail        `json:"order_detail"`
}
type TixPromotionDataResponse struct {
	OrderResponse       *OrderResponse       `json:"order"`
	OrderDetailResponse *OrderDetailResponse `json:"order_detail"`
	Campaign            *Campaign            `json:"campaign"`
}
