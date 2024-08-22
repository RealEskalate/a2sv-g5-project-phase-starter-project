'use client';

import React from 'react';
import { ListItem, ListItemIcon, ListItemText, Typography } from '@mui/material';
import Link from 'next/link';
import { usePathname } from 'next/navigation';
import MenuIcons from '../icons/MenuIcons';

interface SideListProps {
  id: string;
  title: string;
  link: string;
  icon: React.ReactNode;
  available: boolean;
}

interface SideLProps {
  sideitem: SideListProps;
  onDrawerClose?: () => void;
}

const SideList = ({ sideitem, onDrawerClose }: SideLProps) => {
  const pathname = usePathname();

  // Check if list item is active
  const isActive = sideitem.link === pathname;


  const activeClass = 'text-blue-600'; 
  const defaultClass = sideitem.available ? 'text-gray-600' : 'text-gray-400'; 
  const hoverClass = 'hover:bg-gray-200 hover:text-blue-600';



  return (
    <div>
    
    <ListItem
      key={sideitem.id}
      sx={{
        position: 'relative',
        '&::before': {
          content: '""',
          position: 'absolute',
          bgcolor: isActive ? 'primary.main' : 'transparent',
          top: 0,
          bottom: 0,
          left: 0,
          width: 6,
          borderTopRightRadius: 10,
          borderBottomRightRadius: 10,
          transition: 'background-color 0.5s ease',
        },
      }}
    >
      <Link href={sideitem.link} passHref style={{ textDecoration: 'none', flex: 1 }}>
        <div
          onClick={onDrawerClose}
          className={`flex items-center gap-5 py-2 px-2 rounded-lg min-w-5 ${isActive ? activeClass : defaultClass} ${hoverClass}`}
          style={{
            color: isActive ? 'primary.main' : sideitem.available ? '#757575' : '#BDBDBD',
            transition: 'color 0.35s ease',
          }}
          onMouseOver={(e) => {
            if (sideitem.available) {
              e.currentTarget.style.backgroundColor = '#f5f5f5';
              e.currentTarget.style.color = '#3f51b5';
            }
          }}
          onMouseOut={(e) => {
            if (sideitem.available) {
              e.currentTarget.style.backgroundColor = 'transparent';
              e.currentTarget.style.color = isActive ? '#3f51b5' : '#757575';
            }
          }}
        >
          <ListItemIcon
            sx={{
              minWidth: 'auto',
              color: isActive ? 'primary.main' : sideitem.available ? 'neutral.dark' : 'action.disabled',
            }}
          >
            {sideitem.icon}
          </ListItemIcon>
          <ListItemText
            primary={
              <Typography
                sx={{
                  fontSize: { xs: 'body1.fontSize', xl: 'h6.fontSize' },
                  fontWeight: 500,
                  textTransform: 'capitalize',
                }}
              >
                {sideitem.title}
              </Typography>
            }
          />
        </div>
      </Link>
    </ListItem>
    </div>
  );
  
};

export default SideList;
