import { auth } from "~/plugins/firebaseinit";

const generateId = user => user.uid.toLowerCase();

export default function({ store, route, error }) {
  return new Promise((resolve, reject) => {
    auth.onAuthStateChanged(user => {
      if (user) {
        store.dispatch("loginUser", {
          id: generateId(user),
          name: user.displayName
        });
      }
      resolve();
    });
  });
}
