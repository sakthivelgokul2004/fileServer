import { PiDownload } from "react-icons/pi";
import { FaDeleteLeft } from "react-icons/fa6";
import { IconContext } from "react-icons";

export default function FilesList() {
  return (
    <div className="">
      <div className="w-full flex ">
        <a className=" flex-grow ">
          <p>file name</p>
        </a>
        <button>
          <IconContext.Provider value={{ className: "text-text-hed" }}>
            <PiDownload />
          </IconContext.Provider>
        </button>
        <button>
          <IconContext.Provider value={{ className: "text-text-hed" }}>
            <FaDeleteLeft />
          </IconContext.Provider>
        </button>
      </div>
    </div>
  );
}
