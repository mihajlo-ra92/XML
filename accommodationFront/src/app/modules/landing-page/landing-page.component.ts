import { Component, OnInit } from '@angular/core';
import { Accommodation } from '../model/accommodation';
import { AccommodationService} from '../service/accommodation.service';
import { BookingService } from '../service/booking.service'
import { Reservation } from '../model/reservation';
import { SearchRequest } from '../model/getAllByUserRequest';

@Component({
  selector: 'app-landing-page',
  templateUrl: './landing-page.component.html',
  styleUrls: ['./landing-page.component.css'],
})
export class LandingPageComponent implements OnInit {
  allAccommodations: Array<Accommodation> = new Array();
  isAdmin = false;
  loggedUserRole = localStorage.getItem('loggedUserType')
  userJwt = localStorage.getItem('token')
  reservation: Reservation = new Reservation
  request: SearchRequest=new SearchRequest()

  constructor(private accommodationService: AccommodationService, private bookingService: BookingService) {}

  ngOnInit(): void {

    this.AllAccommodations();
  }

  // allUsers() {
  //   var loggedUserType = localStorage.getItem('loggedUserType');
  //   console.log(loggedUserType);

  //   if (loggedUserType?.toString() === 'admin') {
  //     console.log(this.isAdmin);

  //     this.isAdmin = true;
  //   }
  //   this.flightService.getAllFlights().subscribe((res) => {
  //     let resJSON = JSON.parse(res);
  //     this.allFlights = resJSON;
  //     // this.allFlights.map((x) => {
  //     //   const myDate = new Date(x.date);
  //     //   // myDate.setHours(myDate.getHours() + 2)
  //     //   console.log(myDate);
  //     //   x.date = myDate.toLocaleString('en-US', {
  //     //     timeZone: 'America/New_York',
  //     //   });
  //     //   console.log(x.date);
  //     // });
  //     console.log(this.allFlights);
  //   });
  // }
AllAccommodations() {
  this.accommodationService.getAllAccommodations().subscribe((res) => {
    let resJSON = JSON.parse(res);
    this.allAccommodations = resJSON.accommodations;
    // this.allFlights.map((x) => {
    //   const myDate = new Date(x.date);
    //   // myDate.setHours(myDate.getHours() + 2)
    //   console.log(myDate);
    //   x.date = myDate.toLocaleString('en-US', {
    //     timeZone: 'America/New_York',
    //   });
    //   console.log(x.date);
    // });
    console.log(this.allAccommodations);
  });
}

ReserveAccommodation(accommodation: Accommodation){

  localStorage.setItem('accommodationPrice', accommodation.price.toString())
  localStorage.setItem('accommodationId', accommodation.id.toString())

  console.log(accommodation)
  console.log(localStorage.getItem('accommodationId'))
  console.log(localStorage.getItem('accommodationPrice'))
  
  window.location.href = '/accommodation-reservation'

}
Refrash(){
  this.AllAccommodations();
  this.request = new SearchRequest()
}
Search(){
  console.log(this.request.end_date)
  if(this.request.guest===1 && this.request.location ==="" && this.request.start_date === undefined && this.request.end_date === undefined ){
    this.accommodationService.getAllAccommodations().subscribe((res) => {
      let resJSON = JSON.parse(res);
      this.allAccommodations = resJSON.accommodations;
      // this.allFlights.map((x) => {
      //   const myDate = new Date(x.date);
      //   // myDate.setHours(myDate.getHours() + 2)
      //   console.log(myDate);
      //   x.date = myDate.toLocaleString('en-US', {
      //     timeZone: 'America/New_York',
      //   });
      //   console.log(x.date);
      // });
      console.log(this.allAccommodations);
      
    });
    return
  }
  if(this.request.start_date !== undefined && this.request.end_date !== undefined ){
    if(this.request.end_date < this.request.start_date){return}
    const currentDate = new Date();
    if(currentDate < this.request.start_date){return}
    
  }

  if(this.request.start_date === undefined && this.request.end_date === undefined ){
    this.accommodationService.searchAccommodation(this.request).subscribe((res) => {
      this.allAccommodations = []
      console.log(this.allAccommodations);
  
      var list = res.accommodations;
      list.forEach((item) => {
        console.log(item);
        console.log(item.accommodation, "Ovde b");
  
        this.allAccommodations.push(item.accommodation)
      });
      
      // this.allFlights.map((x) => {
      //   const myDate = new Date(x.date);
      //   // myDate.setHours(myDate.getHours() + 2)
      //   console.log(myDate);
      //   x.date = myDate.toLocaleString('en-US', {
      //     timeZone: 'America/New_York',
      //   });
      //   console.log(x.date);
      // });
      console.log(this.allAccommodations);
    });}



  if(this.request.start_date === undefined || this.request.end_date === undefined ){
    return
  }
  var startDateString = this.request.start_date.toString();
  var startDateFormated = startDateString + 'T00:00:00.123Z';

  var endDateString = this.request.end_date.toString();
  var endDateFormated = endDateString + 'T00:00:00.123Z';

  this.accommodationService.searchAccommodation({
    "location" : this.request.location,
    "guest" : this.request.guest,
    "start_date" :startDateFormated,
    "end_date" :endDateFormated
}).subscribe((res) => {
    this.allAccommodations = []
    console.log(this.allAccommodations);

    var list = res.accommodations;
    list.forEach((item) => {
      console.log(item);
      console.log(item.accommodation, "Ovde b");

      this.allAccommodations.push(item.accommodation)
    });
    
    // this.allFlights.map((x) => {
    //   const myDate = new Date(x.date);
    //   // myDate.setHours(myDate.getHours() + 2)
    //   console.log(myDate);
    //   x.date = myDate.toLocaleString('en-US', {
    //     timeZone: 'America/New_York',
    //   });
    //   console.log(x.date);
    // });
    console.log(this.allAccommodations);
  });
}
  // deleteFlight(flight: Flight) {
  //   console.log(flight.id);
  //   this.flightService.deleteFlight(flight.id).subscribe((res) => {
  //     console.log(res);
  //     this.allUsers();
  //     window.location.href = '/';
  //   });
  }


//   search(){
//     if(this.startDate != undefined && this.endDate != undefined){
      
//       var startDateString = this.startDate.toString();
//       var startDateFormated = startDateString + 'T00:00:00.123Z';

//       var endDateString = this.endDate.toString0.();
//       var endDateFormated = endDateString + 'T00:00:00.123Z';


//       this.flightService
//         .searchFlights(
//           this.startPlace,
//           this.endPlace,
//           startDateFormated,
//           endDateFormated,
//           this.quantity.toString()
//         )
//         .subscribe((res) => {
//           console.log(res);
//           if (res !== null) {
//             try{
//             let resJSON = JSON.parse(res);
//             // let resJSON = res;
//             this.allFlights = resJSON;
//             // this.allFlights.map((x) => {
//             //   const myDate = new Date(x.date);
//             //   console.log(myDate);
//             //   x.date = myDate.toLocaleString('en-US', {
//             //     timeZone: 'America/New_York',
//             //   });
//             //   console.log(x.date);
//             // });
//             }
//             catch{
//               alert("No flights")
//             }
//           } else {
//             console.log('res je null');

//           }
          
//         });
//     }
//   }

//   switchPlaces() {
//     let newStartPlace: string;
//     newStartPlace = this.endPlace;
//     this.endPlace = this.startPlace;
//     this.startPlace = newStartPlace;
//   }
// }
