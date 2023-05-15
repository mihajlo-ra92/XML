export class CustomPrice{
    jwt: String = "";
    accommodationId: String = "";
    start_date: String = "";
    end_date: String = "";
    price: number = 0;
    priceType: String ="";
    
    public constructor(obj?: any) {
        if (obj) {
            this.jwt = obj.jwt;
            this.accommodationId = obj.accommodationId;
            this.start_date = obj.start_date;
            this.end_date = obj.end_date;
            this.price = obj.price;
            this.priceType = obj.priceType;
        }
    }
}