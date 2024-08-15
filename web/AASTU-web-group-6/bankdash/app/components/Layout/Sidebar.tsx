import React from 'react'
import Image from 'next/image'

const Sidebar = () => {
  return (
    <div className='py-6 px-5 flex flex-col gap-8 ablsoute left-0 border-r border-r-[#E6EFF5] bg-white'>
      <div className='flex gap-2'>
        <Image
          src="/assets/logo.svg"
          alt="logo"
          width={36}
          height={36}
        />
        <h1 className='text-3xl font-bold text-[#343C6A]'>BankDash</h1>
      </div>

      {/* Menu */}
      <div className='px-8 flex flex-col gap-y-8'>

        <div className="flex items-center gap-x-6 ">
          <Image
          src="/assets/home 2.svg"
          alt="logo"
          width={25}
          height={25}
          />
          <h2 className="text-base text-[#b1b1b1] font-light ">Dashboard</h2>
        </div>

        <div className="flex items-center gap-x-6 ">
          <Image
          src="/assets/transfer 1.svg"
          alt="logo"
          width={25}
          height={25}
          />
          <h2 className="text-base text-[#b1b1b1] font-light ">Transactions</h2>
        </div>
        <div className="flex items-center gap-x-6 ">
          <Image
          src="/assets/user 3 1.svg"
          alt="logo"
          width={25}
          height={25}
          />
          <h2 className="text-base text-[#b1b1b1] font-light ">Accounts</h2>
        </div>
        <div className="flex items-center gap-x-6 ">
          <Image
          src="/assets/economic-investment 1.svg"
          alt="logo"
          width={25}
          height={25}
          />
          <h2 className="text-base text-[#b1b1b1] font-light ">Investments</h2>
        </div>
        <div className="flex items-center gap-x-6 ">
          <Image
          src="/assets/credit-card 1.svg"
          alt="logo"
          width={25}
          height={25}
          />
          <h2 className="text-base text-[#b1b1b1] font-light ">Credit Cards</h2>
        </div>
        <div className="flex items-center gap-x-6 ">
          <Image
          src="/assets/loan 1.svg"
          alt="logo"
          width={25}
          height={25}
          />
          <h2 className="text-base text-[#b1b1b1] font-light ">Loans</h2>
        </div>
        <div className="flex items-center gap-x-6 ">
          <Image
          src="/assets/service 1.svg"
          alt="logo"
          width={25}
          height={25}
          />
          <h2 className="text-base text-[#b1b1b1] font-light ">Services</h2>
        </div>
        <div className="flex items-center gap-x-6 ">
          <Image
          src="/assets/econometrics 1.svg"
          alt="logo"
          width={25}
          height={25}
          />
          <h2 className="text-base text-[#b1b1b1] font-light ">My Privileges</h2>
        </div>
        <div className="flex items-center gap-x-6 ">
          <Image
          src="/assets/settings solid 1.svg"
          alt="logo"
          width={25}
          height={25}
          />
          <h2 className="text-base text-[#b1b1b1] font-light ">Setting</h2>
        </div>

      </div>
    </div>
  )
}

export default Sidebar