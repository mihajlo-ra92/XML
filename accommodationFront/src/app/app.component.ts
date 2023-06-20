import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { from } from 'rxjs';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'xmlFrontend';
  loggedUserJwt = localStorage.getItem('token')
  loggedUserRole = localStorage.getItem('loggedUserType')
  loggedUserId = localStorage.getItem('loggedUserId')
  loggedUsername = localStorage.getItem('loggedUsername')
  constructor(private router: Router) {
  }
  ngOnInit(): void {
console.log(this.loggedUserRole)
    
  }

  onLogout() {
    localStorage.removeItem('token');
    localStorage.removeItem('loggedUserType')
    localStorage.removeItem('loggedUserId')
    localStorage.removeItem('loggedUsername')

    console.log(localStorage.getItem('token'))
    this.router.navigate(['/login']);
    location.reload()
  }
}
