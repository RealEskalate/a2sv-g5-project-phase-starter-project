"use client";

import React from 'react';
import { useRouter, usePathname } from 'next/navigation';

interface MenuItem {
  label: string;
  url: string;
}

const Navigation: React.FC = () => {
  const router = useRouter();
  const pathname = usePathname();

  const menuItems: MenuItem[] = [
    { label: 'Edit Profile', url: '/settings/editprofile' },
    { label: 'Preferences', url: '/settings/preference' },
    { label: 'Security', url: '/settings/security' },
  ];

  const isActive = (path: string) => pathname === path;

  return (
    <div className="flex justify-between items-center border-b py-4 px-2 min-h-6 max-h-20">
      <div className="flex space-x-8">
        {menuItems.map((item, index) => (
          <button
            key={index}
            onClick={() => router.push(item.url)}
            className={`pb-2 border-b-4 rounded cursor-pointer ${
              isActive(item.url) ? 'border-blue-800' : 'border-transparent'
            }`}
          >
            {item.label}
          </button>
        ))}
      </div>
    </div>
  );
};

export default Navigation;
