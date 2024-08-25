// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getStorage } from "firebase/storage";
const firebaseConfig = {
  apiKey: "AIzaSyCv_JSIXqdk05YebMCWuDm79VUlWhvuOgA",
  authDomain: "a2sv-wallet.firebaseapp.com",
  projectId: "a2sv-wallet",
  storageBucket: "a2sv-wallet.appspot.com",
  messagingSenderId: "136604332771",
  appId: "1:136604332771:web:8264100cc84c484fbc4bd6",
  measurementId: "G-44PK3KTJJT"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
export const storage = getStorage(app)