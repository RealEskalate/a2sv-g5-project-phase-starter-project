'use client';  
import React, { useEffect } from 'react';
import Head from 'next/head';
import Navbar from '../components/navbar/Navbar';
import Sidebar from '../components/sidebar/Sidebar';
import { Inter } from 'next/font/google';
import { useSession } from 'next-auth/react';
import { useRouter } from 'next/navigation';

const inter = Inter({ subsets: ['latin'] });

const Layout = ({ children, title = 'My Next.js App' }: { children: React.ReactNode; title?: string }) => {
  const { data: session, status } = useSession();
  const router = useRouter();

  useEffect(() => {
    if (status === 'authenticated') {
      // User is authenticated, no action needed
      return;
    }
    
    if (status === 'unauthenticated') {
      // User is not authenticated, redirect to login
      router.push('/login');
    }
  }, [status, router]);

  if (status === 'loading') {
    return <p>Loading...</p>;
  }

  return (
    <>
      <Head>
        <title>{title}</title>
        <meta name="viewport" content="width=device-width, initial-scale=1" />
      </Head>
      <div className={`${inter.className} flex flex-col min-h-screen`}>
        <Navbar />
        <div className="flex flex-1">
          <Sidebar />
          <main className="max-md:pt-[100px] flex-1 p-4 mt-[60px] lg:ml-[240px] sm:ml-[240px] ml-0 bg-[#E6EFF5]">
            {children}
          </main>
        </div>
      </div>
    </>
  );
};

export default Layout;
