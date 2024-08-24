const CardGridSkeleton = () => {
    return (
      <div className="gap-4 rounded-3xl bg-gray-200 px-5 pt-7 pb-9 h-[300px] w-full grid grid-rows-[3fr-1fr] animate-pulse">
        {/* Skeleton for PersonCard Components */}
        <div className="flex gap-4 items-center">
          {[...Array(2)].map((_, index) => (
            <div
              key={index}
              className="cursor-pointer flex flex-col items-center"
            >
              {/* Skeleton for image */}
              <div className="bg-gray-300 w-[50px] h-[50px] rounded-full mb-2"></div>
              {/* Skeleton for full name */}
              <div className="bg-gray-300 w-[100px] h-[15px] rounded-md mb-1"></div>
              {/* Skeleton for job title */}
              <div className="bg-gray-300 w-[80px] h-[12px] rounded-md"></div>
            </div>
          ))}
          {/* Skeleton for button */}
          <div className="w-[50px] h-[50px] rounded-full bg-gray-300 shadow-custom-shadow text-center text-[#718EBF]"></div>
        </div>
        {/* Skeleton for InputMoney Component */}
        <div className="bg-gray-300 w-full h-[50px] rounded-md"></div>
      </div>
    );
  };
export default CardGridSkeleton;