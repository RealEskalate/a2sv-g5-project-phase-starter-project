import React from "react";
import Image from "next/image";

const loans = () => {
  return (
    <>
      <main className="">
        <div className="min-h-screen bg-gray-100 p-6 flex justify-center items-start">
          <div className="max-w-7xl mx-auto">
            {/* Loan Summary Cards */}
            <div className="min-w-full flex justify-between items-center mb-8">
              <div className="flex gap-6 px-10">
                {/* Personal Loans Card */}
                <div className="flex items-center bg-white p-4 rounded-lg shadow-md w-56">
                  <span className="bg-blue-100 text-blue-500 p-2 rounded-full">
                    <Image src="/user 3 2.svg" alt="Personal Loans" className="h-6 w-6" />
                  </span>
                  <div className="ml-3">
                    <span className="block text-gray-700 font-inter font-medium text-sm leading-5">Personal Loans</span>
                    <span className="block text-gray-900 font-bold text-lg">$50,000</span>
                  </div>
                </div>
                {/* Corporate Loans Card */}
                <div className="flex items-center bg-white p-4 rounded-lg shadow-md w-56">
                  <span className="bg-yellow-100 text-yellow-500 p-2 rounded-full">
                    <Image src="/briefcase 1.svg" alt="Corporate Loans" className="h-6 w-6" />
                  </span>
                  <div className="ml-3">
                    <span className="block text-gray-700 font-inter font-medium text-sm leading-5">Corporate Loans</span>
                    <span className="block text-gray-900 font-bold text-lg">$100,000</span>
                  </div>
                </div>
                {/* Business Loans Card */}
                <div className="flex items-center bg-white p-4 rounded-lg shadow-md w-56">
                  <span className="bg-pink-100 text-pink-500 p-2 rounded-full">
                    <Image src="/graph 1.svg" alt="Business Loans" className="h-6 w-6" />
                  </span>
                  <div className="ml-3">
                    <span className="block text-gray-700 font-inter font-medium text-sm leading-5">Business Loans</span>
                    <span className="block text-gray-900 font-bold text-lg">$500,000</span>
                  </div>
                </div>
                {/* Custom Loans Card */}
                <div className="flex items-center bg-white p-4 rounded-lg shadow-md w-56">
                  <span className="bg-green-100 text-green-500 p-2 rounded-full">
                    <Image src="/support 1.svg" alt="Custom Loans" className="h-6 w-6" />
                  </span>
                  <div className="ml-3">
                    <span className="block text-gray-700 font-inter font-medium text-sm leading-5">Custom Loans</span>
                    <span className="block text-gray-900 font-bold text-lg">Choose Money</span>
                  </div>
                </div>
              </div>
            </div>

            
            <div className="ml-10">
              <h2 className="mx-2 m text-xl font-bold mb-4 background: #333B69;
">Active Loans Overview</h2>
            </div>

            {/* Table Section */}
            <div className="bg-white p-6 rounded-lg shadow-md mx-10">
              <table className="min-w-full text-left text-sm">
                <thead>
                  <tr className="border-b">
                    <th className="py-2">SL No.</th>
                    <th className="py-2">Loan Money</th>
                    <th className="py-2">Left to repay</th>
                    <th className="py-2">Duration</th>
                    <th className="py-2">Interest rate</th>
                    <th className="py-2">Installment</th>
                    <th className="py-2">Repay</th>
                  </tr>
                </thead>
                <tbody>
                  {/* Row Template */}
                  {[...Array(8)].map((_, idx) => (
                    <tr key={idx} className="border-b">
                      <td className="py-2">{`0${idx + 1}.`}</td>
                      <td className="py-2">$100,000</td>
                      <td className="py-2">$40,500</td>
                      <td className="py-2">8 Months</td>
                      <td className="py-2">12%</td>
                      <td className="py-2">$2,000 / month</td>
                      <td className="py-2">
                        <button className="border border-blue-500 text-blue-500 py-1 px-3 rounded-full hover:bg-blue-500 hover:text-white transition duration-200">
                          Repay
                        </button>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>

              {/* Totals Row */}
              <div className="mt-6 flex justify-between text-red-500 font-semibold">
                <span>Total</span>
                <span>$125,000</span>
                <span>$750,000</span>
                <span></span>
                <span></span>
                <span>$50,000 / month</span>
              </div>
            </div>
          </div>
        </div>
      </main>
    </>
  );
};

export default loans;
