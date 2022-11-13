import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { TalkgroupDialogComponent } from '../../talkgroup/dialog/dialog.component';
import { ConfigService } from '../config.service';
import { Config } from '../config.type';

@Component({
  selector: 'app-config',
  templateUrl: './config.component.html',
  styleUrls: ['./config.component.scss']
})
export class ConfigComponent implements OnInit {
  config: Config = {} as Config;

  constructor(
    private ConfigService: ConfigService,
    public dialog: MatDialog,
  ) { }

  ngOnInit(): void {
    this.config = this.ConfigService.getConfig();
  }

  openDialog(): void {
    this.dialog.open(TalkgroupDialogComponent, {
      width: '90%',
    });
  }

}
