export class ApproveBooking {
    jwt: String = '';
    bookingId: String = '';
    
    public constructor(obj?: any) {
      if (obj) {
        this.jwt = obj.jwt;
        this.bookingId = obj.bookingId;
      }
    }
  }