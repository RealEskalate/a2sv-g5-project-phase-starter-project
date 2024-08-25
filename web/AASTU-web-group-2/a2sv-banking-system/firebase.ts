// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getStorage } from "firebase/storage";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
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
const storage = getStorage(app);
export {storage}