import { Component, OnInit } from '@angular/core';
import { User } from '../model/user';
import { UserService } from '../service/user.service';
import { Route, Router } from '@angular/router';
import { UserWithJwt } from '../model/userWithJwt'
import { UserForDelete } from '../model/userForDelete';

@Component({
  selector: 'app-edit-user',
  templateUrl: './edit-user.component.html',
  styleUrls: ['./edit-user.component.css']
})
export class EditUserComponent implements OnInit {
  public user: User = new User();
  public userForEdit: UserWithJwt = new UserWithJwt();
  public userForDelete:  UserForDelete = new UserForDelete();
  constructor(private userService : UserService, private router: Router) { }

  ngOnInit(): void {
    let temp = localStorage.getItem("loggedUserId");
    if(temp !=null){
      this.userService.getById(temp).subscribe((res) =>{
        //console.log(res);
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
      this.router.navigate(['/landing-page']);
    })
  }

  Delete(){
    this.userForDelete.id = this.user.id;
    let temp = localStorage.getItem("token");
    if(temp !=null){
      this.userForDelete.jwt = temp;
    }
    console.log(this.userForEdit);
    this.userService.deleteUser(this.userForDelete).subscribe((res) =>{
      console.log(res);
      this.onLogout();
    })
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
