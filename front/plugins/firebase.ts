import { auth } from "~/plugins/firebaseinit";

export default function({ store, route, error }) {
  return new Promise((resolve, reject) => {
    auth.onAuthStateChanged(user => {
      if (user) {
        store.dispatch("loginUser", { id: user.uid, name: user.displayName });
      }
      resolve();
    });
  });
}
