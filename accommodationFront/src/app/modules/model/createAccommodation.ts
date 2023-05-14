export class CreateAccommodation{
    jwt : String = "";
    name : String = "";
    location : String = "";
    benefits: String[] = [];
    pictures: String[] = [];
    minGuest: number = 0;
    maxGuest: number = 0;

    public constructor(obj?: any) {
        if (obj) {
            this.jwt = obj.jwt;
            this.name = obj.name;
            this.location = obj.location;
            this.benefits = obj.benefits;
            this.pictures = obj.pictures;
            this.minGuest = obj.minGuest;
            this.maxGuest = obj.maxGuest;
        }
    }
}