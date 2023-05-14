import { Component, OnInit } from '@angular/core';
import { BookingService } from '../service/booking.service';

@Component({
  selector: 'app-approve-booking',
  templateUrl: './approve-booking.html',
})
export class ApproveBookingComponent implements OnInit {

  constructor(private bookginService: BookingService) {}

    ngOnInit(): void {

      }
    ApproveBooking(){
        
    }

}
