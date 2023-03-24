import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { Observable, catchError, throwError } from 'rxjs';
import { Flight } from '../model/flight';


@Injectable({
  providedIn: 'root'
})
export class FlightService {

  private apiServerUrl = 'http://localhost:8080';
//   private token = localStorage.getItem('token')
  headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json'});
//   headers2: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json', 'Authorization': `Bearer ${this.token}`});
//   headers3: HttpHeaders = new HttpHeaders({'Authorization': `Bearer ${this.token}`});

  constructor(private http: HttpClient) { }

  getAllFlights(): Observable<any> {
    let flights = this.http.get(this.apiServerUrl + '/flight', {headers: this.headers, responseType: 'text'});
    console.log(flights)
    return flights
  }

}