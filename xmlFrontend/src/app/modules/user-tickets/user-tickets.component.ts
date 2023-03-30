import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TicketService } from '../service/ticket.service';
import { Ticket } from '../ticket';

@Component({
  selector: 'app-user-tickets',
  templateUrl: './user-tickets.component.html',
  styleUrls: ['./user-tickets.component.css'],
})
export class UserTicketsComponent implements OnInit {
  public displayedColumns = ['number', 'floor'];
  public tickets: Ticket[] = [];
  public userId: string = '';

  constructor(private ticketService: TicketService, private router: Router) {}

  ngOnInit(): void {
    let userId = localStorage.getItem("loggedUserId")
    if(userId != null){
      this.userId = userId;
      this.ticketService.getAllTickesByUserId(this.userId).subscribe((res) => {
        this.tickets = res;
      });
    }
  }
}
