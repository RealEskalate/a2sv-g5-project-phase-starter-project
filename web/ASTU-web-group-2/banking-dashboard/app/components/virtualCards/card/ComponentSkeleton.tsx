import CardSkeleton from "../creditCard/CardSkeleton";


const ComponentSkeleton = () => {
    return (
      <div className="bg-gray-200 rounded-3xl p-4 h-[260px] w-full animate-pulse">
        <div className="flex flex-wrap gap-4 overflow-x-auto scrollbar-hide">
         {/* Skeleton for Children Elements */}
      <div className="flex flex-wrap gap-4 overflow-x-auto scrollbar-hide">
       
        <CardSkeleton />
      
      </div>
        </div>
      </div>
    );
  };
export default ComponentSkeleton;