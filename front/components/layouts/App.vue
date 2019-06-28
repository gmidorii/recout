<template>
  <section>
    <Header/>
    <v-app dark>
      <v-navigation-drawer v-model="drawer" cliped fixed app :width="drawerWidth">
        <v-list>
          <v-list-tile to="/">
            <v-list-tile-title>{{appName}}</v-list-tile-title>
          </v-list-tile>
        </v-list>
        <v-divider></v-divider>
        <v-list>
          <v-list-tile v-for="content in drawerContents" :key="content.name" :to="content.to">
            <v-list-tile-content>{{ content.name}}</v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-navigation-drawer>
      <v-toolbar app fixed cliped-left>
        <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
        <v-toolbar-title>
          <router-link to="/" class="toolbar-title">{{appName}}</router-link>
        </v-toolbar-title>
      </v-toolbar>
      <v-content>
        <v-container class="content" fluid>
          <slot></slot>
        </v-container>
      </v-content>
    </v-app>
    <Footer class="footer"/>
  </section>
</template>

<script lang="ts">
import Vue from "vue";
import { Component, Getter } from "nuxt-property-decorator";
import Header from "~/components/layouts/Header.vue";
import Footer from "~/components/layouts/Footer.vue";

export interface DrawerContent {
  to: string;
  name: string;
}

@Component({
  components: {
    Header,
    Footer
  }
})
export default class extends Vue {
  @Getter appName;
  drawerWidth = 150;

  home: string = "/";
  drawer: boolean = false;
  drawerContents: DrawerContent[] = [
    { to: "login", name: "Login" },
    { to: "logout", name: "Logout" }
  ];
}
</script>

<style lang="scss" scoped>
.content {
  margin: 0 auto;
  width: 85%;
}
.toolbar-title {
  color: inherit;
  text-decoration: inherit;
}
</style>
