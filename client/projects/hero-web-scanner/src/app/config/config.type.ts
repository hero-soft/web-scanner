export interface Config {
  id: number;
  server: {
    host: string;
  }
  disabledTalkgroups: string[]
  showActiveCalls: boolean;
  playUnknownTalkgroups: boolean;
}
