import { Injectable } from '@angular/core';
import {
  HttpClient,
  HttpHeaders,
} from '@angular/common/http';
import { Observable, catchError, throwError } from 'rxjs';
import { Rating } from '../model/rating';

@Injectable({
  providedIn: 'root',
})
export class RatingService {
  private apiServerUrl = 'http://localhost:8000';
  private token = localStorage.getItem('token');
  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });
  // headers2: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json', 'Authorization': `Bearer ${this.token}`});
  headers3: HttpHeaders = new HttpHeaders({ Bearer: `${this.token}` });

  constructor(private http: HttpClient) {}

  getAllRatings(): Observable<any> {
    return this.http.get(this.apiServerUrl + '/rating', {
      headers: this.headers,
      responseType: 'text',
    });
  }

  createRating(rating: Rating): Observable<any>{
    return this.http.post(this.apiServerUrl + '/rating', rating, {
        headers: this.headers,
        responseType: 'text',
      });
  }

  getRatingByAccommodationAndGuestId(jwt: String, accommodationId: String): Observable<any>{
    var combinedObject = {
      jwt: jwt,
      accommodationId: accommodationId
    };
    console.log(combinedObject)
    return this.http.post(this.apiServerUrl + '/get-rating-by-accomodation-id', combinedObject, {
        headers: this.headers,
        responseType: 'text',
      });
  }

  getRatingByHostAndGuestId(jwt: String, hostId: String): Observable<any>{
    var combinedObject = {
      jwt: jwt,
      hostId: hostId
    };
    console.log(combinedObject)
    return this.http.post(this.apiServerUrl + '/get-rating-by-host-id', combinedObject, {
        headers: this.headers,
        responseType: 'text',
      });
  }

  deleteRating(jwt: String, ratingId: String): Observable<any>{
    var combinedObject = {
      jwt: jwt,
      ratingId: ratingId
    };
    console.log(combinedObject)
    return this.http.post(this.apiServerUrl + '/delete-rating', combinedObject, {
      headers: this.headers,
      responseType: 'text',
    });
  }

}
