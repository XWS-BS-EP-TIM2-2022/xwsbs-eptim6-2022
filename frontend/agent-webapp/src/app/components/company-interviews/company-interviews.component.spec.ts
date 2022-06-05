import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CompanyInterviewsComponent } from './company-interviews.component';

describe('CompanyInterviewsComponent', () => {
  let component: CompanyInterviewsComponent;
  let fixture: ComponentFixture<CompanyInterviewsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CompanyInterviewsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CompanyInterviewsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
