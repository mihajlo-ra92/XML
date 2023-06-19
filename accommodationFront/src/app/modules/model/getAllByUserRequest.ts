import { Accommodation } from "./accommodation";
import { Booking } from "./booking";

export class GetAllByUserRequest{
    id: String = "";
    bookingType: String="";
    
    public constructor(obj?: any) {
        if (obj) {
            this.id = obj.id;
            this.bookingType = obj.bookingType;
        }
    }
}

export class GetAllByUserResponse{
    bookings: Booking[]=[];

    public constructor(obj?: any) {
        if(obj){
            this.bookings = obj.bookings;
        }
    }
}


export class AuthReservationCancelingRequest{
    jwt : string= ""; 
    id : string = "";
    public constructor(obj?: any) {
        if(obj){
            this.jwt = obj.jwt;
            this.id = obj.id;
        }
    }
}

export class AuthReservationCancelingResponse{
    booking :  Booking= new Booking();
    public constructor(obj?: any) {
        if(obj){
            this.booking = obj.booking;
        }
    }

}

export class SearchRequest{
    location :string="";
    guest : number= 1;
    start_date! :Date;
    end_date! :Date;
    min_price : number = 0;
    max_price : number = 0;
    benefits : string[] = []
    is_outstanding : boolean = false;
    public constructor(obj?: any) {
        if(obj){
            this.location = obj.location;
            this.guest = obj.guest;
            this.start_date = obj.start_date;
            this.end_date = obj.end_date;
            this.min_price = obj.min_price;
            this.max_price = obj.max_price;
            this.benefits = obj.benefits;
            this.is_outstanding = obj.is_outstanding;
        }
    }
}

export class AccommodationWithPrice{
    accommodation: Accommodation=new Accommodation();
    price : number=0;
    public constructor(obj?: any) {
        if(obj){
            this.accommodation = obj.accommodations;
            this.price = obj.price;
        }
    }
}

export class SearchResponse{
    accommodations: AccommodationWithPrice[]=[];
    public constructor(obj?: any) {
        if(obj){
            this.accommodations = obj.accommodations;
        }
    }
}