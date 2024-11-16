import { ProvincesInterface } from "./IProvinces";
import { TourDescriptionsInterface } from "./ITourDescriptions";
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
	TourPrices?:		TourPricesInterface[];
	TourImages?:		TourImagesInterface[];
	TourDescriptions?:	TourDescriptionsInterface;
}