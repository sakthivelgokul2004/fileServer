import React from "react";
import { IconContext } from "react-icons";
import { PiUploadSimpleLight } from "react-icons/pi";

interface login {
  login: React.MouseEventHandler<HTMLButtonElement> | undefined;
}

export default function Navbar(login: login) {
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
            <button className="ml-auto text-text-hed" onClick={login.login}>
              {" "}
              Sign In
            </button>
          </div>
        </div>
      </div>
    </nav>
  );
}
