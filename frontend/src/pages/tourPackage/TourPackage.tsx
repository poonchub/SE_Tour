import { useEffect, useState } from "react";
import Navbar from "../../components/navbar/Navbar"
import "./TourPackage.css"
import { GetTourPackages } from "../../services/http";
import { TourPackagesInterface } from "../../interfaces/ITourPackages";
import PackageItem from "../../components/packageItem/PackageItem";

function TourPackage(){

    const [tourPackages, setTourPackages] = useState<TourPackagesInterface[]>([]);

    async function getTourPackages(){
        let res = await GetTourPackages()
        if (res) {
            setTourPackages(res);
        }
    }

    useEffect(()=> {
        getTourPackages()
    }, [])

    const tourElements = tourPackages.map((tour, index) => {
        return <PackageItem key={index} tour={tour}/>
    })

    return (
        <div className="tour-pavkage-page">
            <Navbar page={"tourPackage"}/>
            <section>
                <div className="pic-slice">
                    <img src="./images/sliceshow/pic1.jpg" alt="" />
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
                            <input type="text" placeholder="ค้นหาแพ็กเกจ..."/>
                        </div>
                        <div className="search-option-box">
                            <div className="option1-box option">
                                <span className="text">แพ็กเกจในจังหวัด</span>
                                <select name="" id="">
                                    <option value="">ระนอง</option>
                                    <option value="">สกลนคร</option>
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
                                    <input type="number" min={0} step={500} defaultValue={0}/>
                                    -
                                    <input type="number" min={1000} step={500} defaultValue={1000}/>   
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