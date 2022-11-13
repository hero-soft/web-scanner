import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Call } from '../call.type';
import { ConfigService } from '../../config/config.service';

@Component({
  selector: 'app-call',
  templateUrl: './call.component.html',
  styleUrls: ['./call.component.scss']
})
export class CallComponent implements OnInit {
  @Input() call: Call | undefined;
  @Input() showAvoid: boolean = true;
  @Input() showSkip: boolean = true;

  @Output() avoidEvent = new EventEmitter<null>();
  @Output() unavoidEvent = new EventEmitter<null>();
  @Output() skipEvent = new EventEmitter<null>();

  disabled = false

  constructor(
    private configService: ConfigService,
  ) { }

  ngOnInit(): void {
    this.configService.disabledTalkgroups$.subscribe((disabled) => {
      this.disabled = disabled.includes(this.call?.talkgroup.id as string)
    })
  }

  avoid() {
    this.avoidEvent.emit();
  }

  unavoid() {
    this.unavoidEvent.emit();
  }

  skip() {
    this.skipEvent.emit();
  }

}
