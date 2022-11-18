import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';
import { Settings } from './settings.type';
import { NgxIndexedDBService } from 'ngx-indexed-db';

@Injectable({
  providedIn: 'root'
})
export class SettingsService {
  private settings = {} as Settings
  settings$ = new BehaviorSubject<Settings>(this.settings);

  constructor(
    private db: NgxIndexedDBService,
    private defaultSettings: Settings,
  ) {}

  public loadData(): Promise<Settings> {
   return new Promise<Settings>((resolve, reject) => {
    this.db.getByKey('settings', 1).subscribe((settings: any) => {
      if (settings) {

        this.settings = settings as Settings;
        this.settings$.next(this.settings);

        resolve(this.settings);
      } else {
        this.db.add('settings',this.defaultSettings).subscribe(() => {
          console.log("added settings to database")

          this.settings$.next(this.settings);

          resolve(this.settings);
        })
      }
    })
    })
  }

  private saveSettings() {
    this.db.update('settings', this.settings).subscribe(() => {
      console.log("saved settings")
    })
    this.settings$.next(this.settings);
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

    this.saveSettings();
  }

  enableTalkgroup(t: string) {
    this.settings.disabled_talkgroups = this.settings.disabled_talkgroups.filter((tg) => tg !== t);

    this.saveSettings();
  }

  checkTalkgroup(t: string): boolean {
    return !this.settings.disabled_talkgroups.includes(t);
  }

  setServerURI(uri: string) {
    this.settings.server.uri = uri;

    this.saveSettings();
  }

  setShowActiveCalls(show: boolean) {
    this.settings.show_active_calls = show;

    this.saveSettings();
  }

  setPlayUnknownTalkgroups(play: boolean) {
    this.settings.play_unknown_talkgroups = play;

    this.saveSettings();
  }
}
