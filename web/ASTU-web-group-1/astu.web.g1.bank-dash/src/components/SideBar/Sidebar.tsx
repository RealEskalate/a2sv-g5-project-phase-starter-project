'use client';
import React from 'react';
import SideBarItems from '../SideBarItems/SideBarItems';
import HomeIcon from '../../../public/assets/icons/home-icon.svg';
import Transaction from '../../../public/assets/icons/transaction-icon.svg';
import Accounts from '../../../public/assets/icons/accounts-icon.svg';
import InvestmentIcon from '../../../public/assets/icons/investment-icon.svg';
import CreditCard from '../../../public/assets/icons/credit-card-icon.svg';
import Loans from '../../../public/assets/icons/loans-icon.svg';
import Services from '../../../public/assets/icons/services-icon.svg';
import MyPrivileges from '../../../public/assets/icons/myPrivileges-icon.svg';
import Setting from '../../../public/assets/icons/setting-icon.svg';
import Logo from '../../../public/assets/icons/logo-icon.svg';
import Link from 'next/link';
import { useParams, usePathname } from 'next/navigation';
import { useRouter } from 'next/navigation';

const data = [
  {
    icon: HomeIcon,
    title: 'Dashboard',
    active: true,
    url: '/',
  },
  {
    icon: Transaction,
    title: 'Transactions',
    active: false,
    url: '/transactions',
  },
  {
    icon: Accounts,
    title: 'Accounts',
    active: false,
    url: '/accounts',
  },
  {
    icon: InvestmentIcon,
    title: 'Investment',
    active: false,
    url: '/investments',
  },
  {
    icon: CreditCard,
    title: 'Credit Card',
    active: false,
    url: '/credit-card',
  },
  {
    icon: Loans,
    title: 'Loans',
    active: false,
    url: '/loans',
  },
  {
    icon: Services,
    title: 'Services',
    active: false,
    url: '/services',
  },
  {
    icon: MyPrivileges,
    title: 'My Privileges',
    active: false,
    url: '/my-privileges',
  },
  {
    icon: Setting,
    title: 'Setting',
    active: false,
    url: '/settings',
  },
];

export default function Sidebar() {
  const asPath = usePathname();
  console.log(asPath, 'as');
  return (
    <>
      <aside className='fixed sm:relative top-0 left-0 z-40 w-[215px] transition-transition -translate-x-full sm:translate-x-0 duration-300 '>
        <div className='min-h-screen py-4 overflow-y-auto bg-white '>
          <p className='flex items-center px-3 text-xl font-bold text-navy mb-10 my-3'>
            <Logo className='w-[28px] h-[25px] mx-2' /> BankDash.
          </p>
          {data.map((ele) => (
            <Link href={ele.url} key={ele.title}>
              <SideBarItems Icon={ele.icon} title={ele.title} active={ele.url === asPath} />
            </Link>
          ))}
        </div>
      </aside>
    </>
  );
}
