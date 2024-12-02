import { Link } from "react-router-dom"
import "./RecPackage.css"

function RecPackage() {
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

        </div>
    )
}
export default RecPackage