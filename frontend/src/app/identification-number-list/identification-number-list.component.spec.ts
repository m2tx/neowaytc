import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IdentificationNumberListComponent } from './identification-number-list.component';

describe('IdentificationNumberListComponent', () => {
  let component: IdentificationNumberListComponent;
  let fixture: ComponentFixture<IdentificationNumberListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ IdentificationNumberListComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(IdentificationNumberListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
