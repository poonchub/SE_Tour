import Navbar from "../../components/navbar/Navbar";
import "./Home.css";
import ShowTourSlice from "../../components/showTourSlice/ShowTourSlice";
import RecPackage from "../../components/recPackage/RecPackage";
import ShowPromotion from "../../components/showPromotion/ShowPromotion";
import Footer from "../../components/footer/Footer";
import { Carousel } from "antd";

const texts = [
    {
        title: 'เตรียมตัวให้พร้อม แล้วออกเดินทางกันเลย!',
        description: 'อย่ารอช้า! ทริปดี ๆ รอคุณอยู่ เริ่มต้นการเดินทางในฝันของคุณได้ที่นี่ แล้วพบกันในจุดหมายปลายทาง!'
    },
    {
        title: ' เที่ยวสนุก ทุกจุดหมายปลายทาง',
        description: 'ไม่ว่าคุณจะรักทะเล ภูเขา หรือเมืองใหญ่ เราพร้อมพาคุณไปพบกับความสุขในทุกสถานที่ที่คุณเลือก'
    },
    {
        title: 'เที่ยวได้ทุกฤดู ตอบโจทย์ทุกสไตล์',
        description: 'ไม่ว่าฤดูไหนก็เที่ยวสนุกได้ กับแพ็คเกจที่หลากหลาย พร้อมแนะนำจุดหมายปลายทางที่เหมาะกับคุณที่สุด'
    },
    {
        title: 'ประสบการณ์พรีเมียม ในราคาที่คุณเอื้อมถึง',
        description: 'สัมผัสบริการระดับพรีเมียม ทั้งที่พัก อาหาร และไกด์นำเที่ยว ในราคาสุดพิเศษที่คุณจะประทับใจ'
    },
]

function Home() {
    const textElement = texts.map((text, index) => {
        return (
            <div key={index} className="text-box">
                <h3 className="text-title">{text.title}</h3>
                <p className="text-description">{text.description}</p>
            </div>
        )
    })

    return (
        <div className="home-page">
            <Navbar page="home" />
            <div className="full-bg">
                <div className="show-text-container">
                    <Carousel autoplay autoplaySpeed={5000} speed={2000}>
                        {textElement}
                    </Carousel>
                </div>
            </div>
            <ShowTourSlice />
            <section>
                <RecPackage />
                <ShowPromotion />
            </section>
            <Footer />
        </div>
    );
}

export default Home;
