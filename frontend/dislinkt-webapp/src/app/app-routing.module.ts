import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { PostsViewComponent } from './components/posts-view/posts-view.component';
import { RegistrationComponent } from './components/registration/registration.component';
import { ProfileComponent } from './components/profile/profile.component';
import { AccountActivationInfoComponent } from './components/account-activation-info/account-activation-info.component';
import { ForgotPasswordComponent } from './components/forgot-password/forgot-password.component';
import { SetNewPasswordComponent } from './components/set-new-password/set-new-password.component';
import { AdminGuard } from './auth-guard/admin-guard';
import { UserGuard } from './auth-guard/user-guard';
import { ChangePasswordComponent } from './components/change-password/change-password.component';
import { HomepageComponent } from './components/homepage/homepage.component';
import { AllUsersComponent } from './components/all-users/all-users.component';
import { FollowRequestsComponent } from './components/follow-requests/follow-requests.component';

const routes: Routes = [
  {
    path: 'login',
    component: LoginComponent,
  },
  {
    path: 'forgot-password',
    component: ForgotPasswordComponent,
  },
  {
    path: 'registration',
    component: RegistrationComponent,
  },
  {
    path: 'posts',
    component: PostsViewComponent,
    canActivate: [AdminGuard]
  },
  {
    path: 'profile',
    component: ProfileComponent,
    canActivate: [UserGuard]
  },
  {
    path: 'change-password',
    component: ChangePasswordComponent,
    canActivate: [UserGuard]
  },
  {
    path: 'home',
    component: HomepageComponent,
    canActivate: [UserGuard]
  },
  {
    path: '',
    component: HomepageComponent,
    canActivate: [UserGuard]
  },
  {
    path: 'set-password/:token',
    component: SetNewPasswordComponent
  },
  {
    path: 'account-activation/:token',
    component: AccountActivationInfoComponent
  },
  {
    path: 'passwordless/:token',
    component: PostsViewComponent,
  },
  {
    path: 'users',
    component: AllUsersComponent,
  },
  {
    path: 'follow-requests',
    component: FollowRequestsComponent,
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
