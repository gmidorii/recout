import axios from "axios";
import { APIUser, APINotFound } from "~/types/api";
import { User } from "~/types";
import { APIStatusError } from "./error";

const baseURL = `${process.env.recoutUrl}`;
const resource = "user";
const statusOK = 200;
const statusNotFound = 404;

const Repository = axios.create({
  baseURL
});

export interface user {
  get(accoutId?: string): Promise<APIUser>;
  post(user: User);
}

const repositoryImpl: user = {
  async get(accoutId: string = "gmidorii") {
    try {
      const response = await Repository.get(`${resource}`, {
        params: {
          account_id: accoutId
        }
      });
      if (response.status === statusNotFound) {
        return null;
      }
      return {
        account_id: response.data.account_id,
        pixela_graph: response.data.pixela_graph
      };
    } catch (e) {
      console.log(e);
    }
  },
  async post(user: User) {
    const response = await Repository.post(`${resource}`, {
      account_id: user.id
    });
    if (response.status != statusOK) {
      throw new APIStatusError();
    }
  }
};

export default repositoryImpl;
