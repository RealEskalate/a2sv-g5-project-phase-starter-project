export const CardSkeleton = () => (
    <div className="w-[231px] h-[170px] sm:w-[265px] sm:h-[170px] md:w-[350px] md:h-[235px] rounded-xl relative bg-gray-200 animate-pulse">
      <div className="flex justify-between w-[95%]">
        <div className="mt-1 ml-3 p-2">
          <div className="h-4 bg-gray-300 rounded w-16"></div>
          <div className="h-6 bg-gray-300 rounded w-24 mt-2"></div>
        </div>
        <div className="h-[29px] w-[30px] bg-gray-300 rounded mt-4 mr-2"></div>
      </div>
      <div className="flex justify-between w-[90%] mt-4">
        <div className="ml-3 pl-1.5 md:p-2">
          <div className="h-4 bg-gray-300 rounded w-20"></div>
          <div className="h-6 bg-gray-300 rounded lg:w-32 w-16 mt-2"></div>
        </div>
        <div className="mr-3 md:mr-9 md:p-2">
          <div className="h-4 bg-gray-300 rounded w-16"></div>
          <div className="h-6 bg-gray-300 rounded lg:w-32 w-16 mt-2"></div>
        </div>
      </div>
      <div className="flex justify-between items-center absolute bottom-0 left-0 right-0 bg-gray-200 p-3">
        <div className="h-6 bg-gray-300 rounded w-32"></div>
        <div className="h-6 bg-gray-300 rounded w-10"></div>
      </div>
    </div>
);
  