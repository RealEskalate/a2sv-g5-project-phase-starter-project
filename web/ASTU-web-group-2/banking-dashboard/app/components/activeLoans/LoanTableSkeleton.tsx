const LoanTableSkeleton = () => {
    return (
      <div className="bg-gray-300 rounded-3xl w-full p-3 animate-pulse">
        <div className="flex justify-center">
          <div className="w-full">
            {/* Skeleton for table headers */}
            <div className="grid grid-cols-7 gap-4 mb-2">
              <div className="hidden sm:block bg-gray-400 h-[16px] rounded w-[60px]"></div>
              <div className="bg-gray-400 h-[16px] rounded w-[100px]"></div>
              <div className="bg-gray-400 h-[16px] rounded w-[100px]"></div>
              <div className="hidden sm:block bg-gray-400 h-[16px] rounded w-[80px]"></div>
              <div className="hidden sm:block bg-gray-400 h-[16px] rounded w-[100px]"></div>
              <div className="hidden sm:block bg-gray-400 h-[16px] rounded w-[100px]"></div>
              <div className="bg-gray-400 h-[16px] rounded w-[70px]"></div>
            </div>
  
            {/* Skeleton for table rows */}
            {Array.from({ length: 5 }).map((_, index) => (
              <div className="grid grid-cols-7 gap-4 mb-2" key={index}>
                <div className="hidden sm:block bg-gray-400 h-[16px] rounded w-[60px]"></div>
                <div className="bg-gray-400 h-[16px] rounded w-[100px]"></div>
                <div className="bg-gray-400 h-[16px] rounded w-[100px]"></div>
                <div className="hidden sm:block bg-gray-400 h-[16px] rounded w-[80px]"></div>
                <div className="hidden sm:block bg-gray-400 h-[16px] rounded w-[100px]"></div>
                <div className="hidden sm:block bg-gray-400 h-[16px] rounded w-[100px]"></div>
                <div className="bg-gray-400 h-[30px] rounded w-[70px]"></div>
              </div>
            ))}
          </div>
        </div>
      </div>
    );
  };

export default LoanTableSkeleton;