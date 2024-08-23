export const CardListSkeleton = () => (
    <div className="flex flex-row justify-between w-full bg-gray-200 p-4 mb-4 rounded-lg shadow-md animate-pulse">
      {/* Icon */}
      <div className="h-12 w-12 rounded-xl bg-gray-300"></div>
  
      {/* Card Type */}
      <div className="flex flex-col w-1/5">
        <div className="h-4 bg-gray-300 rounded w-3/4 mb-2"></div>
        <div className="h-3 bg-gray-300 rounded w-2/3"></div>
      </div>
  
      {/* Bank */}
      <div className="flex flex-col w-1/5">
        <div className="h-4 bg-gray-300 rounded w-3/4 mb-2"></div>
        <div className="h-3 bg-gray-300 rounded w-2/3"></div>
      </div>
  
      {/* Card Number - Hide on small screens */}
      <div className="flex flex-col w-1/5 hidden sm:block">
        <div className="h-4 bg-gray-300 rounded w-3/4 mb-2"></div>
        <div className="h-3 bg-gray-300 rounded w-2/3"></div>
      </div>
  
      {/* Card Name - Hide on small screens */}
      <div className="flex flex-col w-1/5 hidden sm:block">
        <div className="h-4 bg-gray-300 rounded w-3/4 mb-2"></div>
        <div className="h-3 bg-gray-300 rounded w-2/3"></div>
      </div>
  
      {/* View Details Link */}
      <div className="h-4 bg-gray-300 rounded w-20"></div>
    </div>
  );
  