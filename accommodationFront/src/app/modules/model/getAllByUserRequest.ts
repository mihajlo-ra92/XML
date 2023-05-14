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