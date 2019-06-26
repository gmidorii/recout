import axios from "axios";

const baseURL = `${process.env.recoutUrl}`;
const resource = "recout";

const Repository = axios.create({
  baseURL
});

export interface recout {
  get(accountId: string, limit?: number);
  post(accountId: string, message: string);
  getContinues(accountId?: string);
}

const repositoryImpl: recout = {
  get(accountId: string, limit: number = 20) {
    return Repository.get(`${resource}`, {
      params: {
        account_id: accountId,
        limit
      }
    });
  },

  post(accountId: string, message: string) {
    return Repository.post(`${resource}`, {
      account_id: accountId,
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
