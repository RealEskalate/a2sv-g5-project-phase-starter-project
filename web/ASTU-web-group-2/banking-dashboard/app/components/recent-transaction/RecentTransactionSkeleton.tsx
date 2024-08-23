const RecentTransactionSkeleton = () => {
    return (
      <div className="flex flex-col flex-initial flex-wrap gap-[10px] bg-gray-300 rounded-[25px] p-[25px] w-full animate-pulse">
        {Array.from({ length: 4 }).map((_, index) => (
          <div key={index} className="flex items-center gap-2">
            <div className="bg-gray-400 w-[60px] h-[50px] rounded-full"></div>
            <div className="flex flex-col gap-1 w-full">
              <div className="bg-gray-400 h-[19px] w-auto rounded"></div>
              <div className="bg-gray-400 h-[18px] w-auto rounded"></div>
            </div>
            <div className="bg-gray-400 h-[22px] w-auto rounded ml-auto"></div>
          </div>
        ))}
      </div>
    );
  };

export default RecentTransactionSkeleton;
  