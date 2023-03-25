import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Ticket } from '../ticket';

@Injectable({
  providedIn: 'root'
})
export class TicketService {

  private apiServerUrl = 'http://localhost:8080';
  headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json'});
  constructor(private http: HttpClient) { }
  
    public getAllTickesByUserId(id: string): Observable<Ticket[]> {
      let tickets = this.http.get<Ticket[]>(this.apiServerUrl + '/get-tickets-by-user-id?id='+id);
      console.log(tickets)
      return tickets;
    }

    createTicket(ticket: any): Observable<any> {
      return this.http.post<any>(this.apiServerUrl + '/ticket', ticket, {headers: this.headers});
    }

}
