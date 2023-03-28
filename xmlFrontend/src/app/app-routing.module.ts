import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CreateFlightPageComponent } from './modules/create-flight-page/create-flight-page.component';
import { LandingPageComponent } from './modules/landing-page/landing-page.component';
import { LoginUserComponent } from './modules/login-user/login-user.component';
import { UserTicketsComponent } from './modules/user-tickets/user-tickets.component';
import { RegisterUserComponent } from './modules/register-user/register-user.component';

const routes: Routes = [
  { path: 'create-flight', component: CreateFlightPageComponent},
  { path: 'landing-page', component: LandingPageComponent },
  { path: '', redirectTo: '/landing-page', pathMatch: 'full' },
  { path: 'login', component: LoginUserComponent },
  { path: 'register', component: RegisterUserComponent },
  { path: 'tickets', component: UserTicketsComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
