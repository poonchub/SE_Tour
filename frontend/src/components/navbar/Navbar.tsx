import { Link } from 'react-router-dom';
import "./Navbar.css"

function Navbar(){
    return (
        <nav className="navbar">
            <div className="logo-box">
                <h1>My Logo</h1>
            </div>
            <div className="link-page-box">
                <Link to="/" className="link-home link">หน้าหลัก</Link>
                <Link to="/tour-package" className="link-tour-package link">ทัวร์แพ็กเกจ</Link>
                <Link to="/profile" className="link-profile link">โปรไฟล์</Link>
            </div>
        </nav>
    )
}
export default Navbar