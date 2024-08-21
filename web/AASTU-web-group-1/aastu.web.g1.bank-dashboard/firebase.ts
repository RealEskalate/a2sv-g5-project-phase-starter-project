// firebaseConfig.ts
import { initializeApp } from 'firebase/app';
import { getStorage } from 'firebase/storage';

const firebaseConfig = {
    apiKey: process.env.FIREBASEKEY,
    authDomain: "bankdash-fed2e.firebaseapp.com",
    projectId: "bankdash-fed2e",
    storageBucket: "bankdash-fed2e.appspot.com",
    messagingSenderId: "129353860588",
    appId: "1:129353860588:web:a3783681ecc17bd384acc9",
    measurementId: "G-20R0E0PR4T"
  };

const app = initializeApp(firebaseConfig);
const storage = getStorage(app);

export { storage };
