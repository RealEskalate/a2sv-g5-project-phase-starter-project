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
import { usePathname } from 'next/navigation';
import { useAppDispatch, useAppSelector } from '@/hooks/hoooks';
import { toggleHamburgerMenu } from '@/lib/redux/slices/uiSlice';

const data = [
  {
    icon: HomeIcon,
    title: 'Dashboard',
    active: true,
    url: '/bank-dash',
  },
  {
    icon: Transaction,
    title: 'Transactions',
    active: false,
    url: '/bank-dash/transactions',
  },
  {
    icon: Accounts,
    title: 'Accounts',
    active: false,
    url: '/bank-dash/accounts',
  },
  {
    icon: InvestmentIcon,
    title: 'Investment',
    active: false,
    url: '/bank-dash/investments',
  },
  {
    icon: CreditCard,
    title: 'Credit Card',
    active: false,
    url: '/bank-dash/credit-card',
  },
  {
    icon: Loans,
    title: 'Loans',
    active: false,
    url: '/bank-dash/loans',
  },
  {
    icon: Services,
    title: 'Services',
    active: false,
    url: '/bank-dash/services',
  },
  {
    icon: MyPrivileges,
    title: 'My Privileges',
    active: false,
    url: '/bank-dash/my-privileges',
  },
  {
    icon: Setting,
    title: 'Setting',
    active: false,
    url: '/bank-dash/settings',
  },
];

export default function Sidebar() {
  const asPath = usePathname();
  const hamburgerMenu = useAppSelector((state) => state.ui.hamburgerMenu);
  const dispatch = useAppDispatch();
  const handleClick = () => {
    dispatch(toggleHamburgerMenu());
  };

  const handleClickOutside = (e: any) => {
    if (e.target.classList.contains('backdrop-blur-sm')) {
      dispatch(toggleHamburgerMenu());
    }
  };

  return (
    <>
      <aside
        className={`fixed ${
          !hamburgerMenu
            ? '-translate-x-full'
            : 'w-screen backdrop-blur-sm bg-opacity-25 bg-gray-700'
        } sm:relative top-0 left-0 z-40 transition-transition  sm:translate-x-0 duration-300`}
        onClick={handleClickOutside}
      >
        <div className='min-h-screen w-[215px] py-4 overflow-y-auto bg-white'>
          <div className='flex sm:block items-center mb-8'>
            <p className='flex items-center px-3 text-xl font-bold text-navy  my-3'>
              <Logo className='w-[28px] h-[25px] mx-2' /> BankDash.
            </p>
          </div>
          {data.map((ele) => (
            <Link href={ele.url} key={ele.title} onClick={handleClick}>
              <SideBarItems Icon={ele.icon} title={ele.title} active={ele.url === asPath} />
            </Link>
          ))}
        </div>
      </aside>
    </>
  );
}
