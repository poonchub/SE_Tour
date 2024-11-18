import { Key, useState, SetStateAction, useEffect } from "react"
import "./Booking.css"
import { CustomersInterface } from "../../interfaces/ICustomers";
import { GetCustomerByID } from "../../services/http";


function Booking(props: { roomTypes: any; tourPackage: any; personTypes: any; }) {

    const { roomTypes, tourPackage, personTypes } = props

    const [customer, setCustomer] = useState<CustomersInterface>();

    const [childAdultSingleCount, setChildAdultSingleCount] = useState(0)
    const [childAdultDoubleCount, setChildAdultDoubleCount] = useState(0)
    const [childAdultThreeCount, setChildAdultThreeCount] = useState(0)
    const [infantAddBedCount, setInfantAddBedCount] = useState(0)
    const [infantNoAddBedCount, setNoInfantAddBedCount] = useState(0)

    const [childAdultSinglePrice, setChildAdultSinglePrice] = useState(0)
    const [childAdultDoublePrice, setChildAdultDoublePrice] = useState(0)
    const [childAdultThreePrice, setChildAdultThreePrice] = useState(0)
    const [infantAddBedPrice, setInfantAddBedPrice] = useState(0)
    const [infantNoAddBedPrice, setNoInfantAddBedPrice] = useState(0)

    const [fName, setFName] = useState<string | undefined>("");
    const [lName, setLName] = useState<string | undefined>("");
    const [phoneNumber, setPhoneNumber] = useState<string | undefined>("");
    const [email, setEmail] = useState<string | undefined>("");

    const [totalPrice, setTotalPrice] = useState<number>(0)

    async function getCustomerByID() {
        let res = await GetCustomerByID(1)
        if (res) {
            setCustomer(res);
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

    function handleCheckBox(e: boolean) {
        if (e) {
            setFName(customer?.FirstName)
            setLName(customer?.LastName)
            setPhoneNumber(customer?.PhoneNumber)
            setEmail(customer?.Email)
        }
        else {
            setFName("")
            setLName("")
            setPhoneNumber("")
            setEmail("")
        }
    }

    useEffect(() => {
        getCustomerByID()
    }, [])

    console.log(totalPrice)

    useEffect(() => {
        if (childAdultSingleCount == 0 && childAdultDoubleCount == 0 && childAdultDoubleCount == 0) {
            setInfantAddBedCount(0)
            setNoInfantAddBedCount(0)
        }
    }, [childAdultSingleCount, childAdultDoubleCount, childAdultDoubleCount])

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
                    index == 3 ? <span style={{ fontWeight: "600", margin: "12px 0px 5px 0px" }}>เด็กเล็กพักกับผู้ใหญ่</span> : <></>
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
                        <div className="total-price">
                            <span>ราคารวม</span>
                            ฿{totalPrice.toLocaleString('th-TH', {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2,
                            })}
                        </div>
                        <hr />
                    </div>
                    <div className="booking-data-card sub-card">
                        <span className="title">โปรดระบุข้อมูลผู้จอง</span>
                        <form action="">
                            <div className="fname data-box">
                                <span className="title-input">ชื่อ</span>
                                <input type="text"
                                    value={fName}
                                    placeholder="โปรดป้อนชื่อ"
                                    onChange={(e) => setFName(e.target.value)} required
                                />
                            </div>
                            <div className="fname data-box">
                                <span className="title-input">นามสกุล</span>
                                <input type="text"
                                    value={lName}
                                    onChange={(e) => setLName(e.target.value)}
                                    placeholder="โปรดป้อนนามสกุล" required
                                />
                            </div>
                            <div className="fname data-box">
                                <span className="title-input">เบอร์โทรศัพท์</span>
                                <input type="text"
                                    value={phoneNumber}
                                    onChange={(e) => setPhoneNumber(e.target.value)}
                                    placeholder="โปรดป้อนเบอร์โทรศัพท์ (000-000-0000)"
                                    pattern="[0-9]{3}-[0-9]{3}-[0-9]{4}" required
                                />
                            </div>
                            <div className="fname data-box">
                                <span className="title-input">อีเมล</span>
                                <input type="email"
                                    value={email}
                                    onChange={(e) => setEmail(e.target.value)}
                                    placeholder="โปรดป้อนอีเมล (sa@gmail.com)" required />
                            </div>
                            <div className="checkbox">
                                <input type="checkbox" className="select-data-customer" onChange={(e) => handleCheckBox(e.target.checked)} />
                                <span>ใช้ข้อมูลของผู้ใช้</span>
                            </div>
                        </form>

                    </div>
                </section>
            </div>
        </div>
    )
}
export default Booking