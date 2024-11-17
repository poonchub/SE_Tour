import { ReactElement, JSXElementConstructor, ReactNode, ReactPortal, Key, useState, ChangeEvent, SetStateAction } from "react"
import "./Booking.css"

function Booking(props: { roomTypes: any; tourPackage: any; personTypes: any; }) {

    const { roomTypes, tourPackage, personTypes } = props

    const [childAdultSingleCount, setChildAdultSingleCount] = useState(0)
    const [childAdultDoubleCount, setChildAdultDoubleCount] = useState(0)
    const [childAdultThreeCount, setChildAdultThreeCount] = useState(0)

    const [infantAddBedCount, setInfantAddBedCount] = useState(0)
    const [infantNoAddBedCount, setNoInfantAddBedCount] = useState(0)

    function handleChange(e: string | SetStateAction<number>, index: Key | null | undefined) {
        switch (index) {
            case 0:
                setChildAdultSingleCount(Number(e))
                break;
            case 1:
                setChildAdultDoubleCount(Number(e))
                break;
            case 2:
                setChildAdultThreeCount(Number(e))
                break;
            case 3:
                setInfantAddBedCount(Number(e))
                break;
            default:
                setNoInfantAddBedCount(Number(e))
                break;
        }
    }

    const priceElement1 = roomTypes?.map((type: { ID: any; TypeName: string | number | boolean | ReactElement<any, string | JSXElementConstructor<any>> | Iterable<ReactNode> | ReactPortal | null | undefined }, index: Key | null | undefined) => {
        const tourPrices = tourPackage?.TourPrices
        var p1
        tourPrices?.forEach((price: { RoomTypeID: any; PersonTypeID: any; Price: { toLocaleString: (arg0: string, arg1: { minimumFractionDigits: number; maximumFractionDigits: number }) => any } }, _: any) => {
            if (price.RoomTypeID === type.ID && price.PersonTypeID) {
                p1 = price.Price?.toLocaleString('th-TH', {
                    minimumFractionDigits: 2,
                    maximumFractionDigits: 2,
                })
            }
        });
        return p1 ? (
            <>  
                {             
                    index == 3 ? <span style={{fontWeight: "600", margin: "12px 0px 5px 0px"}}>เด็กเล็กพักกับผู้ใหญ่</span> : <></>
                }
                <div className="price-box" key={index}>
                    <span className="type-name">{type.TypeName}</span>
                    <input type="number" defaultValue={
                        index == 0 ? childAdultSingleCount : (
                            index == 1 ? childAdultDoubleCount : (
                                index == 2 ? childAdultThreeCount : (
                                    index == 3 ? infantAddBedCount :
                                        infantNoAddBedCount
                                )
                            )
                        )
                    } onChange={(e) => handleChange(e.target.value, index)} min={0} />
                    <span className="price">{
                        index == 0 ? childAdultSingleCount : (
                            index == 1 ? childAdultDoubleCount : (
                                index == 2 ? childAdultThreeCount : (
                                    index == 3 ? infantAddBedCount :
                                        infantNoAddBedCount
                                )
                            )
                        )
                    } × {p1}</span>
                </div>
            </>

        ) : ""
    })

    return (
        <div className="booking-container">
            <div className="card">
                <span className="header">จองแพ็กเกจทัวร์</span>
                <section>
                    <div className="select-persontype-card sub-card">
                        <span className="title">โปรดระบุจำนวนผู้เดินทาง</span>
                        <div className="add-quantity-box">
                            <div className="title-box">
                                <span className="span1">เด็ก/ผู้ใหญ่</span>
                                <span className="span2">จำนวน (คน)</span>
                                <span className="span3">ราคา (บาท)</span>
                            </div>
                            {priceElement1}
                        </div>
                    </div>
                    <div className="booking-data-card sub-card">
                        <span className="title">โปรดระบุข้อมูลผู้จอง</span>


                    </div>
                </section>
            </div>
        </div>
    )
}
export default Booking