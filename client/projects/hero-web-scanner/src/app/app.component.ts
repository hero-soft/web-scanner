import { Component, OnInit } from '@angular/core';
import { Call } from './call/call.type';
import { SettingsService } from './settings/settings.service';
import { PlayerService } from './player/player.service';
// import { KnownTalkgroups, Talkgroup } from './talkgroup.type';
import { ConnectionStatus, WebsocketService } from './websocket.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit{
  title = 'HeroWebScanner';
  status: ConnectionStatus = ConnectionStatus.DISCONNECTED;

  constructor(
    private settingsService: SettingsService,
    private ws: WebsocketService,
    private playerService: PlayerService,
  ) {}

  enabled = true;
  nowPlaying: Call | undefined;
  activeTGs: Call[] = []

  callQueue: Call[] = []
  historyQueue: Call[] = []
  enableButtonText = "Disable Audio"

  ngOnInit() {

    this.ws.connectionStatus$.subscribe(status => {
      this.status = status;
    })

    this.playerService.playing$.subscribe((track) => {
      if (!track) {
        this.nowPlaying = undefined;
        return
      }

      this.callQueue = this.callQueue.filter((call) => {
        if (call.id !== track.id) {
          return true
        }

        this.historyQueue.unshift(call)

        return false
      })

      this.nowPlaying = this.historyQueue.find((call) => {
        return call.id === track.id;
      })

    })

    console.log("connecting to websocket");

    this.ws.messages$.subscribe({
      next: (message) => {
        let messageParsed: any = JSON.parse(message.data);

        console.log("message", messageParsed.type);

        if (messageParsed.type === "calls_active") {
          console.log("calls_active", messageParsed);

          // let calls: Call[] = []

          if (messageParsed.calls === "") {
            this.activeTGs = [];
            // console.log("calls_active", this.calls);
            return
          }

          this.activeTGs = messageParsed.calls

        }

        if (messageParsed.type === "audio") {
          console.log("audio", messageParsed);

          let call: Call = messageParsed.call

          if (!this.enabled) {
            this.historyQueue.unshift(call);
            return
          }

          if (!this.settingsService.checkTalkgroup(call.talkgroup.id)) {
            console.log("skipping call", call.talkgroup.id, call.talkgroup.name);
            return
          }

          this.callQueue.push(call);
          console.log("audio", call);

          this.playerService.enqueue({
            id: call.id,
            file: call.file,
            priority: 1000,
          })

        }
      },
      error: (err) => {
        console.log(err);
      },
      complete: () => {
        console.log('complete');
      }
    })
  }


  avoid(tgid: string | undefined) {
    if (!tgid) {
      return
    }

    if (tgid === this.nowPlaying?.talkgroup.id) {
      this.playerService.dequeue(this.nowPlaying.id);
    }

    this.settingsService.disableTalkgroup(tgid);

    this.callQueue = this.callQueue.filter(call => {
      if (call.talkgroup.id !== tgid) {
        return true
      }

      this.playerService.dequeue(call.id);

      return false
    })
  }

  unavoid(tgid: string | undefined) {
    if (!tgid) {
      return
    }

    this.settingsService.enableTalkgroup(tgid);
  }

  skip(id: string | undefined) {
    if (!id) {
      return
    }

    this.playerService.dequeue(id);

    this.callQueue = this.callQueue.filter(call => {
      if (call.id !== id) {
        return true
      }

      this.historyQueue.unshift(call)

      return false
    })

  }

  enableClick() {
    if (this.enabled) {
      this.playerService.disable()
      this.enabled = false;
      this.enableButtonText = "Enable Audio"

      this.historyQueue.unshift(...this.callQueue);
      this.callQueue = [];
      this.nowPlaying = undefined;

      return
    }

    if (!this.enabled) {
      this.playerService.enable()
      this.enabled = true;
      this.enableButtonText = "Disable Audio"
      return
    }
  }

  sortHistory() {
    // this.historyQueue.sort((a, b) => {
    //   return a.id - b.id;
    // })
  }
}
