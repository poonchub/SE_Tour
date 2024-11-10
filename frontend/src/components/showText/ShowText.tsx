import { useEffect, useState } from "react";
import "./ShowText.css"

function ShowText() {

    const [slideIndex, setSlideIndex] = useState(0)

    const show_box = document.getElementsByClassName("show-box")

    return (
        <div className="show-text-container">
            
            <div className="text-box">
                <div className="show-box1 show-box">
                    <h2 className="title-show">"1สำรวจเมืองไทยไปกับเรา สัมผัสความงดงามที่รอคุณค้นพบ!"</h2>
                    <p className="des-show">เพลิดเพลินกับทัวร์ในประเทศไทยที่จะพาคุณไปสัมผัสธรรมชาติ วัฒนธรรม และอาหารท้องถิ่นที่ไม่เหมือนใคร เรามีแพ็กเกจทัวร์หลากหลาย ทั้งภูเขา น้ำตก ทะเล และเมืองประวัติศาสตร์ พร้อมพาคุณไปสร้างความทรงจำที่น่าประทับใจในทุกการเดินทาง</p>
                </div>
                <div className="show-box2 show-box">
                    <h2 className="title-show">"2สำรวจเมืองไทยไปกับเรา สัมผัสความงดงามที่รอคุณค้นพบ!"</h2>
                    <p className="des-show">เพลิดเพลินกับทัวร์ในประเทศไทยที่จะพาคุณไปสัมผัสธรรมชาติ วัฒนธรรม และอาหารท้องถิ่นที่ไม่เหมือนใคร เรามีแพ็กเกจทัวร์หลากหลาย ทั้งภูเขา น้ำตก ทะเล และเมืองประวัติศาสตร์ พร้อมพาคุณไปสร้างความทรงจำที่น่าประทับใจในทุกการเดินทาง</p>
                </div>
                <div className="show-box3 show-box">
                    <h2 className="title-show">"3สำรวจเมืองไทยไปกับเรา สัมผัสความงดงามที่รอคุณค้นพบ!"</h2>
                    <p className="des-show">เพลิดเพลินกับทัวร์ในประเทศไทยที่จะพาคุณไปสัมผัสธรรมชาติ วัฒนธรรม และอาหารท้องถิ่นที่ไม่เหมือนใคร เรามีแพ็กเกจทัวร์หลากหลาย ทั้งภูเขา น้ำตก ทะเล และเมืองประวัติศาสตร์ พร้อมพาคุณไปสร้างความทรงจำที่น่าประทับใจในทุกการเดินทาง</p>
                </div>
            </div>
            <div className="dot-box">
                <span className="dot"></span>
                <span className="dot"></span>
                <span className="dot"></span>
            </div>
        </div>
    )
}
export default ShowText