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

const data = [
  {
    icon: HomeIcon,
    title: 'Dashboard',
    active: true,
  },
  {
    icon: Transaction,
    title: 'Transactions',
    active: false,
  },
  {
    icon: Accounts,
    title: 'Accounts',
    active: false,
  },
  {
    icon: InvestmentIcon,
    title: 'Investment',
    active: false,
  },
  {
    icon: CreditCard,
    title: 'Credit Card',
    active: false,
  },
  {
    icon: Loans,
    title: 'Loans',
    active: false,
  },
  {
    icon: Services,
    title: 'Services',
    active: false,
  },
  {
    icon: MyPrivileges,
    title: 'My Privileges',
    active: false,
  },
  {
    icon: Setting,
    title: 'Setting',
    active: false,
  },
];

export default function Sidebar() {
  return (
    <>
      <>
        <aside className='fixed sm:relative top-0 left-0 z-40 w-[230px] transition-transition -translate-x-full sm:translate-x-0 duration-300 '>
          <div className='min-h-screen py-4 overflow-y-auto bg-white '>
            <p className='flex items-center px-3 text-18px font-bold text-navy'>
              <Logo className='w-[28px] h-[25px] mx-2' /> BankDash.
            </p>
            {data.map((ele) => (
              <SideBarItems key={ele.title} Icon={ele.icon} title={ele.title} active={ele.active} />
            ))}
          </div>
        </aside>
      </>
    </>
  );
}
