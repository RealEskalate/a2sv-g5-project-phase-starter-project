'use client'
import Sidebar from '../components/Common/Sidebar';
import Header from '../components/Common/Navbar';
import LoanCard from '../components/loancard1';
import CustomizedTables from '../components/loansTable';

const Page = () => {
  return (
    <div className='bg-[#F5F7FA] h-screen flex'>
      
      <div className="flex-1">
        
        
        <div className="border-t border-[#E6EFF5]"></div>
        <div className="p-4 mt-[-10px]">
          <div className='mt-5 ml-3'><LoanCard /></div>
          <h2 className='font-bold text-[22px] ml-4 text-[##333B69] mt-6'>Active Loans</h2>
          <div className='ml-3 mt-3'><CustomizedTables /></div>
        </div>
      </div>
    </div>
  );
};

export default Page;
