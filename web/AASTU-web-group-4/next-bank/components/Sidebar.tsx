import { FC } from 'react';
import Link from 'next/link';
import { sidebarLinks } from '@/constants';
import { usePathname } from 'next/navigation';

type SidebarProps = {
  isOpen: boolean;
  toggleSidebar: () => void;
};

export const Sidebar: FC<SidebarProps> = ({ isOpen, toggleSidebar }) => {
  const pathname = usePathname();

  return (
    <>
      {/* Sidebar for Mobile Screens */}
      <div
        className={`fixed top-0 left-0 z-40 flex flex-col w-64 h-full bg-white shadow-lg p-4 transition-transform duration-300
        ${isOpen ? 'translate-x-0' : '-translate-x-full'} lg:hidden`}
      >
        {/* Close Button - Only visible on mobile */}
        <button
          onClick={toggleSidebar}
          aria-label="Close Sidebar"
          className="absolute top-4 right-4 p-2 text-gray-600 hover:text-gray-800 focus:outline-none"
        >
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="2" stroke="currentColor" className="w-6 h-6">
            <path strokeLinecap="round" strokeLinejoin="round" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>

        <div className="text-2xl font-bold mb-6">Next Bank</div>
        <div className="space-y-4 flex-1">
          {sidebarLinks.map((link) => (
            <Link
              key={link.route}
              href={link.route}
              className={`flex items-center p-2 rounded-lg cursor-pointer
              ${pathname === link.route ? 'text-blue-600 border-l-4 border-blue-600' : 'text-gray-600'}`}
            >
              <link.Icon className="mr-3" size={25} />
              <span>{link.label}</span>
            </Link>
          ))}
        </div>
      </div>

      {/* Sidebar for Large Screens */}
      <div
        className="hidden lg:flex fixed top-0 left-0 z-20 flex flex-col w-64 h-full bg-white shadow-lg p-4"
      >
        <div className="text-2xl font-bold mb-6">Next Bank</div>
        <div className="space-y-4 flex-1">
          {sidebarLinks.map((link) => (
            <Link
              key={link.route}
              href={link.route}
              className={`flex items-center p-2 rounded-lg cursor-pointer
              ${pathname === link.route ? 'text-blue-600 border-l-4 border-blue-600' : 'text-gray-600'}`}
            >
              <link.Icon className="mr-3" size={25} />
              <span>{link.label}</span>
            </Link>
          ))}
        </div>
      </div>
    </>
  );
};
