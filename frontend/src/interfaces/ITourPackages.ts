import { ProvincesInterface } from "./IProvinces";
import { TourImagesInterface } from "./ITourImages";
import { TourPricesInterface } from "./ITourPrices";

export interface TourPackagesInterface {
    ID?:			number;
	PackageCode?:	string;
	TourName?:		string;
	Description?:	string;
	Duration?:		string;
	ProvinceID?:	number;
	Province?:		ProvincesInterface;
	TourPrice?:		TourPricesInterface;
	TourImages?:		TourImagesInterface;
}