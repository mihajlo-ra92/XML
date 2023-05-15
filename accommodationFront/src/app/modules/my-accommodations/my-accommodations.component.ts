import { Component, OnInit } from '@angular/core';
import { AccommodationService } from '../service/accommodation.service';
import { Route, Router } from '@angular/router';
import { Accommodation } from '../model/accommodation';
import { Booking } from '../model/booking';
import { ApproveBooking } from '../model/approveBooking';
import { BookingService } from '../service/booking.service';
import { CustomPrice } from '../model/customPrice';

@Component({
  selector: 'app-my-accommodations',
  templateUrl: './my-accommodations.component.html',
  styleUrls: ['./my-accommodations.component.css'],
})
export class MyAccommodationsComponent implements OnInit {
  public accommodations: Accommodation[] = [];
  public selected: boolean = false;
  public selectedCustomPrice: boolean = false;
  public selectedAccommodation: Accommodation = new Accommodation();
  public pictures: Document = new Document();
  public start: String = '';
  public end: String = '';
  public bookingsForAccommodation: Booking[] = new Array();
  public customPrice: CustomPrice = new CustomPrice();
  public approveBooking: ApproveBooking = new ApproveBooking();
  public imageUrl1: string = '';
  public imageUrl2: string = '';
  public imageUrl3: string = '';

  constructor(
    private accommodationService: AccommodationService,
    private router: Router,
    private bookingService: BookingService
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
    this.selectedCustomPrice = false;
    this.selected = true;
    this.selectedAccommodation = accommodation;
    let jwt = localStorage.getItem('token');
    if (jwt !== null) {
      this.accommodationService
        .getBookibgByAccommodationId(jwt, accommodation.id.toString())
        .subscribe((res) => {
          console.log(res);
          this.bookingsForAccommodation = res.bookings;
        });
    }
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
    reader3.readAsDataURL(imageFile3);
    reader3.onload = () => {
      this.imageUrl3 = reader3.result as string;
    };
  }
  unselect() {
    this.selected = false;
    this.selectedAccommodation = new Accommodation();
    this.pictures = new Document();
    this.bookingsForAccommodation = [];
    this.selectedCustomPrice = false;
  }

  selectCustomPrice() {
    this.selectedCustomPrice = true;
  }

  addCustomPrice() {
    this.customPrice.accommodationId = this.selectedAccommodation.id;
    let temp = localStorage.getItem('token');
    if (temp !== null) {
      this.customPrice.jwt = temp;
    }

    const startDateString = new Date(
      String(this.start) + 'T12:00:42.123Z'
    ).toISOString();

    const endDateString = new Date(
      String(this.end) + 'T12:00:42.123Z'
    ).toISOString();

    this.customPrice.start_date = startDateString;
    this.customPrice.end_date = endDateString;
    this.customPrice.priceType = 'Regular';
    console.log(this.customPrice);
    this.accommodationService
      .defineCustomPrice(this.customPrice)
      .subscribe((res) => {
        console.log(res);
      });
    this.start = '';
    this.end = '';
    this.customPrice = new CustomPrice();
    this.selectedCustomPrice = false;
  }

  Approve(bookingId: String) {
    this.approveBooking.bookingId = bookingId;
    this.approveBooking.jwt = localStorage.getItem('token')!;

    this.bookingService.approve(this.approveBooking).subscribe(
      (res) => {
        console.log(res);
        console.log(this.approveBooking);

        alert('Successfully approved');

        window.location.href = '/my-accommodations';
      },
      (error) => {
        console.log(error);

        alert(JSON.parse(error.error).message);
      }
    );
  }

  Deny(bookingId: String) {
    this.approveBooking.bookingId = bookingId;
    this.approveBooking.jwt = localStorage.getItem('token')!;

    this.bookingService.deny(this.approveBooking).subscribe((res) => {
      console.log(res);
      console.log(this.approveBooking);

      alert('Successfully denied');

      window.location.href = '/my-accommodations';
    });
  }

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
