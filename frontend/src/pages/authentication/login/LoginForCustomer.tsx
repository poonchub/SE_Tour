import { useState } from "react";
import "./LoginForCustomer.css";
import { message } from "antd";
import { SignInInterface } from "../../../interfaces/ISignIn";
import { GetCustomerByID, SignInForCustomer } from "../../../services/http";

function LoginForCustomer() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [errors, setErrors] = useState<{ username?: string; password?: string }>({});

    const [messageApiLogin, contextHolderLogin] = message.useMessage()
    const [focusedInput, setFocusedInput] = useState<string | null>(null)

    const handleFocus = (inputName: string) => {
        setFocusedInput(inputName)
    };

    const handleBlur = () => {
        setFocusedInput(null)
    }

    async function handleSubmit(e: { preventDefault: () => void }) {
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

                localStorage.setItem('customer', JSON.stringify(resGetCustomer))

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
                        <div className="input-box"
                            style={{
                                border: focusedInput === "input1" ? "2px solid var(--yellow)" : "2px solid transparent"
                            }}
                        >
                            <span>Username</span>
                            <input
                                type="text"
                                className="username-input"
                                value={username}
                                autoComplete="off"
                                onChange={(e) => setUsername(e.target.value)}
                                onFocus={() => handleFocus("input1")}
                                onBlur={handleBlur}
                            />
                            {errors.username && <p className="err-text">{errors.username}</p>}
                        </div>
                        <div className="input-box"
                            style={{
                                border: focusedInput === "input2" ? "2px solid var(--yellow)" : "2px solid transparent"
                            }}
                        >
                            <span>Password</span>
                            <input
                                type="password"
                                className="password-input"
                                value={password}
                                autoComplete="off"
                                onChange={(e) => setPassword(e.target.value)}
                                onFocus={() => handleFocus("input2")}
                                onBlur={handleBlur}
                            />
                            {errors.password && <p className="err-text">{errors.password}</p>}
                        </div>
                        <button className="submit-btn btn" type="submit">
                            Log In
                        </button>
                        <button className="signup-btn btn" type="submit">
                            Sign Up
                        </button>
                    </form>
                </div>
            </div>
        </div>
    );
}

export default LoginForCustomer;
