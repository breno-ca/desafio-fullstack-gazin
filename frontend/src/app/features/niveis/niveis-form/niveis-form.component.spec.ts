import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NiveisFormComponent } from './niveis-form.component';

describe('NiveisFormComponent', () => {
  let component: NiveisFormComponent;
  let fixture: ComponentFixture<NiveisFormComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [NiveisFormComponent]
    });
    fixture = TestBed.createComponent(NiveisFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
