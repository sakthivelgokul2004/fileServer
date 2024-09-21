import { useState } from "react";
import FileContainer from "./fileContainer";
import ImageContaniner from "./imageContaniner";

enum sections {
  file = "files",
  image = "images",
}
const DisplayContainer = () => {
  const [section, setSection] = useState<sections>(sections.file);
  return (
    <>
      <div className=" flex justify-center h-full  w-full overflow-hidden">
        <div className="header w-4/5 border-x border-boder border-b ">
          <ul className="flex">
            <li
              className="p-1 text-text-hed text-xl"
              onClick={() => setSection(sections.file)}
            >
              Files
            </li>
            <li
              className="p-1 text-text-hed text-xl"
              onClick={() => setSection(sections.image)}
            >
              Images
            </li>
          </ul>
          {section == "files" ? <FileContainer /> : <ImageContaniner />}
        </div>
      </div>
    </>
  );
};

export default DisplayContainer;
