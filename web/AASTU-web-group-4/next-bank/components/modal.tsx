import { Button } from "@/components/ui/button"
import {currentuser} from "@/services/userupdate"
import {useEffect, useState } from "react"
import { DocumentDuplicateIcon ,PencilIcon } from "@heroicons/react/24/outline"
import Image from "next/image"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"

import { Label } from "@/components/ui/label"
import { FaEye, FaEyeSlash } from "react-icons/fa"

export function DialogDemo() {
  const [info, setinfo] = useState([]);
  const [currentPage, setCurrentPage] = useState(0);
  const [visible , setvisible] = useState(false)
  useEffect(() => {
    const fetch = async () => {
      try {
        const data = await currentuser();
        setinfo(data.data || []);
      } catch (error) {
        console.error("Error:", error);
      }
    };
    fetch();
  } , []);

  const handleVisibility = () => {
    setvisible(!visible);

    
  }
  const dateOfBirth = new Date(info.dateOfBirth);
const options = { month: 'long', day: 'numeric', year: 'numeric' };
const formattedDateOfBirth = dateOfBirth.toLocaleDateString('en-US' , options);
  return (
    <div className="p-0">
      <DialogContent className="sm:max-w-[425px] max-h-full  ">
      <div className="bg-gray-200 h-2/6 rounded-md"> {/* Add the gray container */}
      <div className=" profile-header flex justify-between user-profile-container">
      <div className="">
      <div className="pt-3 px-6">
        <Image
                src="/Images/profilepic.jpeg"
                alt="User Profile"
                width={50}
                height={50}
                className="rounded-full aspect-square object-cover cursor-pointer "
              />
          <div className="pt-3">
            <h2 className="user-name font-bold text-lg">{info.name}</h2>
            <p className="user-username text-sm">@{info.email}</p>
          </div>
      </div>
      
      </div>

      <div className="profile-actions flex-col gap-3 items-center pt-10 ">
        
        {visible && (
         <div className='text-center text-xl font-bold p-4 flex justify-center gap-3 overflow-x-hidden '>
            <p >
              ${info.accountBalance}
            </p>
            <button onClick={handleVisibility}> <FaEyeSlash className='text-gray-400 hover:text-gray-600 text-2xl' /> </button>
          </div>
        )
         }
         {
          !visible && (
            <div className='text-center text-xl font-bold p-4 flex justify-center gap-3  '>
            <p >
              $****
            </p>
            <button onClick={handleVisibility}> <FaEye className='text-gray-400 hover:text-gray-600 text-2xl' /> </button>
          </div>
          )
         }
       <div className="flex gap-1 pr-2">
         {/* <button className="copy-link-button border-2 px-2 rounded-md flex gap-1 items-center text-sm hover:bg-green-600 hover:text-white"> <> <DocumentDuplicateIcon className="w-4 h-4 text-gray-500 hover:text-white"/> Copy link </></button> */}
               <div><button className="view-profile-button border-2 px-2 rounded-md flex gap-1 items-center text-sm hover:bg-yellow-500 hover:text-white"> <> <PencilIcon className="w-4 h-4"/> Edit profile</></button></div>
       </div>
      </div>

    </div>
    </div>
    <div className="profile-actions">
       
      </div>

      <div className="flex items-center justify-evenly pb-6 text-sm">
       
       
        <div className="profile-body-content px-4">
          <div className="profile-body-content-item flex gap-2 py-2">
            <p className="font-bold ">Username:</p>
            <p>{info.username}</p>
          </div>
          <div className="profile-body-content-item flex gap-2 py-2">
            <p className="font-bold">Dateofbirth:</p>
            
            <p>{formattedDateOfBirth}</p>
            
          </div>
          <div className="profile-body-content-item flex gap-2 py-2 ">
            <p className="font-bold">Currency:</p>
            <p>USD</p>
          </div>
         
        </div>
      
      <div className=" w-0.5 h-36 bg-gray-200 rounded-full"></div>
      <div className="profile-body-content px-4 py-4">
          <div className="profile-body-content-item flex gap-2 py-2">
            <p className="font-bold">Country:</p>
            <p>{info.country}</p>
          </div>
          <div className="profile-body-content-item flex gap-2 py-2">
            <p className="font-bold" >City:</p>
            <p>{info.city}</p>
          </div>
          <div className="profile-body-content-item flex gap-2 py-2 ">
            <p className="font-bold">Address:</p>
            <p>{info.presentAddress}</p>
          </div>
         
        </div>
      </div>
          {/* <DialogFooter>
            <Button type="submit" className="">Save changes</Button>
          </DialogFooter> */}
        </DialogContent>
    </div>
  )
}
