import { TestBed } from '@angular/core/testing';

import { DesenvolvedoresService } from './desenvolvedores.service';

describe('DesenvolvedoresService', () => {
  let service: DesenvolvedoresService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(DesenvolvedoresService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
