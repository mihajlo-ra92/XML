import { Component, OnInit } from '@angular/core';
import { User } from '../model/user';
import { UserService } from '../service/user.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register-user',
  templateUrl: './register-user.component.html',
  styleUrls: ['./register-user.component.css'],
})
export class RegisterUserComponent implements OnInit {
  user: User = new User();
  dateStr: string = '';
  constructor(private userService: UserService, private router: Router) {}

  ngOnInit(): void {}

  Register() {
    this.user.userType = 'regular';
    console.log(this.dateStr);
    console.log(Date.parse(this.dateStr));
    this.user.birthDate = String(Math.round(Date.parse(this.dateStr) / 1000));
    console.log(this.user);
    this.userService.register(this.user).subscribe((res) => {
      console.log(res);
      this.router.navigateByUrl('/');

      alert("Successfully registered")

    });
  }
}
