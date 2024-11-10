import Navbar from "../../components/navbar/Navbar"
import "./Home.css"
import backgroundImage from '../../assets/images/background.jpg';
import ShowText from "../../components/showText/ShowText";

function Home(){
    return (
        <div className="home-page">
            <Navbar/>
            <div className="full-bg"
                style={{ backgroundImage: `url(${backgroundImage})` }}
            >
                <ShowText/>
            </div>
        </div>
    )
}
export default Home