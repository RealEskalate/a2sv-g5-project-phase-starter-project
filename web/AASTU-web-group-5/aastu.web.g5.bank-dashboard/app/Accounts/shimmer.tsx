const SkeletonCard = () => {
    return (
      <div className="p-4 bg-white rounded-lg animate-pulse">
        <div className="flex items-center justify-center space-x-4">
          <div className="w-12 h-12 bg-gray-300 rounded-full"></div>
          <div className="flex-1">
            <div className="h-4 bg-gray-300 rounded w-1/2"></div>
            <div className="h-4 bg-gray-300 rounded w-1/4 mt-2"></div>
          </div>
        </div>
      </div>
    );
  };
  
  export default SkeletonCard;
  