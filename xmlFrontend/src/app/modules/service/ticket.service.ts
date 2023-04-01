import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Ticket } from '../ticket';

@Injectable({
  providedIn: 'root',
})
export class TicketService {
  private apiServerUrl = 'http://localhost:8080';
  private token = localStorage.getItem('token');
  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
    Bearer: `${this.token}`,
  });
  constructor(private http: HttpClient) {}

  public getAllTickesByUserId(): Observable<Ticket[]> {
    let tickets = this.http.get<Ticket[]>(
      this.apiServerUrl + '/get-tickets-by-user-id',
      { headers: this.headers }
    );
    console.log(tickets);
    return tickets;
  }

  createTicket(ticket: any): Observable<any> {
    return this.http.post<any>(this.apiServerUrl + '/ticket', ticket, {
      headers: this.headers,
    });
  }
}
