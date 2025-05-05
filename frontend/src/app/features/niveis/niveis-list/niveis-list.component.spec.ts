import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NiveisListComponent } from './niveis-list.component';

describe('NiveisListComponent', () => {
  let component: NiveisListComponent;
  let fixture: ComponentFixture<NiveisListComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [NiveisListComponent]
    });
    fixture = TestBed.createComponent(NiveisListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
