package entity

// AutoOrderRequestParam 定义下单接口参数结构体
type AutoOrderRequestParam struct {
	Source    string `json:"source" binding:"required"`
	SourceId  string `json:"source_id" binding:"required"`
	Mobile    string `json:"mobile" binding:"required"`
	ShareNum  string `json:"share_num" binding:"required"`
	SendMsn   string `json:"send_msn"`
	OrderInfo string `json:"order_info" binding:"required"`
}

// OrderInfoParam 订单详情结构体
type OrderInfoParam struct {
	Tid     string `json:"tid"`
	Address string `json:"address"`
	Orders  string `json:"orders"`
}

// OrdersCourseParam 订单课程信息结构体
type OrdersCourseParam struct {
	Sid       string        `json:"sid"`
	CourseIds []OrderCourse `json:"course_ids"`
}

// OrderCourse 单个订单的课程信息结构体
type OrderCourse struct {
	CourseType string `json:"course_type"`
	GradeId    string `json:"grade_id"`
	SubjectId  string `json:"subject_id"`
	CourseId   string `json:"course_id"`
}

// OrderAddress 订单 - 收货地址结构体信息
type OrderAddress struct {
	ReceiverName     string `json:"receiver_name"`
	ReceiverState    string `json:"receiver_state"`
	ReceiverCity     string `json:"receiver_city"`
	ReceiverDistrict string `json:"receiver_district"`
	ReceiverAddress  string `json:"receiver_address"`
	ProvinceId       string `json:"province_id"`
	CityId           string `json:"city_id"`
	DistrictId       string `json:"district_id"`
}

// PromotionConfig 订单 - 促销信息结构体
type PromotionConfig struct {
	ProductId string `json:"productId"`
	Type      int    `json:"type"`
	Price     int    `json:"price"`
}

// OrderResult 订单 - 下单处理结果结构体
type OrderResult struct {
	Stat int             `json:"stat"`
	Data OrderResultInfo `json:"data"`
}

// OrderResultInfo 订单 - 下单处理结果的订单信息
type OrderResultInfo struct {
	OrderNum     string `json:"orderNum"`
	UserId       string `json:"user_id"`
	IsExistOrder int    `json:"isExistOrder"`
}
