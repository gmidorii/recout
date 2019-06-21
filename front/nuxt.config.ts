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
        content: "Nuxt.js TypeScript project"
      }
    ],
    link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }]
  },
  loading: { color: "#3B8070" },
  css: ["~/assets/css/main.css", "firebaseui-ja/dist/firebaseui.css"],
  build: {},
  modules: ["@nuxtjs/axios", "@nuxtjs/vuetify"],
  axios: {},
  router: {
    middleware: ["auth"]
  },
  plugins: ["~/plugins/firebase.ts"]
};
