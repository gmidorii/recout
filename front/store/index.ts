import { RootState, User } from "~/types";
import { MutationTree, ActionTree } from "vuex";
import { auth } from "~/plugins/firebaseinit";
import { RepositoryFactory } from "~/repositories/RepositoryFactory";
import { user } from "~/repositories/UserRepository";

const UserRepository: user = RepositoryFactory.getUser();

export const state = (): RootState => ({
  appName: "(β) Recout",
  userState: {
    created: false,
    user: null
  }
});

export const getters = {
  isLoggedIn: state => state.userState.user !== null,
  appName: state => state.appName,
  userId: state => {
    if (state.userState.user) {
      return state.userState.user.id;
    }
    return "";
  },
  pixelaGraphUrl: state => {
    if (!state.userState.created) {
      return "";
    }
    const user = state.userState.user;
    return `${process.env.pixelaUrl}/users/${user.pixelaName}/graphs/${
      user.pixelaGraph
    }`;
  }
};

export const mutations: MutationTree<RootState> = {
  setUser: (state, user) => {
    state.userState.user = Object.assign({}, user);
  },
  setUserDetail: (state, { userName, graph }) => {
    state.userState.user = Object.assign(state.userState.user, {
      pixelaName: userName,
      pixelaGraph: graph
    });
    state.userState.created = true;
  },
  resetUser: state => {
    state.userState.user = null;
    state.userState.created = false;
  }
};

export const actions: ActionTree<RootState, RootState> = {
  async resetUser({ state, commit }) {
    await auth.signOut();
    commit("resetUser");
  },
  async loginUser({ state, commit }, user: User) {
    if (state.userState.user === null || state.userState.user.id !== user.id) {
      commit("setUser", user);
    }

    const currentUser = await UserRepository.get(user.id);
    if (currentUser) {
      commit("setUserDetail", {
        userName: currentUser.account_id,
        graph: currentUser.pixela_graph
      });
      return;
    }

    try {
      await UserRepository.post(user);
      const createdUser = await UserRepository.get(user.id);
      commit("setUserDetail", {
        userName: createdUser.account_id,
        graph: createdUser.pixela_graph
      });
    } catch (e) {
      console.log(e);
      return;
    }
  }
};
