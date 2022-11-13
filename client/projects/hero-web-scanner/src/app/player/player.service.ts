import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { environment } from '../../environments/environment';

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
  private baseURL = environment.serverURL
  private enabled = true

  public enabled$ = new BehaviorSubject<boolean>(this.enabled)
  public playing$: BehaviorSubject<Track | undefined> = new BehaviorSubject(this.playing)

  constructor() {
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
        this.player.play();
      }

      return
    }

    if (this.player.paused) {
      this.playing = undefined
      this.playing$.next(this.playing)
    }
  }
}
