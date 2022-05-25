import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { PostsViewComponent } from './components/posts-view/posts-view.component';
import { RegistrationComponent } from './components/registration/registration.component';
import { ProfileComponent } from './components/profile/profile.component';
import { AdminGuard } from './auth-guard/admin-guard';
import { UserGuard } from './auth-guard/user-guard';
import { ChangePasswordComponent } from './components/change-password/change-password.component';
import { HomepageComponent } from './components/homepage/homepage.component';

const routes: Routes = [
  {
    path : 'login',
    component : LoginComponent,
  },
  {
    path : 'registration',
    component : RegistrationComponent,
  },
  {
    path : 'posts',
    component : PostsViewComponent,
    canActivate: [UserGuard]
  },
  {
    path: 'profile',
    component: ProfileComponent,
    canActivate: [UserGuard]
  },
  {
    path : 'change-password',
    component : ChangePasswordComponent,
    canActivate: [UserGuard]
  },
  {
    path : 'homepage',
    component : HomepageComponent,
    canActivate: [UserGuard]
  },
  {
    path : '',
    component : HomepageComponent,
    canActivate: [UserGuard]
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
