import React from 'react'

const ActiveLoans = () => {
  return (
    <div>
        <div>
           <h2 className="text-lg font-bold mb-4">Active Loans Overview</h2>
           </div>

        <div className="bg-white p-6 rounded-lg shadow-md">
            
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
                
                <tr>
                  <td className="py-2"></td>
                  <td className="py-2"></td>
                  <td className="py-2"></td>
                  <td className="py-2"></td>
                  <td className="py-2"></td>
                  <td className="py-2"></td>
                  <td className="py-2">
                    <button className="border border-blue-500 text-black py-1 px-3 rounded-full">
                      Repay
                    </button>
                  </td>
                </tr>
                
              </tbody>
            </table>
            <div className="mt-4 text-right">
            {/* Total */}
            </div>
          </div>
    </div>
  )
}

export default ActiveLoans