import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { environment } from '../../environments/environment';
import { Talkgroup } from './talkgroup.type';

@Injectable({
  providedIn: 'root'
})
export class TalkgroupService {
  private talkgroups: Talkgroup[] = []

  talkgroups$: BehaviorSubject<Talkgroup[]> = new BehaviorSubject(this.talkgroups)

  constructor(
    private http: HttpClient,
  ) {
    this.updateTalkgroups()
  }

  updateTalkgroups(){
    this.http.get<Talkgroup[]>(environment.serverURL + 'talkgroups').subscribe(talkgroups => {

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
