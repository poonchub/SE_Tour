import "./Navbar.css"

function Navbar(){
    return (
        <nav className="navbar">
            <div className="logo-box">
                <h1>My Logo</h1>
            </div>
            <div className="link-page-box">
                <a href="" className="link-home link">หน้าหลัก</a>
                <a href="" className="link-tour-package link">ทัวร์แพ็กเกจ</a>
                <a href="" className="link-profile link">โปรไฟล์</a>
            </div>
        </nav>
    )
}
export default Navbar