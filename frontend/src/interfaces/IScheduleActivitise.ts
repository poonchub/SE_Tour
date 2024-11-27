import { ActivitiesInterface } from "./IActivities";

export interface ScheduleActivities{
    ID?:    number;
    DateTime?:  string;
    ActivityID?:    number;
    TourScheduleID?:    number;
    Activity?:  ActivitiesInterface;
}