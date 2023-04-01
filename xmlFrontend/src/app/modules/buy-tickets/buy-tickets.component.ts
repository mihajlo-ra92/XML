import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params, Router } from '@angular/router';
import { TicketService } from '../service/ticket.service';
import { FlightService } from '../service/flight.service';
import { Ticket } from '../ticket';
import { Flight } from '../model/flight';

@Component({
  selector: 'app-buy-tickets',
  templateUrl: './buy-tickets.component.html',
  styleUrls: ['./buy-tickets.component.css']
})
export class BuyTicketsComponent implements OnInit {
  public ticket : Ticket =new Ticket();
  public flight : Flight[] =[];
  public tPrice :  number = 0;
  public tCapacity : number = 0;
  constructor(private flightService: FlightService, private route: ActivatedRoute,private ticketService: TicketService, private router: Router) { }

  ngOnInit(): void {
    this.route.params.subscribe((params: Params) => {
      this.flightService.getById(params['flightId']).subscribe(res => {
        this.flight = res;
        console.log(this.flight)
        if(this.flight[0]!= null){
          this.ticket.date = new Date(this.flight[0].date);
          this.ticket.endPlace = this.flight[0].endPlace;
          this.ticket.startPlace = this.flight[0].startPlace;
          let temp = localStorage.getItem("loggedUserId");
          if(temp !=null){
          this.ticket.userId = temp;
          }
          this.ticket.price = this.flight[0].price;
          this.ticket.capacity = 1;
          this.ticket.flightId = this.flight[0].id;
          this.tPrice = this.flight[0].price; 
          this.tCapacity = 1;}
      
      })
    });
  }

  BuyTickets():void{
    console.log(this.ticket)
    this.ticketService.createTicket(this.ticket).subscribe((res) => {
      alert("Successful purchase");
      this.router.navigate(['/landing-page'])
    },(error) =>{
      alert("There are not enough seats")
    }
    );
  }
  onInputChange(event: any) {
    console.log('Unos promijenjen:', event);
    this.tPrice = this.flight[0].price * event;
    this.ticket.price = this.tPrice;
    this.ticket.capacity = event;
  }
}
