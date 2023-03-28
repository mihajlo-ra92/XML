import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateFlightPageComponent } from './create-flight-page.component';

describe('CreateFlightPageComponent', () => {
  let component: CreateFlightPageComponent;
  let fixture: ComponentFixture<CreateFlightPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreateFlightPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateFlightPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
