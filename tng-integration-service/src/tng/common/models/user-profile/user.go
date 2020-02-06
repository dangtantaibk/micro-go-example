package user_profile

type User struct {
	ID         int64  `orm:"column(id)"`
	AppID      int32  `orm:"column(appid)"`
	UCode      string `orm:"column(ucode);size(256)"`
	Title      string `orm:"column(title);size(256)"`
	FirstName  string `orm:"column(firstname);size(45)"`
	SurName    string `orm:"column(surname);size(45)"`
	FullName   string `orm:"column(fullname);size(256)"`
	LastName   string `orm:"column(lastname);size(45)"`
	Phone      string `orm:"column(phone);size(45)"`
	HomePhone  string `orm:"column(home_phone);size(45)"`
	Email      string `orm:"column(email);size(256)"`
	Address    string `orm:"column(address);size(256)"`
	SocialID   string `orm:"column(social_id);size(256)"`
	LoginType  string `orm:"column(login_type);size(256)"`
	WardID     string `orm:"column(ward_id);size(10)"`
	DistrictID string `orm:"column(district_id);size(45)"`
	ProvinceID string `orm:"column(province_id);size(45)"`
	CountryID  string `orm:"column(country_id);size(45)"`
	Created    string `orm:"column(created);size(45)"`
	CreatedBy  string `orm:"column(created_by);size(45)"`
	Avatar     string `orm:"column(avatar);size(512)"`
	Status     string `orm:"column(status);size(8)"`
	Lat        string `orm:"column(lat);size(45)"`
	Long       string `orm:"column(long);size(45)"`
}
