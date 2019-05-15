import RecoutRepository from "./RecoutRepository";

const respositories = {
  recout: RecoutRepository
};

export const RepositoryFactory = {
  getRecout: () => {
    return respositories.recout;
  }
};
