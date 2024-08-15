import React from "react";
import Link from "next/link";
const recentlistitems = [
  

     {icon: "/assets/invoicesSent/icon1.svg" ,
    col1: {
      des: "Card Type",
      sub: "Secondary"
    },
    col2: {
      des: "Bank ",
      sub: "DBL-Bank"
    },
    col3: {
      des: "Card Numver",
      sub: "**** **** 5600"
    },
    col4: {
      des: "Namain Card",
      sub: "William "
    },
    col5: {
      des: "View Detail",
    }},
    {icon: "/assets/invoicesSent/icon1.svg" ,
      col1: {
        des: "Card Type",
        sub: "Secondary"
      },
      col2: {
        des: "Bank ",
        sub: "DBL-Bank"
      },
      col3: {
        des: "Card Numver",
        sub: "**** **** 5600"
      },
      col4: {
        des: "Namain Card",
        sub: "William "
      },
      col5: {
        des: "View Detail",
      }},
      {icon: "/assets/invoicesSent/icon1.svg" ,
        col1: {
          des: "Card Type",
          sub: "Secondary"
        },
        col2: {
          des: "Bank ",
          sub: "DBL-Bank"
        },
        col3: {
          des: "Card Numver",
          sub: "**** **** 5600"
        },
        col4: {
          des: "Namain Card",
          sub: "William "
        },
       },
   
   
  

];

const CardList = () => {
  return (
    <div className=" sm:w-[475px] md:w-[730px]  ">
      {recentlistitems.map((item, index) => (
        <div key={index} className="grid grid-flow-col h-[69px] lg:h-[90px] justify-between mb-[10px] sm:mb-[15px] items-center pl-[20px] bg-white rounded-3xl grid-col-12">
          
            <div className="col-span-1">
              <img
                src={item.icon}
                className="lg:w-[60px] w-[45px]"
              />
            </div>
            
            <div className="col-span-2 ">
              <p className="text-[14px] md:[text-[12px]] lg:text-[16px] text-[#333B69] ">
                {item.col1.des} 
              </p>
              <span className="text-[12px] sm:text-[15px] text-[#718EBF]">
                {item.col1.sub}
              </span>
            </div>
      
            <div className=" col-span-[2.5] ">
              <p className="text-[14px] md:[text-[12px]] lg:text-[16px] text-[#333B69]">
                {item.col2.des}
              </p>
              <span className="text-[12px]  md:[text-[12px]] lg:text-[16px] text-[#718EBF]">
                {item.col2.sub}
              </span>
            </div>



            <div className=" hidden col-span-[2.5] sm:block ">
              <p className="text-[14px] md:text-[12px]  lg:text-[16px] text-[#333B69] font-medium">
                {item.col3.des}
              </p>
              <span className="text-[12px] sm:text-[15px] text-[#718EBF]">
                {item.col3.sub}
              </span>
            </div>


            <div className=" hidden col-span-2 sm:block ">
              <p className="text-[14px] sm:text-[16px] text-[#333B69] font-medium">
                {item.col4.des}
              </p>
              <span className="text-[12px] sm:text-[15px] text-[#718EBF]">
                {item.col4.sub}
              </span>
            </div>
    
            <div className=" col-span-2 ">
              <p className='text-[14px] sm:text-[16px] text-[#1814F3] font-medium ' >
                <Link href = "#">
                View Detail
                </Link>
              </p>
            </div>
      
        </div>
      ))}
    </div>
  );
};

export default CardList;
