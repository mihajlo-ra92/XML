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
  isAdmin = true;

  constructor(private flightService: FlightService) { }

  ngOnInit(): void {
    // var token = localStorage.getItem('token')
    // //  if(token.){

    // // }
    // console.log(token);
    
    this.allUsers()
    
  }

  allUsers() {
    this.flightService.getAllFlights().subscribe(res => {
      let resJSON = JSON.parse(res)
      this.allFlights = resJSON
      console.log(this.allFlights)             
    })
  }

  deleteFlight(flight: Flight){
    console.log(flight.id);
     this.flightService.deleteFlight(flight.id).subscribe(res =>{
       console.log(res);
       this.allUsers();
     })
  }

}
