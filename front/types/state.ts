import { UrlWithStringQuery } from "url";

export interface RootState {
  appName: string;
  userState: UserState;
}

export interface UserState {
  created: boolean;
  user: User;
}

export interface User {
  id: string;
  name: string;
  session?: string;
  pixelaName?: string;
  pixelaGraph?: string;
}
