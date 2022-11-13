import { Talkgroup } from "../talkgroup/talkgroup.type";

export interface Call {
  id: string,
  talkgroup: Talkgroup,
  emergency: boolean;
  priority: number;
  file: string;
}
