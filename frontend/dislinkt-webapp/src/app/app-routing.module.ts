import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { PostsViewComponent } from './components/posts-view/posts-view.component';
import { RegistrationComponent } from './components/registration/registration.component';
import { ProfileComponent } from './components/profile/profile.component';

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
  },
  {
    path: 'profile',
    component: ProfileComponent,
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
