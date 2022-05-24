import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AccountActivationInfoComponent } from './components/account-activation-info/account-activation-info.component';
import { ForgotPasswordComponent } from './components/forgot-password/forgot-password.component';
import { LoginComponent } from './components/login/login.component';
import { PostsViewComponent } from './components/posts-view/posts-view.component';
import { RegistrationComponent } from './components/registration/registration.component';
import { SetNewPasswordComponent } from './components/set-new-password/set-new-password.component';

const routes: Routes = [
  {
    path: 'login',
    component: LoginComponent,
  },
  {
    path: 'registration',
    component: RegistrationComponent,
  },
  {
    path: 'posts',
    component: PostsViewComponent,
  },
  {
    path: 'forgot-password',
    component: ForgotPasswordComponent
  },
  {
    path: 'set-password',
    component: SetNewPasswordComponent
  },
  {
    path: 'account-activation',
    component: AccountActivationInfoComponent
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
