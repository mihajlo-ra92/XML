import { Component, OnInit } from '@angular/core';
import { Accommodation } from '../model/accommodation';
import { AccommodationService} from '../service/accommodation.service';
import { BookingService } from '../service/booking.service'
import { Reservation } from '../model/reservation';
import { UserService } from '../service/user.service';
import { RatingService } from '../service/rating.service';
import { Rating } from '../model/rating';
import { AccommodationWithRate } from '../model/accommodationWithRate';
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
  selectedAccommodation : Accommodation = new Accommodation
  rateSelectedAccommodation : number = 0
  rateSelectedHost : number = 0
  accommodationWithRate : AccommodationWithRate = new AccommodationWithRate
  allAccommodationsWithRate: Array<AccommodationWithRate> = new Array();
  request: SearchRequest=new SearchRequest()

  constructor(private accommodationService: AccommodationService, private ratingService: RatingService) {}


  ngOnInit(): void {

    this.AllAccommodations();
  }

  rateAccommodation(id: String) {

    let rating : Rating = new Rating
    rating.jwt = localStorage.getItem('token')!
    rating.accommodationId = id
    rating.rate = this.rateSelectedAccommodation

    if(this.rateSelectedAccommodation != 0){
        this.ratingService.createRating(rating).subscribe((res) => {
          let resJSON = JSON.parse(res);
          console.log(resJSON)
        },
        (error) => {
          // alert(error.error.message)
          let errorJSON = JSON.parse(error.error);
          alert(errorJSON.message)
        });
        alert("Successfully rated this accommodation")
    }
    else{
      alert("You didn't rate this accommodation")

    }
  }

  deleteRateAccommodation(accommodationId: String){

    let rating : Rating = new Rating
    let guestJwt = localStorage.getItem('token')! 
    console.log(guestJwt)
    console.log(accommodationId)


    this.ratingService.getRatingByAccommodationAndGuestId(guestJwt, accommodationId).subscribe((res) => {
      let resJSON = JSON.parse(res);
      rating = resJSON.rating
      console.log(resJSON.rating)
      console.log(resJSON.rating.id)

        this.ratingService.deleteRating(guestJwt, rating.id).subscribe((res) => {
          let resJSON = JSON.parse(res);
          console.log(resJSON)

          alert("Successfully deleted this rate")
          
        },
        (error) => {
          // alert(error.error.message)
          let errorJSON = JSON.parse(error.error);
          alert(errorJSON.message)
        });
    },
    (error) => {
      // alert(error.error.message)
      let errorJSON = JSON.parse(error.error);
      alert(errorJSON.message)
    });
  }

  deleteRateHost(hostId: String){

    let rating : Rating = new Rating
    let guestJwt = localStorage.getItem('token')! 
    console.log(guestJwt)
    console.log(hostId)


    this.ratingService.getRatingByHostAndGuestId(guestJwt, hostId).subscribe((res) => {
      let resJSON = JSON.parse(res);
      rating = resJSON.rating
      console.log(resJSON.rating)
      console.log(resJSON.rating.id)

        this.ratingService.deleteRating(guestJwt, rating.id).subscribe((res) => {
          let resJSON = JSON.parse(res);
          console.log(resJSON)
          alert("Successfully deleted this rate")

        },
        (error) => {
          // alert(error.error.message)
          let errorJSON = JSON.parse(error.error);
          alert(errorJSON.message)
        });
    },
    (error) => {
      // alert(error.error.message)
      let errorJSON = JSON.parse(error.error);
      alert(errorJSON.message)
    });
  }

  rateHost(id: String) {

    let rating : Rating = new Rating
    rating.jwt = localStorage.getItem('token')!
    rating.hostId = id
    rating.rate = this.rateSelectedHost

    if(this.rateSelectedHost != 0){
        this.ratingService.createRating(rating).subscribe((res) => {
          let resJSON = JSON.parse(res);
          console.log(resJSON)
        },
        (error) => {
          // alert(error.error.message)
          let errorJSON = JSON.parse(error.error);
          alert(errorJSON.message)
        });
        alert("Successfully rated this host")

    }
    else{
      alert("You didn't rate this host")

    }
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
