import { Component, OnInit } from '@angular/core';
import { Flight } from '../model/flight';
import { FlightService } from '../service/flight.service';

@Component({
  selector: 'app-landing-page',
  templateUrl: './landing-page.component.html',
  styleUrls: ['./landing-page.component.css'],
})
export class LandingPageComponent implements OnInit {
  allFlights: Array<Flight> = new Array();
  isAdmin = false;
  loggedUserRole = localStorage.getItem('loggedUserType')
  clickBuy = false;
  startPlace: string = '';
  endPlace: string = '';
  startDate: Date | undefined;
  endDate : Date | undefined;
  quantity: number = 0;

  constructor(private flightService: FlightService) {}

  ngOnInit(): void {
    //  var loggedUserType = localStorage.getItem('loggedUserType')
    //  console.log(loggedUserType);

    //     if(loggedUserType?.toString() === 'admin'){
    //       console.log(this.isAdmin);

    //       this.isAdmin = true;
    //    }

    this.allUsers();
  }

  allUsers() {
    var loggedUserType = localStorage.getItem('loggedUserType');
    console.log(loggedUserType);

    if (loggedUserType?.toString() === 'admin') {
      console.log(this.isAdmin);

      this.isAdmin = true;
    }
    this.flightService.getAllFlights().subscribe((res) => {
      let resJSON = JSON.parse(res);
      this.allFlights = resJSON;
      // this.allFlights.map((x) => {
      //   const myDate = new Date(x.date);
      //   // myDate.setHours(myDate.getHours() + 2)
      //   console.log(myDate);
      //   x.date = myDate.toLocaleString('en-US', {
      //     timeZone: 'America/New_York',
      //   });
      //   console.log(x.date);
      // });
      console.log(this.allFlights);
    });
  }

  deleteFlight(flight: Flight) {
    console.log(flight.id);
    this.flightService.deleteFlight(flight.id).subscribe((res) => {
      console.log(res);
      this.allUsers();
      window.location.href = '/';
    });
  }


  search(){
    if(this.startDate != undefined && this.endDate != undefined){
      
      var startDateString = this.startDate.toString();
      var startDateFormated = startDateString + 'T00:00:00.123Z';

      var endDateString = this.endDate.toString();
      var endDateFormated = endDateString + 'T00:00:00.123Z';


      this.flightService
        .searchFlights(
          this.startPlace,
          this.endPlace,
          startDateFormated,
          endDateFormated,
          this.quantity.toString()
        )
        .subscribe((res) => {
          console.log(res);
          if (res !== null) {
            try{
            let resJSON = JSON.parse(res);
            // let resJSON = res;
            this.allFlights = resJSON;
            // this.allFlights.map((x) => {
            //   const myDate = new Date(x.date);
            //   console.log(myDate);
            //   x.date = myDate.toLocaleString('en-US', {
            //     timeZone: 'America/New_York',
            //   });
            //   console.log(x.date);
            // });
            }
            catch{
              alert("No flights")
            }
          } else {
            console.log('res je null');

          }
          
        });
    }
  }

  switchPlaces() {
    let newStartPlace: string;
    newStartPlace = this.endPlace;
    this.endPlace = this.startPlace;
    this.startPlace = newStartPlace;
  }
}
