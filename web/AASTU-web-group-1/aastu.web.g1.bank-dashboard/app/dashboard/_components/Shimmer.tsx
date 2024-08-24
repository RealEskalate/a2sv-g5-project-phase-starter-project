// Shimmer component for skeleton loading effect
export const RecentTransactionShimmer = () => {
  return (
    <div className="animate-pulse flex space-x-4">
      <div className="rounded-full bg-gray-300 h-12 w-12"></div>
      <div className="flex-1 space-y-4 ">
        <div className="h-4 bg-gray-300 rounded w-3/4"></div>
        <div className="h-4 bg-gray-300 rounded"></div>
      </div>
    </div>
  );
};




export const CreditCardShimmer = () => {
  return (
    <div className="animate-pulse flex flex-col items-center   mb-2  shadow-lg rounded-3xl">
      {/* Simulate card with rounded corners, gradient, and shadow */}
      <div className="bg-white h-56 w-80 rounded-2xl p-4 flex flex-col justify-between ">
        {/* Simulate chip */}
        <div className="bg-gray-300 h-6 w-10 rounded-sm mb-2"></div>

        {/* Simulate card number */}
        <div className="space-y-2">
          <div className="h-4 bg-gray-300 rounded w-full"></div>
          <div className="h-4 bg-gray-300 rounded w-full"></div>
          <div className="h-4 bg-gray-300 rounded w-full"></div>
        </div>

        {/* Simulate cardholder name and expiration date */}
        <div className="flex justify-between mt-4">
          <div className="h-4 bg-gray-300 rounded w-1/2"></div>
          <div className="h-4 bg-gray-300 rounded w-1/4"></div>
        </div>
      </div>
    </div>
  );
};



export const ShimmerRow = () => (
  <div className="animate-pulse flex items-center space-x-4 py-2">
    <div className="bg-gray-300 h-6 w-1/4 rounded-md"></div> 
    <div className="bg-gray-300 h-6 w-1/4 rounded-md"></div> 
    <div className="bg-gray-300 h-6 w-1/4 rounded-md"></div> 
    <div className="bg-gray-300 h-6 w-1/4 rounded-md"></div> 
  </div>
);



