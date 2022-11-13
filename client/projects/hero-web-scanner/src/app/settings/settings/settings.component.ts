import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { TalkgroupDialogComponent } from '../../talkgroup/dialog/dialog.component';
import { SettingsService } from '../settings.service';
import { Settings } from '../settings.type';

@Component({
  selector: 'app-config',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.scss']
})
export class SettingsComponent implements OnInit {
  settings: Settings = {} as Settings;

  constructor(
    private settingsService: SettingsService,
    public dialog: MatDialog,
  ) { }

  ngOnInit(): void {
    this.settings = this.settingsService.getSettings();
  }

  openDialog(): void {
    this.dialog.open(TalkgroupDialogComponent, {
      width: '90%',
    });
  }

}
