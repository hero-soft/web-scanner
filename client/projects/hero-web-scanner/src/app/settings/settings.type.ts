export class Settings {
  id!: number;
  server!: {
    host: string;
  }
  disabled_talkgroups!: string[]
  show_active_calls!: boolean;
  play_unknown_talkgroups!: boolean;
}


