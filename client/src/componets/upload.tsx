import {
  getDownloadURL,
  getStorage,
  ref,
  uploadBytesResumable,
} from "firebase/storage";
import { app } from "../firebase";
import { useRef } from "react";

export default function Upload() {
  const imageRef = useRef<HTMLInputElement>(null);
  const uploadImage = () => {
    if (imageRef.current?.files) {
      console.log("started");
      const img = imageRef.current?.files[0];
      const storage = getStorage(app);
      const storageRef = ref(storage, img.name);
      const uploadTask = uploadBytesResumable(storageRef, img);
      let request = {
        urlString: "",
        filetype: img.type,
      };
      console.log(img.name);
      uploadTask.on(
        "state_changed",
        (snapshot) => {
          const progress =
            (snapshot.bytesTransferred / snapshot.totalBytes) * 100;
          // setImageUploadProgress(progress.toFixed(0));
          console.log(progress);
        },
        (error) => {},
        async () => {
          await getDownloadURL(uploadTask.snapshot.ref).then((downloadUrl) => {
            request.urlString = downloadUrl;
            console.log(downloadUrl);
          });
          await fetch("/api/user/addfile", {
            method: "POST",
            credentials: "same-origin",
            body: JSON.stringify(request),
          });
        }
      );
    } else {
      console.log("file not select");
    }
  };
  return (
    <div className="">
      <input type="file" name="file" id="upload" ref={imageRef} />
      <button
        className="text-text-hed text-3xl"
        id="upload"
        onClick={uploadImage}
      >
        {" "}
        upload
      </button>
    </div>
  );
}
