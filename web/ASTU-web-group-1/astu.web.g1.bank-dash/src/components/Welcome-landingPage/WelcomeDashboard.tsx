import Link from 'next/link';

import HeroImage from './HeroImage';
import Navbar from './navbar';
import { authOptions } from '@/app/api/auth/[...nextauth]/options';
import { getServerSession } from 'next-auth';
import AuthProvider from '../Formx/AuthProvider';
import { getSession } from 'next-auth/react';
import GetStarted from './Getstarted';

const WelcomeDashboard = async () => {
  const session: any = await getSession();
  return (
    <div className='bg-white pb-20'>
      <AuthProvider>
        <Navbar />
      </AuthProvider>

      <div className='relative isolate px-6 pt-6 lg:px-8'>
        <div
          aria-hidden='true'
          className='absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80'
        >
          <div
            style={{
              clipPath:
                'polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)',
            }}
            className='relative left-[calc(50%-11rem)] aspect-[1155/678] w-[36.125rem] -translate-x-1/2 rotate-[30deg] bg-gradient-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%-30rem)] sm:w-[72.1875rem]'
          />
        </div>

        <div className='mx-auto max-w-2xl py-12 sm:py-20 lg:py-24'>
          <div className='hidden sm:mb-8 sm:flex sm:justify-center'></div>
          <div className='text-center'>
            <h1 className='text-4xl font-bold tracking-tight text-gray-900 sm:text-6xl'>
              Empowering Your Finances, Securing Your Future.
            </h1>
            <p className='mt-6 text-lg leading-8 text-gray-600'>
              Anim aute id magna aliqua ad ad non deserunt sunt. Qui irure qui lorem cupidatat
              commodo. Elit sunt amet fugiat veniam occaecat fugiat aliqua.
            </p>
            <div className='mt-10 flex items-center justify-center gap-x-6'>
              <AuthProvider>
                <GetStarted />
              </AuthProvider>
            </div>
          </div>
        </div>

        <div
          aria-hidden='true'
          className='absolute inset-x-0 top-[calc(100%-13rem)] -z-10 transform-gpu overflow-hidden blur-3xl sm:top-[calc(100%-30rem)]'
        >
          <div
            style={{
              clipPath:
                'polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)',
            }}
            className='relative left-[calc(50%+3rem)] aspect-[1155/678] w-[36.125rem] -translate-x-1/2 bg-gradient-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%+36rem)] sm:w-[72.1875rem]'
          />
        </div>
      </div>
      <HeroImage />
    </div>
  );
};

export default WelcomeDashboard;
