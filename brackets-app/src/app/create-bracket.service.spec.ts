import { TestBed } from '@angular/core/testing';

import { CreateBracketService } from './create-bracket.service';

describe('CreateBracketService', () => {
  let service: CreateBracketService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(CreateBracketService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
