export interface Talkgroup {
    id: string,
    name: string,
}

export var KnownTalkgroups: Record<string, Talkgroup> =  {
    "33776": {id: "33776", name: "C-PTL-1"},
    "33840": {id: "33840", name: "C-PTL-3"},
    "39152": {id: "39152", name: "MA Environmental Police 1 - East Dispatch"},
    "33168": {id: "33168", name: "North Dispatch - A, A1, A2, A3, A6"},
    "": {id: "", name: ""},
}
