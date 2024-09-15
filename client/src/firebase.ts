// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getStorage } from "firebase/storage";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyDrIfbObHH5oDrRJ_zO4L9aROsrjVjZJM8",
  authDomain: "fileserver-8c567.firebaseapp.com",
  projectId: "fileserver-8c567",
  storageBucket: "fileserver-8c567.appspot.com",
  messagingSenderId: "962517814944",
  appId: "1:962517814944:web:95156e8c763270634e03c5",
  measurementId: "G-935BPXQDBY",
};

// Initialize Firebase
export const app = initializeApp(firebaseConfig);
export const storage = getStorage(app);
