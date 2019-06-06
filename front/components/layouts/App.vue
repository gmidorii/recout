<template>
  <section>
    <v-app dark>
      <v-navigation-drawer v-model="drawer" cliped fixed app></v-navigation-drawer>
      <v-toolbar app fixed cliped-left>
        <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
        <v-toolbar-title>Recout</v-toolbar-title>
      </v-toolbar>
      <v-content>
        <v-container class="content" fluid>
          <slot></slot>
        </v-container>
      </v-content>
    </v-app>
  </section>
</template>

<script lang="ts">
import Vue from "vue";
import { Component } from "nuxt-property-decorator";

@Component({})
export default class extends Vue {}
</script>

<style lang="scss" scoped>
.content {
  margin: 0 auto;
  width: 90%;
}
</style>
