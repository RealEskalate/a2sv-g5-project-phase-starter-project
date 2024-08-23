const BalanceCardSkeleton = () => {
    return (
      <div className="rounded-[20px] bg-gray-300 h-[90px] shadow-md animate-pulse">
        <div className="flex items-center p-2 gap-3 w-full">
          <div className="flex-shrink-0">
            <div className="bg-gray-400 w-[50px] h-[50px] max-lg:w-[30px] max-lg:h-[30px] max-md:w-[50px] max-md:h-[50px] rounded-full"></div>
          </div>
          <div className="flex flex-col justify-center w-full">
            <div className="bg-gray-400 h-[18px] w-[150px] max-lg:w-[100px] max-md:w-[150px] rounded mb-2"></div>
            <div className="bg-gray-400 h-[22px] w-[100px] max-lg:w-[80px] max-md:w-[100px] rounded"></div>
          </div>
        </div>
      </div>
    );
  };
export default BalanceCardSkeleton  