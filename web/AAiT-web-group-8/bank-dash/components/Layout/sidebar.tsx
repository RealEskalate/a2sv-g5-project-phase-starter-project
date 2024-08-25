'use client'

import React, { useState, useEffect, useRef } from 'react';
import Link from 'next/link';
import { MdHome, MdMiscellaneousServices, MdSettings, MdBarChart, MdClose } from "react-icons/md";
import { FaMoneyBillTransfer, FaUser, FaCreditCard, FaHandHoldingDollar, FaLightbulb } from "react-icons/fa6";
import { SidebarProps } from '../../types'

const Sidebar: React.FC<SidebarProps> = ({ isOpen, toggleSidebar }) => {
  const sidebarRef = useRef<HTMLDivElement>(null);

  const handleClickOutside = (event: MouseEvent) => {
    if (sidebarRef.current && !sidebarRef.current.contains(event.target as Node)) {
      toggleSidebar();
    }
  };

  const handleResize = () => {
    if (window.innerWidth >= 768 && isOpen) {
      toggleSidebar();
    }
  };

  useEffect(() => {
    if (isOpen) {
      document.addEventListener('mousedown', handleClickOutside);
    } else {
      document.removeEventListener('mousedown', handleClickOutside);
    }
    window.addEventListener('resize', handleResize);

    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
      window.removeEventListener('resize', handleResize);
    };
  }, [isOpen]);

  return (
    <div ref={sidebarRef} className={`fixed inset-0 z-50 bg-white p-4 transition-transform transform ${isOpen ? 'translate-x-0' : '-translate-x-full'} md:translate-x-0 md:static md:w-auto md:h-auto md:bg-transparent`}>
      <div className="flex justify-between items-center p-6">
        <h1 className="text-2xl font-extrabold text-blue-900">BankDash.</h1>
        <button onClick={toggleSidebar} className="md:hidden">
          <MdClose className="text-blue-900 w-6 h-6" />
        </button>
      </div>
      <ul className="space-y-2 p-4">
        <li className="flex items-center text-gray-500 space-x-4 px-6 py-2 hover:text-blue-600 cursor-pointer">
          <Link href="/" className="flex items-center space-x-4">
            <MdHome className="text-inherit w-6 h-6" />
            <span className="text-inherit font-semibold">Dashboard</span>
          </Link>
        </li>
        <li className="flex items-center text-gray-500 space-x-4 px-6 py-2 hover:text-blue-600 cursor-pointer">
          <Link href="/transactions" className="flex items-center space-x-4">
            <FaMoneyBillTransfer className="text-inherit w-6 h-6" />
            <span className="text-inherit font-semibold">Transactions</span>
          </Link>
        </li>
        <li className="flex items-center text-gray-500 space-x-4 px-6 py-2 hover:text-blue-600 cursor-pointer">
          <Link href="/accounts" className="flex items-center space-x-4">
            <FaUser className="text-inherit w-6 h-6" />
            <span className="text-inherit font-semibold">Accounts</span>
          </Link>
        </li>
        <li className="flex items-center text-gray-500 space-x-4 px-6 py-2 hover:text-blue-600 cursor-pointer">
          <Link href="/investments" className="flex items-center space-x-4">
            <MdBarChart className="text-inherit w-6 h-6" />
            <span className="text-inherit font-semibold">Investments</span>
          </Link>
        </li>
        <li className="flex items-center text-gray-500 space-x-4 px-6 py-2 hover:text-blue-600 cursor-pointer">
          <Link href="/credit-cards" className="flex items-center space-x-4">
            <FaCreditCard className="text-inherit w-6 h-6" />
            <span className="text-inherit font-semibold">Credit Cards</span>
          </Link>
        </li>
        <li className="flex items-center text-gray-500 space-x-4 px-6 py-2 hover:text-blue-600 cursor-pointer">
          <Link href="/loans" className="flex items-center space-x-4">
            <FaHandHoldingDollar className="text-inherit w-6 h-6" />
            <span className="text-inherit font-semibold">Loans</span>
          </Link>
        </li>
        <li className="flex items-center text-gray-500 space-x-4 px-6 py-2 hover:text-blue-600 cursor-pointer">
          <Link href="/services" className="flex items-center space-x-4">
            <MdMiscellaneousServices className="text-inherit w-6 h-6" />
            <span className="text-inherit font-semibold">Services</span>
          </Link>
        </li>
        <li className="flex items-center text-gray-500 space-x-4 px-6 py-2 hover:text-blue-600 cursor-pointer">
          <Link href="/privileges" className="flex items-center space-x-4">
            <FaLightbulb className="text-inherit w-6 h-6" />
            <span className="text-inherit font-semibold">My Privileges</span>
          </Link>
        </li>
        <li className="flex items-center text-gray-500 space-x-4 px-6 py-2 hover:text-blue-600 cursor-pointer">
          <Link href="/settings" className="flex items-center space-x-4">
            <MdSettings className="text-inherit w-6 h-6" />
            <span className="text-inherit font-semibold">Setting</span>
          </Link>
        </li>
      </ul>
    </div>
  );
};

export default Sidebar;