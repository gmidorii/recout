import firebase from "firebase/app";
import "firebase/auth";
import { devConfig, prodConfig } from "~/firebase.config";

if (!firebase.apps.length) {
  if (process.env.NODE_ENV === "production") {
    firebase.initializeApp(prodConfig);
  } else {
    firebase.initializeApp(devConfig);
  }
}

export const providers = [firebase.auth.GithubAuthProvider.PROVIDER_ID];

export const auth = firebase.auth();
