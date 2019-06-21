export interface RootState {
  authUser: User;
}

export interface User {
  name: string;
  session?: string;
}
