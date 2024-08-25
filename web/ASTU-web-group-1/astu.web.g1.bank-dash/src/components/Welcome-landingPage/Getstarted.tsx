'use client';
import React from 'react';
import { useSession } from 'next-auth/react';
import Link from 'next/link';

export default function GetStarted() {
  const session = useSession();

  return (
    <>
      {session?.data ? (
        <Link
          href='/bank-dash'
          className='rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600'
        >
          Go to Dashboard
        </Link>
      ) : (
        <Link
          href='/api/auth/signin'
          className='rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600'
        >
          Get started
        </Link>
      )}
    </>
  );
}
