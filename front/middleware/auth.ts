export default function({ store, error }) {
  if (!store.state.authUser) {
    error({
      message: "Not connected",
      statusCode: 403
    });
  }
}
