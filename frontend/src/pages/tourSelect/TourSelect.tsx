import { useEffect, useState } from "react";
import Navbar from "../../components/navbar/Navbar";
import "./TourSelect.css";
import { TourPackagesInterface } from "../../interfaces/ITourPackages";
import { apiUrl, GetPersonTypes, GetRoomTypes, GetTourPackageByID } from "../../services/http";
import Loading from "../../components/loading/Loading";
import Calendar from "../../components/calendar/Calendar";
import { PersonTypesInterface } from "../../interfaces/IPersonTypes";
import { RoomTypesInterface } from "../../interfaces/IRoomTypes";
import Footer from "../../components/footer/Footer";
import Booking from "../../components/booking/Booking";

function TourSelect() {
    const [tourPackage, setTourPackage] = useState<TourPackagesInterface>();
    const [personTypes, setPersonTypes] = useState<PersonTypesInterface[]>();
    const [roomTypes, setRoomTypes] = useState<RoomTypesInterface[]>();

    const [bigImage, setBigImage] = useState<string>();
    const [isLoading, setIsLoading] = useState<boolean>(true);
    const [dateSelectedFormat, setDateSelectedFormat] = useState<string | null>("โปรดเลือกวันที่ต้องการจองจากปฎิทิน");
    const [bookingPopUp, setBookingPopUp] = useState(<></>);

    async function getTourPackage() {
        const resTourPackage = await GetTourPackageByID(Number(tourPackageID));
        if (resTourPackage) {
            setTourPackage(resTourPackage);
        }
    }

    async function getPersonTypes() {
        const resPersonType = await GetPersonTypes();
        if (resPersonType) {
            setPersonTypes(resPersonType)
        }
    }

    async function getRoomTypes() {
        const resRoomType = await GetRoomTypes();
        if (resRoomType) {
            setRoomTypes(resRoomType)
        }
    }

    async function fetchData() {
        try {
            getTourPackage()
            getPersonTypes()
            getRoomTypes()
        } catch (error) {
            console.error('Failed to fetch data:', error);
        } finally {
            setIsLoading(false);
        }
    }

    useEffect(() => {
        fetchData()
    }, [isLoading]);

    const startPrice = localStorage.getItem("startPrice");
    const tourPackageID = localStorage.getItem("tourPackageID");

    const content1 = document.querySelector(".content1");
    if (content1 && tourPackage?.TourDescriptions?.PackageDetail) {
        content1.innerHTML = tourPackage.TourDescriptions.PackageDetail.replace(/\n/g, "<br>");
    }

    const content2 = document.querySelector(".content2");
    if (content2 && tourPackage?.TourDescriptions?.TripHighlight) {
        content2.innerHTML = tourPackage.TourDescriptions.TripHighlight.replace(/\n/g, "<br>");
    }

    const content3 = document.querySelector(".content3");
    if (content3 && tourPackage?.TourDescriptions?.PlacesHighlight) {
        content3.innerHTML = tourPackage.TourDescriptions.PlacesHighlight.replace(/\n/g, "<br>");
    }

    const imageElement = (tourPackage?.TourImages as any[])?.map(
        (image, index) => (
            <div className="sImage" id={`image${index + 1}`} key={index} onClick={() => setBigImage(image.FilePath)}>
                <img src={`${apiUrl}/${image.FilePath}`} />
            </div>
        )
    );

    const priceElement1 = roomTypes?.map((type, index) => {
        const tourPrices = tourPackage?.TourPrices
        var p1
        tourPrices?.forEach((price, _) => {
            if (price.RoomTypeID === type.ID && price.PersonTypeID === personTypes?.[1]?.ID) {
                p1 = price.Price?.toLocaleString('th-TH', {
                    minimumFractionDigits: 2,
                    maximumFractionDigits: 2,
                })
            }
        });
        return p1 ? (
            <div className="price-box" key={index}>
                <span className="type-name">{type.TypeName}</span>
                <span className="price">฿{p1}</span>
            </div>
        ) : ""
    })

    const priceElement2 = roomTypes?.map((type, index) => {
        const tourPrices = tourPackage?.TourPrices
        var p2
        tourPrices?.forEach((price, _) => {
            if (price.RoomTypeID === type.ID && price.PersonTypeID === personTypes?.[0]?.ID) {
                p2 = price.Price?.toLocaleString('th-TH', {
                    minimumFractionDigits: 2,
                    maximumFractionDigits: 2,
                })
            }
        });
        return p2 ? (
            <div className="price-box" key={index}>
                <span className="type-name">{type.TypeName}</span>
                <span className="price">฿{p2}</span>
            </div>
        ) : ""
    })

    const schedules = tourPackage?.TourSchedules
    let dateTime: string[][] = [];
    let availableDates: string[] = []
    schedules?.forEach((schedule, index) => {
        if (!dateTime[index]) {
          dateTime[index] = [];
        }
        if (schedule.StartDate && schedule.EndDate){
            dateTime[index].push(schedule.StartDate.slice(0,10));
            dateTime[index].push(schedule.EndDate.slice(0,10));
            availableDates.push(schedule.StartDate.slice(0,10))
        }
    });

    return isLoading ? (
        <Loading />
    ) : (
        <div className="tour-select-page">
            {bookingPopUp}
            <Navbar page={"tourSelect"} />
            <section>
                <div className="package-detail">
                    <div className="image-box">
                        <div className="big-image">
                            <img src={`${apiUrl}/${bigImage ? bigImage : tourPackage?.TourImages ? tourPackage?.TourImages[0].FilePath : ""}`} alt="" />
                        </div>
                        <div className="small-image">{imageElement}</div>
                    </div>
                    <div className="description-box">
                        <h2 className="tour-name">{tourPackage?.TourName}</h2>
                        <div className="package-detail-box des-box">
                            <span className="title">รายละเอียดแพ็กเกจ</span>
                            <p className="content1 detail"></p>
                        </div>
                        <div className="price-box des-box">
                            <span className="price-title">ราคาเริ่มต้น</span>
                            <p className="price">฿{startPrice}</p>
                        </div>
                        <hr />
                        <div className="trip-highlight des-box">
                            <span className="title">ไฮไลท์ของทริป</span>
                            <p className="content2 detail"></p>
                        </div>
                        <div className="places-highlight des-box">
                            <span className="title">จุดเด่นของแพ็กเกจ</span>
                            <p className="content3 detail"></p>
                        </div>
                    </div>
                </div>
                <div className="travel-schedule">
                    <div className="title-box">
                        <div className="img-box">
                            <img src="./images/icons/calendar.png" alt="" />
                        </div>
                        <h2 className="title">กำหนดการเดินทาง</h2>
                    </div>
                    <div className="subsection">
                        <div className="calendar-box">
                            <Calendar dateTime={dateTime} setDateSelectedFormat={setDateSelectedFormat} />
                        </div>
                        <div className="travel-schedule-detail">
                            <div className="date-booking-box">
                                <div className="date-booking">{dateSelectedFormat}</div>
                                <div className="booking-btn" onClick={()=>setBookingPopUp(
                                    <Booking roomTypes={roomTypes} tourPackage={tourPackage} personTypes={personTypes}/>
                                )}>จองทัวร์</div>
                            </div>
                            <div className="price-detail">
                                <span className="title">ราคาแพ็กเกจ</span>
                                <div className="person-type-title">
                                    <div className="type-box">
                                        <span className="type-title">เด็ก/ผู้ใหญ่ (บาท/ท่าน)</span>
                                        {priceElement1}
                                    </div>
                                    <div className="type-box">
                                        <span className="type-title">เด็กเล็ก (บาท/ท่าน)</span>
                                        {priceElement2}
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>

                </div>
                <div className="travel-plane">
                    <div className="title-box">
                        <div className="img-box">
                            <img src="./images/icons/plans.png" alt="" />
                        </div>
                        <h2 className="title">แผนการเดินทาง</h2>
                    </div>
                </div>
            </section>
            <Footer />
        </div>
    )
}
export default TourSelect;
