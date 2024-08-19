import React from "react";
import { FaRegArrowAltCircleUp, FaRegArrowAltCircleDown } from "react-icons/fa";

const loans = [
    {
        "sl_no": 1,
        "loan_amount": 10000,
        "left_to_repay": 7500,
        "duration_months": 36,
        "interest_rate": 5.0,
        "monthly_installment": 300
    },
    {
        "sl_no": 2,
        "loan_amount": 25000,
        "left_to_repay": 18000,
        "duration_months": 60,
        "interest_rate": 4.5,
        "monthly_installment": 470
    },
    {
        "sl_no": 3,
        "loan_amount": 15000,
        "left_to_repay": 12000,
        "duration_months": 48,
        "interest_rate": 6.0,
        "monthly_installment": 350
    },
    {
        "sl_no": 4,
        "loan_amount": 50000,
        "left_to_repay": 45000,
        "duration_months": 72,
        "interest_rate": 3.8,
        "monthly_installment": 750
    },
    {
        "sl_no": 5,
        "loan_amount": 8000,
        "left_to_repay": 4000,
        "duration_months": 24,
        "interest_rate": 5.5,
        "monthly_installment": 200
    },
    {
        "sl_no": 6,
        "loan_amount": 30000,
        "left_to_repay": 22500,
        "duration_months": 60,
        "interest_rate": 4.0,
        "monthly_installment": 550
    },
    {
        "sl_no": 7,
        "loan_amount": 12000,
        "left_to_repay": 9000,
        "duration_months": 36,
        "interest_rate": 5.2,
        "monthly_installment": 350
    },
    {
        "sl_no": 8,
        "loan_amount": 20000,
        "left_to_repay": 15000,
        "duration_months": 48,
        "interest_rate": 4.8,
        "monthly_installment": 430
    },
    {
      "sl_no": 9,
      "loan_amount": 20000,
      "left_to_repay": 15000,
      "duration_months": 48,
      "interest_rate": 4.8,
      "monthly_installment": 430
  }, 
  {
    "sl_no": 10,
    "loan_amount": 20000,
    "left_to_repay": 15000,
    "duration_months": 48,
    "interest_rate": 4.8,
    "monthly_installment": 430
}
]




  const formatDate = (dateString:string) => {
    const date = new Date(dateString)
  

    const formattedDate = date.toLocaleDateString('en-US', {
      month: 'short',
      day: 'numeric',
    })

    const formattedTime = date.toLocaleTimeString('en-US', {
      hour: '2-digit',
      minute: '2-digit',
      hour12: true,
    })
  
    return `${formattedDate}, ${formattedTime}`
  }
  

const ActiveLoans = () => {
  return (
    <div className="">
      
  <h1 className="lg:m-10 text-2xl font-semibold my-4">Active Loans Overview</h1>        
    <section className="border-0 rounded-xl bg-white shadow-md lg:mx-10 p-2">
      <div className="grid grid-cols-3 lg:grid-cols-7 font-medium text-sky-300 min-h-7 items-center border-b mt-2 px-2">
        <div className="hidden md:block">Sl NO</div>
        <div>Loan Money</div>
        <div>Left To Repay</div>
        <div className="hidden md:block">Durations </div>
        <div className="hidden md:block">Interest RaTE</div>
        <div className="hidden md:block">Instillment</div>
        <div className="justify-self-center">Repay</div>
      </div>
 

      {loans.map((loan, index) => (
        <div
          key={index}
          className="grid grid-cols-3 lg:grid-cols-7  border-b min-h-12 items-center   "
        >
          <div className="hidden md:block ">
            {loan.sl_no}
          </div>
          <div className="">${loan.loan_amount}</div>
          <div className="">${loan.left_to_repay}</div>
          <div className="hidden md:block">{loan.duration_months} months</div>
          <div className="hidden md:block">{loan.interest_rate}%</div>
          <div className={`hidden md:block`}>
            ${loan.monthly_installment}
          </div>
          <div className=" border border-blue-200 text-center p-1 rounded-lg justify-self-center hover:border-blue-700 w-24">Repay</div>
        </div>
      ))}
    </section>
    </div>
  );
};

export default ActiveLoans;
