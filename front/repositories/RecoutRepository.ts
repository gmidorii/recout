import axios from "axios";

const baseURL = `${process.env.recoutUrl}`;
const resource = "recout";

const Repository = axios.create({
  baseURL
});

export interface recout {
  get(limit?: number);
  post(message: string);
  getContinues(accountId?: string);
}

const repositoryImpl: recout = {
  get(limit: number = 20) {
    return Repository.get(`${resource}`, {
      params: {
        limit
      }
    });
  },

  post(message: string) {
    return Repository.post(`${resource}`, {
      message
    });
  },

  getContinues(accountId: string = "gmidorii") {
    return Repository.get(`${resource}/continues`, {
      params: {
        account_id: accountId
      }
    });
  }
};

export default repositoryImpl;
