import { Component, OnInit } from '@angular/core';
import { User } from '../model/user';
import { UserService } from '../service/user.service';

@Component({
  selector: 'app-register-user',
  templateUrl: './register-user.component.html',
  styleUrls: ['./register-user.component.css'],
})
export class RegisterUserComponent implements OnInit {
  user: User = new User();
  constructor(private userService: UserService) {}

  ngOnInit(): void {}

  Register() {
    this.userService.register(this.user).subscribe((res) => {
      console.log(res);
    });
  }
}
