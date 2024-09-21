import { RiDownloadLine } from "react-icons/ri";
import { MdDelete } from "react-icons/md";
import { IconContext } from "react-icons";

interface file {
  fileName: string;
  fileUrl: string;
  download: Function;
}
export default function FilesList(file: file) {
  return (
    <div className="">
      <div className="w-full flex items-center p-2 border-2 border-boder ">
        <a className=" flex-grow  ">
          <p>file name</p>
        </a>
        <button
          className="p-4 flex items-baseline"
          onClick={() => file.download(file.fileUrl)}
        >
          <IconContext.Provider value={{ className: "text-text-hed" }}>
            <RiDownloadLine />
          </IconContext.Provider>
        </button>
        <button className="p-4 flex items-baseline">
          <IconContext.Provider value={{ className: "text-text-hed" }}>
            <MdDelete />
          </IconContext.Provider>
        </button>
      </div>
    </div>
  );
}
