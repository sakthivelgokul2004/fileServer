import React from "react";
import { IconContext } from "react-icons";
import { PiUploadSimpleLight } from "react-icons/pi";
import Upload from "./upload";
import DynamicButton from "./dynamicButton";

interface loginProps {
  logout: React.MouseEventHandler<HTMLButtonElement>;
  login: React.MouseEventHandler<HTMLButtonElement>;
}
const Navbar: React.FC<loginProps> = ({ login, logout }) => {
  return (
    <nav>
      <div className="top-0 backdrop-blur -z-0 w-full border-b border-boder ">
        <div className="max-w-[90rem]  mx-auto ">
          <div className=" relative w-full  h-14  flex items-center ">
            <div className="flex">
              <IconContext.Provider value={{ className: "text-text-hed" }}>
                <PiUploadSimpleLight />
              </IconContext.Provider>
              <h5 className="text-text-hed">Filestore</h5>
            </div>
            <div className="">
              <Upload />
            </div>
            <DynamicButton logout={logout} login={login} />
          </div>
        </div>
      </div>
    </nav>
  );
};
export default Navbar;
