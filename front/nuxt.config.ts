export default {
  env: {
    pixelaUrl: "https://pixe.la/v1",
    recoutUrl: process.env.RECOUT_URL
  },
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
  css: ["~/assets/css/main.css"],
  build: {},
  modules: ["@nuxtjs/axios", "@nuxtjs/vuetify"],
  axios: {}
};
