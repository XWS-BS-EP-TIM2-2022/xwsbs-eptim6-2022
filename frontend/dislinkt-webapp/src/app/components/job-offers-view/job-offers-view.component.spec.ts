import { ComponentFixture, TestBed } from '@angular/core/testing';

import { JobOffersViewComponent } from './job-offers-view.component';

describe('JobOffersViewComponent', () => {
  let component: JobOffersViewComponent;
  let fixture: ComponentFixture<JobOffersViewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ JobOffersViewComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(JobOffersViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
