import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AdminRequestsComponent } from './components/admin-requests/admin-requests.component';
import { LoginComponent } from './components/login/login.component';
import { NewCompanyComponent } from './components/new-company/new-company.component';
import { NewJobOfferComponent } from './components/new-job-offer/new-job-offer.component';
import { SignupComponent } from './components/signup/signup.component';

const routes: Routes = [
  {
    path: 'login',
    component: LoginComponent,
  },
  {
    path: 'signup',
    component: SignupComponent,
  },
  {
    path: 'register-company',
    component: NewCompanyComponent,
  },
  {
    path: 'create-job-offer',
    component: NewJobOfferComponent,
  },
  {
    path: '',
    component: AdminRequestsComponent,
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
