'use client'
import React from "react";
import Link from "next/link";
import { useSession } from "next-auth/react";
const  icons = [
  "/assets/cardlist/card1.svg",
 "/assets/cardlist/card2.svg",
 "/assets/cardlist/card3.svg",

]

interface Props {
  cardType: string,
  cardHolder : string;
  cardNumber: string,



}
const recentlistitems = [
  

     {icon: "/assets/cardlist/card1.svg" ,
    col1: {
    
      sub: "Secondary"
    },
    col2: {
      sub: "DBL-Bank"
    },
    col3: {
  
      sub: "**** **** 5600"
    },
    col4: {
  
      sub: "William "
    },
    col5: {
      des: "View Detail",
    }},
    {icon: "/assets/cardlist/card2.svg" ,
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
      {icon: "/assets/cardlist/card3.svg" ,
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

const CardList =  async  () => {
  const session = useSession();
  const accessToken = session.data?.user.accessToken || "";

  try {
    const response =  await fetch("https://bank-dashboard-6acc.onrender.com/cards", {
      method: "GET",
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    });

    if (!response.ok) {
      return <div>Error while fetching data</div>;
    }

    const data =  response.json();
    console.log(data);}
    
    catch(error){
      console.log(error)
    }


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
              Card Type
              </p>
              <span className="text-[12px] sm:text-[15px] text-[#718EBF]">
                {item.col1.sub}
              </span>
            </div>
      
            <div className=" col-span-[2.5] ">
              <p className="text-[14px] md:[text-[12px]] lg:text-[16px] text-[#333B69]">
                Bank
              </p>
              <span className="text-[12px]  md:[text-[12px]] lg:text-[16px] text-[#718EBF]">
                {item.col2.sub}
              </span>
            </div>



            <div className=" hidden col-span-[2.5] sm:block ">
              <p className="text-[14px] md:text-[12px]  lg:text-[16px] text-[#333B69] font-medium">
                Card Number
              </p>
              <span className="text-[12px] sm:text-[15px] text-[#718EBF]">
                {item.col3.sub}
              </span>
            </div>


            <div className=" hidden col-span-2 sm:block ">
              <p className="text-[14px] sm:text-[16px] text-[#333B69] font-medium">
                Nomain Card
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
