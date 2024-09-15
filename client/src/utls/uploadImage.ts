import {
  getDownloadURL,
  getStorage,
  ref,
  uploadBytesResumable,
} from "firebase/storage";
import { app } from "../firebase";
const handleImage = () => {
  console.log("Image");
  // const fileName = new Date().getTime() + "-" + img.name;
  const storage = getStorage(app);
  const storageRef = ref(storage, fileName);
  const uploadTask = uploadBytesResumable(storageRef, img);
  uploadTask.on(
    "state_changed",
    (snapshot) => {
      const progress = (snapshot.bytesTransferred / snapshot.totalBytes) * 100;
      // setImageUploadProgress(progress.toFixed(0));
    },
    (error) => {},
    () => {
      getDownloadURL(uploadTask.snapshot.ref).then((downloadUrl) => {});
    }
  );
};

// function setImageUploadProgress(arg0: string) {
//     throw new Error("Function not implemented.");
// }
