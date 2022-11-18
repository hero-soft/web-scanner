import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { TalkgroupDialogComponent } from '../../talkgroup/dialog/dialog.component';
import { SettingsService } from '../settings.service';
import { Settings } from '../settings.type';
import { FormBuilder } from '@angular/forms';

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
    private fb: FormBuilder
  ) { }

  settingsForm = this.fb.group({
    server: this.fb.group({
      uri: [''],
    }),
    show_active_calls: [true],
    play_unknown_talkgroups: [true],
  });

  ngOnInit(): void {
    this.settings = this.settingsService.getSettings();

    this.settingsForm.controls.server.setValue(this.settings.server);
    this.settingsForm.controls.show_active_calls.setValue(this.settings.show_active_calls);
    this.settingsForm.controls.play_unknown_talkgroups.setValue(this.settings.play_unknown_talkgroups);
  }

  openDialog(): void {
    this.dialog.open(TalkgroupDialogComponent, {
      width: '90%',
    });
  }

  save() {
    if (this.settingsForm.valid) {
      this.settingsService.setServerURI(this.settingsForm.controls.server.controls.uri.value as string);
      this.settingsService.setShowActiveCalls(this.settingsForm.controls.show_active_calls.value as boolean);
      this.settingsService.setPlayUnknownTalkgroups(this.settingsForm.controls.play_unknown_talkgroups.value as boolean);
    }
  }

}
