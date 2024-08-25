'use client';
import { Landmark } from 'lucide-react';
import Link from 'next/link';
import { useState } from 'react';
import { useSession } from 'next-auth/react';

const navigation = [
  { name: 'Product', href: '#' },
  { name: 'Features', href: '#' },
  { name: 'Marketplace', href: '#' },
  { name: 'Company', href: '#' },
];

const Navbar = () => {
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false);
  const session = useSession();
  console.log('session from navbar', session);

  return (
    <>
      <header className='absolute inset-x-0 top-0 z-50 md:mx-20'>
        <nav aria-label='Global' className='flex items-center justify-between p-6 lg:px-8'>
          <div className='flex lg:flex-1'>
            <Link href='/' className='-m-1.5 p-1.5'>
              <Landmark />
            </Link>
          </div>
          <div className='flex lg:hidden'>
            <button
              type='button'
              onClick={() => setMobileMenuOpen(true)}
              className='-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700'
            >
              <span className='sr-only'>Open main menu</span>
            </button>
          </div>
          <div className='hidden lg:flex lg:gap-x-12'>
            {navigation.map((item) => (
              <a
                key={item.name}
                href={item.href}
                className='text-sm font-semibold leading-6 text-gray-900'
              >
                {item.name}
              </a>
            ))}
          </div>
          <div className='hidden lg:flex lg:flex-1 lg:justify-end'>
            {session.data ? (
              <Link href='/bank-dash' className='text-sm font-semibold leading-6 text-gray-900'>
                Dashboard <span aria-hidden='true'>&rarr;</span>
              </Link>
            ) : (
              <Link
                href='/api/auth/signin'
                className='text-sm font-semibold leading-6 text-gray-900'
              >
                Log in <span aria-hidden='true'>&rarr;</span>
              </Link>
            )}
          </div>
        </nav>
      </header>
    </>
  );
};

export default Navbar;
