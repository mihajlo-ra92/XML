import { Component, OnInit } from '@angular/core';
import { AccommodationService } from '../service/accommodation.service';
import { Route, Router } from '@angular/router';
import { Accommodation } from '../model/accommodation';

@Component({
  selector: 'app-my-accommodations',
  templateUrl: './my-accommodations.component.html',
  styleUrls: ['./my-accommodations.component.css'],
})
export class MyAccommodationsComponent implements OnInit {
  public accommodations: Accommodation[] = [];
  public selected: boolean = false;
  public selectedAccommodation: Accommodation = new Accommodation();
  public pictures: Document = new Document();
  public imageUrl1: string = '';
  public imageUrl2: string = '';
  public imageUrl3: string = '';
  constructor(
    private accommodationService: AccommodationService,
    private router: Router
  ) {}

  ngOnInit(): void {
    let token = localStorage.getItem('token');
    if (token !== null) {
      this.accommodationService.getMyAccommodation(token).subscribe((res) => {
        this.accommodations = res.accommodations;
        console.log(this.accommodations);
      });
    }
    let temp = localStorage.getItem('loggedUserType');
    if (temp != '1') {
      this.router.navigate(['/landing-page']);
    }
  }

  select(accommodation: Accommodation) {
    this.selected = true;
    this.selectedAccommodation = accommodation;
    this.imageUrl1 = '';
    this.imageUrl2 = '';
    this.imageUrl3 = '';
    const reader1 = new FileReader();
    const imageFile1 = this.saveImage(
      String(this.selectedAccommodation.pictures[0])
    );
    reader1.readAsDataURL(imageFile1);
    reader1.onload = () => {
      this.imageUrl1 = reader1.result as string;
    };

    const reader2 = new FileReader();
    const imageFile2 = this.saveImage(
      String(this.selectedAccommodation.pictures[1])
    );
    reader2.readAsDataURL(imageFile2);
    reader2.onload = () => {
      this.imageUrl2 = reader2.result as string;
    };

    const reader3 = new FileReader();
    const imageFile3 = this.saveImage(
      String(this.selectedAccommodation.pictures[2])
    );
    reader3.readAsDataURL(imageFile2);
    reader3.onload = () => {
      this.imageUrl3 = reader3.result as string;
    };
  }
  unselect() {
    this.selected = false;
    this.selectedAccommodation = new Accommodation();
    this.pictures = new Document();
  }

  selectCustomPrice() {}

  saveImage(base64Image: String) {
    console.log('in save image');
    console.log(base64Image);
    const imageName = 'name.png';
    const imageBlob = this.dataURItoBlob(base64Image);
    const imageFile = new File([imageBlob], imageName, { type: 'image/png' });
    console.log('imageFile');
    console.log(imageFile);
    return imageFile;
  }
  dataURItoBlob(dataURI: any) {
    const byteString = window.atob(dataURI);
    const arrayBuffer = new ArrayBuffer(byteString.length);
    const int8Array = new Uint8Array(arrayBuffer);
    for (let i = 0; i < byteString.length; i++) {
      int8Array[i] = byteString.charCodeAt(i);
    }
    const blob = new Blob([int8Array], { type: 'image/png' });
    return blob;
  }
}
