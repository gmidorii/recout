import { RootState } from "~/types";
import { MutationTree, ActionTree } from "vuex";

export const state = (): RootState => ({
  people: []
});

export const mutations: MutationTree<RootState> = {};

export const actions: ActionTree<RootState, RootState> = {};
