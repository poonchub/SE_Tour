package config

import (
	"fmt"
	"io/ioutil"
	"time"
	"toursystem/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("TourSystems.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db = database
}

func SetupDatabase() {

	db.AutoMigrate(
		&entity.Accommodations{},
		&entity.Activities{},
		&entity.BookingDetails{},
		&entity.BookingRooms{},
		&entity.Bookings{},
		&entity.BookingStatuses{},
		&entity.Customers{},
		&entity.Employees{},
		&entity.EmployeeSchedules{},
		&entity.Hotels{},
		&entity.Locations{},
		&entity.Meals{},
		&entity.MealTypes{},
		&entity.PaymentMethods{},
		&entity.Payments{},
		&entity.PaymentStatuses{},
		&entity.PersonTypes{},
		&entity.Promotions{},
		&entity.PromotionStatuses{},
		&entity.Providers{},
		&entity.Provinces{},
		&entity.Roles{},
		&entity.RoomTypes{},
		&entity.SalesReports{},
		&entity.Slips{},
		&entity.TourImages{},
		&entity.TourPackages{},
		&entity.TourPrices{},
		&entity.TourSchedules{},
		&entity.TourScheduleStatuses{},
		&entity.Transportations{},
		&entity.TravelInsurances{},
		&entity.VehicleTypes{},
	)

	// Create Person Type
	personTypes := []*entity.PersonTypes{
		{
			TypeName: "เด็ก",
		},
		{
			TypeName: "ผู้ใหญ่",
		},
	}
	for _, personType := range personTypes {
		db.FirstOrCreate(personType, &entity.PersonTypes{
			TypeName: personType.TypeName,
		})
	}

	// Create Booking Status
	bookingStatuses := []*entity.BookingStatuses{
		{
			StatusName: "ยกเลิกแล้ว",
		},
		{
			StatusName: "รอการตรวจสอบ",
		},
		{
			StatusName: "จองสำเร็จ",
		},
	}
	for _, status := range bookingStatuses {
		db.FirstOrCreate(status, &entity.BookingStatuses{
			StatusName: status.StatusName,
		})
	}

	// Create Payment Status
	paymentStatuses := []*entity.PaymentStatuses{
		{
			StatusName: "รอการชำระเงิน",
		},
		{
			StatusName: "รอการตรวจสอบ",
		},
		{
			StatusName: "ชำระเงินสำเร็จ",
		},
	}
	for _, status := range paymentStatuses {
		db.FirstOrCreate(status, &entity.PaymentStatuses{
			StatusName: status.StatusName,
		})
	}

	// Create Tour Schedule Status
	tourScheduleStatuses := []*entity.TourScheduleStatuses{
		{
			StatusName: "เต็ม",
		},
		{
			StatusName: "ยังไม่เต็ม",
		},
		{
			StatusName: "ถูกยกเลิก",
		},
	}
	for _, status := range tourScheduleStatuses {
		db.FirstOrCreate(status, &entity.TourScheduleStatuses{
			StatusName: status.StatusName,
		})
	}

	// Create Payment Method
	paymentMethods := []*entity.PaymentMethods{
		{
			MethodName: "Krungthai",
			LogoPath: "images/logoPaymentMethods/Krung_Thai.png",
		},
		{
			MethodName: "TrueMoney",
			LogoPath: "images/logoPaymentMethods/True_Money.png",
		},
		{
			MethodName: "SCB",
			LogoPath: "images/logoPaymentMethods/SCB.png",
		},
	}
	for _, method := range paymentMethods {
		db.FirstOrCreate(method, &entity.PaymentMethods{
			MethodName: method.MethodName,
		})
	}

	// Create Provider
	providers := []*entity.Providers{
		{
			ProviderName: "เมืองไทยประกันภัย",
			LogoPath: "images/logoProviders/เมืองไทยประกันภัย.png",
		},
		{
			ProviderName: "เอ็ม เอส ไอ จี",
			LogoPath: "images/logoProviders/เอ็ม-เอส-ไอ-จี.png",
		},
		{
			ProviderName: "ประกันภัยไทยวิวัฒน์",
			LogoPath: "images/logoProviders/ประกันภัยไทยวิวัฒน์.png",
		},
		{
			ProviderName: "ทิพยประกันภัย",
			LogoPath: "images/logoProviders/ทิพยประกันภัย.png",
		},
	}
	for _, provider := range providers {
		db.FirstOrCreate(provider, &entity.Providers{
			ProviderName: provider.ProviderName,
		})
	}

	// Create Vehicle Type
	vehicleTypes := []*entity.VehicleTypes{
		{
			TypeName: "รถทัวร์",
		},
		{
			TypeName: "เรือ",
		},
	}
	for _, vehicleType := range vehicleTypes {
		db.FirstOrCreate(vehicleType, &entity.VehicleTypes{
			TypeName: vehicleType.TypeName,
		})
	}

	// Create Meal Type
	mealTypes := []*entity.MealTypes{
		{
			TypeName: "อาหารเช้า",
		},
		{
			TypeName: "อาหารกลางวัน",
		},
		{
			TypeName: "อาหารเย็น",
		},
	}
	for _, mealType := range mealTypes {
		db.FirstOrCreate(mealType, &entity.MealTypes{
			TypeName: mealType.TypeName,
		})
	}

	// Create Role
	roles := []*entity.Roles{
		{
			RoleName: "แอดมิน",
		},
		{
			RoleName: "คนขับรถทัวร์",
		},
		{
			RoleName: "ไกด์",
		},
	}
	for _, role := range roles {
		db.FirstOrCreate(role, &entity.Roles{
			RoleName: role.RoleName,
		})
	}

	// Create Promotion Status
	promotionStatuses := []*entity.PromotionStatuses{
		{
			StatusName: "เปิดใช้งาน",
		},
		{
			StatusName: "ปิดใช้งาน",
		},
	}
	for _, status := range promotionStatuses {
		db.FirstOrCreate(status, &entity.PromotionStatuses{
			StatusName: status.StatusName,
		})
	}

	// Create Province Status
	provinces := []*entity.Provinces{
		{
			ProvinceName: "ระนอง",
		},
		{
			ProvinceName: "ระยอง",
		},
		{
			ProvinceName: "ยะลา",
		},
	}
	for _, province := range provinces {
		db.FirstOrCreate(province, &entity.Provinces{
			ProvinceName: province.ProvinceName,
		})
	}

	// Create Location
	locations := []*entity.Locations{
		{
			LocationName: "วัดเกาะพยาม",
			ProvinceID: 1,
		},
		{
			LocationName: "เกาะค้างคาว",
			ProvinceID: 1,
		},
		{
			LocationName: "หาดบางเบน",
			ProvinceID: 1,
		},
	}
	for _, location := range locations {
		db.FirstOrCreate(location, &entity.Locations{
			LocationName: location.LocationName,
			ProvinceID: location.ProvinceID,
		})
	}

	// Create Room Type
	roomtypes := []*entity.RoomTypes{
		{
			TypeName: "พักเดี่ยว",
		},
		{
			TypeName: "พักคู่",
		},
		{
			TypeName: "พักสาม",
		},
	}
	for _, roomtype := range roomtypes {
		db.FirstOrCreate(roomtype, &entity.RoomTypes{
			TypeName: roomtype.TypeName,
		})
	}

	// Create Hotel
	hotels := []*entity.Hotels{
		{
			HotelName: "พักดี",
		},
		{
			HotelName: "สบายกาย",
		},
		{
			HotelName: "สบายใจ",
		},
	}
	for _, hotel := range hotels {
		db.FirstOrCreate(hotel, &entity.Hotels{
			HotelName: hotel.HotelName,
		})
	}

	// Create Employee
	hashedPassword, _ := HashPassword("123456")
	employee := &entity.Employees{
		UserName: "se",
		FirstName:   "SE",
		LastName:    "67",
		Email:       "se67@gmail.com",
		Password:    hashedPassword,
		PhoneNumber: "000-000-0000",
		ProfilePath: "images/profileEmployee/employee1.jpg/",
		RoleID: 3,
	}
	db.FirstOrCreate(employee, &entity.Employees{
		Email: employee.Email,
	})

	// Create Customer
	customer := &entity.Customers{
		UserName:   "Peter",
		FirstName:   "Peter",
		LastName:    "Parker",
		Email:       "spider@gmail.com",
		Password:    hashedPassword,
		PhoneNumber: "098-594-4576",
		ProfilePath: "images/profileCustomer/customer1.jpg",
	}
	db.FirstOrCreate(customer, &entity.Customers{
		Email: customer.Email,
	})

	// Create Tour Package
	tourpackage := &entity.TourPackages{
		PackageCode: "T00001",
		TourName: "แพ็กเกจทัวร์ทะเลระนอง: เปิดประสบการณ์สู่มนต์เสน่ห์แห่งอันดามันใต้",
		Description: "สัมผัสความเงียบสงบและธรรมชาติที่บริสุทธิ์ของทะเลระนอง เปิดโลกการท่องเที่ยวสุดเอ็กซ์คลูซีฟกับแพ็กเกจทัวร์ทะเลระนอง ดินแดนที่ยังคงความงดงามดั้งเดิมของธรรมชาติ ทะเลใส หาดทรายขาว และหมู่เกาะที่ซ่อนตัวอยู่ในความสงบ เหมาะสำหรับผู้ที่ต้องการพักผ่อนและหลีกหนีจากความวุ่นวายของชีวิตประจำวัน",
		Duration: "2 วัน 2 คืน",
	}
	db.FirstOrCreate(tourpackage, &entity.TourPackages{
		PackageCode: tourpackage.PackageCode,
	})

	// Create Tour Schedule
	StartDate1, _ := time.Parse("2006-01-02", "2024-11-20")
	EndDate1, _ := time.Parse("2006-01-02", "2024-11-21")
	StartDate2, _ := time.Parse("2006-01-02", "2024-12-01")
	EndDate2, _ := time.Parse("2006-01-02", "2024-12-02")
	tourSchedules := []*entity.TourSchedules{
		{
			StartDate: StartDate1,
			EndDate: EndDate1,
			AvailableSlots: 50,
			TourPackageID: 1,
			TourScheduleStatusID: 2,
		},
		{
			StartDate: StartDate2,
			EndDate: EndDate2,
			AvailableSlots: 50,
			TourPackageID: 1,
			TourScheduleStatusID: 2,
		},
	}
	for _, tourSchedule := range tourSchedules {
		db.FirstOrCreate(tourSchedule, &entity.TourSchedules{
			StartDate: tourSchedule.StartDate,
			EndDate: tourSchedule.EndDate,
			TourPackageID: tourSchedule.TourPackageID,
		})
	}

	// Create Promotion
	ValidFrom, _ := time.Parse("2006-01-02", "2024-11-01")
	ValidUntil, _ := time.Parse("2006-01-02", "2024-11-30")
	promotions := []*entity.Promotions{
		{
			PromotionCode: "P00001",
			PromotionName: "ส่วนลด 5% เมื่อซื้อแพ็กเกจครบ 3000 บาท",
			DiscountPercentage: 5.0,
			ValidFrom: ValidFrom,
			ValidUntil: ValidUntil,
			Minimum_Price: 3000.00,
			PromotionStatusID: 1,
		},
	}
	for _, promotion := range promotions {
		db.FirstOrCreate(promotion, &entity.Promotions{
			PromotionCode: promotion.PromotionCode,
		})
	}

	// Create Tour Image 
	for i := uint(1); i <= 1; i++ {
		dir := fmt.Sprintf("images/tourImages/tourPackage%d", i)
		count := countFilesInDir(dir)
		for j := 1; j <= count; j++ {
			filePath := fmt.Sprintf("images/tourImages/tourPackage%d/tour0%d.jpg", i, j)
			err := createImage(filePath, i)
			if err != nil {
				panic(err)
			}
		}
	}

	// Create Tour Price
	tourPrices := []*entity.TourPrices{
		{
			Price: 8299,
			TourPackageID: 1,
			PersonTypeID: 1,
		},
		{
			Price: 7299,
			TourPackageID: 1,
			PersonTypeID: 2,
		},
	}
	for _, tourPrice := range tourPrices {
		db.FirstOrCreate(tourPrice, &entity.TourPrices{
			TourPackageID: tourPrice.TourPackageID,
			PersonTypeID: tourPrice.PersonTypeID,
		})
	}

	// Create Transportation
	DepartureTime1, _ := time.Parse("2006-01-02 15:04:05", "2024-11-20 08:00:00")
	DepartureTime2, _ := time.Parse("2006-01-02 15:04:05", "2024-11-20 10:00:00")
	DepartureTime3, _ := time.Parse("2006-01-02 15:04:05", "2024-11-21 13:00:00")
	transportations := []*entity.Transportations{
		{
			DepartureTime: DepartureTime1,
			VehicleTypeID: 1,
			TourPackageID: 1,
			LocationID: 1,
		},
		{
			DepartureTime: DepartureTime2,
			VehicleTypeID: 1,
			TourPackageID: 1,
			LocationID: 2,
		},
		{
			DepartureTime: DepartureTime3,
			VehicleTypeID: 1,
			TourPackageID: 1,
			LocationID: 3,
		},
	}
	for _, transportation := range transportations {
		db.FirstOrCreate(transportation, &entity.Transportations{
			DepartureTime: transportation.DepartureTime,
			TourPackageID: transportation.TourPackageID,
		})
	}

	// Create Activity
	StartTime1, _ := time.Parse("2006-01-02 15:04:05", "2024-11-20 08:20:00")
	EndTime1, _ := time.Parse("2006-01-02 15:04:05", "2024-11-20 10:00:00")
	StartTime2, _ := time.Parse("2006-01-02 15:04:05", "2024-11-21 11:00:00")
	EndTime2, _ := time.Parse("2006-01-02 15:04:05", "2024-11-21 13:00:00")
	activities := []*entity.Activities{
		{
			ActivityName: "เที่ยวชมวัดเกาะพยาม",
			Description: "เที่ยวชมวัดเกาะพยาม",
			StartTime: StartTime1,
			EndTime: EndTime1,
			LocationID: 1,
			TourPackageID: 1,
		},
		{
			ActivityName: "พักรับประทานอาหาร",
			Description: "พักรับประทานอาหาร",
			StartTime: StartTime2,
			EndTime: EndTime2,
			LocationID: 2,
			TourPackageID: 1,
		},
	}
	for _, activity := range activities {
		db.FirstOrCreate(activity, &entity.Activities{
			ActivityName: activity.ActivityName,
			StartTime: activity.StartTime,
			EndTime: activity.EndTime,
			LocationID: activity.LocationID,
		})
	}

	// Create Accommodation
	CheckInDate, _ := time.Parse("2006-01-02", "2024-11-20")
	CheckOutDate, _ := time.Parse("2006-01-02", "2024-11-21")
	accommodations := []*entity.Accommodations{
		{
			CheckInDate: CheckInDate,
			CheckOutDate: CheckOutDate,
			TourPackageID: 1,
			HotelID: 1,
		},
	}
	for _, accommodation := range accommodations {
		db.FirstOrCreate(accommodation, &entity.Accommodations{
			CheckInDate: accommodation.CheckInDate,
			CheckOutDate: accommodation.CheckOutDate,
			TourPackageID: accommodation.TourPackageID,
			HotelID: accommodation.HotelID,
		})
	}

	// Create Booking Room
	bookingRooms := []*entity.BookingRooms{
		{
			RoomQuantity: 0,
			AccommodationID: 1,
			RoomTypeID: 1,
		},
		{
			RoomQuantity: 0,
			AccommodationID: 1,
			RoomTypeID: 2,
		},
	}
	for _, bookingRoom := range bookingRooms {
		db.FirstOrCreate(bookingRoom, &entity.BookingRooms{
			AccommodationID: bookingRoom.AccommodationID,
			RoomTypeID: bookingRoom.RoomTypeID,
		})
	}


}


func createImage(filePath string, id uint) error {

	image := entity.TourImages{FilePath: filePath, TourPackageID: id}

	if err := db.Where("file_path = ?", &image.FilePath).FirstOrCreate(&image).Error; err != nil {
		return err
	}
	return nil
}

func countFilesInDir(dir string) int {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return 0
	}

	fileCount := 0
	for _, file := range files {
		if !file.IsDir() {
			fileCount++
		}
	}

	return fileCount
}