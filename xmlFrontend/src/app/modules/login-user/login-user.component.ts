import { Component, OnInit } from '@angular/core';
import { LoginUser } from '../model/loginUser';
import { UserService } from '../service/user.service';

@Component({
  selector: 'app-login-user',
  templateUrl: './login-user.component.html',
  styleUrls: ['./login-user.component.css'],
})
export class LoginUserComponent implements OnInit {
  user: LoginUser = new LoginUser();
  constructor(private userService: UserService) {}

  ngOnInit(): void {}

  Login() {
    this.userService.login(this.user).subscribe((res) => {
      localStorage.setItem('token', res.body.Bearer);
      console.log(localStorage.getItem('token'));
    });
  }
}
