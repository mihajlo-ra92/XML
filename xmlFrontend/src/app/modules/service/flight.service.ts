import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { Observable, catchError, throwError } from 'rxjs';
import { Flight } from '../model/flight';


@Injectable({
  providedIn: 'root'
})
export class FlightService {

  private apiServerUrl = 'http://localhost:8080';
  private token = localStorage.getItem('token')
  headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json'});
  // headers2: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json', 'Authorization': `Bearer ${this.token}`});
  headers3: HttpHeaders = new HttpHeaders({'Bearer': `${this.token}`});

  constructor(private http: HttpClient) { }

  getAllFlights(): Observable<any> {
    let flights = this.http.get(this.apiServerUrl + '/flight', {headers: this.headers, responseType: 'text'});
    return flights
  }

  deleteFlight(id: String): Observable<any>{
    return this.http.delete(this.apiServerUrl + '/flight/' + id,{headers:this.headers, responseType: 'text'})
  }
    create(flight: Flight): Observable<any> {
    return this.http.post(this.apiServerUrl + '/flight', flight, {headers: this.headers3, responseType: 'text'});
  }

  getById(id: string): Observable<Flight[]>{
    return this.http.get<Flight[]>(this.apiServerUrl + '/get-flight-by-id?id='+id,{headers: this.headers});
 searchFlights(startPlace: string, endPlace: String, startDateString: string, endDateString: string):Observable<any> {
    return this.http.get(this.apiServerUrl + '/flight/search?startPlace=' + startPlace + '&endPlace=' + endPlace + '&startDate=' + startDateString + '&endDate=' + endDateString,{headers: this.headers, responseType: 'text'})
  }

}