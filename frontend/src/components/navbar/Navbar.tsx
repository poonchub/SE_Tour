import { Link } from 'react-router-dom';
import "./Navbar.css"
import { apiUrl } from '../../services/http';

function Navbar(props: { page: string; }) {

    const { page } = props
    const storedCustomer = localStorage.getItem('customer')

    const customer = storedCustomer ? JSON.parse(storedCustomer) : ""

    const isLoggedIn = localStorage.getItem("isLogin") === "true";
    const imageUrl = `${apiUrl}/${customer.ProfilePath}`

    return (
        <nav className="navbar"
            style={{ backgroundColor: page == "home" ? "transparent" : "var(--yellow)" }}
        >
            <div className="logo-box">
                <h1>My Logo</h1>
            </div>
            <div className="link-page-box">
                <Link to="/" className="link-home link">หน้าหลัก</Link>
                <Link to="/tour-package" className="link-tour-package link">ทัวร์แพ็กเกจ</Link>
                <Link to="" className="link-promotion link">โปรโมชั่น</Link>
                <Link to="/profile" className="link-promotion link">โปรไฟล์</Link>
            </div>
            <Link to={isLoggedIn ? "/profile" : "/login-customer"}>
                <div className="login-box">
                    <div className='text'>{customer ? `${customer.UserName}` : "Sign In"}</div>
                    <div className="img-box">
                        <img src={customer ? imageUrl : "./images/icons/log-in.png"} alt="" />
                    </div>
                </div>
            </Link>

        </nav>
    )
}
export default Navbar