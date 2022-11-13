import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { Settings } from './settings.type';
import { NgxIndexedDBService } from 'ngx-indexed-db';

@Injectable({
  providedIn: 'root'
})
export class SettingsService {
  // private settings: Settings = {
  //   id: 1,
  //   server: {
  //     host: "localhost:8080"
  //   },
  //   showActiveCalls: true,
  //   playUnknownTalkgroups: true,
  //   disabledTalkgroups: [],
  // } as Settings;

  disabledTalkgroups$ = new BehaviorSubject<string[]>(this.settings.disabled_talkgroups);

  constructor(
    private db: NgxIndexedDBService,
    private settings: Settings,
  ) {

    this.db.getByKey('settings', 1).subscribe((settings: any) => {
      if (settings) {
        this.settings = settings as Settings;
        this.disabledTalkgroups$.next(this.settings.disabled_talkgroups);
      } else {
        this.db.add('settings',this.settings).subscribe(() => {
          console.log("added settings to database")
        })
      }
    })
  }

  private saveSettings() {
    this.db.update('settings', this.settings).subscribe(() => {
      console.log("saved settings")
    })
  }

  getSettings(): Settings {
    return this.settings;
  }

  disableTalkgroup(t: string) {
    if (this.settings.disabled_talkgroups.includes(t)) {
      return
    }

    this.settings.disabled_talkgroups.push(t);

    console.log("disabled talkgroups", this.settings.disabled_talkgroups);

    this.disabledTalkgroups$.next(this.settings.disabled_talkgroups);

    this.saveSettings();
  }

  enableTalkgroup(t: string) {
    this.settings.disabled_talkgroups = this.settings.disabled_talkgroups.filter((tg) => tg !== t);

    this.disabledTalkgroups$.next(this.settings.disabled_talkgroups);

    this.saveSettings();
  }

  checkTalkgroup(t: string): boolean {
    return !this.settings.disabled_talkgroups.includes(t);
  }
}
