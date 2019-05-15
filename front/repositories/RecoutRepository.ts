import axios from "axios";

const baseURL = `${process.env.recoutUrl}`;
const resource = "recout";

const Repository = axios.create({
  baseURL
});

export interface recout {
  get(limit?: number);
  post(message: string);
  getContinues();
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

  getContinues() {
    return Repository.get(`${resource}/continues`, {
      params: {
        // TODO: get account_id
        account_id: "gmidorii"
      }
    });
  }
};

export default repositoryImpl;
