import { Injectable } from '@angular/core';
import {
  HttpClient,
  HttpHeaders,
  HttpErrorResponse,
} from '@angular/common/http';
import { Observable, catchError, throwError } from 'rxjs';

@Injectable({
  providedIn: 'root',
})

export class BookingService {
  private apiServerUrl = 'http://localhost:8000';
  private token = localStorage.getItem('token');
  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });
  // headers2: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json', 'Authorization': `Bearer ${this.token}`});
  headers3: HttpHeaders = new HttpHeaders({ Bearer: `${this.token}` });

  constructor(private http: HttpClient) {}

  get(accommodationId: String): Observable<any> {
    return this.http.post(this.apiServerUrl + '/booking/' + accommodationId, accommodationId, {
      headers: this.headers3,
      responseType: 'text',
    });
  }
}
