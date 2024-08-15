'use client'

import { FC, useState } from 'react';
import { Sidebar } from '@/components/Sidebar';
import { Navbar } from '@/components/Navbar';
import { usePathname } from 'next/navigation';
import { sidebarLinks, user } from '@/constants';



const RootLayout: FC<{ children: React.ReactNode }> = ({ children }) => {
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);
  const toggleSidebar = () => setIsSidebarOpen(!isSidebarOpen);

  const pathname = usePathname();
  const pageTitle = pathname === "/"
    ? "Overview"
    : sidebarLinks.find(link => link.route === pathname)?.label || "";

  return (
<<<<<<< HEAD
    <main className="flex h-screen w-full font-inter">
      <Sidebar user={{
        $id: "",
        email: "",
        userId: "",
        dwollaCustomerUrl: "",
        dwollaCustomerId: "",
        firstName: "",
        lastName: "",
        name: "",
        address1: "",
        city: "",
        state: "",
        postalCode: "",
        dateOfBirth: "",
        ssn: ""
      }}/>

      <div className="flex size-full flex-col">
        <div className="root-layout">
          <Image src="/icons/logo.svg" width={30} height={30} alt="logo" />
          <div>
            <MobileNav user={{
              $id: "",
              email: "",
              userId: "",
              dwollaCustomerUrl: "",
              dwollaCustomerId: "",
              firstName: "",
              lastName: "",
              name: "",
              address1: "",
              city: "",
              state: "",
              postalCode: "",
              dateOfBirth: "",
              ssn: ""
            }}/>
          </div>
        </div>
        {children}
=======
    <div className="flex">
      <Sidebar isOpen={isSidebarOpen} toggleSidebar={toggleSidebar} />
      <div className="flex-1 flex flex-col">
        <Navbar pageTitle={pageTitle} toggleSidebar={toggleSidebar} userProfileImage={user.profileImage} />
        <main className="flex-grow">{children}</main>
>>>>>>> ec6c4ad143429aa0ad205c2a553e5d11744d4b00
      </div>
    </div>
  );
};

export default RootLayout;
