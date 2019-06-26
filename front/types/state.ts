export interface RootState {
  authUser: User;
}

export interface User {
  id: string;
  name: string;
  session?: string;
}
