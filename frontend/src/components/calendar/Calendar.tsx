import { useState } from "react"
import "./Calendar.css"

function Calendar() {
    const [date, setDate] = useState<string>("พฤษจิกายน 2567")

    return (
        <div className="calendar-container">
            <span className="calendar-title">{date}</span>
            <div className="day-box">
                

            </div>
        </div>
    )
}
export default Calendar