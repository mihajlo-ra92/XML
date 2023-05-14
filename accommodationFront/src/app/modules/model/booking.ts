export class Booking{
    id: String = "";
    accommodationId: String = "";
    guestId: String = "";
    price: number=0;
    priceType: String ="";
    numberOfGuests: number=0;
    bookingType: String="";
    startDate!: Date;
    endDate!: Date;
    
    public constructor(obj?: any) {
        if (obj) {
            this.id = obj.id;
            this.accommodationId = obj.accommodationId;
            this.guestId = obj.guestId;
            this.price = obj.price;
            this.priceType = obj.priceType;
            this.numberOfGuests = obj.numberOfGuests;
            this.bookingType = obj.bookingType;
            this.startDate = obj.startDate;
            this.endDate = obj.endDate;
        }
    }
}