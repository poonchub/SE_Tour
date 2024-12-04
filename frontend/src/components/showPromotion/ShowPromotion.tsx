import { Link } from "react-router-dom"
import "./ShowPromotion.css"

function ShowPromotion(){
    return (
        <div className="show-promotion-container">
            <div className="card-title">
                <div className="title">
                    <div className="icon-box">
                        <img src="/images/icons/promotion.png" alt="" />
                    </div>
                    <div className="text-box">โปรโมชันพิเศษประจำเดือนนี้</div>
                </div>
                <Link to={"/tour-package"}>
                    <button className="more-promotion-btn">{"โปรโมชั่นทั้งหมด>>"}</button>
                </Link>
            </div>
        </div>
    )
}
export default ShowPromotion