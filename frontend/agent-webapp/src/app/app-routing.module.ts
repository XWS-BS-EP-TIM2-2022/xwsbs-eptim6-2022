import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AdminRequestsComponent } from './components/admin-requests/admin-requests.component';
import { AllCompaniesComponent } from './components/all-companies/all-companies.component';
import { CompanyProfileComponent } from './components/company-profile/company-profile.component';
import { JobOffersComponent } from './components/job-offers/job-offers.component';
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
    path: 'admin',
    component: AdminRequestsComponent,
  },
  { 
    path: 'company/:id',
    component: CompanyProfileComponent,
  },
  {
    path: 'job-offers',
    component: JobOffersComponent,
  },
  {
    path: 'all-companies',
    component: AllCompaniesComponent,
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
