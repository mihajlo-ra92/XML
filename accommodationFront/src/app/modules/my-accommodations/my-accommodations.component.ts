import { Component, OnInit } from '@angular/core';
import { AccommodationService } from '../service/accommodation.service';
import { Route, Router } from '@angular/router';
import { Accommodation } from '../model/accommodation';
import { Booking } from '../model/booking';
import { CustomPrice } from '../model/CustomPrice';

@Component({
  selector: 'app-my-accommodations',
  templateUrl: './my-accommodations.component.html',
  styleUrls: ['./my-accommodations.component.css']
})
export class MyAccommodationsComponent implements OnInit {
  public accommodations: Accommodation[] = [];
  public selected : boolean = false;
  public selectedCustomPrice : boolean = false;
  public selectedAccommodation: Accommodation = new Accommodation();
  public pictures: Document = new Document;
  public start : String = "";
  public end : String = "";
  public bookingsForAccommodation: Booking[] = new Array;
  public customPrice: CustomPrice = new CustomPrice();
  constructor(private accommodationService : AccommodationService, private router: Router) { }

  ngOnInit(): void {
    let token =localStorage.getItem("token");
    if(token !== null){
      this.accommodationService.getMyAccommodation(token).subscribe((res) =>{
        this.accommodations = res.accommodations;
        console.log(this.accommodations);
      })
    }
    let temp = localStorage.getItem("loggedUserType");
    if(temp != "1"){
    this.router.navigate(['/landing-page']);    
    }
  }

  select(accommodation: Accommodation){
    this.selectedCustomPrice = false;
    this.selected = true;
    this.selectedAccommodation = accommodation;
    let jwt = localStorage.getItem("token")
    if(jwt !== null){
      this.accommodationService.getBookibgByAccommodationId(jwt,accommodation.id.toString()).subscribe((res) =>{
        console.log(res);
        this.bookingsForAccommodation = res.bookings;
      })
    }
  }
  unselect(){
    this.selected = false;
    this.selectedAccommodation = new Accommodation();
    this.pictures  = new Document;
    this.bookingsForAccommodation = [];
    this.selectedCustomPrice = false;
  }

  selectCustomPrice(){
   this.selectedCustomPrice = true;
  }

  addCustomPrice(){
    this.customPrice.accommodationId = this.selectedAccommodation.id;
    let temp = localStorage.getItem("token")
    if(temp!==null){
      this.customPrice.jwt = temp;
    }

    const startDateString = new Date(
      String(this.start) + 'T12:00:42.123Z'
    ).toISOString();

    const endDateString = new Date(
      String(this.end) + 'T12:00:42.123Z'
      ).toISOString();

      this.customPrice.start_date = startDateString;
      this.customPrice.end_date = endDateString;
    this.customPrice.priceType = "Regular";
    console.log(this.customPrice);
    this.accommodationService.defineCustomPrice(this.customPrice).subscribe((res) =>{
      console.log(res);
    })
    this.start = "";
    this.end = "";
    this.customPrice = new CustomPrice();
    this.selectedCustomPrice = false;
  }
}
