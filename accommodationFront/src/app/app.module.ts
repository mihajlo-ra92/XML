import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginUserComponent } from './modules/login-user/login-user.component';
import { LandingPageComponent } from './modules/landing-page/landing-page.component';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { RegisterComponent } from './modules/register/register.component';
import { EditUserComponent } from './modules/edit-user/edit-user.component';
import { AccommodationReserationComponent } from './modules/accommodation-reservation/accommodation-reservation';


@NgModule({
  declarations: [
    AppComponent,
    LoginUserComponent,
    LandingPageComponent,
    RegisterComponent,
    EditUserComponent,
    AccommodationReserationComponent
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
