import { initializeApp } from "firebase/app";

const firebaseConfig = {
  apiKey: "AIzaSyCaL8GhM5hhJ3XkPwXR4pJ_iqW_Ab-KBXI",
  authDomain: "fir-nextjs-de759.firebaseapp.com",
  projectId: "fir-nextjs-de759",
  storageBucket: "fir-nextjs-de759.appspot.com",
  messagingSenderId: "1025415682300",
  appId: "1:1025415682300:web:75e23b42600336d802df79"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);

export default app;
