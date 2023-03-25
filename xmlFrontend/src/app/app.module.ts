import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginUserComponent } from './modules/login-user/login-user.component';
import { LandingPageComponent } from './modules/landing-page/landing-page.component';
import { RegisterUserComponent } from './register-user/register-user.component';
import { UserTicketsComponent } from './modules/user-tickets/user-tickets.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginUserComponent,
    LandingPageComponent,
    RegisterUserComponent,
    UserTicketsComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
