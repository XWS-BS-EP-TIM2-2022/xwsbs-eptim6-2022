import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HomepageComponent } from './components/homepage/homepage.component';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatCardModule } from '@angular/material/card';
import { MatInputModule } from '@angular/material/input';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { PostsViewComponent } from './components/posts-view/posts-view.component';
import { LoginComponent } from './components/login/login.component';
import { RegistrationComponent } from './components/registration/registration.component';
import { TokenInterceptor } from './interceptor/token-interceptor';
import { ProfileComponent } from './components/profile/profile.component';
import { ChangePasswordComponent } from './components/change-password/change-password.component';
import { ForgotPasswordComponent } from './components/forgot-password/forgot-password.component';
import { SetNewPasswordComponent } from './components/set-new-password/set-new-password.component';
import { AccountActivationInfoComponent } from './components/account-activation-info/account-activation-info.component';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { AllUsersComponent } from './components/all-users/all-users.component';
import { FollowRequestsComponent } from './components/follow-requests/follow-requests.component';
import {MatListModule} from '@angular/material/list';
import {MatGridListModule} from '@angular/material/grid-list';
import { JobOffersViewComponent } from './components/job-offers-view/job-offers-view.component';




@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    HomepageComponent,
    PostsViewComponent,
    LoginComponent,
    RegistrationComponent,
    ProfileComponent,
    ChangePasswordComponent,
    ForgotPasswordComponent,
    SetNewPasswordComponent,
    AccountActivationInfoComponent,
    AllUsersComponent,
    FollowRequestsComponent,
    JobOffersViewComponent

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
    MatSnackBarModule,
    MatCardModule,
    MatSnackBarModule,
    MatProgressSpinnerModule,
    MatListModule,
    MatGridListModule
  ],
  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      useClass: TokenInterceptor,
      multi: true
    },
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
