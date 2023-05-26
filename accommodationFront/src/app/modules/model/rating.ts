export class Rating{
    id: String = "";
    jwt : String = "";
    hostId : String = "";
    accommodationId: String = "";
    guestId: String = "";
    rate: number = 0;
    
    public constructor(obj?: any) {
        if (obj) {
            this.id = obj.id;
            this.hostId = obj.hostId;
            this.accommodationId = obj.accommodationId;
            this.guestId = obj.guestId;
            this.jwt = obj.jwt;
        }
    }
}