import { useState } from "react";
import "./PackageItem.css";
import { Link } from "react-router-dom";
import { apiUrl } from "../../services/http";

function PackageItem(props: { tour: any }) {
    const { tour } = props;

    const [startPrice, setStartPrice] = useState(999999);

    const imageUrl = `${apiUrl}/${tour.TourImages[0].FilePath}`

    tour.TourPrices.forEach((price: { Price: any; "": any }) => {
        if (price.Price < startPrice){
            setStartPrice(price.Price.toLocaleString('th-TH', {
                minimumFractionDigits: 2,
                maximumFractionDigits: 2,
            }))
        }
    });

    return (
        <div className="package-item-container">
            <div className="title-bg"></div>
            <div className="detail-box">
                <div className="title"><h3>{tour.TourName}</h3></div>
                <p className="description">{tour.Description}</p>
                <div className="duration-province-box">
                    <div className="sub-box">
                        <div className="icon-box">
                            <img src="./images/icons/calendar.png" alt="" />
                        </div>
                        <p className="detail">{tour.Duration}</p>
                    </div>
                    <div className="sub-box">
                        <div className="icon-box">
                            <img src="./images/icons/map.png" alt="" />
                        </div>
                        <p className="detail">{tour.Province?.ProvinceName}</p>
                    </div>
                </div>
                <div className="price-tourDetail-box">
                    <div className="sub-box">
                        <span className="title">ราคาเริ่มต้น</span>
                        <span className="price">฿{startPrice}</span>
                    </div>
                    <div className="sub-box">
                        <Link to="">
                            <button className="btn-detail">รายละเอียดทัวร์</button>
                        </Link>
                    </div>
                </div>
            </div>
            <div className="image-box">
                <img src={imageUrl} alt="" />
            </div>
        </div>
    );
}
export default PackageItem;
