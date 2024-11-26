import { useEffect, useState } from "react";
import Navbar from "../../components/navbar/Navbar"
import "./TourPackage.css"
import { GetProvinces, GetTourPackages } from "../../services/http";
import { TourPackagesInterface } from "../../interfaces/ITourPackages";
import PackageItem from "../../components/packageItem/PackageItem";
import Loading from "../../components/loading/Loading";
import { ProvincesInterface } from "../../interfaces/IProvinces";
import Footer from "../../components/footer/Footer";

function TourPackage() {

    const [tourPackages, setTourPackages] = useState<TourPackagesInterface[]>([]);
    const [provinces, setProvinces] = useState<ProvincesInterface[]>([]);

    const [isLoading, setIsLoading] = useState<boolean>(true);
    const [currentIndex, setCurrentIndex] = useState(0);

    const [searchText, setSearchText] = useState("")
    const [provinceID, setProvinceID] = useState("")
    const [startDate, setStartDate] = useState("")
    const [endDate, setEndDate] = useState("")
    const [minPrice, setMinPrice] = useState(0)
    const [maxPrice, setMaxPrice] = useState(10000)
    const [sortOption, setSortOption] = useState(1)

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
        const today = new Date();
        const formattedDate = today.toISOString().split('T')[0];
        setStartDate(formattedDate);
    }, [])

    const filteredTours = tourPackages.filter((tour) => {
        let Price = 999999
        tour?.TourPrices?.forEach((price) => {
            if (price.PersonTypeID != 1 && price.Price && price.Price < Price) {
                Price = price.Price
            }
        });

        const hasValidSchedule = (startDate != "" && endDate != "") ? (
            tour?.TourSchedules?.some((schedule) => {
                const std_ = schedule.StartDate?.slice(0, 10)
                const ed_ = schedule.EndDate?.slice(0, 10)
                return (ed_ && std_ && (std_ >= startDate) && (ed_ <= endDate))
            })
        ) : (startDate != "") ? (
            tour?.TourSchedules?.some((schedule) => {
                const std_ = schedule.StartDate?.slice(0, 10)
                return (std_ && (std_ >= startDate))
            })
        ) : (endDate != "") ? (
            tour?.TourSchedules?.some((schedule) => {
                const ed_ = schedule.EndDate?.slice(0, 10)
                return (ed_ && (ed_ <= endDate))
            })
        ) : (true)

        return provinceID != "" ? (
            (tour?.TourName?.toLowerCase().includes(searchText.toLowerCase())) && (Price >= minPrice && Price <= maxPrice) && (tour.ProvinceID == Number(provinceID)) && hasValidSchedule
        ) : (
            (tour?.TourName?.toLowerCase().includes(searchText.toLowerCase())) && (Price >= minPrice && Price <= maxPrice) && hasValidSchedule
        )
    })
    const sortedTours = (sortOption===1) ? (
        filteredTours.sort((a, b) => {
            const priceA = a?.TourPrices?.length ? 
                a.TourPrices.reduce((min, price) => {
                    return (price.PersonTypeID !== 1 && price.Price && price.Price < min) ? price.Price : min
                }, 999999) : 999999
        
            const priceB = b?.TourPrices?.length ? 
                b.TourPrices.reduce((min, price) => {
                    return (price.PersonTypeID !== 1 && price.Price && price.Price < min) ? price.Price : min
                }, 999999) : 999999
            return priceA - priceB
        })
    ) : (
        filteredTours.sort((a, b) => {
            const priceA = a?.TourPrices?.length ? 
                a.TourPrices.reduce((min, price) => {
                    return (price.PersonTypeID !== 1 && price.Price && price.Price < min) ? price.Price : min
                }, 999999) : 999999
        
            const priceB = b?.TourPrices?.length ? 
                b.TourPrices.reduce((min, price) => {
                    return (price.PersonTypeID !== 1 && price.Price && price.Price < min) ? price.Price : min
                }, 999999) : 999999
            return priceB - priceA
        })
    )

    const tourElements = sortedTours.map((tour, index) => {
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
                    <div className="search-package-box">
                        <h3 className="title">ค้นหาแพ็กเกจทัวร์ที่คุณต้องการได้อย่างง่ายดาย✨</h3>
                        <div className="search-box">
                            <div className="img-box">
                                <img src="./images/icons/search.png" alt="" />
                            </div>
                            <input
                                type="text"
                                placeholder="ค้นหาแพ็กเกจ..."
                                onChange={(e) => setSearchText(e.target.value)}
                            />
                        </div>
                        <div className="search-option-box">
                            <div className="option1-box option">
                                <span className="text">แพ็กเกจในจังหวัด</span>
                                <select name="" id="" onChange={(e) => setProvinceID(e.target.value)}>
                                    <option value="">ทุกจังหวัด</option>
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
                                    <input type="date" value={startDate} min={new Date().toISOString().split('T')[0]} onChange={(e) => setStartDate(e.target.value)} />
                                    -
                                    <input type="date" onChange={(e) => setEndDate(e.target.value)} />
                                </div>
                            </div>
                            <div className="option3-box option">
                                <span className="text">ช่วงราคา</span>
                                <div className="input-box">
                                    <input type="number"
                                        min={0}
                                        step={500}
                                        value={minPrice}
                                        onChange={(e) => setMinPrice(Number(e.target.value))}
                                    />
                                    -
                                    <input type="number"
                                        min={1000}
                                        step={500}
                                        value={maxPrice}
                                        onChange={(e) => setMaxPrice(Number(e.target.value))}
                                    />
                                </div>
                            </div>
                            <div className="option4-box option">
                                <span className="text">จัดเรียงตาม</span>
                                <select name="" id="" onChange={(e)=>setSortOption(Number(e.target.value))}>
                                    <option value={1}>ราคาต่ำ-สูง</option>
                                    <option value={2}>ราคาสูง-ต่ำ</option>
                                </select>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="tour-element-container">
                    {tourElements}
                </div>
            </section>
            <Footer/>
        </div>
    )
}
export default TourPackage