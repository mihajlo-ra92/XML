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
      this.allFlights.map((x) => {
        const myDate = new Date(x.date);
        // myDate.setHours(myDate.getHours() + 2)
        console.log(myDate);
        x.date = myDate.toLocaleString('en-US', {
          timeZone: 'America/New_York',
        });
        console.log(x.date);
      });
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
}
