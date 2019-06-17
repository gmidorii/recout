import { RootState } from "~/types";
import { MutationTree, ActionTree } from "vuex";

export const state = (): RootState => ({
  authUser: null
});

export const mutations: MutationTree<RootState> = {
  setUser: (state, user) => {
    state.authUser = user;
  }
};

export const actions: ActionTree<RootState, RootState> = {};
