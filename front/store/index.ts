import { RootState } from "~/types";
import { MutationTree, ActionTree } from "vuex";
import { auth } from "~/plugins/firebaseinit";

export const state = (): RootState => ({
  authUser: null
});

export const getters = {
  isLoggedIn: state => state.authUser !== null
};

export const mutations: MutationTree<RootState> = {
  setUser: (state, user) => {
    state.authUser = user;
  },
  resetUser: state => {
    state.authUser = null;
  }
};

export const actions: ActionTree<RootState, RootState> = {
  async resetUser({ state, commit }) {
    await auth.signOut();
    commit("resetUser");
  }
};
