import { useEffect, useState } from "react"
import "./ShowTourSlice.css"
import { TourPackagesInterface } from "../../interfaces/ITourPackages";
import { apiUrl, GetTourPackages } from "../../services/http";
import { Carousel } from "antd";

function ShowTourSlice() {
    const [tourPackages, setTourPackages] = useState<TourPackagesInterface[]>([]);

    const [imageBg, setimageBg] = useState()
    const [isLoading, setIsLoading] = useState<boolean>(true)
    const [currentSlide, setCurrentSlide] = useState(0);

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

    useEffect(() => {
        fetchData()
    }, [])

    console.log(tourPackages?.[0]?.TourDescriptions?.Intro)

    const imageElement = tourPackages.map((tour, index) => {
        const imageUrl = `${apiUrl}/${tour?.TourImages?.[0]?.FilePath}`
        return (
            <div className="image-box" key={index}>
                <img src={imageUrl} alt="" />
            </div>
        )
    })

    return (
        <div className="show-tour-slice-container">
            <div className="image-bg">
                <img src={"/images/backgrounds/background.jpg"} alt="" />
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