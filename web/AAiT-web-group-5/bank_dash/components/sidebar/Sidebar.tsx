import React from 'react';
import SideList from './SideList'; // Adjust the path based on your file structure
import MenuIcons from '../icons/MenuIcons';
import { Box, Toolbar } from '@mui/material';
import Image from 'next/image';
import { title } from 'process';

const SideBar = () => {
  const menuItems = [
    {
      id: 'dashboard',
      title: 'Dashboard',
      link: '/',
      icon: <MenuIcons.Homeicon />,
      available: true
    },
    {
      id: 'transaction',
      title: 'Transaction',
      link: '/transaction',
      icon: <MenuIcons.Transfericon />,
      available: true
    },
    {
      id: 'account',
      title: 'Account',
      link: '/account',
      icon: <MenuIcons.Usericon />,
      available: true
    },
    {
      id: 'investment',
      title: 'Investment',
      link: '/investment',
      icon: <MenuIcons.Investicon />,
      available: true
    },
    {
      id: 'credit-card',
      title: 'Credit Card',
      link: '/credit-card',
      icon: <MenuIcons.CreditCardicon />,
      available: true
    },
    {
      id: 'services',
      title: 'Services',
      link: '/services',
      icon: <MenuIcons.Serviceicon />,
      available: true
    },
    {
      id: 'loan',
      title: 'Loan',
      link: '/loan',
      icon: <MenuIcons.Loanicon />,
      available: true
    },
    {
      id: 'my-privileges',
      title: 'My Privileges',
      link: '/my-privileges',
      icon: <MenuIcons.Privilegeicon />,
      available: true
    },
    {
      id: 'settings',
      title: 'Settings',
      link: '/settings',
      icon: <MenuIcons.Settingsicon />,
      available: true
    },
  ];

  return (
    <div className='flex flex-col gap-5'>
        <Toolbar sx={{ gap: 1, minHeight: 100, cursor: 'pointer' }}>
            <MenuIcons.Logo />
        </Toolbar>

        <Box component="nav" className='border-r min-h-screen w-64 flex flex-col' >
        {menuItems.map((item) => (
            <SideList key={item.id} sideitem={item} />
        ))}
    </Box>

    </div>
    
  );
};

export default SideBar;


