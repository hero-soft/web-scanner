import { Component, OnInit } from '@angular/core';
import { KnownTalkgroups, Talkgroup } from './talkgroup.type';
import { WebsocketService } from './websocket.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit{
  title = 'HeroWebScanner';

  constructor(private ws: WebsocketService) {}

  messages: any[] = []
  calls: Talkgroup[] = []

  callQueue: Talkgroup[] = []

  ngOnInit() {
    console.log("connecting to websocket");
    this.ws.connect('ws://localhost:8080/ws/client').subscribe({
      next: (message) => {

        let messageParsed: any = JSON.parse(message.data);
        console.log("message", messageParsed.type);

        if (messageParsed.type === "calls_active") {
          // console.log("calls", messageParsed.calls);

          let calls: Talkgroup[] = []

          if (messageParsed.calls === "") {
            this.calls = [];
            // console.log("calls_active", this.calls);
            return
          }

          for (let call of messageParsed.calls) {
            let known = KnownTalkgroups[call.talkgroup];

            if (known == undefined) {
              calls.push({id: call.talkgroup, name: "Unknown"});
            } else {
              calls.push(known);
            }
          }

          this.calls = calls
          // console.log("calls_active", this.calls);
        } else if (messageParsed.type === "call_start") {
          //  console.log("call_active", messageParsed.call);

          let call = messageParsed.call
          let known = KnownTalkgroups[call.talkgroup];

          if (known == undefined) {
            this.callQueue.push({id: call.talkgroup, name: "Unknown"});
          } else {
            this.callQueue.push(known);
          }
        } else if (messageParsed.type === "call_end") {
          console.log("!!!END", messageParsed.call);
        } else {
          // console.log("message", messageParsed.type, messageParsed);
        }
      },
      error: (err) => {
        console.log(err);
      },
      complete: () => {
        console.log('complete');
      }
  }
  )}
}
