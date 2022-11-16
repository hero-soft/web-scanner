// src\app\services\websocket.service.ts
import { Injectable } from "@angular/core";
import { BehaviorSubject, Observable, Observer } from 'rxjs';
import { AnonymousSubject, Subject } from 'rxjs/internal/Subject';
import { SettingsService } from "./settings/settings.service";

export interface Message {
    source: string;
    content: string;
}

export enum ConnectionStatus {
    CONNECTED = "Connected",
    DISCONNECTED = "Disconnected",
    CONNECTING = "Connecting",
}

@Injectable()
export class WebsocketService {
    // private subject: AnonymousSubject<MessageEvent> | undefined;
    // public messages: Subject<Message>;

    private connectionStatus: ConnectionStatus = ConnectionStatus.DISCONNECTED;
    public connectionStatus$ = new BehaviorSubject<ConnectionStatus>(this.connectionStatus);
    public messages$ = new Subject<any>();

    constructor(
      private settings: SettingsService
    ) {
      this.connect(this.settings.getSettings().server.host)
    }

    // public connect(url: string): AnonymousSubject<MessageEvent> {
    //     if (!this.subject) {
    //         this.subject = this.create(url);
    //     }
    //     return this.subject;
    // }

    private connect(url: string) {
        let ws = new WebSocket(url);

        this.connectionStatus = ConnectionStatus.CONNECTING;
        this.connectionStatus$.next(this.connectionStatus);

        ws.addEventListener('open', () => {
            this.connectionStatus = ConnectionStatus.CONNECTED;
            this.connectionStatus$.next(this.connectionStatus);
            console.log("Websocket open: " + url);
        })

        ws.addEventListener('close', () => {
            this.connectionStatus = ConnectionStatus.CONNECTING;
            this.connectionStatus$.next(this.connectionStatus);
            console.log("Websocket closed: " + url);

            setTimeout(() => {
                if (this.connectionStatus !== ConnectionStatus.CONNECTED) {
                  this.connect(url)
                }
            }, 5000)
        })

        ws.addEventListener('message', (message) => {
            this.messages$.next(message);
        })


        // let observable = new Observable((obs: Observer<MessageEvent>) => {
        //     ws.onmessage = obs.next.bind(obs);
        //     ws.onerror = obs.error.bind(obs);
        //     ws.onclose = obs.complete.bind(obs);
        //     return ws.close.bind(ws);
        // });

        // let observer = {
        //     error: (err: any) => {
        //       console.log(err);
        //     },
        //     complete: () => {},
        //     next: (data: Object) => {
        //         console.log('Message sent to websocket: ', data);
        //         if (ws.readyState === WebSocket.OPEN) {
        //             ws.send(JSON.stringify(data));
        //         }
        //     }
        // };

        // console.log("Successfully created websocket");
        // return new AnonymousSubject<MessageEvent>(observer, observable);
    }
}
