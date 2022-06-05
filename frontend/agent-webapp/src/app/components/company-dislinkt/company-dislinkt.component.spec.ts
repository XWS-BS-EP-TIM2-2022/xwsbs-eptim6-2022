import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CompanyDislinktComponent } from './company-dislinkt.component';

describe('CompanyDislinktComponent', () => {
  let component: CompanyDislinktComponent;
  let fixture: ComponentFixture<CompanyDislinktComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CompanyDislinktComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CompanyDislinktComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
