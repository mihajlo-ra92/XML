export class AccommodationWithRate{
    id: String = "";
    jwt : String = "";
    hostId : string = "";
    name: String = "";
    location: String = "";
    benefits: String[] = [];
    pictures: String[] = [];
    minGuests: number = 0;
    maxGuests: number = 0;
    price: number = 0;
    rates: number[] = [];
    
    public constructor(obj?: any) {
        if (obj) {
            this.id = obj.id;
            this.hostId = obj.hostId;
            this.name = obj.name;
            this.location = obj.location;
            this.benefits = obj.benefits;
            this.price = obj.Price;
            this.pictures = obj.pictures;
            this.minGuests = obj.minGuests;
            this.maxGuests = obj.maxGuests;
            this.jwt = obj.jwt;
        }
    }
}