/* eslint-disable react/no-children-prop */
import Image from "next/image";
import ImageComponent from "../components/ImageComponent"
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { getSession } from 'next-auth/react';
import Refresh from "../../api/auth/[...nextauth]/token/RefreshToken";
import { getQuickTransfers } from "@/lib/api/transactionController";
import { PropsWithChildren } from 'react';
import { PostTransactionRequest } from "@/types/transactionController.interface";
import { postTransaction } from "@/lib/api/transactionController";


import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "@/components/ui/carousel"

type infoType = {
  id: string;
  name: string;
  username: string;
  city: string;
  country: string;
  profilePicture: string;
}
type Data = {
  access_token: string;
  data: string;
  refresh_token: string;
};

type SessionDataType = {
  user: Data;
};
export default function Home() {
  const ShimmerEffect = () => (
    <div className="animate-pulse flex flex-col gap-6 px-5 py-6 dark:border dark:border-[#333B69]">
      <p className="text-[#343C6A] font-bold mx-3 py-3 text-xl md:hidden bg-gray-300 rounded w-3/4 h-6"></p>
  
      {/* Carousel Shimmer */}
      <div className="flex py-4 gap-6 justify-center items-center [&::-webkit-scrollbar]:hidden">
        <Carousel className="w-full max-w-sm">
          <CarouselContent className="flex gap-6">
            {[...Array(3)].map((_, index) => (
              <CarouselItem
                key={index}
                className="flex flex-shrink-0 basis-1/3 pr-3 border-transparent"
              >
                <div className="bg-gray-300 rounded-full w-20 h-20"></div>
              </CarouselItem>
            ))}
          </CarouselContent>
        </Carousel>
      </div>
  
      <div className="flex justify-between items-center text-nowrap mt-6 gap-8">
        <p className="bg-gray-300 rounded w-1/5 h-4 mx-8"></p>
        <div className="flex gap-8 rounded-full bg-gray-300 w-1/2 mx-8 h-8"></div>
      </div>

    </div>
  );
  
  const [session, setSession] = useState<Data | null>(null);
  const [access_token, setAccess_token] = useState("");
  const router = useRouter();
  const [loading, setLoading] = useState(true);
  const [transfer, setQuickTransfer] = useState<infoType[]>([]);
  const [activeIndex, setActiveIndex] = useState(-1);
  const [amount, setAmount] = useState<string>("");
  const [showDialog, setShowDialog] = useState(false);

  useEffect(() => {
    const fetchSession = async () => {
      try {
        const sessionData = (await getSession()) as SessionDataType | null;
        setAccess_token(await Refresh());
        if (sessionData && sessionData.user) {
          setSession(sessionData.user);
        } else {
          router.push(`./api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")}`);
        }
      } catch (error) {
        console.error("Error fetching session:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchSession();
  }, [router]);

  useEffect(() => {
    const addingData = async () => {
      if (!access_token) return;
      if (access_token) {
        const transfers = await getQuickTransfers(100, access_token);
        console.log("Fetching Completeddddd", transfers);

        setQuickTransfer(transfers.data); // Set the content array
      }
    };
    addingData();
  }), [];

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
  
    if (activeIndex < 0 || !transfer[activeIndex]) {
      alert("You have to choose a user and Enter data.")
      console.error('Invalid index or no transfer data available');
      return;
    }
    else{
      console.log(amount)
    }
  
    const transactionDetails: PostTransactionRequest = {
      type: 'transfer', // Or dynamically set based on your application logic
      description: `Transfer to ${transfer[activeIndex].username}`,
      amount: parseInt(amount),
      receiverUserName: transfer[activeIndex].username,
    };
  
    try {
      console.log(transactionDetails);
      const response = await postTransaction(transactionDetails, access_token);
      if (response.success === true){
        // setShowDialog(true); // Show the dialog on success
        console.log()
        alert("success")
        console.log('Transaction successful:', response);
      }
      else{
        alert('NO'); // Display success message
        console.log(response.message)
      }
    } catch (error) {
      console.error('Error posting transaction:', error);
    }
  };

  const newLocal = "M1 1L7.5 7.5L1 14";
  return (
    <div className = " border rounded-3xl my-4 mx-4 bg-white dark:bg-[#020817] dark:border dark:border-[#333B69]">
      {
        loading ? (
          <ShimmerEffect/>
        ): (
          <div>
          <p className="text-[#343C6A] font-bold mx-3 py-3 text-xl md:hidden">Quick Transfers</p>
      <div className="flex flex-col gap-3 px-5 py-5">
        {/*  Image Component  */}
            <div className="flex py-2 gap-6 justify-center items-center [&::-webkit-scrollbar]:hidden">
            <Carousel       opts={{
        align: "start",
      }}
      className="w-full max-w-sm">
              <CarouselContent className="flex gap-3   ">
                {transfer.map((item,index) => (
                  <CarouselItem key={item.id}           className={`flex flex-shrink-0 basis-1/3 pr-3 ${activeIndex === index ? 'border-black border-2 rounded-3xl dark:border dark:border-white' : 'border-transparent'} `}
                  onClick={()=>{setActiveIndex(index)}}>
                    <ImageComponent
                      src = "data:image/png;base64,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"
                      alt={`${item.name}'s profile picture`}
                      width={80} // Replace with your desired width
                      height={80} // Replace with your desired height
                      name={item.name}
                      role={item.username} // Assuming `role` can be replaced with `username
                    />
                  </CarouselItem>
                ))}
                
              </CarouselContent>
              <CarouselPrevious children={undefined} className=" md:rounded-full dark:bg-white overflow-hidden"/>
              <CarouselNext children={undefined} className=" md:rounded-full dark:bg-white overflow-hidden"/>
            </Carousel>

            
            </div> 

            <div className="flex  text-[#718EBF] text-xs justify-between items-center text-nowrap ">
              <p>Write Amount</p>
              <div className="flex gap-6  rounded-full ">
              <div>
                <form onSubmit={handleSubmit} className="flex flex-1 items-center">
                  <div className="bg-[#EDF1F7] rounded-full flex items-center flex-1">
                    <input
                      className="bg-[#EDF1F7] rounded-full text-center border-none outline-none"
                      placeholder="0.00"
                      name="amount"
                      type="text"
                      value={amount}
                      onChange={(e) => setAmount(e.target.value)}
                    />
                    <button
                      type="submit"
                      className="flex bg-[#1814F3] rounded-full px-5 py-3 items-center gap-1"
                    >
                      <p className="text-white text-xs">Send</p>
                      <svg
                        width="15"
                        height="10"
                        viewBox="0 0 26 23"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                      >
                        <path
                          d="M25.9824 0.923369C26.1091 0.333347 25.5307 -0.164153 24.9664 0.0511577L0.490037 9.39483C0.195457 9.50731 0.000610804 9.78965 1.43342e-06 10.105C-0.000607937 10.4203 0.193121 10.7034 0.487294 10.817L7.36317 13.4726V21.8369C7.36317 22.1897 7.60545 22.4963 7.94873 22.5779C8.28972 22.659 8.64529 22.4967 8.80515 22.1796L11.6489 16.5364L18.5888 21.6868C19.011 22.0001 19.6178 21.8008 19.7714 21.2974C26.251 0.0528342 25.9708 0.97674 25.9824 0.923369ZM19.9404 3.60043L8.01692 12.092L2.88664 10.1106L19.9404 3.60043ZM8.8866 13.3428L19.2798 5.94118C10.3366 15.3758 10.8037 14.8792 10.7647 14.9317C10.7067 15.0096 10.8655 14.7058 8.8866 18.6327V13.3428ZM18.6293 19.8197L12.5206 15.2862L23.566 3.63395L18.6293 19.8197Z"
                          fill="white"
                        />
                      </svg>
                    </button>
                  </div>
                </form>
              </div>

            </div>
          </div>  
      </div>
      </div>
        )
      }
      
    </div>

  );
  
}
