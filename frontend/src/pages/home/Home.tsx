import Navbar from "../../components/navbar/Navbar"
import "./Home.css"
import backgroundImage from '../../assets/images/background/background.jpg';
import ShowText from "../../components/showText/ShowText";
import ShowTourSlice from "../../components/showTourSlice/ShowTourSlice";
import RecPackage from "../../components/recPackage/RecPackage";
import ShowPromotion from "../../components/showPromotion/ShowPromotion";
import Footer from "../../components/footer/Footer";

function Home(){
    return (
        <div className="home-page">
            <Navbar/>
            <div className="full-bg"
                style={{ backgroundImage: `url(${backgroundImage})` }}
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