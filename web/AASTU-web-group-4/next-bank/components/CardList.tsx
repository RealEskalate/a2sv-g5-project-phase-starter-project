import React from "react";
import { colors,sidebarLinks } from "../constants/index";
import { CgCreditCard } from "react-icons/cg";

const CardList = () => {
  const cards = [
    {
      cardType: "Secondary",
      bank: "DBL Bank",
      cardNumber: "**** **** 5600",
      cardName: "William",
    },
    {
      cardType: "Secondary",
      bank: "BRC Bank",
      cardNumber: "**** **** 4300",
      cardName: "Michel",
    },
    {
      cardType: "Secondary",
      bank: "ABM Bank",
      cardNumber: "**** **** 7560",
      cardName: "Edward",
    },
    {
      cardType: "Primary",
      bank: "XYZ Bank",
      cardNumber: "**** **** 1234",
      cardName: "Alice",
    },
    {
      cardType: "Secondary",
      bank: "LMN Bank",
      cardNumber: "**** **** 5678",
      cardName: "Bob",
    },
    // Add more cards as needed
  ];

 


  return (
    <div className="max-h-[400px] lg:w-[730px] md:w-[487px] w-[325] overflow-y-scroll pr-6 py-4 scrollbar-thin scrollbar-thumb-rounded scrollbar-thumb-blue-400">
      {cards.map((card, index) => (
       <div
       key={index}
       className="flex flex-row justify-between w-full bg-white p-4 mb-4 rounded-lg shadow-md "
     >
       {/* Icon */}
       <div className="h-12 w-12 rounded-xl bg-[#E7EDFF] flex items-center justify-center">
         {/* Replace with actual icon */}
         <CgCreditCard className="w-7 h-7" />
         

         
       </div>
     
       {/* Card Type */}
       <div className="flex flex-col">
         <p className="text-sm font-medium">Card Type</p>
         <p className={`${colors.textgray} text-xs `}>{card.cardType}</p>
       </div>
     
       {/* Bank */}
       <div className="flex flex-col">
         <p className="text-sm font-medium">Bank</p>
         <p className={`${colors.textgray} text-xs`}>{card.bank}</p>
       </div>
     
       {/* Card Number - Hide on small screens */}
       <div className="flex flex-col">
         <p className="text-sm font-medium hidden sm:block">Card Number</p>
         <p className={`${colors.textgray} text-xs hidden sm:block`}>
           {card.cardNumber}
         </p>
       </div>
     
       {/* Card Name - Hide on small screens */}
       <div className="flex flex-col">
         <p className="text-sm font-medium hidden sm:block">Card Name</p>
         <p className={`${colors.textgray} text-xs hidden sm:block`}>
           {card.cardName}
         </p>
       </div>
     
       {/* View Details Link */}
       <a href="#" className="text-[#1814F3] font-semibold text-sm sm:text-xs">
         View Details
       </a>
     </div>
     
      ))}
    </div>
  );
};

export default CardList;

/* Group 343 */




