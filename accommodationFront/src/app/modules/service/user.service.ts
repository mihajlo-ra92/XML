import { Injectable } from '@angular/core';
import {
  HttpClient,
  HttpHeaders,
  HttpErrorResponse,
} from '@angular/common/http';
import { Observable, catchError, throwError } from 'rxjs';
import { LoginUser } from '../model/loginUser';
import { User } from '../model/user';
import { RegisterUser } from '../model/registerUser';
import { UserWithJwt } from '../model/userWithJwt';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  private apiServerUrl = 'http://localhost:8000';
  private token = localStorage.getItem('token');
  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });
  headers2: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
    Bearer: 'token',
  });
  headers3: HttpHeaders = new HttpHeaders({
    Authorization: `Bearer ${this.token}`,
  });

  constructor(private http: HttpClient) {}

  login(user: LoginUser): Observable<any> {
    console.log(user);
    let response = this.http.post(this.apiServerUrl + '/login', user, {
      headers: this.headers,
      observe: 'response',
      responseType: 'json',
    });
    console.log(this.headers2);
    return response;
  }

  register(user: RegisterUser): Observable<any> {
    console.log(user);
    const response = this.http.post(this.apiServerUrl + '/user', user, {
      headers: this.headers,
    });
    return response;
  }

  getById(id: string): Observable<any> {
    return this.http.get(this.apiServerUrl + '/user/' + id, {headers: this.headers});
  }

  editUser(user: UserWithJwt): Observable<any> {
    return this.http.post(this.apiServerUrl + '/update-user',user,{headers: this.headers})
  }
}
