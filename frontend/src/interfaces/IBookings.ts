import { TourSchedulesInterface } from "./ITourSchedules";

export interface BookingsInterface {
    ID?:            number;
    TotalPrice?:    number;
    CustomerID?:    number;
    TourScheduleID?:    number;
    BookingStatusID?:   number;
    PromotionID?:   number;
    TourSchedule?:  TourSchedulesInterface;
}