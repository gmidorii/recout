import RecoutRepository from "./RecoutRepository";
import UserRepository, { user } from "./UserRepository";

const respositories = {
  recout: RecoutRepository,
  user: UserRepository
};

export const RepositoryFactory = {
  getRecout: () => {
    return respositories.recout;
  },
  getUser: () => {
    return respositories.user;
  }
};
