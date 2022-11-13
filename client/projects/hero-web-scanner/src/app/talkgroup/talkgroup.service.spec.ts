import { TestBed } from '@angular/core/testing';

import { TalkgroupService } from './talkgroup.service';

describe('TalkgroupService', () => {
  let service: TalkgroupService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TalkgroupService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
