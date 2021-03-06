export default {
  env: {
    pixelaUrl: "https://pixe.la/v1",
    recoutUrl: process.env.RECOUT_URL
  },
  mode: "spa",
  head: {
    title: "recout",
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      {
        hid: "description",
        name: "description",
        content: "Record + Output = Recout"
      }
    ],
    link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }]
  },
  loading: { color: "#3B8070" },
  css: ["~/assets/css/main.css", "firebaseui-ja/dist/firebaseui.css"],
  build: {
    quiet: false
  },
  modules: ["@nuxtjs/axios", "@nuxtjs/vuetify"],
  axios: {},
  router: {
    middleware: ["auth"]
  },
  plugins: ["~/plugins/firebase.ts"],
  vuetify: {
    theme: {
      primary: "#00695C",
      secondary: "#3f51b5",
      accent: "#ff5722j",
      error: "#f44336",
      warning: "#ff9800",
      info: "#03a9f4",
      success: "#4caf50"
    }
  }
};
