"use client";
import React from "react";
import { redirect } from "next/navigation";
import { useSession } from "next-auth/react";
import Image from "next/image";

const Page = () => {
  const { data: session } = useSession({
    required: true,
    onUnauthenticated() {
      redirect("/api/auth/signin?calbackUrl=/login");
    },
  });
  const loans = [
    {
      id: 1,
      loanMoney: "$100,000",
      leftToRepay: "$40,500",
      duration: "8 Months",
      interestRate: "12%",
      installment: "$2,000 / month",
    },
    {
      id: 2,
      loanMoney: "$500,000",
      leftToRepay: "$250,000",
      duration: "36 Months",
      interestRate: "10%",
      installment: "$8,000 / month",
    },
    {
      id: 3,
      loanMoney: "$900,000",
      leftToRepay: "$40,500",
      duration: "12 Months",
      interestRate: "12%",
      installment: "$5,000 / month",
    },
    {
      id: 4,
      loanMoney: "$50,000",
      leftToRepay: "$40,500",
      duration: "25 Months",
      interestRate: "5%",
      installment: "$2,000 / month",
    },
    {
      id: 5,
      loanMoney: "$50,000",
      leftToRepay: "$40,500",
      duration: "5 Months",
      interestRate: "16%",
      installment: "$10,000 / month",
    },
    {
      id: 6,
      loanMoney: "$80,000",
      leftToRepay: "$25,500",
      duration: "14 Months",
      interestRate: "8%",
      installment: "$2,000 / month",
    },
    {
      id: 7,
      loanMoney: "$12,000",
      leftToRepay: "$5,000",
      duration: "9 Months",
      interestRate: "13%",
      installment: "$500 / month",
    },
    {
      id: 8,
      loanMoney: "$160,000",
      leftToRepay: "$100,800",
      duration: "3 Months",
      interestRate: "12%",
      installment: "$900 / month",
    },
  ];

  return (
    <>
      <main className="">
        <div className="min-h-screen bg-gray-100 p-6">
          <div className="max-w-7xl mx-auto">
            <div className="flex justify-between items-center mb-6 w-[1110px] h-[120px] ml-[70px]  mr-[10px]">
              <div className="flex space-x-4 w-full">
                <div className="w-1/4  items-center justify-between bg-white p-4 rounded-lg shadow-md">
                  <div className="flex items-center">
                    <span className="bg-blue-100 text-blue-500 p-2 rounded-full">
                      <Image
                        src="/user 3 2.svg"
                        alt="Personal Loans"
                        className="h-6 w-6"
                      />
                    </span>
                    <span className="ml-3">Personal Loans</span>
                  </div>
                  <span className="ml-6 text-lg font-bold">$50,000</span>
                </div>

                <div className="w-1/4 items-center justify-between bg-white p-4 rounded-lg shadow-md">
                  <div className="flex items-center">
                    <span className="bg-yellow-100 text-yellow-500 p-2 rounded-full">
                      {/* <img
                        src="/briefcase 1.svg"
                        alt="Corporate Loans"
                        className="h-6 w-6"
                      /> */}
                      <Image
                        src="/user 3 2.svg"
                        alt="Corporate Loans"
                        className="h-6 w-6"
                      />
                    </span>
                    <span className="ml-3">Corporate Loans</span>
                  </div>
                  <span className="ml-6 text-lg font-bold">$100,000</span>
                </div>

                <div className="w-1/4  items-center justify-between bg-white p-4 rounded-lg shadow-md">
                  <div className="flex items-center">
                    <span className="bg-pink-100 text-pink-500 p-2 rounded-full">
                      <Image
                        src="/user 3 2.svg"
                        alt="Business Loans"
                        className="h-6 w-6"
                      />
                    </span>
                    <span className="ml-3">Business Loans</span>
                  </div>
                  <span className="ml-6 text-lg font-bold">$500,000</span>
                </div>

                <div className="w-1/4  items-center justify-between bg-white p-4 rounded-lg shadow-md">
                  <div className="flex items-center">
                    <span className="bg-green-100 text-green-500 p-2 rounded-full">
                      {/* <img
                        src="/support 1.svg"
                        alt="Custom Loans"
                        className="h-6 w-6"
                      /> */}
                      <Image
                        src="/user 3 2.svg"
                        alt="Custom Loans"
                        className="h-6 w-6"
                      />
                    </span>
                    <span className="ml-3">Custom Loans</span>
                  </div>
                  <span className="ml-6 text-lg font-bold">Choose Money</span>
                </div>
              </div>
            </div>

            <div>
              <h2 className="text-lg font-bold mb-4 ml-[70px]">
                Active Loans Overview
              </h2>
            </div>

            <div className="bg-white p-6 rounded-lg shadow-md w-[1110px] h-[625px] ml-[70px]">
              <div>
                <table className="min-w-full bg-white">
                  <thead>
                    <tr>
                      <th className="py-2 text-left">SL No.</th>
                      <th className="py-2 text-left">Loan Money</th>
                      <th className="py-2 text-left">Left to repay</th>
                      <th className="py-2 text-left">Duration</th>
                      <th className="py-2 text-left">Interest rate</th>
                      <th className="py-2 text-left">Installment</th>
                      <th className="py-2 text-left">Repay</th>
                    </tr>
                  </thead>
                  <tbody>
                    {loans.map((loan) => (
                      <tr key={loan.id}>
                        <td className="py-2">{loan.id}</td>
                        <td className="py-2">{loan.loanMoney}</td>
                        <td className="py-2">{loan.leftToRepay}</td>
                        <td className="py-2">{loan.duration}</td>
                        <td className="py-2">{loan.interestRate}</td>
                        <td className="py-2">{loan.installment}</td>
                        <td className="py-2">
                          <button className="border border-blue-500 text-black py-1 px-3 rounded-full w-[100px] h-[35px]">
                            Repay
                          </button>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                  <tfoot>
                    <tr>
                      <td className="py-2 font-bold text-red-500">Total</td>
                      <td className="py-2 font-bold text-red-500">
                        $1,125,000
                      </td>
                      <td className="py-2 font-bold text-red-500">$750,000</td>
                      <td className="py-2"></td>
                      <td className="py-2"></td>
                      <td className="py-2 font-bold text-red-500">
                        $50,000 / month
                      </td>
                      <td className="py-2"></td>
                    </tr>
                  </tfoot>
                </table>
              </div>
            </div>
          </div>
        </div>
      </main>
    </>
  );
};

export default Page;
