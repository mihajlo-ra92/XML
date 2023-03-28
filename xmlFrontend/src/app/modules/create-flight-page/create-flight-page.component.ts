import { Component, OnInit } from '@angular/core';
import { Flight } from '../model/flight';
import { FlightService } from '../service/flight.service';

@Component({
  selector: 'app-create-flight-page',
  templateUrl: './create-flight-page.component.html',
  styleUrls: ['./create-flight-page.component.css']
})
export class CreateFlightPageComponent implements OnInit {

  today: Date = new Date
  flight: Flight = new Flight
  flightDate: Date = new Date
  constructor(private flightService: FlightService) { }

  ngOnInit(): void {
  }

  Refresh() {
    window.location.reload()
  }

  createFlight() {
    this.flight.capacity = Number(this.flight.capacity)
    this.flight.freeSeats = this.flight.capacity
    this.flight.price = Number(this.flight.price)
    // console.log(this.flightDate)

    const flightDateString = new Date(this.flightDate).toISOString()
    // let flightDateString = new Date(this.flightDate).toISOString()
    
    console.log(flightDateString)
    this.flight.date = flightDateString

    console.log(this.flight)

    this.flightService.create(this.flight).subscribe(res => {
      // let resJSON = JSON.parse(res)
      console.log(res)
      // console.log(resJSON)
    })

    window.location.href = '/'

  }
}
