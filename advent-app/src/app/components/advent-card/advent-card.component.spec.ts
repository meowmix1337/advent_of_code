import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AdventCardComponent } from './advent-card.component';

describe('AdventCardComponent', () => {
  let component: AdventCardComponent;
  let fixture: ComponentFixture<AdventCardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AdventCardComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AdventCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
