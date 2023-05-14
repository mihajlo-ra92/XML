import { Component, OnInit } from '@angular/core';
import { Booking } from '../model/booking';
import { CancelingReservationService } from '../service/canceling-reservation.service';
import { ActivatedRoute, Router } from '@angular/router';
import { GetAllByUserRequest } from '../model/getAllByUserRequest';

@Component({
  selector: 'app-canceling-reservation',
  templateUrl: './canceling-reservation.component.html',
  styleUrls: ['./canceling-reservation.component.css']
})
export class CancelingReservationComponent implements OnInit {
  public request: GetAllByUserRequest= new GetAllByUserRequest()
  public bookings : Booking[] =[];
  public displayedColumns = ['number', 'floor'];

  
  constructor(private cancelingService: CancelingReservationService, private route: ActivatedRoute, private router: Router) { }

  ngOnInit(): void {
    let pom = new Booking()
    pom.guestId="1"
    pom.numberOfGuests=5
    pom.price = 50
    this.bookings.push(pom)
    this.request.bookingType = "Booked"
      this.request.id ="guest1Id"
      this.cancelingService.getAllReservationByUserId(this.request).subscribe(res => {

        console.log(this.bookings)
      });
  }
  onButtonClick(item: any) {
    this.bookings = this.bookings.filter(booking => booking !== item);
  }
}
