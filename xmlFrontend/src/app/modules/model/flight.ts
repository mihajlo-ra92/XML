export class Flight{
    flightId: String = "";
    date!: string;
    endPlace: String = "";
    startPlace: String = "";
    capacity: number = 0;
    price: number = 0;
    freeSeats: number = 0;
    
    public constructor(obj?: any) {
        if (obj) {
            this.flightId = obj.FlightId;
            this.date = obj.Date;
            this.endPlace = obj.EndPlace;
            this.startPlace = obj.StartPlace;
            this.capacity = obj.Capacity;
            this.price = obj.Price;
            this.freeSeats = obj.FreeSeats;
        }
    }
}