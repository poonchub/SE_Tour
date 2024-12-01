import { useEffect, useState } from "react"
import Navbar from "../../components/navbar/Navbar"
import { apiUrl, GetBookingByCustomerID } from "../../services/http"
import "./Profile.css"
import { BookingsInterface } from "../../interfaces/IBookings"
import { Steps } from "antd"

function Profile() {

    const [bookings, setBookings] = useState<BookingsInterface[]>()
    const storedCustomer = localStorage.getItem('customer')

    const customer = storedCustomer ? JSON.parse(storedCustomer) : ""

    const [statusIsClicked, setStatusIsClicked] = useState(false)
    const [elementClicked, setElementClicked] = useState<number>()

    async function getBookings() {
        const res = await GetBookingByCustomerID(customer.ID)
        if (res) {
            setBookings(res)
        }
    }

    function fetchData() {
        try {
            getBookings()
        } catch (error) {
            console.error('Failed to fetch data:', error);
        }
    }

    function toPayment(id: number | undefined) {
        localStorage.setItem("booking-id", String(id))
        setTimeout(() => {
            location.href = "/payment";
        });
    }

    function handleClick(id: number | undefined) {
        if (!id) return
        if (elementClicked === id) {
            setStatusIsClicked((prev) => !prev)
        } else {
            setElementClicked(id)
            setStatusIsClicked(true)
        }
    }

    useEffect(() => {
        fetchData()
    }, [])

    const items = [
        {
            title: 'จองแพ็กเกจ',
        },
        {
            title: 'ชำระเงิน',
        },
        {
            title: 'รอการตรวจสอบ',
        },
        {
            title: 'การจองเสร็จสิ้น',
        },
    ];

    const imageUrl = `${apiUrl}/${customer.ProfilePath}`
    const bookingElement = bookings?.map((booking, index) => {
        return (
            <div className="booking-box" key={index}>
                <div className="container">
                    <span className="tour-name">{booking.TourSchedule?.TourPackage?.TourName}</span>
                    <div className="btn-box">
                        <button className="check-detail-btn check-btn" onClick={() => toPayment(booking?.ID)}>รายละเอียด</button>
                        <button className="check-status-btn check-btn" onClick={() => handleClick(booking.ID)}>สถานะการจอง</button>
                    </div>
                </div>
                {
                    elementClicked==booking?.ID && statusIsClicked ? (
                        <div className="status-box">
                            {
                                booking.BookingStatusID!=4 ? (
                                    <Steps current={
                                        booking.BookingStatusID == 1 ? 1 :
                                            booking.BookingStatusID == 2 ? 2 :
                                                booking.BookingStatusID == 2 ? 3 : 0
                                    }
                                        labelPlacement="vertical"
                                        items={items}
                                    />
                                ) : (
                                    <Steps current={0}
                                        labelPlacement="vertical"
                                        items={[{title: 'ถูกยกเลิกแล้ว'}]}
                                    />
                                )
                            }
                            
                        </div>
                    ) : (<></>)
                }
            </div>
        )
    })
    return (
        <div className="profile-page">
            <Navbar page={"profile"} />
            <section>
                <div className="card-dashbord">
                    <div className="img-box">
                        <img src={imageUrl} alt="" />
                    </div>
                    <div className="profile-detail">
                        <span className="username">{customer.UserName}</span>
                        <span className="name">{`${customer.FirstName} ${customer.LastName}`}</span>
                        <span className="email">{customer.Email}</span>
                    </div>
                    <div className="nav-btn-box">
                        <div className="top-box">
                            <button className="your-booking-btn btn active">Your Booking</button>
                            <button className="payment-history-btn btn"> Payment History</button>
                            <button className="edit-profile-btn btn">Edit Profile</button>
                        </div>
                        <div className="bottom-box">
                            <button className="logout-btn btn">Log Out</button>
                        </div>
                    </div>
                </div>
                <div className="card-detail">
                    <div className="your-booking-container">
                        <h5 className="title">Your Booking</h5>
                        {bookingElement}
                    </div>
                    <div className="payment-history-container">

                    </div>
                    <div className="edit-profile-container">

                    </div>
                </div>
            </section>
        </div>
    )
}
export default Profile