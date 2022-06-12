import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IdentificationNumberFormComponent } from './identification-number-form.component';

describe('IdentificationNumberFormComponent', () => {
  let component: IdentificationNumberFormComponent;
  let fixture: ComponentFixture<IdentificationNumberFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ IdentificationNumberFormComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(IdentificationNumberFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
