import { Key, useState, SetStateAction, useEffect } from "react"
import "./Booking.css"
import { CustomersInterface } from "../../interfaces/ICustomers";
import { CreateBooking, GetCustomerByID, GetPromotionByCode } from "../../services/http";
import { PromotionsInterface } from "../../interfaces/IPromotions";
import { useDateContext } from "../../context/DateContext";
import { BookingsInterface } from "../../interfaces/IBookings";


function Booking(props: { roomTypes: any; tourPackage: any; personTypes: any; setPopUp: any; messageApi: any; }) {

    const { dateSelectedFormat, dateID } = useDateContext()

    const { roomTypes, tourPackage, personTypes, setPopUp, messageApi } = props

    const [customer, setCustomer] = useState<CustomersInterface>();
    const [promotion, setPromotion] = useState<PromotionsInterface>();

    const [childAdultSingleCount, setChildAdultSingleCount] = useState(0)
    const [childAdultDoubleCount, setChildAdultDoubleCount] = useState(0)
    const [childAdultThreeCount, setChildAdultThreeCount] = useState(0)
    const [infantAddBedCount, setInfantAddBedCount] = useState(0)
    const [infantNoAddBedCount, setNoInfantAddBedCount] = useState(0)
    const [totalPeople, setTotalPeople] = useState(0)

    const [childAdultSinglePrice, setChildAdultSinglePrice] = useState(0)
    const [childAdultDoublePrice, setChildAdultDoublePrice] = useState(0)
    const [childAdultThreePrice, setChildAdultThreePrice] = useState(0)
    const [infantAddBedPrice, setInfantAddBedPrice] = useState(0)
    const [infantNoAddBedPrice, setNoInfantAddBedPrice] = useState(0)

    const [fName, setFName] = useState<string | undefined>("");
    const [lName, setLName] = useState<string | undefined>("");
    const [phoneNumber, setPhoneNumber] = useState<string | undefined>("");
    const [email, setEmail] = useState<string | undefined>("");

    const [promotionCode, setPromotionCode] = useState<string | undefined>("")

    const [totalPrice, setTotalPrice] = useState<number>(0)

    const [isDisabled, setIsDisabled] = useState(true);

    async function getCustomerByID() {
        let res = await GetCustomerByID(1)
        if (res) {
            setCustomer(res);
        }
    }

    // async function getPromotionByCode() {
    //     let res = await GetPromotionByCode(promotionCode)
    //     if (res) {
    //         setCustomer(res);
    //     }
    // }

    async function fetchData() {
        try {
            getCustomerByID()
        } catch (error) {
            console.error('Failed to fetch data:', error);
        }
    }

    function handleChange(e: string, index: Key, p: number | undefined) {
        switch (index) {
            case 0:
                setChildAdultSingleCount(Number(e))
                setChildAdultSinglePrice(p ? p * Number(e) : 0)
                break;
            case 1:
                setChildAdultDoubleCount(Number(e))
                setChildAdultDoublePrice(p ? p * Number(e) : 0)
                break;
            case 2:
                setChildAdultThreeCount(Number(e))
                setChildAdultThreePrice(p ? p * Number(e) : 0)
                break;
            case 3:
                setInfantAddBedCount(Number(e))
                setInfantAddBedPrice(p ? p * Number(e) : 0)
                break;
            default:
                setNoInfantAddBedCount(Number(e))
                setNoInfantAddBedPrice(p ? p * Number(e) : 0)
                break;
        }
    }

    function handleCancle() {
        setIsDisabled(true)
        setFName(customer?.FirstName)
    }

    async function handleCreateBooking(){
        try{

            const booking: BookingsInterface = {
                TotalPrice: totalPrice,
                CustomerID: 1,
                TourScheduleID: dateID,
                // PromotionID: 
            }
            const res = await CreateBooking(booking)
            if (res){
                messageApi.open({
                    type: "success",
                    content: "ทำการจองแพ็กเกจทัวร์เรียบร้อยแล้ว",
                });
            }
            else {
                messageApi.open({
                    type: "error",
                    content: "เกิดข้อผิดพลาดในการจองแพ็กเกจทัวร์",
                });
            }
        } catch (error) {
            console.error("Error creating order:", error);
            messageApi.open({
                type: "error",
                content: "เกิดข้อผิดพลาดในการจองแพ็กเกจทัวร์",
            });
        }
        // setTimeout(() => {
        //     location.href = "/Payment";
        //     setIsButtonDisabled(false);
        // }, 1800);
    }

    useEffect(() => {
        fetchData()
    }, [])

    useEffect(() => {
        if (childAdultSingleCount == 0 && childAdultDoubleCount == 0 && childAdultDoubleCount == 0) {
            setInfantAddBedCount(0)
            setNoInfantAddBedCount(0)
            setInfantAddBedPrice(0)
            setNoInfantAddBedPrice(0)
        }
    }, [childAdultSingleCount, childAdultDoubleCount, childAdultDoubleCount])

    useEffect(() => {
        const count = childAdultSingleCount + childAdultDoubleCount + childAdultThreeCount + infantAddBedCount + infantNoAddBedCount
        setTotalPeople(count)
    }, [childAdultSingleCount, childAdultDoubleCount, childAdultThreeCount, infantAddBedCount, infantNoAddBedCount])

    useEffect(() => {
        const total = childAdultSinglePrice + childAdultDoublePrice + childAdultThreePrice + infantAddBedPrice + infantNoAddBedPrice
        setTotalPrice(total)
    }, [childAdultSinglePrice, childAdultDoublePrice, childAdultThreePrice, infantAddBedPrice, infantNoAddBedPrice])

    const priceElement = roomTypes?.map((type: any, index: Key) => {
        const tourPrices = tourPackage?.TourPrices
        var p: number | undefined
        var pfm: number | undefined
        tourPrices?.forEach((price: any, _: Key) => {
            if (price.RoomTypeID === type.ID && price.PersonTypeID) {
                pfm = price.Price?.toLocaleString('th-TH', {
                    minimumFractionDigits: 2,
                    maximumFractionDigits: 2,
                })
                p = price.Price
            }
        });

        return pfm ? (
            <div key={index}>
                {
                    index == 3 ? <span className="infant-title">เด็กเล็กพักกับผู้ใหญ่</span> : <></>
                }
                <div className="price-box">
                    <span className="type-name">{type.TypeName}</span>
                    <input type="number" value={
                        index == 0 ? childAdultSingleCount : (
                            index == 1 ? childAdultDoubleCount : (
                                index == 2 ? childAdultThreeCount : (
                                    index == 3 ? infantAddBedCount :
                                        infantNoAddBedCount
                                )
                            )
                        )
                    } onChange={(e) => (handleChange(e.target.value, index, p))}
                        min={0} />
                    <span className="price">{
                        index == 0 ? childAdultSingleCount : (
                            index == 1 ? childAdultDoubleCount : (
                                index == 2 ? childAdultThreeCount : (
                                    index == 3 ? infantAddBedCount :
                                        infantNoAddBedCount
                                )
                            )
                        )
                    } × {pfm}</span>
                </div>
            </div>
        ) : ""
    })

    return (
        <div className="booking-container">
            <div className="card">
                <span className="header">จองแพ็กเกจทัวร์</span>
                <section className="section-cover">
                    <div className="select-persontype-card sub-card">
                        <span className="title">โปรดระบุจำนวนผู้เดินทาง</span>
                        <div className="add-quantity-box">
                            <div className="title-box">
                                <span className="span1">เด็ก/ผู้ใหญ่</span>
                                <span className="span2">จำนวน (คน)</span>
                                <span className="span3">ราคา (บาท)</span>
                            </div>
                            {priceElement}
                        </div>
                        <hr />
                        <div className="total-price-box">
                            <span className="title-total-price">ราคารวม</span>
                            <span className="total-people">{totalPeople}</span>
                            <span className="total-price">
                                ฿{totalPrice.toLocaleString('th-TH', {
                                    minimumFractionDigits: 2,
                                    maximumFractionDigits: 2,
                                })}
                            </span>
                        </div>
                        <hr />
                        <div className="promotion-box">
                            <div className="box-for-input">
                                <span className="title-promotion">โค้ดส่วนลด (ถ้ามี)</span>
                                <div className="sub-box-input">
                                    <button className="promotion-btn">ใช้โค้ด</button>
                                </div>
                            </div>
                            <div className="discount-box">
                                <span className="description"></span>
                                <span className="discount"></span>
                            </div>
                        </div>
                    </div>
                    <div className="card-box">
                        <div className="booking-data-card sub-card">
                            <span className="title">ตรวจสอบข้อมูลผู้จอง</span>
                            <form action="">
                                <div className="fname data-box">
                                    <span className="title-input">ชื่อ</span>
                                    <input type="text"
                                        defaultValue={customer?.FirstName}
                                        placeholder="โปรดป้อนชื่อ"
                                        disabled={isDisabled}
                                        onChange={(e) => setFName(e.target.value)} required
                                    />
                                </div>
                                <div className="fname data-box">
                                    <span className="title-input">นามสกุล</span>
                                    <input type="text"
                                        defaultValue={customer?.LastName}
                                        placeholder="โปรดป้อนนามสกุล"
                                        disabled={isDisabled}
                                        onChange={(e) => setLName(e.target.value)} required
                                    />
                                </div>
                                <div className="fname data-box">
                                    <span className="title-input">เบอร์โทรศัพท์</span>
                                    <input type="text"
                                        defaultValue={customer?.PhoneNumber}
                                        placeholder="โปรดป้อนเบอร์โทรศัพท์ (000-000-0000)"
                                        disabled={isDisabled}
                                        pattern="[0-9]{3}-[0-9]{3}-[0-9]{4}"
                                        onChange={(e) => setPhoneNumber(e.target.value)} required
                                    />
                                </div>
                                <div className="fname data-box">
                                    <span className="title-input">อีเมล</span>
                                    <input type="email"
                                        defaultValue={customer?.Email}
                                        placeholder="โปรดป้อนอีเมล (sa@gmail.com)"
                                        disabled={isDisabled}
                                        onChange={(e) => setEmail(e.target.value)} required />
                                </div>
                                <div className="btn-box">
                                    {
                                        isDisabled ? (
                                            <button className="edit-btn" onClick={() => setIsDisabled(false)}>แก้ไขข้อมูล</button>
                                        ) : (
                                            <div className="sub-btn-box">
                                                <div className="cancel-btn" onClick={() => handleCancle()}>ยกเลิก</div>
                                                <div className="confirm-btn">บันทึก</div>
                                            </div>
                                        )
                                    }

                                </div>
                            </form>

                        </div>
                        <div className="confirm-booking-card sub-card">
                            <div className="detail">
                                <span className="title">ตรวจสอบข้อมูลการจอง</span>
                                <div className="tour-name-box sub-box">
                                    <div className="img-box">
                                        <img src="./images/icons/tour.png" alt="" />
                                    </div>
                                    {tourPackage?.TourName}
                                </div>
                                <div className="tour-date-box sub-box">
                                    <div className="img-box">
                                        <img src="./images/icons/calendar.png" alt="" />
                                    </div>
                                    {dateSelectedFormat}
                                </div>
                                <div className="people-quantity-box sub-box">
                                    <div className="img-box">
                                        <img src="./images/icons/bag.png" alt="" />
                                    </div>
                                    ผู้เดินทางจำนวน {totalPeople} ท่าน
                                </div>
                                <div className="btn-box">
                                    <button className="cancel-btn btn" onClick={()=>setPopUp(<></>)}>ยกเลิก </button>
                                    <button className="confirm-btn btn" 
                                        disabled={totalPeople!=0 ? false : true}
                                        onClick={handleCreateBooking}
                                        style={{
                                            pointerEvents: totalPeople!=0 ? "auto" : "none",
                                            opacity: totalPeople!=0 ? "1" : "0.6"
                                        }}
                                    >ยืนยันการจอง</button>
                                </div>
                            </div>
                            <div className="picture-box">
                                <img src="./images/backgrounds/travel.jpg" alt="" />
                            </div>
                        </div>
                    </div>

                </section>
            </div>
        </div>
    )
}
export default Booking