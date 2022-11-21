import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { environment } from '../../environments/environment';
import { SettingsService } from '../settings/settings.service';
import { Talkgroup } from './talkgroup.type';

@Injectable({
  providedIn: 'root'
})
export class TalkgroupService {
  private talkgroups: Talkgroup[] = []

  talkgroups$: BehaviorSubject<Talkgroup[]> = new BehaviorSubject(this.talkgroups)
  currentURI: string = ""

  constructor(
    private http: HttpClient,
    private settingsService: SettingsService,
  ) {
    this.updateTalkgroups(this.settingsService.getSettings().server.uri)

    this.settingsService.settings$.subscribe(settings => {
      if (settings.server.uri !== this.currentURI) {
        this.currentURI = settings.server.uri
        this.updateTalkgroups(this.currentURI)
      }

      this.talkgroups.forEach(talkgroup => {
        talkgroup.disabled = !this.settingsService.checkTalkgroup(talkgroup.id)
      })

    })
  }

  updateTalkgroups(base_uri: string){
    this.http.get<Talkgroup[]>("http://" + base_uri + '/talkgroups').subscribe(talkgroups => {

      talkgroups.forEach(talkgroup => {
        talkgroup.disabled = !this.settingsService.checkTalkgroup(talkgroup.id)
      })

      talkgroups.sort((a, b) => {
        if (a.name < b.name) {
          return -1;
        }
        if (a.name > b.name) {
          return 1;
        }
        return 0;
      })

      this.talkgroups = talkgroups
      this.talkgroups$.next(this.talkgroups)
    })
  }
}
