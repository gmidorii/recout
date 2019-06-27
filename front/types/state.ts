export interface RootState {
  appName: string;
  authUser: User;
}

export interface User {
  id: string;
  name: string;
  session?: string;
}
