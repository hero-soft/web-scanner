import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { Config } from './config.type';
import { NgxIndexedDBService } from 'ngx-indexed-db';

@Injectable({
  providedIn: 'root'
})
export class ConfigService {
  private config: Config = {
    id: 1,
    server: {
      host: "localhost:8080"
    },
    showActiveCalls: true,
    playUnknownTalkgroups: true,
    disabledTalkgroups: [],
  } as Config;

  disabledTalkgroups$ = new BehaviorSubject<string[]>(this.config.disabledTalkgroups);

  constructor(
    private db: NgxIndexedDBService,
  ) {

    this.db.getByKey('config', 1).subscribe((config: any) => {
      if (config) {
        this.config = config as Config;
        this.disabledTalkgroups$.next(this.config.disabledTalkgroups);
      } else {
        this.db.add('config',this.config).subscribe(() => {
          console.log("added config to database")
        })
      }
    })
  }

  private saveConfig() {
    this.db.update('config', this.config).subscribe(() => {
      console.log("saved config")
    })
  }

  getConfig(): Config {
    return this.config;
  }

  disableTalkgroup(t: string) {
    if (this.config.disabledTalkgroups.includes(t)) {
      return
    }

    this.config.disabledTalkgroups.push(t);

    console.log("disabled talkgroups", this.config.disabledTalkgroups);

    this.disabledTalkgroups$.next(this.config.disabledTalkgroups);

    this.saveConfig();
  }

  enableTalkgroup(t: string) {
    this.config.disabledTalkgroups = this.config.disabledTalkgroups.filter((tg) => tg !== t);

    this.disabledTalkgroups$.next(this.config.disabledTalkgroups);

    this.saveConfig();
  }

  checkTalkgroup(t: string): boolean {
    return !this.config.disabledTalkgroups.includes(t);
  }
}
