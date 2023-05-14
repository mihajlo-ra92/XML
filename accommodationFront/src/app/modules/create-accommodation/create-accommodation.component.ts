import { Component, OnInit } from '@angular/core';
import { UserService } from '../service/user.service';
import { Route, Router } from '@angular/router';
import { CreateAccommodation } from '../model/createAccommodation';
import { User } from '../model/user';
import { AccommodationService } from '../service/accommodation.service';
import { PriceType } from '../model/priceType';

@Component({
  selector: 'app-create-accommodation',
  templateUrl: './create-accommodation.component.html',
  styleUrls: ['./create-accommodation.component.css']
})
export class CreateAccommodationComponent implements OnInit {
  public accommodation:CreateAccommodation = new CreateAccommodation(); 
  public benefite: string = "";
  selectedFile: File = new File([],"");
  public nameOfPictures: string[] = [];
  public picturesAsFile : File[] = [];  
  //public picture: ImageBitmap = new ImageBitmap();
  constructor(private accommodationService : AccommodationService, private router: Router) { }

  ngOnInit(): void {
    let temp = localStorage.getItem("loggedUserType");
    if(temp != "1"){
    this.router.navigate(['/landing-page']);    
    }
  }

  async  Create(){
    let temp = localStorage.getItem("token")
    this.accommodation.priceType = PriceType.Regular;
    if(temp != null){
      this.accommodation.jwt = temp;
    }
    for(let i = 0; i< this.picturesAsFile.length; ++i){
       let picture = this.imageToBase64(this.picturesAsFile[i]);
      //    this.accommodation.pictures.push(picture)
       
      //    console.log(picture);

      const base64String: string = await this.imageToBase64(this.picturesAsFile[i]);// Prikaz rezulta
      this.accommodation.pictures.push(base64String)
    }
    //this.accommodation.pictures = this.nameOfPictures;
    this.accommodationService.createAccommodation(this.accommodation).subscribe((res) =>{
      console.log(res);
      this.accommodation = new CreateAccommodation();
      this.nameOfPictures = [];
      this.picturesAsFile = [];
      this.selectedFile = new File([],"");
    })
  }

  addBenefit(){
    this.accommodation.benefits.push(this.benefite);
    this.benefite = ""
  }

  onFileSelected(event: any) {
    this.selectedFile = event.target.files[0];
    this.picturesAsFile.push(this.selectedFile)
    this.nameOfPictures.push(this.selectedFile.name)
  }

  async imageToBase64(imageFile: File): Promise<string> {
    return new Promise<string>((resolve, reject) => {
      const reader = new FileReader();
      reader.onloadend = () => {
        const base64String = reader.result as string;
        resolve(base64String.substring(base64String.indexOf(',') + 1)); // Uklanjanje "data:image/jpeg;base64," deo
      };
      reader.onerror = reject;
      reader.readAsDataURL(imageFile);
    });
  }
  

}
