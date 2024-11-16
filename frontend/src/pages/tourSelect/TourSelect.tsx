import { useEffect, useState } from "react";
import Navbar from "../../components/navbar/Navbar";
import "./TourSelect.css";
import { TourPackagesInterface } from "../../interfaces/ITourPackages";
import { apiUrl, GetPersonTypes, GetRoomTypes, GetTourPackageByID } from "../../services/http";
import Loading from "../../components/loading/Loading";
import Calendar from "../../components/calendar/Calendar";
import { PersonTypesInterface } from "../../interfaces/IPersonTypes";
import { RoomTypesInterface } from "../../interfaces/IRoomTypes";

function TourSelect() {
    const [tourPackage, setTourPackage] = useState<TourPackagesInterface>();
    const [personTypes, setPersonTypes] = useState<PersonTypesInterface[]>();
    const [roomTypes, setRoomTypes] = useState<RoomTypesInterface[]>();

    const [bigImage, setBigImage] = useState<string>();
    const [isLoading, setIsLoading] = useState<boolean>(true);

    async function getTourPackage() {
        try {
            const resTourPackage = await GetTourPackageByID(Number(tourPackageID));
            if (resTourPackage) {
                setTourPackage(resTourPackage);
            }

            const resPersonType = await GetPersonTypes();
            if (resPersonType) {
                setPersonTypes(resPersonType)
            }

            const resRoomType = await GetRoomTypes();
            if (resRoomType) {
                setRoomTypes(resRoomType)
            }
        } catch (error) {
            console.error('Failed to fetch tour package:', error);
        } finally {
            setIsLoading(false);
        }
    }

    useEffect(() => {
        getTourPackage();

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



    const priceElement = roomTypes?.map((type, index) => {
        const tourPrices = tourPackage?.TourPrices
        var p
        tourPrices?.forEach((price, _) => {
            if(price.RoomTypeID==type.ID){
                p = price.Price?.toLocaleString('th-TH', {
                    minimumFractionDigits: 2,
                    maximumFractionDigits: 2,
                })
            }
        });
        return (
            <div className="price-box" key={index}>
                <span className="type">{type.TypeName}</span>
                <span className="price">{p}</span>
            </div>
        )
    })

    console.log(tourPackage?.TourPrices)
    console.log(roomTypes)

    return isLoading ? (
        <Loading />
    ) : (
        <div className="tour-select-page">
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
                            <Calendar />
                        </div>
                        <div className="travel-schedule-detail">
                            <div className="date-booking-box">
                                <div className="date-booking">วันที่ 22-23 พฤศจิกายน 2567</div>
                                <div className="booking-btn">จองทัวร์</div>
                            </div>
                            <div className="price-detail">
                                <span className="title">ราคาแพ็กเกจ</span>
                                <div className="person-type-title">
                                    <span className="type">เด็ก/ผู้ใหญ่ (บาท/ท่าน)</span>
                                    <span className="type">เด็กเล็ก (บาท/ท่าน)</span>
                                </div>
                                <div className="price-all-box">
                                    {priceElement}
                                </div>
                            </div>
                        </div>
                    </div>

                </div>
            </section>
        </div>
    )
}
export default TourSelect;
