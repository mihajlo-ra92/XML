import { Component, OnInit } from '@angular/core';
import { AccommodationService } from '../service/accommodation.service';
import { Route, Router } from '@angular/router';
import { Accommodation } from '../model/accommodation';

@Component({
  selector: 'app-my-accommodations',
  templateUrl: './my-accommodations.component.html',
  styleUrls: ['./my-accommodations.component.css']
})
export class MyAccommodationsComponent implements OnInit {
  public accommodations: Accommodation[] = [];
  public selected : boolean = false;
  public selectedAccommodation: Accommodation = new Accommodation();
  public pictures: Document = new Document;
  constructor(private accommodationService : AccommodationService, private router: Router) { }

  ngOnInit(): void {
    let token =localStorage.getItem("token");
    if(token !== null){
      this.accommodationService.getMyAccommodation(token).subscribe((res) =>{
        this.accommodations = res.accommodations;
        console.log(this.accommodations);
      })
    }
    let temp = localStorage.getItem("loggedUserType");
    if(temp != "1"){
    this.router.navigate(['/landing-page']);    
    }
  }

  select(accommodation: Accommodation){
    this.selected = true;
    this.selectedAccommodation = accommodation;
    this.pictures = this.decodePicture(this.selectedAccommodation.pictures[0].toString());
    console.log(this.decodePicture(this.selectedAccommodation.pictures[0].toString()));
    
  }
  unselect(){
    this.selected = false;
    this.selectedAccommodation = new Accommodation();
    this.pictures  = new Document;;
  }

  decodePicture(base64String: string):any{
    const byteArray = Uint8Array.from(atob(base64String), c => c.charCodeAt(0));

// Kreirajte Blob objekat od ByteArray
const blob = new Blob([byteArray], { type: "image/jpeg" });

// Kreirajte URL za prikaz slike
const imageUrl = URL.createObjectURL(blob);

// Kako biste prikazali sliku u HTML-u, možete koristiti <img> tag
const img = document.createElement("img");
img.src = imageUrl;
document.body.appendChild(img);
return document;
  }
  selectCustomPrice(){
    
  }
}
