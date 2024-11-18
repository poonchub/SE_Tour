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
		&entity.TourDescriptions{},
		&entity.TourImages{},
		&entity.TourPackages{},
		&entity.TourPrices{},
		&entity.TourSchedules{},
		&entity.TourScheduleStatuses{},
		&entity.Transportations{},
		&entity.TravelInsurances{},
		&entity.Vehicles{},
		&entity.VehicleTypes{},
	)

	// Create Person Type
	personTypes := []*entity.PersonTypes{
		{
			TypeName: "เด็กเล็ก (อายุ 1-3 ปี)",
		},
		{
			TypeName: "เด็ก (อายุ 4-12 ปี) หรือ ผู้ใหญ่",
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
			LogoPath:   "images/logoPaymentMethods/Krung_Thai.png",
		},
		{
			MethodName: "TrueMoney",
			LogoPath:   "images/logoPaymentMethods/True_Money.png",
		},
		{
			MethodName: "SCB",
			LogoPath:   "images/logoPaymentMethods/SCB.png",
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
			LogoPath:     "images/logoProviders/เมืองไทยประกันภัย.png",
		},
		{
			ProviderName: "เอ็ม เอส ไอ จี",
			LogoPath:     "images/logoProviders/เอ็ม-เอส-ไอ-จี.png",
		},
		{
			ProviderName: "ประกันภัยไทยวิวัฒน์",
			LogoPath:     "images/logoProviders/ประกันภัยไทยวิวัฒน์.png",
		},
		{
			ProviderName: "ทิพยประกันภัย",
			LogoPath:     "images/logoProviders/ทิพยประกันภัย.png",
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

	// Create Vehicle Type
	vehicles := []*entity.Vehicles{
		{
			VehicleName:   "รถบัสประจำทัวร์ 1",
			VehicleTypeID: 1,
		},
		{
			VehicleName:   "เรือประจำทัวร์ 1",
			VehicleTypeID: 2,
		},
	}
	for _, vehicle := range vehicles {
		db.FirstOrCreate(vehicle, &entity.Vehicles{
			VehicleName: vehicle.VehicleName,
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
			ProvinceName: "กระบี่",
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
		},
		{
			LocationName: "เกาะค้างคาว",
		},
		{
			LocationName: "หาดบางเบน",
		},
		{
			LocationName: "สนามบิน",
		},
		{
			LocationName: "ท่าเรือ",
		},
		{
			LocationName: "ที่พัก",
		},
	}
	for _, location := range locations {
		db.FirstOrCreate(location, &entity.Locations{
			LocationName: location.LocationName,
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
		{
			TypeName: "เพิ่มเตียงเสริม",
		},
		{
			TypeName: "ไม่เพิ่มเตียงเสริม",
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
		UserName:    "se",
		FirstName:   "SE",
		LastName:    "67",
		Email:       "se67@gmail.com",
		Password:    hashedPassword,
		PhoneNumber: "000-000-0000",
		ProfilePath: "images/profileEmployee/employee1.jpg/",
		RoleID:      3,
	}
	db.FirstOrCreate(employee, &entity.Employees{
		Email: employee.Email,
	})

	// Create Customer
	customer := &entity.Customers{
		UserName:    "Peter",
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
	tourPackages := []*entity.TourPackages{
		{
			PackageCode: "T00001",
			TourName:    "แพ็กเกจทัวร์ทะเลระนอง: เปิดประสบการณ์สู่มนต์เสน่ห์แห่งอันดามันใต้",
			Duration:    "2 วัน 2 คืน",
			ProvinceID:  1,
		},
		{
			PackageCode: "T00002",
			TourName:    "แพ็กเกจทัวร์ทะเลกระบี่: สวรรค์แห่งอันดามันที่ต้องสัมผัส",
			Duration:    "4 วัน 3 คืน",
			ProvinceID:  2,
		},
	}
	for _, tourPackage := range tourPackages {
		db.FirstOrCreate(tourPackage, &entity.TourPackages{
			PackageCode: tourPackage.PackageCode,
		})
	}

	// Create Tour Description
	tourDescriptions := []*entity.TourDescriptions{
		{
			Intro: "สัมผัสความเงียบสงบและธรรมชาติที่บริสุทธิ์ของทะเลระนอง เปิดโลกการท่องเที่ยวสุดเอ็กซ์คลูซีฟกับแพ็กเกจทัวร์ทะเลระนอง ดินแดนที่ยังคงความงดงามดั้งเดิมของธรรมชาติ ทะเลใส หาดทรายขาว และหมู่เกาะที่ซ่อนตัวอยู่ในความสงบ เหมาะสำหรับผู้ที่ต้องการพักผ่อนและหลีกหนีจากความวุ่นวายของชีวิตประจำวัน",
			PackageDetail: "จังหวัด: ระนอง\nระยะเวลา: 2 วัน 2 คืน\nที่พัก: รีสอร์ทหรูระดับ 4 ดาว พร้อมวิวทะเลส่วนตัว\nรวมอาหาร: อาหารครบทุกมื้อ เน้นซีฟู้ดสดใหม่จากทะเลระนอง\nการเดินทาง: รถรับ-ส่งสนามบินระนอง พร้อมเรือสปีดโบ๊ทส่วนตัวสำหรับทัวร์เกาะ",
			TripHighlight:	"🏝️ ทัวร์หมู่เกาะกำและเกาะค้างคาว สัมผัสหาดทรายขาวละเอียดและน้ำทะเลใสราวคริสตัลที่หมู่เกาะกำ เพลิดเพลินกับการดำน้ำตื้น(Snorkeling) เพื่อชมปะการังหลากสีสัน และปลาสวยงามนานาชนิด\n🌅 ชิลล์บนเกาะพยาม สัมผัสกับวิถีชีวิตเรียบง่ายบนเกาะพยาม เกาะที่ขึ้นชื่อว่า \"มัลดีฟส์เมืองไทย\" ชมความงดงามของชายหาดที่เงียบสงบและดื่มด่ำกับพระอาทิตย์ตกดินสุดโรแมนติก\n🛶 ล่องเรือชมป่าชายเลนและน้ำตกปุญญบาล ตื่นตาตื่นใจกับความสมบูรณ์ของป่าชายเลนที่อุดมสมบูรณ์ สำรวจน้ำตกปุญญบาลที่มีน้ำใสไหลเย็นตลอดปี\n✨ บ่อน้ำพุร้อนรักษะวาริน ผ่อนคลายกับการแช่น้ำแร่ร้อนธรรมชาติที่บ่อน้ำพุร้อนรักษะวาริน สถานที่ท่องเที่ยวเพื่อสุขภาพที่ไม่ควรพลาด",
			PlacesHighlight: "🌊 ดำน้ำดูปะการังและปลาสวยงาม: สัมผัสประสบการณ์ดำน้ำที่น่าตื่นเต้นในน้ำทะเลใสแจ๋ว\n🌴 พักผ่อนท่ามกลางธรรมชาติ: ที่พักสุดหรูที่โอบล้อมด้วยธรรมชาติที่งดงาม\n🍹 ปาร์ตี้ริมชายหาด: เพลิดเพลินกับปาร์ตี้บาร์บีคิวซีฟู้ดและเครื่องดื่มริมทะเลยามค่ำคืน",
			TourPackageID: 1,
		},
		{
			Intro: "สัมผัสความงามทะเลอันดามัน ดินแดนสวรรค์ของนักท่องเที่ยว ค้นพบเสน่ห์ของทะเลกระบี่ด้วยแพ็กเกจทัวร์สุดพิเศษ ที่จะพาคุณไปดื่มด่ำกับธรรมชาติที่สวยงาม น้ำทะเลใส หาดทรายขาวละเอียด และเกาะแก่งน้อยใหญ่ที่รอคอยการมาเยือนของคุณ กระบี่พร้อมมอบประสบการณ์การพักผ่อนที่จะทำให้คุณลืมความวุ่นวายในชีวิตประจำวัน",
			PackageDetail: "จังหวัด: กระบี่\nระยะเวลา: 4 วัน 3 คืน\nที่พัก: รีสอร์ทระดับ 4-5 ดาว ริมชายหาดพร้อมวิวทะเลทุกห้องพัก\nรวมอาหาร: อาหารเช้า, กลางวัน และมื้อค่ำแบบซีฟู้ดสดใหม่ทุกวัน\nการเดินทาง: รถรับ-ส่งจากสนามบินกระบี่ พร้อมเรือสปีดโบ๊ทส่วนตัวสำหรับทัวร์เกาะ",
			TripHighlight:	"🏝️ ทัวร์ 4 เกาะสุดฮิต เยือนเกาะปอดะ, ทะเลแหวก, เกาะไก่ และเกาะทับ สถานที่ท่องเที่ยวที่ขึ้นชื่อที่สุดในกระบี่ เพลิดเพลินกับการดำน้ำชมปะการังและฝูงปลาหลากสีสัน\n🌅 ล่องเรือหางยาวชมพระอาทิตย์ตก ดื่มด่ำกับความงามของพระอาทิตย์ตกดินที่หาดอ่าวนาง พร้อมบรรยากาศสุดโรแมนติก\n🛶 พายเรือคายัคที่อ่าวท่าเลน สำรวจป่าโกงกางและถ้ำหินปูนที่ซ่อนตัวอยู่ในป่าชายเลน ด้วยการพายเรือคายัคที่สนุกและท้าทาย\n✨ เที่ยวหมู่เกาะพีพี ดื่มด่ำกับความสวยงามของเกาะพีพีเล ชมอ่าวมาหยา สถานที่ถ่ายทำภาพยนตร์ชื่อดัง “The Beach” และอ่าวปิเละที่มีน้ำทะเลใสเหมือนกระจก",
			PlacesHighlight: "🌊 ดำน้ำตื้น (Snorkeling) ที่อ่าวลิง: สัมผัสชีวิตใต้ทะเลกับปลาหลากสีสันและปะการังที่สมบูรณ์\n🌴 ชิลล์ริมชายหาดที่ไร่เลย์: พักผ่อนบนหาดทรายขาวละเอียด ท่ามกลางหน้าผาหินปูนสูงตระหง่าน\n💆 สปาผ่อนคลาย: ผ่อนคลายกับการนวดแผนไทยที่สปาริมชายหาด",
			TourPackageID: 2,
		},
	}
	for _, tourDescription := range tourDescriptions {
		db.FirstOrCreate(tourDescription, &entity.TourDescriptions{
			TourPackageID: tourDescription.TourPackageID,
		})
	}

	// Create Tour Schedule
	StartDate1, _ := time.Parse("2006-01-02", "2024-11-10")
	EndDate1, _ := time.Parse("2006-01-02", "2024-11-11")
	StartDate2, _ := time.Parse("2006-01-02", "2024-11-26")
	EndDate2, _ := time.Parse("2006-01-02", "2024-11-27")
	StartDate3, _ := time.Parse("2006-01-02", "2024-12-01")
	EndDate3, _ := time.Parse("2006-01-02", "2024-12-02")
	tourSchedules := []*entity.TourSchedules{
		{
			StartDate:            StartDate1,
			EndDate:              EndDate1,
			AvailableSlots:       50,
			TourPackageID:        1,
			TourScheduleStatusID: 2,
		},
		{
			StartDate:            StartDate2,
			EndDate:              EndDate2,
			AvailableSlots:       50,
			TourPackageID:        1,
			TourScheduleStatusID: 2,
		},
		{
			StartDate:            StartDate3,
			EndDate:              EndDate3,
			AvailableSlots:       50,
			TourPackageID:        1,
			TourScheduleStatusID: 2,
		},
	}
	for _, tourSchedule := range tourSchedules {
		db.FirstOrCreate(tourSchedule, &entity.TourSchedules{
			StartDate:     tourSchedule.StartDate,
			EndDate:       tourSchedule.EndDate,
			TourPackageID: tourSchedule.TourPackageID,
		})
	}

	// Create Promotion
	ValidFrom, _ := time.Parse("2006-01-02", "2024-11-01")
	ValidUntil, _ := time.Parse("2006-01-02", "2024-11-30")
	promotions := []*entity.Promotions{
		{
			PromotionCode:      "P00001",
			PromotionName:      "ส่วนลด 5% เมื่อซื้อแพ็กเกจครบ 3000 บาท",
			DiscountPercentage: 5.0,
			ValidFrom:          ValidFrom,
			ValidUntil:         ValidUntil,
			Minimum_Price:      3000.00,
			PromotionStatusID:  1,
		},
	}
	for _, promotion := range promotions {
		db.FirstOrCreate(promotion, &entity.Promotions{
			PromotionCode: promotion.PromotionCode,
		})
	}

	// Create Tour Image
	for i := uint(1); i <= 2; i++ {
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
			Price: 500,
			TourPackageID: 1,
			PersonTypeID: 1,	// เด็กเล็ก
			RoomTypeID: 4,		// เพิ่มเตียง
		},
		{
			Price: 0,
			TourPackageID: 1,
			PersonTypeID: 1,	// เด็กเล็ก
			RoomTypeID: 5,		// ไม่เพิ่มเตียง
		},
		{
			Price: 8290,
			TourPackageID: 1,
			PersonTypeID: 2,	// เด็ก/ผู้ใหญ่
			RoomTypeID: 1,		// พักเดี่ยว
		},
		{
			Price: 6990,
			TourPackageID: 1,
			PersonTypeID: 2,	// เด็ก/ผู้ใหญ่
			RoomTypeID: 2,		// คู่
		},
		{
			Price: 6990,
			TourPackageID: 1,
			PersonTypeID: 2,	// เด็ก/ผู้ใหญ่
			RoomTypeID: 3,		// พักสาม
		},
		{
			Price: 500,
			TourPackageID: 2,
			PersonTypeID: 1,	// เด็กเล็ก
			RoomTypeID: 4,		// เพิ่มเตียง
		},
		{
			Price: 0,
			TourPackageID: 2,
			PersonTypeID: 1,	// เด็กเล็ก
			RoomTypeID: 5,		// ไม่เพิ่มเตียง
		},
		{
			Price: 6590,
			TourPackageID: 2,
			PersonTypeID: 2,	// เด็ก/ผู้ใหญ่
			RoomTypeID: 1,		// พักเดี่ยว
		},
		{
			Price: 5190,
			TourPackageID: 2,
			PersonTypeID: 2,	// เด็ก/ผู้ใหญ่
			RoomTypeID: 2,		// คู่
		},
		{
			Price: 5190,
			TourPackageID: 2,
			PersonTypeID: 2,	// เด็ก/ผู้ใหญ่
			RoomTypeID: 3,		// พักสาม
		},
	}
	for _, tourPrice := range tourPrices {
		db.FirstOrCreate(tourPrice, &entity.TourPrices{
			TourPackageID: tourPrice.TourPackageID,
			PersonTypeID:  tourPrice.PersonTypeID,
			RoomTypeID: tourPrice.RoomTypeID,
		})
	}

	// Create Transportation
	DepartureTime1, _ := time.Parse("2006-01-02 15:04:05", "2024-11-20 08:00:00")
	ArrivalTime1, _ := time.Parse("2006-01-02 15:04:05", "2024-11-20 08:30:00")
	DepartureTime2, _ := time.Parse("2006-01-02 15:04:05", "2024-11-20 10:00:00")
	ArrivalTime2, _ := time.Parse("2006-01-02 15:04:05", "2024-11-20 10:30:00")
	DepartureTime3, _ := time.Parse("2006-01-02 15:04:05", "2024-11-21 13:00:00")
	ArrivalTime3, _ := time.Parse("2006-01-02 15:04:05", "2024-11-20 13:30:00")
	transportations := []*entity.Transportations{
		{
			DepartureTime: DepartureTime1,
			ArrivalTime:   ArrivalTime1,
			VehicleID:     1,
			TourPackageID: 1,
			LocationID:    1,
		},
		{
			DepartureTime: DepartureTime2,
			ArrivalTime:   ArrivalTime2,
			VehicleID:     2,
			TourPackageID: 1,
			LocationID:    2,
		},
		{
			DepartureTime: DepartureTime3,
			ArrivalTime:   ArrivalTime3,
			VehicleID:     1,
			TourPackageID: 1,
			LocationID:    3,
		},
	}
	for _, transportation := range transportations {
		db.FirstOrCreate(transportation, &entity.Transportations{
			DepartureTime: transportation.DepartureTime,
			TourPackageID: transportation.TourPackageID,
		})
	}

	// Create Activity
	DateTime1, _ := time.Parse("2006-01-02 15:04:05", "2024-11-20 08:00:00")
	DateTime2, _ := time.Parse("2006-01-02 15:04:05", "2024-11-20 09:00:00")
	DateTime3, _ := time.Parse("2006-01-02 15:04:05", "2024-11-20 10:30:00")
	DateTime4, _ := time.Parse("2006-01-02 15:04:05", "2024-11-21 07:30:00")
	activities := []*entity.Activities{
		{
			ActivityName:  "ถึงจุดนัดรับ",
			Description:   "ทีมงานต้อนรับที่สนามบินระนอง (หรือจุดนัดพบ) พร้อมบริการรับส่งด้วยรถตู้ปรับอากาศ",
			DateTime: DateTime1,
			LocationID:    4,
			TourPackageID: 1,
		},
		{
			ActivityName:  "เดินทางไปยังท่าเรือ",
			Description:   "เดินทางไปยังท่าเรือเพื่อขึ้นเรือสปีดโบ๊ท มุ่งหน้าสู่ เกาะพยาม (ใช้เวลาประมาณ 45 นาที) เพลิดเพลินกับบรรยากาศทะเลสวยงามและทิวทัศน์ระหว่างการเดินทาง",
			DateTime: DateTime2,
			LocationID:    5,
			TourPackageID: 1,
		},
		{
			ActivityName:  "ถึงเกาะพยาม",
			Description:   "ถึงเกาะพยาม และเยี่ยมชมวัดเกาะพยาม เที่ยวชมสถานที่สำคัญต่าง ๆ เช่น อ่าวเขาควาย และ อ่าวใหญ่",
			DateTime: DateTime3,
			LocationID:    1,
			TourPackageID: 1,
		},
		{
			ActivityName:  "เช็คเอาท์จากที่พัก",
			Description:   "รับประทานอาหารเช้าที่รีสอร์ท และเช็คเอาท์จากที่พัก",
			DateTime: DateTime4,
			LocationID:    6,
			TourPackageID: 1,
		},
	}
	for _, activity := range activities {
		db.FirstOrCreate(activity, &entity.Activities{
			ActivityName: activity.ActivityName,
			DateTime: activity.DateTime,
			LocationID:   activity.LocationID,
			TourPackageID: activity.TourPackageID,
		})
	}

	// Create Accommodation
	CheckInDate, _ := time.Parse("2006-01-02", "2024-11-20")
	CheckOutDate, _ := time.Parse("2006-01-02", "2024-11-21")
	accommodations := []*entity.Accommodations{
		{
			CheckInDate:   CheckInDate,
			CheckOutDate:  CheckOutDate,
			TourPackageID: 1,
			HotelID:       1,
		},
	}
	for _, accommodation := range accommodations {
		db.FirstOrCreate(accommodation, &entity.Accommodations{
			CheckInDate:   accommodation.CheckInDate,
			CheckOutDate:  accommodation.CheckOutDate,
			TourPackageID: accommodation.TourPackageID,
			HotelID:       accommodation.HotelID,
		})
	}

	// Create Booking Room
	bookingRooms := []*entity.BookingRooms{
		{
			RoomQuantity:    0,
			AccommodationID: 1,
			RoomTypeID:      1,
		},
		{
			RoomQuantity:    0,
			AccommodationID: 1,
			RoomTypeID:      2,
		},
	}
	for _, bookingRoom := range bookingRooms {
		db.FirstOrCreate(bookingRoom, &entity.BookingRooms{
			AccommodationID: bookingRoom.AccommodationID,
			RoomTypeID:      bookingRoom.RoomTypeID,
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
