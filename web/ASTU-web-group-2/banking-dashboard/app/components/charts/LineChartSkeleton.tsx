const LineChartSkeleton = () => {
    return (
      <div className="text-gray-500 border rounded-[22px] bg-gray-200 p-5 w-full h-auto animate-pulse">
        {/* Skeleton for the Line Chart */}
        <div className="w-full h-[95%] bg-gray-300 rounded-lg flex items-center justify-center">
          {/* Simulated line chart */}
          <svg width="100%" height="100%" viewBox="0 0 200 100">
            <polyline
              fill="none"
              stroke="#aaaaaa"
              strokeWidth="2"
              points="10,80 40,60 70,65 100,45 130,50 160,30 190,40"
            />
            <circle cx="10" cy="80" r="3" fill="#aaaaaa" />
            <circle cx="40" cy="60" r="3" fill="#aaaaaa" />
            <circle cx="70" cy="65" r="3" fill="#aaaaaa" />
            <circle cx="100" cy="45" r="3" fill="#aaaaaa" />
            <circle cx="130" cy="50" r="3" fill="#aaaaaa" />
            <circle cx="160" cy="30" r="3" fill="#aaaaaa" />
            <circle cx="190" cy="40" r="3" fill="#aaaaaa" />
          </svg>
        </div>
      </div>
    );
  };
  
  export default LineChartSkeleton;
  