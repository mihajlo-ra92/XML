import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CancelingReservationComponent } from './canceling-reservation.component';

describe('CancelingReservationComponent', () => {
  let component: CancelingReservationComponent;
  let fixture: ComponentFixture<CancelingReservationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CancelingReservationComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CancelingReservationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
