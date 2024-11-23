import { useEffect, useState } from "react"
import Navbar from "../../components/navbar/Navbar"
import "./Payment.css"
import { BookingsInterface } from "../../interfaces/IBookings"
import { GetBookingByID } from "../../services/http"
import Loading from "../../components/loading/Loading"

import QRCode from 'react-qr-code';

import { Steps } from 'antd';
import generatePayload from 'promptpay-qr';


function Payment() {
    const bookingID = localStorage.getItem("booking-id")

    const [booking, setBooking] = useState<BookingsInterface>()

    const [isLoading, setIsLoading] = useState<boolean>(true);

    const [phoneNumber, setPhoneNumber] = useState("084-786-6591");
    const [amount, setAmount] = useState(10000.00);
    const [qrCode, setqrCode] = useState("sample");

    function handleQR() {
        setqrCode(generatePayload(phoneNumber, { amount }));
    }
    console.log(qrCode)

    async function getBookingByID() {
        const resBooking = await GetBookingByID(1);
        if (resBooking) {
            setBooking(resBooking)
        }
    }

    async function fetchData() {
        try {
            getBookingByID()
        } catch (error) {
            console.error('Failed to fetch data:', error);
        } finally {
            setIsLoading(false);
        }
    }

    useEffect(() => {
        fetchData()
        // generateQrCode()
        handleQR()
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


    return isLoading ? (
        <Loading />
    ) : (
        <div className="payment-page">
            <Navbar page="payment" />
            <section className="payment-box-card">
                <div className="booking-step-card card">
                    <Steps current={1} labelPlacement="vertical" items={items} />
                </div>
                <div className="sub-section-card">
                    <div className="booking-detail-card card">

                    </div>
                    <div className="qr-code-card card">
                        <QRCode value={qrCode} size={256} level="M" />
                    </div>
                </div>

            </section>
        </div>
    )
}
export default Payment