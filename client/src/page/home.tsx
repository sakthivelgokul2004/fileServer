import { useEffect, useRef, useState } from "react";
import Navbar from "../componets/navbar";
import gsap from "gsap";
import { useGSAP } from "@gsap/react";
import { Login } from "../componets/Login";

export const Home = () => {
  const animation = useRef(gsap.timeline({ paused: true }));
  const elementRef = useRef<HTMLDivElement>(null);
  gsap.registerPlugin(useGSAP);

  const [state, setState] = useState(false);
  useEffect(() => {
    animation.current = gsap.timeline({ paused: true }).fromTo(
      elementRef.current,
      {
        opacity: 0,
        x: 20,
        display: "none",
        duration: 0.3,
        autoAlpha: 0,
      },
      {
        duration: 0.3,
        opacity: 1,
        display: "block",
        x: 0,
        ease: "power1.in",
        autoAlpha: 1,
      }
    );
  }, []);

  const loginAtten = () => {
    if (state) {
      // Animate out and set visibi
      animation.current.reverse();
      setState(false);
    } else {
      // Set visibility to true and play the animation forward
      setState(true);
      animation.current.play();
    }
  };

  return (
    <div className="z-40 relative flex flex-col h-screen">
      <Navbar login={loginAtten} />
      {/* <h1 className="text-text-hed text-3xl flex  justify-center  ">
        store excess retive files
      </h1>
      <button className="text-text-hed" onClick={() => loginAtten()}>
        sign in
      </button>
      <Upload /> */}
      <Login refElement={elementRef} />
      <div className=" flex justify-center h-full  w-full overflow-hidden">
        <div className="header w-4/5 border-x border-boder border-b ">
          <ul className="flex">
            <li className="p-1 text-text-hed text-xl">Images</li>
            <li className="p-1 text-text-hed text-xl">Files</li>
          </ul>

          <div className=" grid grid-flow-row grid-cols-4  grid-rows-5 gap-2 h-5/6 flex-grow overflow-hidden   ">
            <div className="col-span-2 row-span-3  ">
              <img
                className="object-cover object-center w-full  max-w-full rounded-lg"
                src="/d.png"
                alt=""
              />
            </div>
            <div className="row-span-3 ">
              <img
                className="object-cover  object-center w-full aspect-[2/3]  max-w-full rounded-lg"
                src="../../public/mount.jpeg"
                alt=""
              />
            </div>
            <div className=" row-span-3">
              <img
                className="object-cover aspect-[9/16] object-center w-full h-full max-w-full  rounded-lg"
                src="../../public/mount.jpeg"
                alt=""
              />
            </div>
            <div className=" row-span-2 ">
              <img
                className="object-cover  object-center w-full h-full max-w-full  rounded-lg"
                src="../../public/mount.jpeg"
                alt=""
              />
            </div>
            <div className="  col-span-2 row-span-2">
              <img
                className="object-cover  object-center w-full h-full max-w-full  rounded-lg"
                src="../../public/mount.jpeg"
                alt=""
              />
            </div>

            <div className="row-span-2  ">
              <img
                className="object-cover object-center w-full aspect-[2/3] max-w-full rounded-lg"
                src="../../public/mount.jpeg"
                alt=""
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
