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