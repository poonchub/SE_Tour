package config

import (
	"fmt"
	"io/ioutil"
	"toursystem/entity"

	"time"

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
		&entity.ScheduleActivities{},
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
			StatusName: "รอการชำระเงิน",
		},
		{
			StatusName: "รอการตรวจสอบ",
		},
		{
			StatusName: "จองสำเร็จ",
		},
		{
			StatusName: "ยกเลิกแล้ว",
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
			TypeName: "บัส",
		},
		{
			TypeName: "เรือ",
		},
		{
			TypeName: "รถตู้",
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
		{
			VehicleName:   "รถตู้ประจำทัวร์ 1",
			VehicleTypeID: 3,
		},
	}
	for _, vehicle := range vehicles {
		db.FirstOrCreate(vehicle, &entity.Vehicles{
			VehicleName: vehicle.VehicleName,
		})
	}

	// Create Vehicle Image
	for i := uint(1); i <= 2; i++ {
		dir := fmt.Sprintf("images/vehicleImages/vehicle%d", i)
		count := countFilesInDir(dir)
		for j := 1; j <= count; j++ {
			filePath := fmt.Sprintf("images/vehicleImages/vehicle%d/vehicle0%d.jpg", i, j)
			err := createImageVehicle(filePath, i)
			if err != nil {
				panic(err)
			}
		}
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

	// Create Meal
	meals := []*entity.Meals{
		{
			MenusDetail:   "ข้าวต้มกุ้ง, ไข่เจียวหมูสับ, ผลไม้ตามฤดูกาล",
			TourPackageID: 1,
			MealTypeID:    1, // อาหารเช้า
		},
		{
			MenusDetail:   "ข้าวมันไก่, ต้มยำกุ้ง, น้ำสมุนไพร",
			TourPackageID: 1,
			MealTypeID:    2, // อาหารกลางวัน
		},
		{
			MenusDetail:   "ปลาทอดสมุนไพร, ผัดผักรวมมิตร, ข้าวหอมมะลิ, ของหวาน",
			TourPackageID: 1,
			MealTypeID:    3, // อาหารเย็น
		},
		{
			MenusDetail:   "ข้าวต้มไก่, ไข่ลวก, น้ำส้มคั้น",
			TourPackageID: 2,
			MealTypeID:    1, // อาหารเช้า
		},
		{
			MenusDetail:   "ส้มตำไก่ย่าง, ข้าวเหนียว, ขนมไทย",
			TourPackageID: 2,
			MealTypeID:    2, // อาหารกลางวัน
		},
		{
			MenusDetail:   "แกงเขียวหวานไก่, ผัดไทยกุ้งสด, ผลไม้",
			TourPackageID: 2,
			MealTypeID:    3, // อาหารเย็น
		},
	}
	for _, meal := range meals {
		db.FirstOrCreate(meal, &entity.Meals{
			MenusDetail:   meal.MenusDetail,
			TourPackageID: meal.TourPackageID,
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

	// Create Province
	provinces := []string{
		"กระบี่", "กรุงเทพมหานคร", "กาญจนบุรี", "กาฬสินธุ์", "กำแพงเพชร", "ขอนแก่น", "จันทบุรี", "ฉะเชิงเทรา",
		"ชลบุรี", "ชัยนาท", "ชัยภูมิ", "ชุมพร", "เชียงราย", "เชียงใหม่", "ตรัง", "ตราด", "ตาก", "นครนายก",
		"นครปฐม", "นครพนม", "นครราชสีมา", "นครศรีธรรมราช", "นครสวรรค์", "นนทบุรี", "นราธิวาส", "น่าน",
		"บึงกาฬ", "บุรีรัมย์", "ปทุมธานี", "ประจวบคีรีขันธ์", "ปราจีนบุรี", "ปัตตานี", "พระนครศรีอยุธยา",
		"พะเยา", "พังงา", "พัทลุง", "พิจิตร", "พิษณุโลก", "เพชรบุรี", "เพชรบูรณ์", "แพร่", "ภูเก็ต",
		"มหาสารคาม", "มุกดาหาร", "แม่ฮ่องสอน", "ยโสธร", "ยะลา", "ร้อยเอ็ด", "ระนอง", "ระยอง", "ราชบุรี",
		"ลพบุรี", "ลำปาง", "ลำพูน", "เลย", "ศรีสะเกษ", "สกลนคร", "สงขลา", "สตูล", "สมุทรปราการ",
		"สมุทรสงคราม", "สมุทรสาคร", "สระแก้ว", "สระบุรี", "สิงห์บุรี", "สุโขทัย", "สุพรรณบุรี", "สุราษฎร์ธานี",
		"สุรินทร์", "หนองคาย", "หนองบัวลำภู", "อ่างทอง", "อำนาจเจริญ", "อุดรธานี", "อุตรดิตถ์", "อุทัยธานี",
		"อุบลราชธานี",
	}
	for _, province := range provinces {
		db.FirstOrCreate(&entity.Provinces{}, entity.Provinces{ProvinceName: province})
	}

	// Create Location
	locations := []*entity.Locations{
		{
			LocationName: "วัดเกาะพยาม",
			ProvinceID:   49,
		},
		{
			LocationName: "เกาะค้างคาว",
			ProvinceID:   49,
		},
		{
			LocationName: "หาดบางเบน",
			ProvinceID:   49,
		},
		{
			LocationName: "สนามบิน",
			ProvinceID:   49,
		},
		{
			LocationName: "ท่าเรือ",
			ProvinceID:   49,
		},
		{
			LocationName: "ที่พัก",
			ProvinceID:   49,
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
		ProfilePath: "images/profileEmployees/employee1.jpg/",
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
		ProfilePath: "images/profileCustomers/customer1.jpg",
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
			ProvinceID:  47, // ระนอง
		},
		{
			PackageCode: "T00002",
			TourName:    "แพ็กเกจทัวร์ทะเลกระบี่: สวรรค์แห่งอันดามันที่ต้องสัมผัส",
			Duration:    "4 วัน 3 คืน",
			ProvinceID:  1, // กระบี่
		},
		{
			PackageCode: "T00003",
			TourName:    "แพ็กเกจทัวร์ภูเก็ต: เกาะสวรรค์อันดามัน",
			Duration:    "3 วัน 2 คืน",
			ProvinceID:  42, // ภูเก็ต
		},
		{
			PackageCode: "T00004",
			TourName:    "แพ็กเกจทัวร์เชียงใหม่: ธรรมชาติและวัฒนธรรม",
			Duration:    "5 วัน 4 คืน",
			ProvinceID:  13, // เชียงใหม่
		},
		{
			PackageCode: "T00005",
			TourName:    "แพ็กเกจทัวร์กรุงเทพมหานคร: สัมผัสเมืองหลวง",
			Duration:    "1 วัน",
			ProvinceID:  2, // กรุงเทพมหานคร
		},
		{
			PackageCode: "T00006",
			TourName:    "แพ็กเกจทัวร์นครศรีธรรมราช: แดนธรรมะแห่งใต้",
			Duration:    "3 วัน 2 คืน",
			ProvinceID:  21, // นครศรีธรรมราช
		},
		{
			PackageCode: "T00007",
			TourName:    "แพ็กเกจทัวร์แม่ฮ่องสอน: ผจญภัยเหนือขุนเขา",
			Duration:    "4 วัน 3 คืน",
			ProvinceID:  44, // แม่ฮ่องสอน
		},
		{
			PackageCode: "T00008",
			TourName:    "แพ็กเกจทัวร์ชลบุรี: ทะเลใกล้กรุง",
			Duration:    "2 วัน 1 คืน",
			ProvinceID:  9, // ชลบุรี
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
			Intro:           "สัมผัสความเงียบสงบและธรรมชาติที่บริสุทธิ์ของทะเลระนอง เปิดโลกการท่องเที่ยวสุดเอ็กซ์คลูซีฟกับแพ็กเกจทัวร์ทะเลระนอง ดินแดนที่ยังคงความงดงามดั้งเดิมของธรรมชาติ ทะเลใส หาดทรายขาว และหมู่เกาะที่ซ่อนตัวอยู่ในความสงบ เหมาะสำหรับผู้ที่ต้องการพักผ่อนและหลีกหนีจากความวุ่นวายของชีวิตประจำวัน",
			PackageDetail:   "จังหวัด: ระนอง\nระยะเวลา: 2 วัน 2 คืน\nที่พัก: รีสอร์ทหรูระดับ 4 ดาว พร้อมวิวทะเลส่วนตัว\nรวมอาหาร: อาหารครบทุกมื้อ เน้นซีฟู้ดสดใหม่จากทะเลระนอง\nการเดินทาง: รถรับ-ส่งสนามบินระนอง พร้อมเรือสปีดโบ๊ทส่วนตัวสำหรับทัวร์เกาะ",
			TripHighlight:   "🏝️ ทัวร์หมู่เกาะกำและเกาะค้างคาว สัมผัสหาดทรายขาวละเอียดและน้ำทะเลใสราวคริสตัลที่หมู่เกาะกำ เพลิดเพลินกับการดำน้ำตื้น(Snorkeling) เพื่อชมปะการังหลากสีสัน และปลาสวยงามนานาชนิด\n🌅 ชิลล์บนเกาะพยาม สัมผัสกับวิถีชีวิตเรียบง่ายบนเกาะพยาม เกาะที่ขึ้นชื่อว่า \"มัลดีฟส์เมืองไทย\" ชมความงดงามของชายหาดที่เงียบสงบและดื่มด่ำกับพระอาทิตย์ตกดินสุดโรแมนติก\n🛶 ล่องเรือชมป่าชายเลนและน้ำตกปุญญบาล ตื่นตาตื่นใจกับความสมบูรณ์ของป่าชายเลนที่อุดมสมบูรณ์ สำรวจน้ำตกปุญญบาลที่มีน้ำใสไหลเย็นตลอดปี\n✨ บ่อน้ำพุร้อนรักษะวาริน ผ่อนคลายกับการแช่น้ำแร่ร้อนธรรมชาติที่บ่อน้ำพุร้อนรักษะวาริน สถานที่ท่องเที่ยวเพื่อสุขภาพที่ไม่ควรพลาด",
			PlacesHighlight: "🌊 ดำน้ำดูปะการังและปลาสวยงาม: สัมผัสประสบการณ์ดำน้ำที่น่าตื่นเต้นในน้ำทะเลใสแจ๋ว\n🌴 พักผ่อนท่ามกลางธรรมชาติ: ที่พักสุดหรูที่โอบล้อมด้วยธรรมชาติที่งดงาม\n🍹 ปาร์ตี้ริมชายหาด: เพลิดเพลินกับปาร์ตี้บาร์บีคิวซีฟู้ดและเครื่องดื่มริมทะเลยามค่ำคืน",
			TourPackageID:   1,
		},
		{
			Intro:           "สัมผัสความงามทะเลอันดามัน ดินแดนสวรรค์ของนักท่องเที่ยว ค้นพบเสน่ห์ของทะเลกระบี่ด้วยแพ็กเกจทัวร์สุดพิเศษ ที่จะพาคุณไปดื่มด่ำกับธรรมชาติที่สวยงาม น้ำทะเลใส หาดทรายขาวละเอียด และเกาะแก่งน้อยใหญ่ที่รอคอยการมาเยือนของคุณ กระบี่พร้อมมอบประสบการณ์การพักผ่อนที่จะทำให้คุณลืมความวุ่นวายในชีวิตประจำวัน",
			PackageDetail:   "จังหวัด: กระบี่\nระยะเวลา: 4 วัน 3 คืน\nที่พัก: รีสอร์ทระดับ 4-5 ดาว ริมชายหาดพร้อมวิวทะเลทุกห้องพัก\nรวมอาหาร: อาหารเช้า, กลางวัน และมื้อค่ำแบบซีฟู้ดสดใหม่ทุกวัน\nการเดินทาง: รถรับ-ส่งจากสนามบินกระบี่ พร้อมเรือสปีดโบ๊ทส่วนตัวสำหรับทัวร์เกาะ",
			TripHighlight:   "🏝️ ทัวร์ 4 เกาะสุดฮิต เยือนเกาะปอดะ, ทะเลแหวก, เกาะไก่ และเกาะทับ สถานที่ท่องเที่ยวที่ขึ้นชื่อที่สุดในกระบี่ เพลิดเพลินกับการดำน้ำชมปะการังและฝูงปลาหลากสีสัน\n🌅 ล่องเรือหางยาวชมพระอาทิตย์ตก ดื่มด่ำกับความงามของพระอาทิตย์ตกดินที่หาดอ่าวนาง พร้อมบรรยากาศสุดโรแมนติก\n🛶 พายเรือคายัคที่อ่าวท่าเลน สำรวจป่าโกงกางและถ้ำหินปูนที่ซ่อนตัวอยู่ในป่าชายเลน ด้วยการพายเรือคายัคที่สนุกและท้าทาย\n✨ เที่ยวหมู่เกาะพีพี ดื่มด่ำกับความสวยงามของเกาะพีพีเล ชมอ่าวมาหยา สถานที่ถ่ายทำภาพยนตร์ชื่อดัง “The Beach” และอ่าวปิเละที่มีน้ำทะเลใสเหมือนกระจก",
			PlacesHighlight: "🌊 ดำน้ำตื้น (Snorkeling) ที่อ่าวลิง: สัมผัสชีวิตใต้ทะเลกับปลาหลากสีสันและปะการังที่สมบูรณ์\n🌴 ชิลล์ริมชายหาดที่ไร่เลย์: พักผ่อนบนหาดทรายขาวละเอียด ท่ามกลางหน้าผาหินปูนสูงตระหง่าน\n💆 สปาผ่อนคลาย: ผ่อนคลายกับการนวดแผนไทยที่สปาริมชายหาด",
			TourPackageID:   2,
		},
		{
			Intro:           "สัมผัสกับมนต์เสน่ห์ของเกาะภูเก็ต ที่เต็มไปด้วยความสนุกสนานและผ่อนคลายจากชายหาดสุดหรู พร้อมกับการเที่ยวชมสถานที่ท่องเที่ยวที่มีชื่อเสียงระดับโลก.",
			PackageDetail:   "จังหวัด: ภูเก็ต\nระยะเวลา: 3 วัน 2 คืน\nที่พัก: รีสอร์ทหรูระดับ 5 ดาว ใกล้หาดป่าตอง\nรวมอาหาร: อาหารครบทุกมื้อ รวมเมนูอาหารทะเลสดๆ\nการเดินทาง: รถรับ-ส่งจากสนามบินภูเก็ต และทริปเที่ยวเกาะต่างๆ",
			TripHighlight:   "🏖️ ชมวิวที่หาดป่าตองและหาดกะตะที่มีชื่อเสียง\n🌴 ท่องเที่ยวเกาะไข่และเกาะราชา เพลิดเพลินกับกิจกรรมดำน้ำตื้นและตกปลา\n🌅 ชมพระอาทิตย์ตกดินที่หาดกะรน\n🍽️ ลิ้มรสอาหารทะเลสดใหม่ที่ร้านอาหารริมทะเล",
			PlacesHighlight: "🌊 ดำน้ำดูปะการัง: สัมผัสกับชีวิตใต้ทะเลที่มีสีสันสวยงาม\n🌴 ท่องเที่ยวเกาะ: เยี่ยมชมเกาะต่างๆ ด้วยเรือเร็ว\n🍹 ชมพระอาทิตย์ตกดิน: เพลิดเพลินกับวิวทะเลที่สวยงาม",
			TourPackageID:   3,
		},
		{
			Intro:           "เดินทางสู่เชียงใหม่ เมืองแห่งธรรมชาติและวัฒนธรรมที่น่าค้นหา ทัวร์นี้จะพาคุณสัมผัสกับภูเขา วิวสวยๆ และวัดโบราณที่เก่าแก่",
			PackageDetail:   "จังหวัด: เชียงใหม่\nระยะเวลา: 5 วัน 4 คืน\nที่พัก: โรงแรมหรูระดับ 4 ดาว\nรวมอาหาร: อาหารครบทุกมื้อ รวมอาหารพื้นเมืองเชียงใหม่\nการเดินทาง: รถรับ-ส่งจากสนามบินเชียงใหม่ และทริปท่องเที่ยวในเมือง",
			TripHighlight:   "🌄 เดินป่าที่ดอยอินทนนท์ ชมวิวสวยๆ ของยอดเขาที่สูงที่สุดในประเทศไทย\n⛩️ เที่ยววัดพระธาตุดอยสุเทพ วัดที่มีความสำคัญทางศาสนาและวัฒนธรรม\n🌺 ชมดอกไม้เมืองหนาวที่สวนดอกไม้ในช่วงฤดูหนาว",
			PlacesHighlight: "🌲 เดินป่า: ผจญภัยในธรรมชาติของดอยอินทนนท์\n⛩️ วัดพระธาตุดอยสุเทพ: ชมวัดที่มีความสำคัญทางศาสนา\n🏵️ สวนดอกไม้: ชมความสวยงามของดอกไม้ในช่วงฤดูหนาว",
			TourPackageID:   4,
		},
		{
			Intro:           "ทัวร์สุดหรูในเมืองหลวงกรุงเทพมหานคร พบกับความหลากหลายของสถานที่ท่องเที่ยวทั้งที่เป็นศิลปวัฒนธรรม และการช้อปปิ้งในห้างสรรพสินค้าชั้นนำ แนะนำกิจกรรมที่เหมาะสำหรับการท่องเที่ยวแบบวันเดียวที่คุณไม่ควรพลาด",
			PackageDetail:   "จังหวัด: กรุงเทพมหานคร\nระยะเวลา: 1 วัน\nที่พัก: ไม่มีที่พัก (ทัวร์วันเดียว)\nรวมอาหาร: อาหารกลางวันและเย็น เน้นอาหารไทย\nการเดินทาง: รถรับ-ส่งจากโรงแรมในกรุงเทพฯ พร้อมไกด์ทัวร์",
			TripHighlight:   "🏯 เยี่ยมชมวัดพระแก้วและพระบรมมหาราชวัง สถานที่สำคัญทางประวัติศาสตร์และศาสนา\n🛳️ ล่องเรือแม่น้ำเจ้าพระยา ชมทัศนียภาพริมฝั่งแม่น้ำ\n🏙️ ช้อปปิ้งที่ห้างสรรพสินค้าหรู สยามพารากอนและเอ็มบาสซี",
			PlacesHighlight: "🌇 วัดพระแก้ว: สัมผัสความงดงามของพระอุโบสถและพระบรมมหาราชวัง\n🛥️ ล่องเรือเจ้าพระยา: ชมวิวริมแม่น้ำในบรรยากาศโรแมนติก\n🛍️ ช้อปปิ้งในกรุงเทพ: ช้อปปิ้งสินค้าหรูที่ห้างสรรพสินค้าชั้นนำ",
			TourPackageID:   5,
		},
		{
			Intro:           "สัมผัสความสงบและเสน่ห์ของนครศรีธรรมราช เมืองที่เต็มไปด้วยธรรมชาติและประวัติศาสตร์ ทัวร์นี้จะพาคุณสำรวจสถานที่ท่องเที่ยวทั้งในและนอกเมือง เช่น วัดพระมหาธาตุ และน้ำตกที่สวยงาม",
			PackageDetail:   "จังหวัด: นครศรีธรรมราช\nระยะเวลา: 3 วัน 2 คืน\nที่พัก: รีสอร์ทท่ามกลางธรรมชาติ\nรวมอาหาร: อาหารท้องถิ่นและซีฟู้ด\nการเดินทาง: รถรับ-ส่งสนามบินนครศรีธรรมราช พร้อมทัวร์ในพื้นที่",
			TripHighlight:   "🏯 เยี่ยมชมวัดพระมหาธาตุเมืองนครศรีธรรมราช\n🏞️ เดินทางไปยังน้ำตกหินตก น้ำตกที่สวยงามและเหมาะสำหรับการพักผ่อน\n🚶‍♂️ ท่องเที่ยวสวนสมเด็จพระนเรศวรและป่าสนเขา",
			PlacesHighlight: "🌳 น้ำตกหินตก: เพลิดเพลินกับการเที่ยวชมธรรมชาติในน้ำตกที่สวยงาม\n🍲 อาหารพื้นเมือง: ลิ้มรสอาหารท้องถิ่น เช่น แกงไก่คั่วกลิ้ง\n⛰️ ท่องเที่ยวภูเขา: ผจญภัยในป่าสนเขาของนครศรีธรรมราช",
			TourPackageID:   6,
		},
		{
			Intro:           "ดินแดนที่ตั้งอยู่ท่ามกลางขุนเขา แม่ฮ่องสอนคือสถานที่ท่องเที่ยวที่ผสมผสานธรรมชาติอันสวยงามและวิถีชีวิตที่เงียบสงบ ทัวร์นี้พาคุณไปสัมผัสประสบการณ์การท่องเที่ยวในพื้นที่ที่เต็มไปด้วยวัฒนธรรมและธรรมชาติ",
			PackageDetail:   "จังหวัด: แม่ฮ่องสอน\nระยะเวลา: 4 วัน 3 คืน\nที่พัก: รีสอร์ทกลางธรรมชาติ\nรวมอาหาร: อาหารพื้นเมืองและซีฟู้ด\nการเดินทาง: รถรับ-ส่งสนามบินแม่ฮ่องสอน พร้อมทริปท่องเที่ยว",
			TripHighlight:   "🌄 เดินทางไปยังหมู่บ้านชาวเขาและชมวิถีชีวิตท้องถิ่น\n🏞️ สัมผัสความสวยงามของทะเลหมอกที่ดอยแม่สะลอง\n🚶‍♂️ เดินป่าที่อุทยานแห่งชาติแม่ฮ่องสอน",
			PlacesHighlight: "🏞️ ทะเลหมอกดอยแม่สะลอง: ชมทะเลหมอกสวยงามในตอนเช้า\n🌄 หมู่บ้านชาวเขา: เรียนรู้วิถีชีวิตของชนเผ่าพื้นเมือง\n🍽️ อาหารพื้นเมือง: ลิ้มรสอาหารท้องถิ่นจากจังหวัดแม่ฮ่องสอน",
			TourPackageID:   7,
		},
		{
			Intro:           "สัมผัสบรรยากาศทะเลใกล้กรุงเทพฯ กับแพ็กเกจทัวร์ชลบุรี ที่จะพาคุณไปสัมผัสความสวยงามของทะเล หาดทรายขาว และกิจกรรมต่างๆ ที่เหมาะสำหรับการพักผ่อนสุดสัปดาห์",
			PackageDetail:   "จังหวัด: ชลบุรี\nระยะเวลา: 2 วัน 1 คืน\nที่พัก: รีสอร์ทติดทะเล\nรวมอาหาร: อาหารท้องถิ่นและซีฟู้ด\nการเดินทาง: รถรับ-ส่งจากกรุงเทพฯ พร้อมทริปท่องเที่ยว",
			TripHighlight:   "🏖️ ชมชายหาดพัทยาและเกาะล้าน สัมผัสกับกิจกรรมทางน้ำต่างๆ เช่น ดำน้ำ, บานาน่าโบ๊ท\n🍤 ลิ้มรสอาหารทะเลสดใหม่ที่พัทยา\n🌅 ชมพระอาทิตย์ตกดินที่เกาะล้าน",
			PlacesHighlight: "🌊 ดำน้ำดูปะการัง: เพลิดเพลินกับกิจกรรมดำน้ำและชมปะการัง\n🍽️ อาหารทะเล: ลิ้มรสอาหารทะเลสดที่ร้านอาหารริมทะเล\n🏝️ เกาะล้าน: เที่ยวเกาะที่มีทะเลสวยงามและกิจกรรมต่างๆ",
			TourPackageID:   8,
		},
	}
	for _, tourDescription := range tourDescriptions {
		db.FirstOrCreate(tourDescription, &entity.TourDescriptions{
			TourPackageID: tourDescription.TourPackageID,
		})
	}

	// Create Tour Schedule
	StartDate1, _ := time.Parse("2006-01-02", "2025-01-10")
	StartDate2, _ := time.Parse("2006-01-02", "2025-01-15")
	StartDate3, _ := time.Parse("2006-01-02", "2025-01-20")
	StartDate4, _ := time.Parse("2006-01-02", "2025-01-25")
	StartDate5, _ := time.Parse("2006-01-02", "2025-02-01")
	StartDate6, _ := time.Parse("2006-01-02", "2025-02-05")
	StartDate7, _ := time.Parse("2006-01-02", "2025-02-10")
	StartDate8, _ := time.Parse("2006-01-02", "2025-02-15")
	StartDate9, _ := time.Parse("2006-01-02", "2025-01-12")
	StartDate10, _ := time.Parse("2006-01-02", "2025-01-17")
	StartDate11, _ := time.Parse("2006-01-02", "2025-01-22")
	StartDate12, _ := time.Parse("2006-01-02", "2025-01-27")
	StartDate13, _ := time.Parse("2006-01-02", "2025-02-01")
	StartDate14, _ := time.Parse("2006-01-02", "2025-02-06")
	StartDate15, _ := time.Parse("2006-01-02", "2025-02-11")
	StartDate16, _ := time.Parse("2006-01-02", "2025-02-16")
	StartDate17, _ := time.Parse("2006-01-02", "2025-01-05")
	StartDate18, _ := time.Parse("2006-01-02", "2025-01-10")
	StartDate19, _ := time.Parse("2006-01-02", "2025-01-15")
	StartDate20, _ := time.Parse("2006-01-02", "2025-01-20")
	StartDate21, _ := time.Parse("2006-01-02", "2025-01-25")
	StartDate22, _ := time.Parse("2006-01-02", "2025-02-01")
	StartDate23, _ := time.Parse("2006-01-02", "2025-02-05")
	StartDate24, _ := time.Parse("2006-01-02", "2025-02-10")
	tourSchedules := []*entity.TourSchedules{
		// Tour Package 1 Schedules (January and February) - Duration 2 days 2 nights
	{
		StartDate:            StartDate1,
		EndDate:              StartDate1.AddDate(0, 0, 2),
		AvailableSlots:       50,
		TourPackageID:        1,
		TourScheduleStatusID: 2,
	},
	{
		StartDate:            StartDate2,
		EndDate:              StartDate2.AddDate(0, 0, 2),
		AvailableSlots:       50,
		TourPackageID:        1,
		TourScheduleStatusID: 2,
	},
	{
		StartDate:            StartDate3,
		EndDate:              StartDate3.AddDate(0, 0, 2),
		AvailableSlots:       50,
		TourPackageID:        1,
		TourScheduleStatusID: 2,
	},
	// Tour Package 2 Schedules (January and February) - Duration 4 days 3 nights
	{
		StartDate:            StartDate4,
		EndDate:              StartDate4.AddDate(0, 0, 4),
		AvailableSlots:       50,
		TourPackageID:        2,
		TourScheduleStatusID: 2,
	},
	{
		StartDate:            StartDate5,
		EndDate:              StartDate5.AddDate(0, 0, 4),
		AvailableSlots:       50,
		TourPackageID:        2,
		TourScheduleStatusID: 2,
	},
	{
		StartDate:            StartDate6,
		EndDate:              StartDate6.AddDate(0, 0, 4),
		AvailableSlots:       50,
		TourPackageID:        2,
		TourScheduleStatusID: 2,
	},
	// Tour Package 3 Schedules (January and February) - Duration 3 days 2 nights
	{
		StartDate:            StartDate7,
		EndDate:              StartDate7.AddDate(0, 0, 3),
		AvailableSlots:       50,
		TourPackageID:        3,
		TourScheduleStatusID: 2,
	},
	{
		StartDate:            StartDate8,
		EndDate:              StartDate8.AddDate(0, 0, 3),
		AvailableSlots:       50,
		TourPackageID:        3,
		TourScheduleStatusID: 2,
	},
	{
		StartDate:            StartDate9,
		EndDate:              StartDate9.AddDate(0, 0, 3),
		AvailableSlots:       50,
		TourPackageID:        3,
		TourScheduleStatusID: 2,
	},
	// Tour Package 4 Schedules (January and February) - Duration 5 days 4 nights
	{
		StartDate:            StartDate10,
		EndDate:              StartDate10.AddDate(0, 0, 5),
		AvailableSlots:       50,
		TourPackageID:        4,
		TourScheduleStatusID: 2,
	},
	{
		StartDate:            StartDate11,
		EndDate:              StartDate11.AddDate(0, 0, 5),
		AvailableSlots:       50,
		TourPackageID:        4,
		TourScheduleStatusID: 2,
	},
	{
		StartDate:            StartDate12,
		EndDate:              StartDate12.AddDate(0, 0, 5),
		AvailableSlots:       50,
		TourPackageID:        4,
		TourScheduleStatusID: 2,
	},
	// Tour Package 5 Schedules (January and February) - Duration 1 day
	{
		StartDate:            StartDate13,
		EndDate:              StartDate13.AddDate(0, 0, 1),
		AvailableSlots:       50,
		TourPackageID:        5,
		TourScheduleStatusID: 2,
	},
	{
		StartDate:            StartDate14,
		EndDate:              StartDate14.AddDate(0, 0, 1),
		AvailableSlots:       50,
		TourPackageID:        5,
		TourScheduleStatusID: 2,
	},
	{
		StartDate:            StartDate15,
		EndDate:              StartDate15.AddDate(0, 0, 1),
		AvailableSlots:       50,
		TourPackageID:        5,
		TourScheduleStatusID: 2,
	},
	// Tour Package 6 Schedules (January and February) - Duration 3 days 2 nights
	{
		StartDate:            StartDate16,
		EndDate:              StartDate16.AddDate(0, 0, 3),
		AvailableSlots:       50,
		TourPackageID:        6,
		TourScheduleStatusID: 2,
	},
	{
		StartDate:            StartDate17,
		EndDate:              StartDate17.AddDate(0, 0, 3),
		AvailableSlots:       50,
		TourPackageID:        6,
		TourScheduleStatusID: 2,
	},
	{
		StartDate:            StartDate18,
		EndDate:              StartDate18.AddDate(0, 0, 3),
		AvailableSlots:       50,
		TourPackageID:        6,
		TourScheduleStatusID: 2,
	},
	// Tour Package 7 Schedules (January and February) - Duration 4 days 3 nights
	{
		StartDate:            StartDate19,
		EndDate:              StartDate19.AddDate(0, 0, 4),
		AvailableSlots:       50,
		TourPackageID:        7,
		TourScheduleStatusID: 2,
	},
	{
		StartDate:            StartDate20,
		EndDate:              StartDate20.AddDate(0, 0, 4),
		AvailableSlots:       50,
		TourPackageID:        7,
		TourScheduleStatusID: 2,
	},
	{
		StartDate:            StartDate21,
		EndDate:              StartDate21.AddDate(0, 0, 4),
		AvailableSlots:       50,
		TourPackageID:        7,
		TourScheduleStatusID: 2,
	},
	// Tour Package 8 Schedules (January and February) - Duration 2 days 1 nights
	{
		StartDate:            StartDate22,
		EndDate:              StartDate22.AddDate(0, 0, 2),
		AvailableSlots:       50,
		TourPackageID:        8,
		TourScheduleStatusID: 2,
	},
	{
		StartDate:            StartDate23,
		EndDate:              StartDate23.AddDate(0, 0, 2),
		AvailableSlots:       50,
		TourPackageID:        8,
		TourScheduleStatusID: 2,
	},
	{
		StartDate:            StartDate24,
		EndDate:              StartDate24.AddDate(0, 0, 2),
		AvailableSlots:       50,
		TourPackageID:        8,
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
	for i := uint(1); i <= 8; i++ {
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
		// Tour Package 1 Prices
		{
			Price:         500,
			TourPackageID: 1,
			PersonTypeID:  1, // เด็กเล็ก
			RoomTypeID:    4, // เพิ่มเตียง
		},
		{
			Price:         0,
			TourPackageID: 1,
			PersonTypeID:  1, // เด็กเล็ก
			RoomTypeID:    5, // ไม่เพิ่มเตียง
		},
		{
			Price:         8290,
			TourPackageID: 1,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    1, // พักเดี่ยว
		},
		{
			Price:         6990,
			TourPackageID: 1,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    2, // คู่
		},
		{
			Price:         6990,
			TourPackageID: 1,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    3, // พักสาม
		},
		// Tour Package 2 Prices
		{
			Price:         500,
			TourPackageID: 2,
			PersonTypeID:  1, // เด็กเล็ก
			RoomTypeID:    4, // เพิ่มเตียง
		},
		{
			Price:         0,
			TourPackageID: 2,
			PersonTypeID:  1, // เด็กเล็ก
			RoomTypeID:    5, // ไม่เพิ่มเตียง
		},
		{
			Price:         6590,
			TourPackageID: 2,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    1, // พักเดี่ยว
		},
		{
			Price:         5190,
			TourPackageID: 2,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    2, // คู่
		},
		{
			Price:         5190,
			TourPackageID: 2,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    3, // พักสาม
		},
		// Tour Package 3 Prices
		{
			Price:         600,
			TourPackageID: 3,
			PersonTypeID:  1, // เด็กเล็ก
			RoomTypeID:    4, // เพิ่มเตียง
		},
		{
			Price:         0,
			TourPackageID: 3,
			PersonTypeID:  1, // เด็กเล็ก
			RoomTypeID:    5, // ไม่เพิ่มเตียง
		},
		{
			Price:         7990,
			TourPackageID: 3,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    1, // พักเดี่ยว
		},
		{
			Price:         6790,
			TourPackageID: 3,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    2, // คู่
		},
		{
			Price:         6790,
			TourPackageID: 3,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    3, // พักสาม
		},
		// Tour Package 4 Prices
		{
			Price:         700,
			TourPackageID: 4,
			PersonTypeID:  1, // เด็กเล็ก
			RoomTypeID:    4, // เพิ่มเตียง
		},
		{
			Price:         0,
			TourPackageID: 4,
			PersonTypeID:  1, // เด็กเล็ก
			RoomTypeID:    5, // ไม่เพิ่มเตียง
		},
		{
			Price:         8990,
			TourPackageID: 4,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    1, // พักเดี่ยว
		},
		{
			Price:         7490,
			TourPackageID: 4,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    2, // คู่
		},
		{
			Price:         7490,
			TourPackageID: 4,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    3, // พักสาม
		},
		// Tour Package 5 Prices
		{
			Price:         400,
			TourPackageID: 5,
			PersonTypeID:  1, // เด็กเล็ก
			RoomTypeID:    4, // เพิ่มเตียง
		},
		{
			Price:         0,
			TourPackageID: 5,
			PersonTypeID:  1, // เด็กเล็ก
			RoomTypeID:    5, // ไม่เพิ่มเตียง
		},
		{
			Price:         5990,
			TourPackageID: 5,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    1, // พักเดี่ยว
		},
		{
			Price:         4990,
			TourPackageID: 5,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    2, // คู่
		},
		{
			Price:         4990,
			TourPackageID: 5,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    3, // พักสาม
		},
		// Tour Package 6 Prices
		{
			Price:         700,
			TourPackageID: 6,
			PersonTypeID:  1, // เด็กเล็ก
			RoomTypeID:    4, // เพิ่มเตียง
		},
		{
			Price:         0,
			TourPackageID: 6,
			PersonTypeID:  1, // เด็กเล็ก
			RoomTypeID:    5, // ไม่เพิ่มเตียง
		},
		{
			Price:         9490,
			TourPackageID: 6,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    1, // พักเดี่ยว
		},
		{
			Price:         7990,
			TourPackageID: 6,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    2, // คู่
		},
		{
			Price:         7990,
			TourPackageID: 6,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    3, // พักสาม
		},
		// Tour Package 7 Prices
		{
			Price:         550,
			TourPackageID: 7,
			PersonTypeID:  1, // เด็กเล็ก
			RoomTypeID:    4, // เพิ่มเตียง
		},
		{
			Price:         0,
			TourPackageID: 7,
			PersonTypeID:  1, // เด็กเล็ก
			RoomTypeID:    5, // ไม่เพิ่มเตียง
		},
		{
			Price:         7490,
			TourPackageID: 7,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    1, // พักเดี่ยว
		},
		{
			Price:         6490,
			TourPackageID: 7,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    2, // คู่
		},
		{
			Price:         6490,
			TourPackageID: 7,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    3, // พักสาม
		},
		// Tour Package 8 Prices
		{
			Price:         800,
			TourPackageID: 8,
			PersonTypeID:  1, // เด็กเล็ก
			RoomTypeID:    4, // เพิ่มเตียง
		},
		{
			Price:         0,
			TourPackageID: 8,
			PersonTypeID:  1, // เด็กเล็ก
			RoomTypeID:    5, // ไม่เพิ่มเตียง
		},
		{
			Price:         9990,
			TourPackageID: 8,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    1, // พักเดี่ยว
		},
		{
			Price:         8490,
			TourPackageID: 8,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    2, // คู่
		},
		{
			Price:         8490,
			TourPackageID: 8,
			PersonTypeID:  2, // เด็ก/ผู้ใหญ่
			RoomTypeID:    3, // พักสาม
		},
	}
	for _, tourPrice := range tourPrices {
		db.FirstOrCreate(tourPrice, &entity.TourPrices{
			TourPackageID: tourPrice.TourPackageID,
			PersonTypeID:  tourPrice.PersonTypeID,
			RoomTypeID:    tourPrice.RoomTypeID,
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
	activities := []*entity.Activities{
		{
			ActivityName: "เดินทางไปยังท่าเรือ",
			Description:  "เดินทางไปยังท่าเรือเพื่อขึ้นเรือสปีดโบ๊ท มุ่งหน้าสู่ เกาะพยาม (ใช้เวลาประมาณ 45 นาที) เพลิดเพลินกับบรรยากาศทะเลสวยงามและทิวทัศน์ระหว่างการเดินทาง",
			LocationID:   5,
		},
		{
			ActivityName: "ถึงจุดนัดรับ",
			Description:  "ทีมงานต้อนรับที่สนามบินระนอง (หรือจุดนัดพบ) พร้อมบริการรับส่งด้วยรถตู้ปรับอากาศ",
			LocationID:   4,
		},
		{
			ActivityName: "ถึงเกาะพยาม",
			Description:  "ถึงเกาะพยาม และเยี่ยมชมวัดเกาะพยาม เที่ยวชมสถานที่สำคัญต่าง ๆ เช่น อ่าวเขาควาย และ อ่าวใหญ่",
			LocationID:   1,
		},
		{
			ActivityName: "เช็คเอาท์จากที่พัก",
			Description:  "รับประทานอาหารเช้าที่รีสอร์ท และเช็คเอาท์จากที่พัก",
			LocationID:   6,
		},
	}
	for _, activity := range activities {
		db.FirstOrCreate(activity, &entity.Activities{
			ActivityName: activity.ActivityName,
			Description:  activity.Description,
			LocationID:   activity.LocationID,
		})
	}

	// Create ScheduleActivities
	DateTime1, _ := time.Parse("2006-01-02 15:04:05", "2025-01-10 08:00:00")
	DateTime2, _ := time.Parse("2006-01-02 15:04:05", "2025-01-10 09:00:00")
	DateTime3, _ := time.Parse("2006-01-02 15:04:05", "2025-01-10 10:30:00")
	DateTime4, _ := time.Parse("2006-01-02 15:04:05", "2025-01-11 07:30:00")

	DateTime5, _ := time.Parse("2006-01-02 15:04:05", "2025-01-23 08:00:00")
	DateTime6, _ := time.Parse("2006-01-02 15:04:05", "2025-01-23 10:30:00")
	DateTime7, _ := time.Parse("2006-01-02 15:04:05", "2025-01-24 07:30:00")
	scheduleAcs := []*entity.ScheduleActivities{
		{
			DateTime:       DateTime2,
			ActivityID:     1,
			TourScheduleID: 1,
		},
		{
			DateTime:       DateTime1,
			ActivityID:     2,
			TourScheduleID: 1,
		},

		{
			DateTime:       DateTime3,
			ActivityID:     3,
			TourScheduleID: 1,
		},
		{
			DateTime:       DateTime4,
			ActivityID:     4,
			TourScheduleID: 1,
		},
		{
			DateTime:       DateTime5,
			ActivityID:     2,
			TourScheduleID: 2,
		},

		{
			DateTime:       DateTime6,
			ActivityID:     3,
			TourScheduleID: 2,
		},
		{
			DateTime:       DateTime7,
			ActivityID:     4,
			TourScheduleID: 2,
		},
	}
	for _, scheAc := range scheduleAcs {
		db.FirstOrCreate(scheAc, &entity.ScheduleActivities{
			DateTime:       scheAc.DateTime,
			ActivityID:     scheAc.ActivityID,
			TourScheduleID: scheAc.TourScheduleID,
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

func createImageVehicle(filePath string, id uint) error {
	image := &entity.VehicleImages{FilePath: filePath, VehicleID: id}
	if err := db.Where("file_path = ?", &image.FilePath).FirstOrCreate(&image).Error; err != nil {
		return err
	}
	return nil
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
