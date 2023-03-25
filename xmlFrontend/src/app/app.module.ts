import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginUserComponent } from './modules/login-user/login-user.component';
import { LandingPageComponent } from './modules/landing-page/landing-page.component';
import { RegisterUserComponent } from './modules/register-user/register-user.component';
import { UserTicketsComponent } from './modules/user-tickets/user-tickets.component';
import { HttpClientModule } from '@angular/common/http';
import { FlightService } from './modules/service/flight.service';
import { FormsModule } from '@angular/forms';

@NgModule({
  declarations: [
    AppComponent,
    LoginUserComponent,
    LandingPageComponent,
    RegisterUserComponent,
    UserTicketsComponent,
  ],
  imports: [BrowserModule, AppRoutingModule, HttpClientModule, FormsModule],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
