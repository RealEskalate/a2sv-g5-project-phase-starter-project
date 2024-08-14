import React from 'react';
import Image from 'next/image';

const NavBar = () => {
  return (
    <div className='w-full flex flex-row justify-between bg-white pl-4'>
      <div>
        <h1 className='text-3xl font-semibold text-[#343C6A]'>Overview</h1>
      </div>
      <div className='flex justify-between gap-2'>
        
        {/* Search */}
        <div className='p-4 flex gap-2 bg-[#F5F7FA] rounded-3xl'>
          <Image
            src="/assets/search.svg"
            alt="search"
            width={32}
            height={32}
          />
          <input 
          className='bg-[#F5F7FA] '
          type="text" 
          name="search" 
          id="search" 
          placeholder='Search for something'  />
        </div>

        {/* User tool */}
            <button>
              <Image
                src="/assets/settings 1.svg"
                alt="setting"
                width={32}
                height={32}
              />
            </button>
            <button>
              <Image
                src="/assets/notification.svg"
                alt="notification"
                width={32}
                height={32}
              />
            </button>
            <button>
              <Image
                src="/assets/user.svg"
                alt="user image"
                width={32}
                height={32}
              />
            </button>
      </div>
    </div>
  );
}

export default NavBar;
