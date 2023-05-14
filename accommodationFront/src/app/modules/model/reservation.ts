export class Reservation{
    jwt: String = "";
    accommodationId : String = "";
    price: number = 0;
    priceType: number = 0;
    numberOfGuests: number = 0;
    bookingType: number = 1;
    start_date: String = "";
    end_date: String = "";
    
    public constructor(obj?: any) {
        if (obj) {
            this.jwt = obj.jwt;
            this.accommodationId = obj.accommodationId;
            this.price = obj.price;
            this.priceType = obj.priceType;
            this.numberOfGuests = obj.numberOfGuests;
            this.bookingType = obj.bookingType;
            this.start_date = obj.start_date;
            this.end_date = obj.end_date;
        }
    }
}