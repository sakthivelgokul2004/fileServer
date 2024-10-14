import { RiDownloadLine } from "react-icons/ri";
import { MdDelete } from "react-icons/md";
import { IconContext } from "react-icons";
import { useMutation } from "@tanstack/react-query";

interface file {
  fileName: string;
  fileUrl: string;
  download: Function;
  id: string;
}
const delet = async (id: string) => {
  const requset = {
    fileid: id,
  };
  let response = await fetch("/api/user/delete", {
    method: "DELETE",
    body: JSON.stringify(requset),
  });
  return response.json;
};
export default function FilesList(file: file) {
  const mutation = useMutation({
    mutationKey: ["delete"],
    mutationFn: (id: string) => delet(id),
  });
  console.log(file);
  return (
    <div className="">
      <div className="w-full flex items-center p-2 border-2 border-boder ">
        <a className=" flex-grow  ">
          <p>{file.fileName} </p>
        </a>
        <button
          className="p-4 flex items-baseline"
          onClick={() => file.download(file.fileUrl, file.fileName)}
        >
          <IconContext.Provider value={{ className: "text-text-hed" }}>
            <RiDownloadLine />
          </IconContext.Provider>
        </button>
        <button
          className="p-4 flex items-baseline"
          onClick={() => mutation.mutate(file.id)}
        >
          <IconContext.Provider value={{ className: "text-text-hed" }}>
            <MdDelete />
          </IconContext.Provider>
        </button>
      </div>
    </div>
  );
}
