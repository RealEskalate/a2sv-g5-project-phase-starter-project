import BarSkeleton from "./BarSkeleton";


const BarChartSkeleton = () => {
    return (
      <div className="bg-gray-300 rounded-[22px] lg:h-[322px] h-[261px] animate-pulse">
        <div className="flex flex-row justify-end gap-2">
          <div className="flex flex-row mx-5 mt-5 gap-1">
            <div className="w-[12px] h-[12px] mt-[6px] border rounded-full bg-gray-400"></div>
            <div className="bg-gray-400 h-[16px] w-[70px] rounded"></div>
          </div>
          <div className="flex flex-row mx-5 mt-5 gap-1">
            <div className="w-[12px] h-[12px] mt-[6px] border rounded-full bg-gray-400"></div>
            <div className="bg-gray-400 h-[16px] w-[80px] rounded"></div>
          </div>
        </div>
        <div className="h-[75%] mx-5 mb-5 flex items-center justify-center">
        <BarSkeleton />
        </div>
      </div>
    );
  };
export default BarChartSkeleton  