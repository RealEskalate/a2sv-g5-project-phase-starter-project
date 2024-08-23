// firebase.ts
import { initializeApp } from 'firebase/app';
import { getFirestore } from 'firebase/firestore';
import { getAuth } from 'firebase/auth';

export const firebaseConfig = {
  apiKey: "AIzaSyDdTXFHR-2XWqXc7NKNhckPhxSB9Nf5LvE",
  authDomain: "next-bank-a27e3.firebaseapp.com",
  projectId: "next-bank-a27e3",
  storageBucket: "next-bank-a27e3.appspot.com",
  messagingSenderId: "984662193003",
  appId: "1:984662193003:web:ae9f16b9d21b598c05122e",
  measurementId: "G-0QH0RTQLQZ"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);

// Initialize Firestore
export const db = getFirestore(app);

// Initialize Firebase Auth (if needed)
export const auth = getAuth(app);
