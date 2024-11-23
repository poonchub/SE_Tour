import { useEffect, useState } from "react";
import Navbar from "../../components/navbar/Navbar"
import "./TourPackage.css"
import { GetProvinces, GetTourPackages } from "../../services/http";
import { TourPackagesInterface } from "../../interfaces/ITourPackages";
import PackageItem from "../../components/packageItem/PackageItem";
import Loading from "../../components/loading/Loading";
import { ProvincesInterface } from "../../interfaces/IProvinces";

function TourPackage() {

    const [tourPackages, setTourPackages] = useState<TourPackagesInterface[]>([]);
    const [provinces, setProvinces] = useState<ProvincesInterface[]>([]);

    const [isLoading, setIsLoading] = useState<boolean>(true);
    const [currentIndex, setCurrentIndex] = useState(0);

    async function getTourPackages() {
        let res = await GetTourPackages()
        if (res) {
            setTourPackages(res);
        }
    }

    async function getTourProvinces() {
        let res = await GetProvinces()
        if (res) {
            setProvinces(res);
        }
    }

    async function fetchData() {
        try {
            getTourPackages()
            getTourProvinces()
        } catch (error) {
            console.error('Failed to fetch data:', error);
        } finally {
            setIsLoading(false);
        }
    }

    const images = [
        './images/sliceshow/pic1.jpg',
        './images/sliceshow/pic2.jpg',
        './images/sliceshow/pic3.jpg',
        './images/sliceshow/pic4.jpg',
        './images/sliceshow/pic5.jpg',
        './images/sliceshow/pic6.jpg',
    ];

    useEffect(() => {
        const interval = setInterval(() => {
            setCurrentIndex((prevIndex) =>
                prevIndex === images.length - 1 ? 0 : prevIndex + 1
            );
        }, 3000);
        return () => clearInterval(interval);
    }, [currentIndex]);

    useEffect(() => {
        fetchData()
    }, [])

    const tourElements = tourPackages.map((tour, index) => {
        return <PackageItem key={index} tour={tour} />
    })

    return isLoading ? (
        <Loading />
    ) : (
        <div className="tour-package-page">
            <Navbar page={"tourPackage"} />
            <section>
                <div className="slideshow-container">
                    <div className="slideshow-wrapper" style={{
                        transform: `translateX(-${currentIndex * 100}%)`,
                    }}>
                        {images.map((image, index) => (
                            <img
                                key={index}
                                src={image}
                                alt={`Slide ${index}`}
                                className="slide-image"
                            />
                        ))}
                    </div>
                </div>

                <div className="subsection">
                    <div className="show-mini-promotion">

                    </div>
                    <form className="search-package-box">
                        <h3 className="title">ค้นหาแพ็กเกจทัวร์ที่คุณต้องการได้อย่างง่ายดาย✨</h3>
                        <div className="search-box">
                            <div className="img-box">
                                <img src="./images/icons/search.png" alt="" />
                            </div>
                            <input type="text" placeholder="ค้นหาแพ็กเกจ..." />
                        </div>
                        <div className="search-option-box">
                            <div className="option1-box option">
                                <span className="text">แพ็กเกจในจังหวัด</span>
                                <select name="" id="">
                                    {
                                        provinces.map((province, index) => (
                                            <option value={province.ID} key={index}>{province.ProvinceName}</option>
                                        ))
                                    }
                                </select>
                            </div>
                            <div className="option2-box option">
                                <span className="text">ช่วงเวลา</span>
                                <div className="input-box">
                                    <input type="date" />
                                    -
                                    <input type="date" />
                                </div>
                            </div>
                            <div className="option3-box option">
                                <span className="text">ช่วงราคา</span>
                                <div className="input-box">
                                    <input type="number" min={0} step={500} defaultValue={0} />
                                    -
                                    <input type="number" min={1000} step={500} defaultValue={1000} />
                                </div>
                            </div>
                            <div className="option4-box option">
                                <span className="text">จัดเรียงตาม</span>
                                <select name="" id="">
                                    <option value="">ราคาต่ำ-สูง</option>
                                    <option value="">ราคาสูง-ต่ำ</option>
                                </select>
                            </div>
                        </div>
                    </form>
                </div>
                {tourElements}
            </section>
        </div>
    )
}
export default TourPackage