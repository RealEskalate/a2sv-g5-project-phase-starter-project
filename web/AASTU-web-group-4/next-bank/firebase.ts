// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
import { getStorage, ref, uploadBytes, getDownloadURL } from 'firebase/storage';
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyCRZCppJFqbBCicovpjulqHwa3z-IEo5k4",
  authDomain: "next-bank-3c29e.firebaseapp.com",
  projectId: "next-bank-3c29e",
  storageBucket: "next-bank-3c29e.appspot.com",
  messagingSenderId: "523285153466",
  appId: "1:523285153466:web:f0b3a893e7fee3f036a483",
  measurementId: "G-F9GY603MSN"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);
const storage = getStorage(app);

export {storage}