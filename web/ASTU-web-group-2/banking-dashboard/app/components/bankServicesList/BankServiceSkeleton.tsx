import BalanceCardSkeleton from "../infobox/BalanceCardSkeleton";

const BankServiceSkeleton = () => {
    const SkeletonCard = () => (
      <div>
        <BalanceCardSkeleton />
      </div>
    );
  
    return (
      <div>
     
        <div className="flex flex-col gap-5 max-md:hidden">
          {[...Array(4)].map((_, index) => (
            <SkeletonCard key={index} />
          ))}
        </div>
  
      
        <div className="flex flex-col gap-5 md:hidden">
          {[...Array(4)].map((_, index) => (
            <SkeletonCard key={index} />
          ))}
        </div>
      </div>
    );
  };
  
  export default BankServiceSkeleton;
  