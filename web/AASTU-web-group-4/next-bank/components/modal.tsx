import { Button } from "@/components/ui/button"
import { currentuser } from "@/services/userupdate"
import { useEffect, useState } from "react"
import { DocumentDuplicateIcon, PencilIcon } from "@heroicons/react/24/outline"
import Image from "next/image"
import { UserData } from "@/types"

import SettingsPage from "@/app/(root)/setting/page"
import Link from "next/link"
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"

import { Label } from "@/components/ui/label"
import { FaEye, FaEyeSlash } from "react-icons/fa"


// Add Tailwind CSS classes for styling
export function DialogDemo() {
  const [accountBalance, setAccountBalance ] = useState(0);
    const [info, setinfo] = useState<UserData>();
    const [visible , setvisible] = useState(false)
    useEffect(() => {
      const fetch = async () => {
        try {
          const data = await currentuser();
          setinfo(data.data || []);
          setAccountBalance(data.data.accountBalance);
        } catch (error) {
          console.error("Error:", error);
        }
      };
      fetch();
    } , []);
  
    const handleVisibility = () => {
      setvisible(!visible);
  
      
    }
    let dateOfBirth: Date;
  if (info?.dateOfBirth) {
    dateOfBirth = new Date(info.dateOfBirth);
  } else {
    dateOfBirth = new Date();
  }
  const options:any = { month: 'long', day: 'numeric', year: 'numeric' };
  const formattedDateOfBirth = dateOfBirth.toLocaleDateString('en-US' , options);
  // ... (rest of the code is the same as the previous response)

  return (
    <div className="p-0">
      <DialogContent className="sm:max-w-[425px] max-h-full">
        <div className="bg-gray-100 dark:bg-black-2 p-6 rounded-lg shadow-md"> 
          <div className="flex flex-col items-center mb-6"> {/* Center align the top section */}
            <Image
              src="/Images/profilepic.jpeg"
              alt="User Profile"
              width={80} // Increased image size
              height={80}
              className="rounded-full mb-3 aspect-square object-cover"
            />
            <h2 className="text-xl font-bold mb-1">{info?.name}</h2> {/* Reduced font size */}
            <p className="text-gray-500 text-sm">@{info?.email}</p> {/* Reduced font size */}
          </div>

          <div className="flex justify-between items-center mb-6"> {/* Added spacing */}
            <div className="text-center">
              {visible ? (
                <div className="flex items-center justify-center gap-2">
                  <p className="text-lg font-bold">${info?.accountBalance}</p> {/* Slightly larger font */}
                  <button onClick={handleVisibility}>
                    <FaEyeSlash className="text-gray-400 hover:text-gray-600 text-xl" /> 
                  </button>
                </div>
              ) : (
                <div className="flex items-center justify-center gap-2">
                  <p className="text-lg font-bold">$****</p>
                  <button onClick={handleVisibility}>
                    <FaEye className="text-gray-400 hover:text-gray-600 text-xl" />
                  </button>
                </div>
              )}
            </div>

            <div>
              <Link href="../setting" passHref>
              <button className="bg-yellow-500 hover:bg-yellow-600 text-white font-bold py-2 px-4 rounded flex items-center">
                <PencilIcon className="w-4 h-4 mr-2" /> Edit profile
              </button>
              </Link>
              
            </div>
          </div>

          <div className="grid grid-cols-2 gap-4 text-sm"> 
            <div>
              <div className="flex items-center mb-3"> {/* Added spacing */}
                <p className="font-medium mr-2">Username:</p> {/* Reduced font weight */}
                <p className="text-sm">{info?.username}</p> {/* Reduced font size */}
              </div>
              <div className="flex items-center mb-3">
                <p className="font-medium mr-2">DoB:</p>
                <p className="text-sm">{formattedDateOfBirth}</p>
              </div>
              <div className="flex items-center mb-3">
                <p className="font-medium mr-2">Currency:</p>
                <p className="text-sm">USD</p>
              </div>
            </div>
            <div>
              <div className="flex items-center mb-3">
                <p className="font-medium mr-2">Country:</p>
                <p className="text-sm">{info?.country}</p>
              </div>
              <div className="flex items-center mb-3">
                <p className="font-medium mr-2">City:</p>
                <p className="text-sm">{info?.city}</p>
              </div>
              <div className="flex items-center mb-3">
                <p className="font-medium mr-2">Address:</p>
                <p className="text-sm">{info?.presentAddress}</p>
              </div>
            </div>
          </div>
        </div>
      </DialogContent>
    </div>
  )
}

// ... rest of your code