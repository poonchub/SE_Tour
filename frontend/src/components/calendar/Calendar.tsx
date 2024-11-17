import { useState } from "react"
import "./Calendar.css"

function Calendar(props: { dateTime: any; setDateSelectedFormat: any; }) {
    const { dateTime, setDateSelectedFormat } = props
    const [currentDate, setCurrentDate] = useState(new Date());

    const [dateSelected, setDateSelected] = useState<string | null>();

    const changeMonth = (offset: number) => {
        const newDate = new Date(currentDate.setMonth(currentDate.getMonth() + offset));
        setCurrentDate(newDate);
    };

    const getCalendarDays = () => {

        const startOfMonth = new Date(currentDate.getFullYear(), currentDate.getMonth(), 1);
        const endOfMonth = new Date(currentDate.getFullYear(), currentDate.getMonth() + 1, 0);
        const daysInMonth = endOfMonth.getDate();
        const startDayOfWeek = startOfMonth.getDay();

        const weeks: JSX.Element[][] = [];
        let daysArray: JSX.Element[] = [];

        for (let i = 0; i < startDayOfWeek; i++) {
            daysArray.push(<td key={`empty-${i}`}></td>);
        }

        for (let day = 1; day <= daysInMonth; day++) {
            const dateStr = `${currentDate.getFullYear()}-${String(currentDate.getMonth() + 1).padStart(2, '0')}-${String(day).padStart(2, '0')}`;

            let isAvailable = false
            let index = 0
            for (let i = 0 ; i < dateTime.length ; i++){
                isAvailable = dateTime[i]?.[0]===dateStr ? true : false
                if (isAvailable){
                    index = i
                    break
                }
            }
            
            const endTime = dateTime[index]?.[1].slice(8, 10);
            const endTimeFormat = endTime?.startsWith("0") ? endTime?.slice(1, 2) : endTime;

            const dateObj = new Date(currentDate.getFullYear(), currentDate.getMonth(), day);

            const months = [
                "มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน",
                "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม"
            ];

            const dateStrFormat = `วันที่ ${dateObj.getDate()} - ${endTimeFormat} ${months[dateObj.getMonth()]} ${dateObj.getFullYear() + 543}`;

            daysArray.push(
                <td
                    key={day}
                    className={`calendar-day ${isAvailable ? 'available' : ''} ${dateSelected === dateStr ? 'selected' : ''}`}
                    onClick={() => {
                        if (isAvailable) {
                            setDateSelectedFormat(dateStrFormat);
                            setDateSelected(dateStr);
                        }
                    }}
                >
                    {day}
                </td>
            );

            if ((daysArray.length) % 7 === 0 || day === daysInMonth) {
                weeks.push([...daysArray]);
                daysArray = [];
            }
        }

        return weeks.map((week, index) => <tr key={`week-${index}`}>{week}</tr>);
    };

    const dayOfWeeks = ["อาทิตย์", "จันทร์", "อังคาร", "พุธ", "พฤหัส", "ศุกร์", "เสาร์"]

    return (
        <div className="calendar-container">
            <div className="calendar-header">
                <button className="prev-btn btn" onClick={() => changeMonth(-1)}>«</button>
                <span>
                    {currentDate.toLocaleString('default', { month: 'long' })} {currentDate.getFullYear()+543}
                </span>
                <button className="next-btn btn" onClick={() => changeMonth(1)}>»</button>
            </div>
            <div className="date-box">
                <table className="calendar-table">
                    <thead>
                        <tr>
                            {dayOfWeeks.map((day, index) => <th key={index}>{day}</th>)}
                        </tr>
                    </thead>
                    <tbody>
                        {getCalendarDays()}
                    </tbody>
                </table>
            </div>
            
        </div>
    );
}
export default Calendar