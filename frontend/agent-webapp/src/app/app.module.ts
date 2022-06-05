import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './components/login/login.component';
import { SignupComponent } from './components/signup/signup.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatCardModule } from '@angular/material/card';
import { MatInputModule } from '@angular/material/input';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NewCompanyComponent } from './components/new-company/new-company.component';
import { NewJobOfferComponent } from './components/new-job-offer/new-job-offer.component';
import {MatDatepickerModule} from '@angular/material/datepicker';
import { MatNativeDateModule } from '@angular/material/core';
import { TokenInterceptor } from './interceptor/token-interceptor';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { AdminRequestsComponent } from './components/admin-requests/admin-requests.component';
import {MatCheckboxModule} from '@angular/material/checkbox';
import { CompanyProfileComponent } from './components/company-profile/company-profile.component';
import { CompanyInfoComponent } from './components/company-info/company-info.component';
import { CompanyCommentsComponent } from './components/company-comments/company-comments.component';
import { CompanySalaryComponent } from './components/company-salary/company-salary.component';
import { CompanyInterviewsComponent } from './components/company-interviews/company-interviews.component';
import { JobOffersComponent } from './components/job-offers/job-offers.component';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatDividerModule} from '@angular/material/divider';
import { CompanyDislinktComponent } from './components/company-dislinkt/company-dislinkt.component';



@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    SignupComponent,
    NavbarComponent,
    NewCompanyComponent,
    NewJobOfferComponent,
    AdminRequestsComponent,
    CompanyProfileComponent,
    CompanyInfoComponent,
    CompanyCommentsComponent,
    CompanySalaryComponent,
    CompanyInterviewsComponent,
    JobOffersComponent,
    CompanyDislinktComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatFormFieldModule,
    MatInputModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    MatCardModule,
    MatDatepickerModule,
    MatNativeDateModule,
    MatSnackBarModule,
    MatCheckboxModule,
    MatGridListModule,
    MatDividerModule
  ],
  providers: [ {
    provide: HTTP_INTERCEPTORS,
    useClass: TokenInterceptor,
    multi: true
  },],
  bootstrap: [AppComponent]
})
export class AppModule { }
