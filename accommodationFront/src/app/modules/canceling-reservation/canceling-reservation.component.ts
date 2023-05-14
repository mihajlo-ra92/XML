import { Component, OnInit } from '@angular/core';
import { Booking } from '../model/booking';
import { CancelingReservationService } from '../service/canceling-reservation.service';
import { ActivatedRoute, Router } from '@angular/router';
import { GetAllByUserRequest } from '../model/getAllByUserRequest';
import { AuthReservationCancelingRequest } from "../model/getAllByUserRequest"

@Component({
  selector: 'app-canceling-reservation',
  templateUrl: './canceling-reservation.component.html',
  styleUrls: ['./canceling-reservation.component.css']
})
export class CancelingReservationComponent implements OnInit {
  public request: GetAllByUserRequest= new GetAllByUserRequest()
  public bookings : Booking[] =[];
  public canceling: AuthReservationCancelingRequest = new AuthReservationCancelingRequest

  
  constructor(private cancelingService: CancelingReservationService, private route: ActivatedRoute, private router: Router) { }

  ngOnInit(): void {
  
      this.cancelingService.getAllReservationByUserId({
        "id" : localStorage.getItem("loggedUserId"),
        "bookingType": "Reserved"
    }).subscribe((res) => {
      this.bookings = res.bookings;
      console.log(this.bookings)
      });
  }
  onButtonClick(item: any) {
    if(localStorage.getItem("loggedUserType") == "0"){
      this.cancelingService.AuthReservationCanceling({"jwt":localStorage.getItem("token"),"id" : item.id}).subscribe((res) => {
      alert("Successfull");
      //this.router.navigate(['/landing-page'])
      
      this.cancelingService.getAllReservationByUserId({
        "id" : localStorage.getItem("loggedUserId"),
        "bookingType": "Reserved"
    }).subscribe((res) => {
      this.bookings = res.bookings;
      console.log(this.bookings)
      });
    },(error) =>{
      alert("Reservation is start")
    });  

  }
  }
}
