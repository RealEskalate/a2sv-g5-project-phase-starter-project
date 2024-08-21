'use client'

import React, { useEffect, useState } from 'react';
import { Carousel, CarouselItem, CarouselContent } from '@/components/ui/carousel';
import { loanCardMapping } from '@/constants/index';
import Link from 'next/link';
import { getMyLoans, getLoanDetailData } from '@/services/activeloan';
import CustomLoans from '@/public/icons/CustomLoans';

interface LoanCard {
  icon: (props: any) => JSX.Element; // Define the icon type
  title: string;
  loanAmount: string;
}



const LoanCard: React.FC<{ icon: React.FC<React.SVGProps<SVGSVGElement>>; title: string; description: string }> = ({
  icon: Icon,
  title,
  description,
}) => (
  <div className="flex items-center space-x-4">
    <Icon className="pl-1 w-12 h-12" aria-hidden="true" />
    <div>
      <h3 className="text-sm text-gray-500">{title}</h3> {/* Smaller, grayer title */}
      <p className="text-lg font-extrabold">{description}</p> {/* Extra bold and larger description */}
    </div>
  </div>
);

const LoansPage: React.FC = () => {
  
  const [loanCards, setLoanCards] = useState<{ icon: React.FC; title: string; loanAmount: string }[]>([]);  
  const [loans, setLoans] = useState([])

  useEffect(() => {
    const fetchLoans = async () => {

      try {
        const response = await getMyLoans();
        // console.log("Response: ", response)
        setLoans(response.data)
      } catch (error) {
        console.error("Error fetching the loans: ", error)
      }
    }
    fetchLoans();
  }, [])

// Fetching loan card data
useEffect(() => {
  const fetchLoanCard = async () => {
    try {
      const response = await getLoanDetailData();
      const { data } = response;

      // Map the API data to the loan cards with icons and titles
      const loanCard = loanCardMapping.map((loan) => ({
        icon: loan.icon,
        title: loan.title,
        loanAmount: `$${data[loan.descriptionKey]?.toLocaleString() || 0}`,
      }));

      // Add the constant Custom Loans card
      loanCard.push({
        icon: CustomLoans,
        title: 'Custom Loans',
        loanAmount: 'Choose Loans', // Or any default value you want for the constant card
      });

      setLoanCards(loanCard);
    } catch (error) {
      console.error('Error fetching the loan card data:', error);
    }
  };

  fetchLoanCard();
}, []);


  const totalLoanMoney = loans.reduce((total, loan: any) => total + parseInt(loan.loanAmount), 0);
  const totalLeftToRepay = loans.reduce((total, loan: any) => total + parseInt(loan.amountLeftToRepay), 0);
  const totalInstallment = loans.reduce((total, loan: any) => total + parseInt(loan.installment), 0);

  return (
    <div className=" flex lg:ml-72 lg:max-w-[100%]">

      {/* Mobile and Tablet View */}
      <div className="w-[100%] block lg:hidden">
        <Carousel>
          <CarouselContent className="p-6 ">
            {loanCards.map((loanItem, index) => (
              <CarouselItem key={index} className="w-[240px] h-[85px] mx-auto mr-4 flex-none">
                <div className="shadow-lg p-4 rounded-md flex items-center h-full">
                  <LoanCard icon={loanItem.icon} title={loanItem.title} description={loanItem.loanAmount} />
                </div>
              </CarouselItem>
            ))}
          </CarouselContent>
        </Carousel>

        {/* Active Loans Overview - Mobile View */}
        <div className="w-[100%] flex flex-col items-center mt-8 text-sm">
          <h2 className="text-lg font-bold mb-4 ml-5">Active Loans Overview</h2>
          <div className="w-[100%] flex justify-center ">
            <table className=" w-[70%] md:w-[60%]  h-[85px] bg-white rounded-lg shadow-md text-[12px]">
              <thead>
                <tr>
                  <th className="px-4 py-2 text-left font-semibold text-gray-400">Loan Money</th>
                  <th className="px-4 py-2 text-left font-semibold text-gray-400">Left to Repay</th>
                  <th className="px-4 py-2 text-left font-semibold text-gray-400">Repay</th>
                </tr>
              </thead>
              <tbody>
                {loans.map((loan: any, index) => (
                  <tr key={index}>
                    <td className="border-t px-4 py-2">{loan.loanAmount}</td>
                    <td className="border-t px-4 py-2">{loan.amountLeftToRepay}</td>
                    <td className="border-t px-4 py-2">
                      <Link href="#" className="text-purple-900 border border-purple-900 rounded-full px-4 py-1">
                        Repay
                      </Link>
                    </td>
                  </tr>
                ))}
                <tr className="font-bold text-red-500">
                  <td className="border-t px-4 py-2">${totalLoanMoney.toLocaleString()}</td>
                  <td className="border-t px-4 py-2">${totalLeftToRepay.toLocaleString()}</td>
                  <td className="border-t px-4 py-2"></td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      {/* Desktop and Tablet View */}
      <div className="hidden lg:block lg:w-[100%] ">
        <div className=" w-[100%] flex gap-6 pr-6  py-8">
          {loanCards.map((loanItem, index) => (
            <div key={index} className="w-[100%] h-[120px] shadow-lg  rounded-lg  flex items-center">
              <div className="w-[100%]">
                <LoanCard icon={loanItem.icon} title={loanItem.title} description={loanItem.loanAmount} />
              </div>
            </div>
          ))}
        </div>

        {/* Active Loans Overview - Desktop/Table View */}
        <div className="mt-8">
          <h2 className="text-lg font-bold mb-4">Active Loans Overview</h2>
          <div className="overflow-x-auto">
            <table className="w-[100%] bg-white rounded-2xl shadow-md table-fixed">
              <thead>
                <tr>
                  <th className="px-4 py-2 text-left font-semibold text-gray-400">SL No</th>
                  <th className="px-4 py-2 text-left font-semibold text-gray-400">Loan Money</th>
                  <th className="px-4 py-2 text-left font-semibold text-gray-400">Left to Repay</th>
                  <th className="px-4 py-2 text-left font-semibold text-gray-400">Duration</th>
                  <th className="px-4 py-2 text-left font-semibold text-gray-400">Interest Rate</th>
                  <th className="px-4 py-2 text-left font-semibold text-gray-400">Installment</th>
                  <th className="px-4 py-2 text-left font-semibold text-gray-400">Repay</th>
                </tr>
              </thead>
              <tbody>
                {loans.map((loan: any, index) => (
                  <tr key={index} className="text-sm leading-6">
                    <td className="border-t px-4 py-2">{index + 1}</td>
                    <td className="border-t px-4 py-2">{loan.loanAmount}</td>
                    <td className="border-t px-4 py-2">{loan.amountLeftToRepay}</td>
                    <td className="border-t px-4 py-2">{loan.duration}</td>
                    <td className="border-t px-4 py-2">{loan.interestRate}</td>
                    <td className="border-t px-4 py-2">{ `${Number(loan.installment ?? '0').toFixed(2)?.toLocaleString() || 0} $`   }</td>
                    <td className="border-t px-4 py-2">
                      <Link href="#" className="text-purple-900 border border-purple-900 rounded-full px-4 py-1">
                        Repay
                      </Link>
                    </td>
                  </tr>
                ))}
                <tr className="font-bold text-red-500">
                  <td className="border-t px-4 py-2">Total</td>
                  <td className="border-t px-4 py-2">${totalLoanMoney.toLocaleString()}</td>
                  <td className="border-t px-4 py-2">${totalLeftToRepay.toLocaleString()}</td>
                  <td className="border-t px-4 py-2"></td>
                  <td className="border-t px-4 py-2"></td>
                  <td className="border-t px-4 py-2">${totalInstallment.toLocaleString()}</td>
                  <td className="border-t px-4 py-2"></td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  );
};

export default LoansPage;
