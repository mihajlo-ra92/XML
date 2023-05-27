import { Injectable } from '@angular/core';
import {
  HttpClient,
  HttpHeaders,
  HttpErrorResponse,
} from '@angular/common/http';
import { Observable, catchError, throwError } from 'rxjs';
import { Accommodation } from '../model/accommodation';
import { Reservation } from '../model/reservation';
import { SearchRequest, SearchResponse } from '../model/getAllByUserRequest';
import { CreateAccommodation } from '../model/createAccommodation';
import { Jwt } from '../model/jwt';
import { AccommodationIdWithJwt } from '../model/accommodationIdWithJwt';
import { CustomPrice } from '../model/customPrice';

@Injectable({
  providedIn: 'root',
})
export class AccommodationService {
  private apiServerUrl = 'http://localhost:8000';
  private token = localStorage.getItem('token');
  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });
  // headers2: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json', 'Authorization': `Bearer ${this.token}`});
  headers3: HttpHeaders = new HttpHeaders({ Bearer: `${this.token}` });

  constructor(private http: HttpClient) {}

  getAllAccommodations(): Observable<any> {
    let accommodations = this.http.get(this.apiServerUrl + '/accommodation', {
      headers: this.headers,
      responseType: 'text',
    });
    // console.log(accommodations)

    return accommodations;
  }

  reserve(reservation: Reservation): Observable<any> {
    return this.http.post(
      this.apiServerUrl + '/accomodation-reserve',
      reservation,
      {
        headers: this.headers3,
        responseType: 'text',
      }
    );
  }
  createAccommodation(accommodation: CreateAccommodation): Observable<any> {
    return this.http.post(this.apiServerUrl + '/accomodation', accommodation, {
      headers: this.headers,
    });
  }

  getMyAccommodation(jwtString: string): Observable<any> {
    let token = new Jwt();
    token.jwt = jwtString;
    return this.http.post(
      this.apiServerUrl + '/get-accommodations-by-host',
      token,
      { headers: this.headers }
    );
  }
  getBookibgByAccommodationId(
    jwtString: string,
    accomodationId: string
  ): Observable<any> {
    let request = new AccommodationIdWithJwt();
    request.accommodationId = accomodationId;
    request.jwt = jwtString;
    return this.http.post(
      this.apiServerUrl + '/get-bookings-by-accommodation',
      request,
      { headers: this.headers }
    );
  }
  defineCustomPrice(customPrice: CustomPrice): Observable<any> {
    return this.http.post(
      this.apiServerUrl + '/custom-price-define',
      customPrice,
      { headers: this.headers }
    );
  }

  getById(accomodationId : String):Observable<any>{
    return this.http.get(this.apiServerUrl +'/accommodation/' + accomodationId,{headers: this.headers})
  }

  searchAccommodation(request: any): Observable<SearchResponse> {
    return this.http.post<SearchResponse>(
      this.apiServerUrl + '/accommodation/search',
      request,
      { headers: this.headers }
    );
  }

  // getById(id: string): Observable<Flight[]> {
  //   return this.http.get<Flight[]>(
  //     this.apiServerUrl + '/get-flight-by-id?id=' + id,
  //     { headers: this.headers }
  //   );
  // }

  // searchFlights(
  //   startPlace: string,
  //   endPlace: String,
  //   startDateString: string,
  //   endDateString: string,
  //   quantity : string
  // ): Observable<any> {
  //   return this.http.get(
  //     this.apiServerUrl +
  //       '/flight/search?startPlace=' +
  //       startPlace +
  //       '&endPlace=' +
  //       endPlace +
  //       '&startDate=' +
  //       startDateString +
  //       '&endDate=' +
  //       endDateString +
  //       '&quantity=' +
  //        quantity,
  //     { headers: this.headers, responseType: 'text' }
  //   );
  // }
}
