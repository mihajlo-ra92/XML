import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LandingPageComponent } from './modules/landing-page/landing-page.component';
import { LoginUserComponent } from './modules/login-user/login-user.component';
import { RegisterComponent } from './modules/register/register.component';
import { EditUserComponent } from './modules/edit-user/edit-user.component';

const routes: Routes = [
  { path: 'landing-page', component: LandingPageComponent },
  { path: '', redirectTo: '/landing-page', pathMatch: 'full' },
  { path: 'login', component: LoginUserComponent },
  { path: 'register', component: RegisterComponent},
  { path: 'edit-user', component: EditUserComponent},
  
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
