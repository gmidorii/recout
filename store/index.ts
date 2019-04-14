import { RootState, Person } from "~/types";
import { MutationTree, ActionTree } from "vuex";

export const state = (): RootState => ({
  people: []
});

export const mutations: MutationTree<RootState> = {
  setPeople(state: RootState, people: Person[]): void {
    state.people = people;
  }
};

export const actions: ActionTree<RootState, RootState> = {};
