import React from 'react'

const recentlistitems = [
    {   
        title: {
            des : "Apple Store",
            sub : "E-commerce, Marketplace"

        },
        percent: {
            isDeposited : true,
            des : "+16%",
            sub : "Return Value"

        },
        amount: {
            des : "$54,000",
            sub : "Envestment Value"

        },
        icon: "/assets/myinvestment/card1.svg"
    },
    {   
        title: {
            des : "Samsung mobile",
            sub : "E-commerce, Marketplace"

        },
        percent: {
            isDeposited : false,
            des : "-4%",
            sub : "Return Value"

        },
        amount: {
            des : "$54,000",
            sub : "Envestment Value"

        },
        icon: "/assets/myinvestment/card2.svg"
    },
    {   
        title: {
            des : "Tesla Motor",
            sub : "E-commerce, Marketplace"

        },
        percent: {
            isDeposited : true,
            des : "+16%",
            sub : "Return Value"

        },
        amount: {
            isDeposited : true,
            des : "$54,000",
            sub : "samsung motor"

        },
        icon: "/assets/myinvestment/card3.svg"
    }
   
  ];
  const MyInvestment = () => {
    return (
      <div className=" sm:w-[475px] md:w-[635px]  ">
        {recentlistitems.map((item, index) => (
          <div key={index} className="grid grid-flow-col h-[69px] lg:h-[90px] justify-between mb-[10px] sm:mb-[15px] items-center pl-[20px] bg-white rounded-3xl grid-col-12">
      
              <div className="col-span-2">
                <img
                  src={item.icon}
                  className="lg:w-[60px] w-[45px]"
                />
              </div>
        
              <div className="col-span-4 ">
                <p className="text-[14px] sm:text-[16px] text-[#333B69] font-medium">
                  {item.title.des}
                </p>
                <span className="text-[12px] sm:text-[15px] text-[#718EBF]">
                  {item.title.sub}
                </span>
              </div>
    
              <div className=" hidden col-span-3 sm:block ">
                <p className="text-[14px] sm:text-[16px] text-[#333B69] font-medium">
                  {item.amount.des}
                </p>
                <span className="text-[12px] sm:text-[15px] text-[#718EBF]">
                  {item.amount.sub}
                </span>
              </div>
      
              <div className=" col-span-3 ">
                <p className='text-[14px] sm:text-[16px]  font-medium ' style={{ color: item.percent.isDeposited ? '#41D4A8' : '#FF4B4A' }} >
                  {item.percent.des}
                </p>
                <span className="hidden sm:block text-[12px] sm:text-[15px] text-[#718EBF]">
                  {item.percent.sub}
                </span>
              </div>
        
          </div>
        ))}
      </div>
    );
  };
  
  export default MyInvestment;