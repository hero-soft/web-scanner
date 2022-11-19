import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AutoplayDialogComponent } from './autoplay-dialog.component';

describe('AutoplayDialogComponent', () => {
  let component: AutoplayDialogComponent;
  let fixture: ComponentFixture<AutoplayDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AutoplayDialogComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AutoplayDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
