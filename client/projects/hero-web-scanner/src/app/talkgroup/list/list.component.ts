import { Component, Input, OnInit } from '@angular/core';
import { SettingsService } from '../../settings/settings.service';
import { Talkgroup } from '../talkgroup.type';

@Component({
  selector: 'talkgroup-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.scss']
})
export class TalkgroupListComponent implements OnInit {
  @Input() talkgroups: Talkgroup[] = []

  constructor(private settingsService: SettingsService) { }

  ngOnInit(): void {

  }

  changeState(tgID: string, state: boolean){
    if (!state){
      this.settingsService.enableTalkgroup(tgID)
    } else {
      this.settingsService.disableTalkgroup(tgID)
    }
  }

}
