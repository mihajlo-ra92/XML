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
    guest : number= 0;
    start_date! :Date;
    end_date! :Date;
    public constructor(obj?: any) {
        if(obj){
            this.location = obj.location;
            this.guest = obj.guest;
            this.start_date = obj.start_date;
            this.end_date = obj.end_date;
        }
    }
}

export class SearchResponse{
    accommodations: Accommodation[]=[];
    public constructor(obj?: any) {
        if(obj){
            this.accommodations = obj.accommodations;
        }
    }
}