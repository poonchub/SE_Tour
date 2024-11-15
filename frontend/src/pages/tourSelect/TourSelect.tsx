import { useEffect, useState } from "react";
import Navbar from "../../components/navbar/Navbar";
import "./TourSelect.css";
import { TourPackagesInterface } from "../../interfaces/ITourPackages";
import { apiUrl, GetTourPackageByID } from "../../services/http";

function TourSelect() {
    const [tourPackage, setTourPackage] = useState<TourPackagesInterface>();

    const [bigImage, setBigImage] = useState()

    const startPrice = localStorage.getItem("startPrice");
    const tourPackageID = localStorage.getItem("tourPackageID");

    async function getTourPackage() {
        let res = await GetTourPackageByID(Number(tourPackageID));
        if (res) {
            setTourPackage(res);
        }
    }

    useEffect(() => {
        getTourPackage();
    }, []);

    const content1 = document.querySelector(".content1");
    if (content1 && tourPackage?.TourDescriptions?.PackageDetail) {
        content1.innerHTML = tourPackage.TourDescriptions.PackageDetail.replace(/\n/g,"<br>");
    }

    const content2 = document.querySelector(".content2");
    if (content2 && tourPackage?.TourDescriptions?.TripHighlight) {
        content2.innerHTML = tourPackage.TourDescriptions.TripHighlight.replace(/\n/g,"<br>");
    }

    const content3 = document.querySelector(".content3");
    if (content3 && tourPackage?.TourDescriptions?.PlacesHighlight) {
        content3.innerHTML = tourPackage.TourDescriptions.PlacesHighlight.replace(/\n/g,"<br>");
    }

    const imageElement = (tourPackage?.TourImages as any[])?.map(
        (image, index) => (
            <div className="sImage" id={`image${index+1}`} key={index} onClick={()=>setBigImage(image.FilePath)}>
                <img src={`${apiUrl}/${image.FilePath}`}/>
            </div>
        )
    );

    console.log(tourPackage?.TourImages);

    return (
        <div className="tour-select-page">
            <Navbar page={"tourSelect"} />
            <section>
                <div className="package-detail">
                    <div className="image-box">
                        <div className="big-image">
                            <img src={`${apiUrl}/${bigImage}`} alt="" />
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
            </section>
        </div>
    );
}
export default TourSelect;
