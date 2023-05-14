import { Component, OnInit } from '@angular/core';
import { User } from '../model/user';
import { UserService } from '../service/user.service';
import { Route, Router } from '@angular/router';
import { UserWithJwt } from '../model/userWithJwt'

@Component({
  selector: 'app-edit-user',
  templateUrl: './edit-user.component.html',
  styleUrls: ['./edit-user.component.css']
})
export class EditUserComponent implements OnInit {
  public user: User = new User();
  public userForEdit: UserWithJwt = new UserWithJwt();
  constructor(private userService : UserService, private router: Router) { }

  ngOnInit(): void {
    let temp = localStorage.getItem("loggedUserId");
    if(temp !=null){
      this.userService.getById(temp).subscribe((res) =>{
        console.log(res);
        this.user = res.user;
        
      })
    }else{
      this.router.navigate(['/login']);

    }
  }
  Edit(){
    console.log("editujemo");
    this.userForEdit.user = this.user;
    let temp = localStorage.getItem("token");
    if(temp !=null){
      this.userForEdit.jwt = temp;
    }
    this.userService.editUser(this.userForEdit).subscribe((res) =>{
      console.log(res);
      
    })
  }

}
