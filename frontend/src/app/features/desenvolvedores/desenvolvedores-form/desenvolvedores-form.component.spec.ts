import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DesenvolvedoresFormComponent } from './desenvolvedores-form.component';

describe('DesenvolvedoresFormComponent', () => {
  let component: DesenvolvedoresFormComponent;
  let fixture: ComponentFixture<DesenvolvedoresFormComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [DesenvolvedoresFormComponent]
    });
    fixture = TestBed.createComponent(DesenvolvedoresFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
