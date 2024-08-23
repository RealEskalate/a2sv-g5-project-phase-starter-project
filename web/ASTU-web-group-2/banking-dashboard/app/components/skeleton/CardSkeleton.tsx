const CardSkeleton = () => {
  return (
    <div className="flex flex-col justify-around items-center bg-gray-200 h-[175px] w-[360px] animate-pulse rounded-2xl overflow-clip p-4">
      <div className="h-[30%] flex justify-between  w-[100%]">
        <div className=" h-[100%] w-[60%] bg-grey-100 rounded-2xl"></div>
        <div className=" h-[100%] w-[30%] bg-grey-100 rounded-2xl"></div>
      </div>
      <div className=" h-[30%] w-[100%] bg-grey-100 rounded-2xl"></div>
      <div className=" h-[30%] w-[100%] bg-grey-100 rounded-2xl"></div>
    </div>
  );
};

export default CardSkeleton;
