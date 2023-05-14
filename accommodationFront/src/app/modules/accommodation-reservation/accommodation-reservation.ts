import { Component, OnInit } from '@angular/core';
import { Reservation } from '../model/reservation';
import { AccommodationService } from '../service/accommodation.service';

@Component({
  selector: 'app-accommodation-reservation',
  templateUrl: './accommodation-reservation.html',
})
export class AccommodationReserationComponent implements OnInit {

    reservation: Reservation = new Reservation()
    start: string = "";
    end: string = "";
    loggedUserRole = localStorage.getItem('loggedUserType')

  constructor(private accommodationService: AccommodationService) {}

    ngOnInit(): void {

      }

    Reserve(){

       this.reservation.price = Number(localStorage.getItem('accommodationPrice'))
       this.reservation.accommodationId = localStorage.getItem('accommodationId')!
       this.reservation.jwt = localStorage.getItem('token')!
       this.reservation.bookingType = 1
       this.reservation.numberOfGuests = Number(this.reservation.numberOfGuests)

    const startDateString = new Date(
        String(this.start) + 'T12:00:42.123Z'
      ).toISOString();

    const endDateString = new Date(
    String(this.end) + 'T12:00:42.123Z'
    ).toISOString();

    this.reservation.start_date = startDateString
    this.reservation.end_date = endDateString

        console.log(this.reservation)
        console.log(localStorage.getItem('accommodationId'))
        console.log(localStorage.getItem('accommodationPrice'))
        console.log(localStorage.getItem('token'))
       this.accommodationService.reserve(this.reservation).subscribe((res) => {

        console.log(res);
        console.log(this.reservation);

      alert("Successfully reserved from" + " " + this.start + " " + "to" + " " + this.end)

      });
    }
}
