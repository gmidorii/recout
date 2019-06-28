<template>
  <section>
    <link rel="apple-touch-icon" sizes="120x120" href="/apple-touch-icon.png">
    <link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png">
    <link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png">
    <link rel="manifest" href="/site.webmanifest">
    <link rel="mask-icon" href="/safari-pinned-tab.svg" color="#5bbad5">
    <meta name="msapplication-TileColor" content="#da532c">
    <meta name="theme-color" content="#ffffff">

    <v-navigation-drawer v-model="drawer" cliped fixed app :width="drawerWidth">
      <v-list>
        <v-list-tile to="/">
          <v-list-tile-title>{{appName}}</v-list-tile-title>
        </v-list-tile>
      </v-list>
      <v-divider></v-divider>
      <v-list>
        <v-list-tile v-for="content in getDrawerContent" :key="content.name" :to="content.to">
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
  </section>
</template>

<script lang="ts">
import Vue from "vue";
import { Component, Getter } from "nuxt-property-decorator";
import { DrawerContent } from "~/types";

@Component({})
export default class extends Vue {
  @Getter appName;
  @Getter isLoggedIn;
  drawerWidth = 150;

  home: string = "/";
  drawer: boolean = false;
  drawerContents: DrawerContent[] = [
    { to: "login", name: "Login" },
    { to: "logout", name: "Logout" }
  ];

  get getDrawerContent() {
    const contents = [...this.drawerContents];
    if (this.isLoggedIn) {
      contents.push({ to: "recout", name: "Mypage" });
    }
    return contents;
  }
}
</script>

<style lang="scss" scoped>
.toolbar-title {
  color: inherit;
  text-decoration: inherit;

  .title {
    position: relative;
    .image {
      position: absolute;
      bottom: 0;
    }
  }
}
</style>
