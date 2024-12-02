import { useEffect, useState } from "react"
import Navbar from "../../components/navbar/Navbar"
import { apiUrl, GetBookingByCustomerID } from "../../services/http"
import "./Profile.css"
import { BookingsInterface } from "../../interfaces/IBookings"
import { Button, message, Modal, Steps } from "antd"

const { confirm } = Modal;

function Profile() {

    const [bookings, setBookings] = useState<BookingsInterface[]>()
    const storedCustomer = localStorage.getItem('customer')

    const customer = storedCustomer ? JSON.parse(storedCustomer) : ""

    const [statusIsClicked, setStatusIsClicked] = useState(false)
    const [elementClicked, setElementClicked] = useState<number>()

    const [btnIsClicked, setBtnIsClicked] = useState<number>(1)

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

    function handleElementClick(id: number | undefined) {
        if (!id) return
        if (elementClicked === id) {
            setStatusIsClicked((prev) => !prev)
        } else {
            setElementClicked(id)
            setStatusIsClicked(true)
        }
    }

    function handleBtnClick(index: number) {
        setBtnIsClicked(index)
    }

    const showConfirm = () => {
        confirm({
            title: "Log out?",
            content: "Do you want to log out?",
            okText: "confirm",
            cancelText: "cancel",
            centered: true,
            onOk() {
                localStorage.clear();
                setTimeout(() => {
                    location.href = "/login-customer";
                }, 2000);
                message.success('Log out successful');
            },
        });
    };

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
                    <div className="text-box">
                        <span className="tour-name">{booking.TourSchedule?.TourPackage?.TourName}</span>
                        <span className="tourID">{booking.TourSchedule?.TourPackage?.PackageCode}</span>
                    </div>
                    <div className="btn-box">
                        <button className="check-detail-btn check-btn" onClick={() => toPayment(booking?.ID)}>รายละเอียด</button>
                        <button className="check-status-btn check-btn" onClick={() => handleElementClick(booking.ID)}>สถานะการจอง</button>
                    </div>
                </div>
                {
                    elementClicked == booking?.ID && statusIsClicked ? (
                        <div className="status-box">
                            {
                                booking.BookingStatusID != 4 ? (
                                    <Steps current={
                                        booking.BookingStatusID == 1 ? 1 :
                                            booking.BookingStatusID == 2 ? 2 :
                                                booking.BookingStatusID == 2 ? 3 : 0
                                    }
                                        labelPlacement="vertical"
                                        items={items}
                                    />
                                ) : (
                                    <div className="step-error">
                                        <Steps status="error"
                                            labelPlacement="vertical"
                                            items={[{ title: 'ถูกยกเลิกแล้ว' }]}
                                        />
                                    </div>
                                )
                            }
                        </div>
                    ) : (<></>)
                }
            </div>
        )
    })

    useEffect(() => {
        const your_booking_container = document.querySelector(".your-booking-container")
        const payment_history_container = document.querySelector(".payment-history-container")
        const edit_profile_container = document.querySelector(".edit-profile-container")
        if (btnIsClicked == 1) {
            payment_history_container?.classList.remove("active")
            edit_profile_container?.classList.remove("active")
            your_booking_container?.classList.add("active")
        }
        else if (btnIsClicked == 2) {
            your_booking_container?.classList.remove("active")
            edit_profile_container?.classList.remove("active")
            payment_history_container?.classList.add("active")
        }
        else if (btnIsClicked == 3) {
            payment_history_container?.classList.remove("active")
            your_booking_container?.classList.remove("active")
            edit_profile_container?.classList.add("active")
        }
    }, [btnIsClicked])

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
                            <button
                                className={`your-booking-btn btn ${btnIsClicked == 1 ? 'active' : ''}`}
                                onClick={() => handleBtnClick(1)}
                            >Your Booking</button>
                            <button
                                className={`payment-history-btn btn ${btnIsClicked == 2 ? 'active' : ''}`}
                                onClick={() => handleBtnClick(2)}
                            >Payment History</button>
                            <button
                                className={`edit-profile-btn btn ${btnIsClicked == 3 ? 'active' : ''}`}
                                onClick={() => handleBtnClick(3)}
                            >Edit Profile</button>
                        </div>
                        <div className="bottom-box">
                            <Button className="logout-btn btn" type="primary" onClick={showConfirm}>Log Out</Button>
                        </div>
                    </div>
                </div>
                <div className="card-detail">
                    <div className="your-booking-container con">
                        <h5 className="title">Your Booking</h5>
                        {bookingElement}
                    </div>
                    <div className="payment-history-container con">
                        <h5 className="title">Payment History</h5>
                    </div>
                    <div className="edit-profile-container con">
                        <h5 className="title">Edit Profile</h5>
                    </div>
                </div>
            </section>
        </div>
    )
}
export default Profile