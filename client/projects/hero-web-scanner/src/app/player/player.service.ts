import { MatDialog } from '@angular/material/dialog';
import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { environment } from '../../environments/environment';
import { AutoplayDialogComponent } from './autoplay-dialog/autoplay-dialog.component';
import { SettingsService } from '../settings/settings.service';

export interface Track {
  id: string;
  file: string;
  priority: number;
}

@Injectable({
  providedIn: 'root'
})
export class PlayerService {
  private player = new Audio()
  private queue: Track[] = []
  private playing: Track | undefined = undefined
  private baseURL = ""
  private enabled = true
  private autoplayDialogShown = false

  public enabled$ = new BehaviorSubject<boolean>(this.enabled)
  public playing$: BehaviorSubject<Track | undefined> = new BehaviorSubject(this.playing)

  constructor(
    private dialog: MatDialog,
    private settings: SettingsService
  ) {
    this.settings.settings$.subscribe(settings => {
      this.baseURL = settings.server.uri
    })

    this.player.addEventListener('ended', () => {
      this.tryNext()
    });
    this.player.addEventListener('error', () => {
      this.tryNext()
    });
    this.player.addEventListener('abort', () => {
      this.tryNext()
    });
  }

  disable() {
    console.log("Player disabled")

    this.player.pause()
    this.queue = []
    this.enabled = false
    this.enabled$.next(this.enabled)
  }

  enable() {
    console.log("Player enabled")

    this.enabled = true
    this.enabled$.next(this.enabled)
  }

  enqueue(t: Track) {
    console.log("enqueue", t)
    if (!this.enabled) {
      return
    }

    this.queue.push(t)

    this.tryNext()
  }

  dequeue(id: string) {
    console.log("dequeue", id)

    this.queue = this.queue.filter((t) => t.id !== id)

    if (this.playing?.id === id) {
      this.player.pause()
      this.tryNext()
    }
  }

  enqueueNext(t: Track) {
    if (!this.enabled) {
      return
    }

    this.queue.unshift(t)

    this.tryNext()
  }

  clear() {
    if (!this.enabled) {
      return
    }

    this.queue = []
    this.player.pause()
  }

  // skip(id: string) {
  //   if (!this.enabled) {
  //     return
  //   }

  //   if (this.playing?.id === id) {
  //     this.player.pause()
  //   }

  //   this.queue = this.queue.filter(t => t.id !== id)

  //   this.tryNext()
  // }

  private tryNext() {
    if (!this.enabled) {
      return
    }

    if (this.queue.length > 0 && this.player.paused) {
      let next = this.queue.shift();

      if (next) {
        this.playing = next
        this.playing$.next(this.playing)

        this.player.src = this.baseURL + next.file;
        this.player.play().catch((e) => {
          if (e.name === "NotAllowedError") {
            console.log("error playing audio", e);

            if (!this.autoplayDialogShown) {
              this.autoplayDialogShown = true
              const dialog = this.dialog.open(AutoplayDialogComponent)

              dialog.afterClosed().subscribe((result) => {
                this.player.play()
              })
            }
          }
        });
      }

      return
    }

    if (this.player.paused) {
      this.playing = undefined
      this.playing$.next(this.playing)
    }
  }
}
