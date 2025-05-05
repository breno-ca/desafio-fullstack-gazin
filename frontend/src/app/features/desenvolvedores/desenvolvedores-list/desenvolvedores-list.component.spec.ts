import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DesenvolvedoresListComponent } from './desenvolvedores-list.component';

describe('DesenvolvedoresListComponent', () => {
  let component: DesenvolvedoresListComponent;
  let fixture: ComponentFixture<DesenvolvedoresListComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [DesenvolvedoresListComponent]
    });
    fixture = TestBed.createComponent(DesenvolvedoresListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
