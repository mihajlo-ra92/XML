import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LandingPageComponent } from './modules/landing-page/landing-page.component';
import { LoginUserComponent } from './modules/login-user/login-user.component';
import { AccommodationReserationComponent } from './modules/accommodation-reservation/accommodation-reservation';

const routes: Routes = [
  { path: 'landing-page', component: LandingPageComponent },
  { path: '', redirectTo: '/landing-page', pathMatch: 'full' },
  { path: 'login', component: LoginUserComponent },
  { path: 'accommodation-reservation', component: AccommodationReserationComponent },
  
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
