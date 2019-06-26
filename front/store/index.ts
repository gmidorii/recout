import { RootState, User } from "~/types";
import { MutationTree, ActionTree } from "vuex";
import { auth } from "~/plugins/firebaseinit";
import { RepositoryFactory } from "~/repositories/RepositoryFactory";
import { user } from "~/repositories/UserRepository";

const UserRepository: user = RepositoryFactory.getUser();

export const state = (): RootState => ({
  authUser: null
});

export const getters = {
  isLoggedIn: state => state.authUser !== null,
  userId: state => state.authUser.id
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
  },
  async loginUser({ state, commit }, user: User) {
    if (state.authUser === null || state.authUser.id !== user.id) {
      commit("setUser", user);
    }

    const currentUser = await UserRepository.get(user.id);
    if (!currentUser) {
      try {
        await UserRepository.post(user);
      } catch (e) {
        console.log(e);
        return;
      }
    }
  }
};
