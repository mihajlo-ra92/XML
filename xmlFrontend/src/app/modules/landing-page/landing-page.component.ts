import { Component, OnInit } from '@angular/core';
import { Flight } from '../model/flight';
import { FlightService } from '../service/flight.service';

@Component({
  selector: 'app-landing-page',
  templateUrl: './landing-page.component.html',
  styleUrls: ['./landing-page.component.css']
})
export class LandingPageComponent implements OnInit {

  allFlights : Array<Flight> = new Array

  constructor(private flightService: FlightService) { }

  ngOnInit(): void {

    this.allUsers()
    
  }

  allUsers() {
    this.flightService.getAllFlights().subscribe(res => {
      let resJSON = JSON.parse(res)
      this.allFlights = resJSON
      this.allFlights.map((x) => {
        const myDate = new Date(x.date)
        // myDate.setHours(myDate.getHours() + 2) 
        console.log(myDate)  
        x.date = myDate.toLocaleString()
        console.log(x.date)  

      }
      )
      console.log(this.allFlights)             
    })
  }

}
