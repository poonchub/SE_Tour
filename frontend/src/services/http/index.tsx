import { BookingDetailsInterface } from "../../interfaces/IBookingDetails";
import { BookingsInterface } from "../../interfaces/IBookings";

export const apiUrl = "http://localhost:8000";

// async function SignInForCustomer(data: SignInInterface) {
//     const requestOptions = {
//         method: "POST",
//         headers: { "Content-Type": "application/json" },
//         body: JSON.stringify(data),
//     };

//     let res = await fetch(`${apiUrl}/signin-customer`, requestOptions).then((res) => {
//         if (res.status == 200) {
//             return res.json();
//         } else {
//             return false;
//         }
//     });

//     return res;
// }

// async function SignInForOwner(data: SignInInterface) {
//   const requestOptions = {
//       method: "POST",
//       headers: { "Content-Type": "application/json" },
//       body: JSON.stringify(data),
//   };

//   let res = await fetch(`${apiUrl}/signin-owner`, requestOptions).then((res) => {
//       if (res.status == 200) {
//           return res.json();
//       } else {
//           return false;
//       }
//   });

//   return res;
// }

async function GenerateQrCode(amount: number){
    const requestOptions = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({amount}),
    };

    let res = await fetch(`${apiUrl}/generate-qr`, requestOptions).then((res) => {
        if (res.status == 200) {
            return res.json();
        } else {
            return false;
        }
    });

    return res;
}

// BookingDetail
async function GetBookingDetails() {
    const requestOptions = {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/booking-details`, requestOptions)
        .then((res) => {
            if (res.status == 200) {
                return res.json();
            } else {
                return false;
            }
        });

    return res;
}
async function CreateBookingDetail(data: BookingDetailsInterface) {
    const requestOptions = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
    };

    let res = await fetch(`${apiUrl}/booking-detail`, requestOptions).then((res) => {
        if (res.status == 201) {
            return res.json();
        } else {
            return false;
        }
    });

    return res;
}


// Booking
async function GetBookings() {
    const requestOptions = {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/bookings`, requestOptions)
        .then((res) => {
            if (res.status == 200) {
                return res.json();
            } else {
                return false;
            }
        });

    return res;
}
async function GetBookingByID(id: Number) {
    const requestOptions = {
        method: "GET",
    };

    let res = await fetch(`${apiUrl}/booking/${id}`, requestOptions).then(
        (res) => {
            if (res.status == 200) {
                return res.json();
            } else {
                return false;
            }
        }
    );

    return res;
}
async function GetBookingByCustomerID(customerid: Number) {
    const requestOptions = {
        method: "GET",
    };

    let res = await fetch(`${apiUrl}/booking/${customerid}`, requestOptions).then(
        (res) => {
            if (res.status == 200) {
                return res.json();
            } else {
                return false;
            }
        }
    );

    return res;
}
async function CreateBooking(data: BookingsInterface) {
    const requestOptions = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
    };

    let res = await fetch(`${apiUrl}/booking`, requestOptions).then((res) => {
        if (res.status == 201) {
            return res.json();
        } else {
            return false;
        }
    });

    return res;
}

// Customer
async function GetCustomerByID(id: Number) {
    const requestOptions = {
        method: "GET",
    };

    let res = await fetch(`${apiUrl}/customer/${id}`, requestOptions).then(
        (res) => {
            if (res.status == 200) {
                return res.json();
            } else {
                return false;
            }
        }
    );

    return res;
}

// PersonType
async function GetPersonTypes() {
    const requestOptions = {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/person-types`, requestOptions)
        .then((res) => {
            if (res.status == 200) {
                return res.json();
            } else {
                return false;
            }
        });

    return res;
}

// Promotion
async function GetPromotionByCode(code: string | undefined) {
    const requestOptions = {
        method: "GET",
    };

    let res = await fetch(`${apiUrl}/promotion/${code}`, requestOptions).then(
        (res) => {
            if (res.status == 200) {
                return res.json();
            } else {
                return false;
            }
        }
    );

    return res;
}

// Province
async function GetProvinces() {
    const requestOptions = {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/provinces`, requestOptions)
        .then((res) => {
            if (res.status == 200) {
                return res.json();
            } else {
                return false;
            }
        });

    return res;
}

// RoomType
async function GetRoomTypes() {
    const requestOptions = {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/room-types`, requestOptions)
        .then((res) => {
            if (res.status == 200) {
                return res.json();
            } else {
                return false;
            }
        });

    return res;
}

// TourPackages
async function GetTourPackages() {
    const requestOptions = {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    };

    let res = await fetch(`${apiUrl}/tour-packages`, requestOptions)
        .then((res) => {
            if (res.status == 200) {
                return res.json();
            } else {
                return false;
            }
        });

    return res;
}
async function GetTourPackageByID(id: Number | undefined) {
    const requestOptions = {
        method: "GET",
    };

    let res = await fetch(`${apiUrl}/tour-package/${id}`, requestOptions).then(
        (res) => {
            if (res.status == 200) {
                return res.json();
            } else {
                return false;
            }
        }
    );

    return res;
}

// TourImages
async function GetTourImageByTourPackageID(id: Number | undefined) {
    const requestOptions = {
        method: "GET",
    };

    let res = await fetch(`${apiUrl}/tour-image/${id}`, requestOptions).then(
        (res) => {
            if (res.status == 200) {
                return res.json();
            } else {
                return false;
            }
        }
    );

    return res;
}

export {

    GenerateQrCode,

    // BookingDetails
    GetBookingDetails,
    CreateBookingDetail,
    
    // Bookings
    GetBookings,
    GetBookingByID,
    GetBookingByCustomerID,
    CreateBooking,

    // Customers
    GetCustomerByID,

    // PersonTypes
    GetPersonTypes,

    // Promotions
    GetPromotionByCode,

    // Province
    GetProvinces,

    // RoomTypes
    GetRoomTypes,

    // TourPackages
    GetTourPackages,
    GetTourPackageByID,

    // TourImages
    GetTourImageByTourPackageID,
}