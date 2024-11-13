package config

import (
	"fmt"
	// "io/ioutil"
	// "time"
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
		&entity.Bookings{},
		&entity.BookingStatuses{},
		&entity.Customers{},
		&entity.Employees{},
		&entity.EmployeeSchedules{},
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
		&entity.TourAccommodations{},
		&entity.TourActivities{},
		&entity.TourImages{},
		&entity.TourPackages{},
		&entity.TourPrices{},
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

	// Create Payment Method
	paymentMethods := []*entity.PaymentMethods{
		{
			MethodName: "Krungthai",
		},
		{
			MethodName: "TrueMoney",
		},
		{
			MethodName: "SCB",
		},
	}
	for _, method := range paymentMethods {
		db.FirstOrCreate(method, &entity.PaymentMethods{
			MethodName: method.MethodName,
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

	// Create Vehicle Type
	vehicleTypes := []*entity.VehicleTypes{
		{
			TypeName: "แอดมิน",
		},
		{
			TypeName: "คนขับรถทัวร์",
		},
		{
			TypeName: "ไกด์",
		},
	}
	for _, vehicleType := range vehicleTypes {
		db.FirstOrCreate(vehicleType, &entity.VehicleTypes{
			TypeName: vehicleType.TypeName,
		})
	}

	// Create Employee
	hashedPassword, _ := HashPassword("123456")
	employee := &entity.Employees{
		UserName: "NamNam",
		FirstName:   "Poonchub",
		LastName:    "Nanawan",
		Email:       "poonchubnanawan310@gmail.com",
		Password:    hashedPassword,
		PhoneNumber: "000-000-0000",
		ProfilePath: "images/profile/owner/owner_id01.jpg",
		RoleID: 1,
	}
	db.FirstOrCreate(employee, &entity.Employees{
		Email: employee.Email,
	})

	// Create Customer
	customers := []*entity.Customers{
		{	
			UserName:   "Poonchub",
			FirstName:   "Poonchub",
			LastName:    "Nanawan",
			Email:       "poonchubnanawan320@gmail.com",
			Password:    hashedPassword,
			PhoneNumber: "098-594-4576",
			ProfilePath: "images/profile/customer/",
		},
		{	
			UserName:   "Peter",
			FirstName:   "Peter",
			LastName:    "Parker",
			Email:       "spider@gmail.com",
			Password:    hashedPassword,
			PhoneNumber: "098-594-4576",
			ProfilePath: "images/profile/customer/",
		},
	}
	for _, customer := range customers {
		db.FirstOrCreate(customer, &entity.Customers{
			Email: customer.Email,
		})
	}

	

	

	// Create Image 
	// for i := uint(1); i <= 14; i++ {
	// 	dir := fmt.Sprintf("images/product/product%d", i)
	// 	count := countFilesInDir(dir)
	// 	for j := 1; j <= count; j++ {
	// 		filePath := fmt.Sprintf("images/product/product%d/p0%d.jpg", i, j)
	// 		err := createImage(filePath, i)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	}
	// }
}


// func createImage(filePath string, id uint) error {

// 	image := entity.Image{FilePath: filePath, ProductID: id}

// 	if err := db.Where("file_path = ?", &image.FilePath).FirstOrCreate(&image).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func countFilesInDir(dir string) int {
// 	files, err := ioutil.ReadDir(dir)
// 	if err != nil {
// 		return 0
// 	}

// 	fileCount := 0
// 	for _, file := range files {
// 		if !file.IsDir() {
// 			fileCount++
// 		}
// 	}

// 	return fileCount
// }