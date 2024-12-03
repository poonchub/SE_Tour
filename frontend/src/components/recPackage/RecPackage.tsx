import { Link } from "react-router-dom"
import "./RecPackage.css"
import { TourPackagesInterface } from "../../interfaces/ITourPackages"
import { apiUrl } from "../../services/http"
import { useEffect, useState } from "react"

function RecPackage(props: { tourPackages: any }) {
    const { tourPackages } = props

    const [startPriceFormat, setStartPriceFormat] = useState<string>("");

    function handleSetStartPrice() {
        let startPrice = 999999;
        tourPackages[0]?.TourPrices?.forEach((price: any) => {
            if (price.PersonTypeID !== 1 && price.Price && price.Price < startPrice) {
                startPrice = price.Price;
            }
        });
        setStartPriceFormat(startPrice.toLocaleString('th-TH', {
            minimumFractionDigits: 2,
            maximumFractionDigits: 2,
        }))
    }

    const tourElement = tourPackages.map((tour: TourPackagesInterface, index: number)=>{
        const imageUrl = `${apiUrl}/${tour?.TourImages?.[0]?.FilePath}`
        return (
            <div className="tour-box" key={index}>
                <img className="image-bg" src={imageUrl} alt="" />
                <div className="text-box">
                    <div className="tour-title">
                        {tour.TourName}
                    </div>
                    <div className="tour-price">
                        {startPriceFormat}
                    </div>
                </div>
            </div>
        )
    })

    useEffect(() => {
        handleSetStartPrice()
    }, [tourPackages]);

    return (
        <div className="rec-package-container">
            <div className="card-title">
                <div className="title">
                    <div className="icon-box">
                        <img src="/images/icons/island.png" alt="" />
                    </div>
                    <div className="text-box">แนะนำแพ็กเกจทัวร์สุดพิเศษ</div>
                </div>
                <Link to={"/tour-package"}>
                    <button className="more-tour-btn">แพ็กเกจทั้งหมด</button>
                </Link>
            </div>
            <section className="tour-section">
                {tourElement}
            </section>
        </div>
    )
}
export default RecPackage