const LoanTableSkeleton = () => {
    return (
      <div className="bg-gray-300 rounded-3xl w-full p-3 animate-pulse">
      <div className="flex justify-center">
        <div className="w-full">
        
          <div className="grid md:grid-cols-7 grid-cols-3 gap-4 mb-4">
            <div className="hidden sm:block text-gray-500 text-[12px] md:text-[16px] font-medium">
              SL No
            </div>
            <div className="text-gray-500 text-[12px] md:text-[16px] font-medium">
              Loan Money
            </div>
            <div className="text-gray-500 text-[12px] md:text-[16px] font-medium">
              Left to repay
            </div>
            <div className="hidden sm:block text-gray-500 text-[12px] md:text-[16px] font-medium">
              Duration
            </div>
            <div className="hidden sm:block text-gray-500 text-[12px] md:text-[16px] font-medium">
              Interest rate
            </div>
            <div className="hidden sm:block text-gray-500 text-[12px] md:text-[16px] font-medium">
              Installment
            </div>
            <div className="text-gray-500 text-[12px] md:text-[16px] font-medium">
              Repay
            </div>
          </div>
    
          {/* Skeleton for table rows */}
          {Array.from({ length: 5 }).map((_, index) => (
            <div className="grid md:grid-cols-7 grid-cols-3 gap-4 mb-2" key={index}>
              <div className="hidden sm:block bg-gray-400 h-[16px] rounded w-[60px]"></div>
              <div className="bg-gray-400 h-[16px] rounded w-[100px]"></div>
              <div className="bg-gray-400 h-[16px] rounded w-[100px]"></div>
              <div className="hidden sm:block bg-gray-400 h-[16px] rounded w-[80px]"></div>
              <div className="hidden sm:block bg-gray-400 h-[16px] rounded w-[100px]"></div>
              <div className="hidden sm:block bg-gray-400 h-[16px] rounded w-[100px]"></div>
              <div className="bg-gray-400 h-[20px] rounded-lg w-[70px]"></div>
            </div>
          ))}
        </div>
      </div>
    </div>
    
    );
  };

export default LoanTableSkeleton;