import { Component, OnInit } from '@angular/core';
import { LoginUser } from '../model/loginUser';
import { UserService } from '../service/user.service';

@Component({
  selector: 'app-login-user',
  templateUrl: './login-user.component.html',
  styleUrls: ['./login-user.component.css'],
})
export class LoginUserComponent implements OnInit {
  user: LoginUser = new LoginUser()
  token : String = ""
  constructor( private userService: UserService) {}

  ngOnInit(): void {}

  Login() {
    this.userService.login(this.user).subscribe((res) => {
      console.log(res.body.jwt)
      this.token = res.body.jwt

      localStorage.setItem('token', res.body.jwt);
      // console.log(localStorage.getItem('token'));

      const decodedJWT = JSON.parse(window.atob(this.token.split('.')[1]));
      localStorage.setItem('loggedUserId', decodedJWT.userId);
      localStorage.setItem('loggedUserType', decodedJWT.userType);
      localStorage.setItem('loggedUsername', decodedJWT.username);

      console.log(localStorage.getItem('loggedUserType'))
      console.log(localStorage.getItem('loggedUserId'))
      console.log(localStorage.getItem('loggedUsername'))
      
      // this.toastr.success('Successfully logged in');
      alert("Successfully logged in")

      window.location.href = '/'
    },
    (error) => {
      // this.toastr.error("Invalid email/password");
      // console.log(error);
      alert("Invalid email/password")

    }
    );
  }
}
