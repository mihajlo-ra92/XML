import { TestBed } from '@angular/core/testing';

import { CancelingReservationService } from './canceling-reservation.service';

describe('CancelingReservationService', () => {
  let service: CancelingReservationService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(CancelingReservationService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
