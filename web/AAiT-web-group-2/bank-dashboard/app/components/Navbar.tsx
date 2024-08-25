"use client";
import React from "react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { FiMenu } from "react-icons/fi";
import { useSession } from "next-auth/react";
import { useGetCurrentUserQuery } from "@/lib/redux/api/bankApi";

const Navbar = ({ onMenuClick }: { onMenuClick: () => void }) => {
  const path = usePathname();
  const session = useSession();
  const access_token = session.data?.access_token;

  const { isLoading, isError, error, data } = useGetCurrentUserQuery(access_token as string);

  let initials = '';

  if (data) {
    const name_parts = data.data.name.split(' ');
    initials = name_parts[0][0].toUpperCase();
    if (name_parts.length >= 2) {
      initials = initials.concat(name_parts[1][0].toUpperCase());
    }
  }

  return (
<<<<<<< Updated upstream
    <div className="flex flex-col justify-center h-24 gap-2 px-3 ">
=======
    <div className="flex flex-col justify-center h-28 sm:h-24 gap-2 px-3 ">
>>>>>>> Stashed changes
      <div className="flex justify-between items-center">
        <div className="text-2xl sm:hidden ">
          <FiMenu onClick={onMenuClick} />
        </div>
        {session.status === "unauthenticated" && 
        <div className="flex gap-4 text-lg ">
          <Link className='text-custom-purple font-semibold text-lg hover:text-blue-500' href='/auth/signup'>Signup</Link>
          <Link className='text-custom-purple font-semibold text-lg hover:text-blue-500' href='/auth/signin'>Signin</Link>
        </div>
        }

        {session.status === "authenticated" && <>
        <div className="flex gap-[50px] items-center">
          <p className="text-wrap font-[600] text-[28px] text-[#343C6A]">
            {path.slice(1, 2).toUpperCase() + path.slice(2)}
          </p>
        </div>
        <div className="flex gap-2  md:gap-2 lg:gap-5 items-center">
          <div className=" hidden  sm:flex sm:gap-1  md:gap-2  lg:gap-3 ">
            <label className="flex items-center gap-3 h-[40px] bg-[#F5F7FA] rounded-3xl cursor-pointer px-5">
              <img src="/search.png" alt="" />
              <input
                className="bg-inherit w-20 md:w-full p-1 focus:outline-none "
                type="search"
                placeholder="search for something"
              />
            </label>
            <Link
              href="/settings"
              className="h-[40px] w-[40px] flex items-center justify-center rounded-full bg-[#F5F7FA]"
            >
              <img className="w-6 h-6" src="/settings.png" alt="" />
            </Link>
            <Link
              href="/notifications"
              className="h-[40px] w-[40px] flex items-center justify-center rounded-full bg-[#F5F7FA]"
            >
              <img className="w-6 h-6" src="/notification.png" alt="" />
            </Link>
          </div>
          {session.status === "authenticated" && <div className="flex justify-center ">
            <div className="flex justify-center items-center">
                {data && data.data.profilePicture &&
                <Link  href="/settings">
                  <img className="bg-yellow-50 w-20 h-20 rounded-full" src={data.data.profilePicture} alt="profile" />
                  </Link>
                }
                {data && data.data.name && !data.data.profilePicture &&
                  <div className='w-14  h-14 text-3xl font-medium rounded-full bg-gray-300 flex justify-center items-center'>
                    <p>{initials}</p>
                  </div>  
                }
              </div>
          </div>}
        </div>
        </>
        }
      </div>
      
    </div>
  );
};

export default Navbar;
