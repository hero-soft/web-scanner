import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Call } from '../call.type';
import { SettingsService } from '../../settings/settings.service';

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
    private settingsService: SettingsService,
  ) { }

  ngOnInit(): void {
    this.settingsService.settings$.subscribe((settings) => {
      if (!settings) {
        this.disabled = false
        return
      }
      this.disabled = settings.disabled_talkgroups.includes(this.call?.talkgroup.id as string)
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
