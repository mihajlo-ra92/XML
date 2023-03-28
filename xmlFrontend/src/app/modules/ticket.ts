export class Ticket{
    id: number = 0;
    date!: Date;
    endPlace: String = "";
    startPlace: String = "";
    capacity: number = 0;
    price: number = 0;
    flightId: String = "";
    userId: string = "";
    
    public constructor(obj?: any) {
        if (obj) {
            this.id = obj.id;
            this.date = obj.date;
            this.endPlace = obj.endPlace;
            this.startPlace = obj.startPlace;
            this.capacity = obj.capacity;
            this.price = obj.price;
            this.flightId = obj.flightId;
            this.userId = obj.userId;
        }
    }
}

