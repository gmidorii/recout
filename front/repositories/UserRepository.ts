import axios from "axios";

const baseURL = `${process.env.recoutUrl}`;
const resource = "user";

const Repository = axios.create({
  baseURL
});

export interface user {
  get(accoutId?: string);
}

const repositoryImpl: user = {
  get(accoutId: string = "gmidorii") {
    return Repository.get(`${resource}`, {
      params: {
        account_id: accoutId
      }
    });
  }
};

export default repositoryImpl;
