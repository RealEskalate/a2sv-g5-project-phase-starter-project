
'use client';
import Sidebar from '../components/Common/Sidebar';
import UserProfile from '../components/setting';
import Header from '../components/Common/Navbar';

const Page = () => {
  return (
    <div className='bg-[#F5F7FA] min-h-screen w-[100%]'>
        <div className="flex border-r border-[#E6EFF5] h-screen">
        {/* <Sidebar /> */}
        {/* <div className="border-r border-[#E6EFF5] h-screen"></div> */}
        <div className="flex-1">
            <Header/>
            <div className="border-t border-[#E6EFF5]"></div>
            <div className="p-4 mt-[-10px]">
            <UserProfile />
            </div>
        </div>
        </div>
    </div>
  );
};

export default Page;