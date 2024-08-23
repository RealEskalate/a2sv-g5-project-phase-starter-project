const BarSkeleton = () => {
    return (
    
          <div className="w-[80%] h-[60%] bg-gray-300 flex justify-between items-end">
            {[...Array(7)].map((_, index) => (
              <div
                key={index}
                className="w-[5%] bg-gray-100 rounded-t-lg"
                style={{ height: `${20 + index * 15}%` }} // Varying heights to mimic bars
              ></div>
            ))}
          </div>
   
      
    );
  };
  
  export default BarSkeleton;
  