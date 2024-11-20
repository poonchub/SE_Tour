import Navbar from "../../components/navbar/Navbar"
import "./Home.css"
import ShowText from "../../components/showText/ShowText";
import ShowTourSlice from "../../components/showTourSlice/ShowTourSlice";
import RecPackage from "../../components/recPackage/RecPackage";
import ShowPromotion from "../../components/showPromotion/ShowPromotion";
import Footer from "../../components/footer/Footer";

function Home(){
    return (
        <div className="home-page">
            <Navbar page={"home"}/>
            <div className="full-bg"
                style={{ backgroundImage: `url(./images/backgrounds/background.jpg)` }}
            >
                <ShowText/>
            </div>
            <ShowTourSlice/>
            <section>
                <RecPackage/>
                <ShowPromotion/>
            </section>
            <Footer/>
        </div>
    )
}
export default Home