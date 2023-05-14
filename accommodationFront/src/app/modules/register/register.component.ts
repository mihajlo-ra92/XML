import { Component, OnInit } from '@angular/core';
import { User } from '../model/user';
import { RegisterUser} from '../model/registerUser'
import { UserService } from '../service/user.service';
import { Route, Router } from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  user: RegisterUser = new RegisterUser();
  dateStr: string = '';
  selectedColor : string = '';
  constructor(private userService : UserService, private router: Router) { }

  ngOnInit(): void {
  }
  Register(){
    console.log(this.user);
    this.userService.register(this.user).subscribe((res) =>{
      // console.log(res);
      this.router.navigate(['/login']);

    })
  }
}
