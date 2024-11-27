import { useState } from "react";
import "./LoginForCustomer.css";
import { message } from "antd";
import { SignInInterface } from "../../../interfaces/ISignIn";
import { GetCustomerByID, SignInForCustomer } from "../../../services/http";

function LoginForCustomer() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [errors, setErrors] = useState<{ username?: string; password?: string }>({});

    const [messageApiLogin, contextHolderLogin] = message.useMessage();

    async function handleSubmit(e: { preventDefault: () => void }){
        e.preventDefault();
        let hasError = false;
        const newErrors: { username?: string; password?: string } = {};
        if (!username.trim()) {
            newErrors.username = "กรุณากรอกชื่อผู้ใช้";
            hasError = true;
        }
        if (!password.trim()) {
            newErrors.password = "กรุณากรอกรหัสผ่าน";
            hasError = true;
        }
        setErrors(newErrors);
        if (!hasError) {
            const data: SignInInterface = {
                Username: username,
                Password: password,
            };
            let resSignin = await SignInForCustomer(data);
            if (resSignin) {
                messageApiLogin.success("Sign-in successful");
                localStorage.setItem("isLogin", "true");
                localStorage.setItem("token_type", resSignin.token_type);
                localStorage.setItem("token", resSignin.token);
                localStorage.setItem("id", resSignin.id);

                let resGetCustomer = await GetCustomerByID(resSignin.id);

                localStorage.setItem("firstName", resGetCustomer.FirstName);
                localStorage.setItem("lastName", resGetCustomer.LastName);
                localStorage.setItem("profilePath", resGetCustomer.ProfilePath);

                setTimeout(() => {
                    location.href = "/";
                }, 2000);
            } else {
                messageApiLogin.error("Email or Password is Incorrect");
            }
        }
    };

    return (
        <div className="login-for-customer-page">
            {contextHolderLogin}
            <video className="bg-video" autoPlay muted loop>
                <source src="./images/backgrounds/beach.mp4" type="video/mp4" />
            </video>
            <div className="form-container">
                <div className="box">
                    <form className="login-form" onSubmit={handleSubmit}>
                        <span className="title">Sign In</span>
                        <div className="input-box">
                            <span>Username</span>
                            <input
                                type="text"
                                className="username-input"
                                value={username}
                                autoComplete="off"
                                onChange={(e) => setUsername(e.target.value)}
                            />
                            {errors.username && <p className="err-text">{errors.username}</p>}
                        </div>
                        <div className="input-box">
                            <span>Password</span>
                            <input
                                type="password"
                                className="password-input"
                                value={password}
                                autoComplete="off"
                                onChange={(e) => setPassword(e.target.value)}
                            />
                            {errors.password && <p className="err-text">{errors.password}</p>}
                        </div>
                        <button className="submit-btn" type="submit">
                            Login
                        </button>
                    </form>
                </div>
            </div>
        </div>
    );
}

export default LoginForCustomer;
