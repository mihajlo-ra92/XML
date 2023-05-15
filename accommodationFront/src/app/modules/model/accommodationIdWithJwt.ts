export class AccommodationIdWithJwt{
    jwt : String = "";
    accommodationId: String = "";
    
    public constructor(obj?: any) {
        if (obj) {
            this.jwt = obj.jwt;
            this.accommodationId = obj.accommodationId;
        }
    }
}