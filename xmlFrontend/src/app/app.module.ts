import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginUserComponent } from './modules/login-user/login-user.component';
import { LandingPageComponent } from './modules/landing-page/landing-page.component';
import { RegisterUserComponent } from './modules/register-user/register-user.component';
import { UserTicketsComponent } from './modules/user-tickets/user-tickets.component';
import { HttpClientModule } from '@angular/common/http';
import { FlightService } from './modules/service/flight.service';
import { FormsModule } from '@angular/forms';
import { CreateFlightPageComponent } from './modules/create-flight-page/create-flight-page.component';
import { BuyTicketsComponent } from './modules/buy-tickets/buy-tickets.component';


@NgModule({
  declarations: [
    AppComponent,
    LoginUserComponent,
    LandingPageComponent,
    RegisterUserComponent,
    CreateFlightPageComponent,
    UserTicketsComponent,
    BuyTicketsComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
