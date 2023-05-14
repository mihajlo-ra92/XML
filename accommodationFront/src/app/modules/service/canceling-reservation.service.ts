import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Booking } from '../model/booking';
import { Observable } from 'rxjs';
import { AuthReservationCancelingResponse, GetAllByUserResponse } from '../model/getAllByUserRequest';


@Injectable({
  providedIn: 'root'
})
export class CancelingReservationService {
  private apiServerUrl = 'http://localhost:8000';
  //private token = localStorage.getItem('token');
  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });
  constructor(private http: HttpClient) {}

  public getAllReservationByUserId(request :any): Observable<GetAllByUserResponse> {
    return this.http.post<GetAllByUserResponse>(this.apiServerUrl + '/booking/byUser',request,{ headers: this.headers });
  }

  public AuthReservationCanceling(request :any): Observable<AuthReservationCancelingResponse> {
    return this.http.post<AuthReservationCancelingResponse>(this.apiServerUrl + '/reservation-canceling',request,{ headers: this.headers });
  }

}
