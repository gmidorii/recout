import { auth } from "~/plugins/firebaseinit";

export default function({ store, route, error }) {
  if (!store.getters.isLoggedIn) {
    const ignores = ["index", "login", "logout"];
    if (!ignores.includes(route.name)) {
      error({
        message: "Not connected",
        statusCode: 403
      });
    }
  }
}
