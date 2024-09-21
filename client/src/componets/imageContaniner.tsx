const ImageContaniner = () => {
  return (
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
  );
};

export default ImageContaniner;
