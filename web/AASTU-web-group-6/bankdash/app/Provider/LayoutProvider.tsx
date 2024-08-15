"use client";

import React, { ReactNode } from 'react';
import { usePathname } from 'next/navigation';
import Sidebar from '../components/Layout/Sidebar';
import NavBar from '../components/Layout/NavBar';

const LayoutProvider = ({children}:{children: ReactNode}) => {

    const pathname = usePathname();

  const noLayoutPaths = ["/login", "/signup"];

  const shouldRenderLayout = !noLayoutPaths.includes(pathname);


  return (
    <div>
        {shouldRenderLayout ? (
          <div className="w-full flex">
            <Sidebar />
            <div className="w-full py-6 px-5">
              <NavBar />
              {children}
            </div>
          </div>
        ) : (
          <div>{children}</div> 
        )}
    </div>
  )
}

export default LayoutProvider