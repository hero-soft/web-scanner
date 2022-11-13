import { Component, Input, OnInit } from '@angular/core';
import { Talkgroup } from '../talkgroup.type';

@Component({
  selector: 'talkgroup-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.scss']
})
export class TalkgroupListComponent implements OnInit {
  @Input() talkgroups: Talkgroup[] = []

  constructor() { }

  ngOnInit(): void {
  }

}
