import { useEffect, useState } from "react"
import "./ShowTourSlice.css"
import { TourPackagesInterface } from "../../interfaces/ITourPackages";
import { apiUrl, GetTourPackages } from "../../services/http";
import { Carousel } from "antd";

function ShowTourSlice() {
    const [tourPackages, setTourPackages] = useState<TourPackagesInterface[]>([]);

    const [isLoading, setIsLoading] = useState<boolean>(true)
    const [currentSlide, setCurrentSlide] = useState(0);
    const [startPriceFormat, setStartPriceFormat] = useState<string>("");

    async function getTourPackages() {
        let res = await GetTourPackages()
        if (res) {
            setTourPackages(res);
        }
    }

    async function fetchData() {
        try {
            getTourPackages()
        } catch (error) {
            console.error('Failed to fetch data:', error);
        } finally {
            setIsLoading(false);
        }
    }

    function handleSetStartPrice() {
        let startPrice = 999999;
        tourPackages[currentSlide]?.TourPrices?.forEach((price: any) => {
            if (price.PersonTypeID !== 1 && price.Price && price.Price < startPrice) {
                startPrice = price.Price;
            }
        });
        setStartPriceFormat(startPrice.toLocaleString('th-TH', {
            minimumFractionDigits: 2,
            maximumFractionDigits: 2,
        }))
    }

    function setPackageData(id: number | undefined) {
        localStorage.setItem("startPrice", startPriceFormat)
        localStorage.setItem("tourPackageID", String(id))
        setTimeout(() => {
            location.href = "/tour-select";
        });
    }

    console.log(startPriceFormat)

    useEffect(() => {
        handleSetStartPrice()
    }, [currentSlide, tourPackages]);

    useEffect(() => {
        fetchData()
    }, [])

    console.log(tourPackages?.[0]?.TourDescriptions?.Intro)

    const bigImageUrl = `${apiUrl}/${tourPackages[currentSlide]?.TourImages?.[0]?.FilePath}`

    const imageElement = tourPackages.map((tour, index) => {
        const imageUrl = `${apiUrl}/${tour?.TourImages?.[0]?.FilePath}`
        return (
            <div className="image-box" key={index} onClick={()=>setPackageData(tour.ID)}>
                <img src={imageUrl} alt="" />
            </div>
        )
    })

    return (
        <div className="show-tour-slice-container">
            <div className="image-bg">
                <img src={bigImageUrl} alt="" />
            </div>
            <div className="tour-container">
                <div className="image-box">
                    <Carousel afterChange={(current) => setCurrentSlide(current)} autoplay autoplaySpeed={5000} speed={2000}>
                        {imageElement}
                    </Carousel>
                </div>
                <div className="text-box">
                    <h3 className="tour-name">
                        {tourPackages?.[currentSlide]?.TourName}
                    </h3>
                    <p className="tour-description">
                        {tourPackages?.[currentSlide]?.TourDescriptions?.Intro}
                    </p>
                    <p className="title-text">รายละเอียดแพ็กเกจ</p>
                    <p className="tour-detail subtext">
                        {tourPackages?.[currentSlide]?.TourDescriptions?.PackageDetail}
                    </p>
                    <p className="title-text">ไฮไลท์ของทริป</p>
                    <p className="tour-detail subtext">
                        {tourPackages?.[currentSlide]?.TourDescriptions?.TripHighlight}
                    </p>
                </div>
            </div>

        </div>
    )
}
export default ShowTourSlice