import { Component, OnInit } from '@angular/core';
import { TalkgroupService } from '../talkgroup.service';
import { Talkgroup } from '../talkgroup.type';

@Component({
  selector: 'app-dialog',
  templateUrl: './dialog.component.html',
  styleUrls: ['./dialog.component.scss']
})
export class TalkgroupDialogComponent implements OnInit {
  talkgroups: Talkgroup[] = []

  constructor(
    private tgService: TalkgroupService,
  ) { }

  ngOnInit(): void {
    this.tgService.talkgroups$.subscribe(talkgroups => {
      this.talkgroups = talkgroups
    })
  }

}
