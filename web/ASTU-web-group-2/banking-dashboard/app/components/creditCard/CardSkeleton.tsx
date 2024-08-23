const CardSkeleton = () => {
    return (
      <div className="flex flex-col animate-pulse">
        <div className="w-[350px] h-[165px] bg-gray-300 rounded-t-[25px]">
          <div className="pl-[26px] pt-[24px] flex gap-[202px]">
            <div>
              <div className="bg-gray-400 h-[14.4px] w-[60px] mb-[10px] rounded"></div>
              <div className="bg-gray-400 h-[24px] w-[100px] rounded"></div>
            </div>
            <div className="bg-gray-400 h-[34.77px] w-[34.77px] rounded-full"></div>
          </div>
          <div className="flex justify-between pt-[33px] pl-[24px] pr-[24px]">
            <div>
              <div className="bg-gray-400 h-[14.4px] w-[80px] mb-[10px] rounded"></div>
              <div className="bg-gray-400 h-[18px] w-[120px] rounded"></div>
            </div>
            <div>
              <div className="bg-gray-400 h-[14.4px] w-[60px] mb-[10px] rounded"></div>
              <div className="bg-gray-400 h-[18px] w-[80px] rounded"></div>
            </div>
          </div>
        </div>
        <div className="w-[350px] h-[70px] bg-gray-300 rounded-b-[25px] flex gap-[32px] items-center justify-between px-[26px]">
          <div className="bg-gray-400 h-[26.4px] w-[200px] rounded"></div>
          <div className="bg-gray-400 h-[44px] w-[44px] rounded-full"></div>
        </div>
      </div>
    );
  };

export default CardSkeleton