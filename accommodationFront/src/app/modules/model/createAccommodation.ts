import { PriceType } from "./priceType";

export class CreateAccommodation{
    jwt : String = "";
    name : String = "";
    location : String = "";
    benefits: String[] = [];
    pictures: String[] = [];
    minGuests: number = 0;
    maxGuests: number = 0;
    price: number = 0;
    priceType : PriceType = PriceType.Regular;

    public constructor(obj?: any) {
        if (obj) {
            this.jwt = obj.jwt;
            this.name = obj.name;
            this.location = obj.location;
            this.benefits = obj.benefits;
            this.pictures = obj.pictures;
            this.minGuests = obj.minGuest;
            this.maxGuests = obj.maxGuest;
            this.price = obj.price;
        }
    }
}