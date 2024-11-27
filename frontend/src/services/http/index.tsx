import { BookingDetailsInterface } from "../../interfaces/IBookingDetails";
import { BookingsInterface } from "../../interfaces/IBookings";
import { SignInInterface } from "../../interfaces/ISignIn";
import { TourSchedulesInterface } from "../../interfaces/ITourSchedules";

export const apiUrl = "http://localhost:8000";

// SignInForCustomer
async function SignInForCustomer(data: SignInInterface) {
    const requestOptions = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
    };

    let res = await fetch(`${apiUrl}/signin-customer`, requestOptions).then((res) => {
        if (res.status == 200) {
            return res.json();
        } else {
            return false;
        }
    });

    return res;
}

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

// BookingDetails
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

// Bookings
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

// Customers
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

// PersonTypes
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

// Promotions
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

// Provinces
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

// RoomTypes
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

// TourSchedules
async function GetTourScheduleByID(id: Number | undefined) {
    const requestOptions = {
        method: "GET",
    };

    let res = await fetch(`${apiUrl}/tour-schedule/${id}`, requestOptions).then(
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
async function UpdateTourScheduleByID(data: TourSchedulesInterface, id: Number | undefined) {
    const requestOptions = {
        method: "PATCH",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
    };

    let res = await fetch(`${apiUrl}/tour-schedule/${id}`, requestOptions)
        .then((res) => {
            if (res.status == 200) {
                return res.json();
            } else {
                return false;
            }
        });

    return res;
}

// ScheduleActivities
async function GetScheduleActivityByTourScheduleID(id: Number | undefined) {
    const requestOptions = {
        method: "GET",
    };

    let res = await fetch(`${apiUrl}/schedule-activity/${id}`, requestOptions).then(
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
    // SignInForCustomer
    SignInForCustomer,

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

    // TourSchedules
    GetTourScheduleByID,
    UpdateTourScheduleByID,

    // ScheduleActivities
    GetScheduleActivityByTourScheduleID,
}