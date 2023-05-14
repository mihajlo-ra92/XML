import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LandingPageComponent } from './modules/landing-page/landing-page.component';
import { LoginUserComponent } from './modules/login-user/login-user.component';
import { CancelingReservationComponent } from './modules/canceling-reservation/canceling-reservation.component';
import { RegisterComponent } from './modules/register/register.component';
import { EditUserComponent } from './modules/edit-user/edit-user.component';
import { AccommodationReserationComponent } from './modules/accommodation-reservation/accommodation-reservation';
import { CreateAccommodationComponent } from './modules/create-accommodation/create-accommodation.component';
import { MyAccommodationsComponent } from './modules/my-accommodations/my-accommodations.component';

const routes: Routes = [
  { path: 'landing-page', component: LandingPageComponent },
  { path: '', redirectTo: '/landing-page', pathMatch: 'full' },
  { path: 'login', component: LoginUserComponent },
  { path: 'canceling-reservation', component: CancelingReservationComponent },
  { path: 'register', component: RegisterComponent},
  { path: 'edit-user', component: EditUserComponent},
  { path: 'accommodation-reservation', component: AccommodationReserationComponent },
  { path: 'create-accommodation', component: CreateAccommodationComponent},
  { path: 'my-accommodations', component: MyAccommodationsComponent},
  
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
